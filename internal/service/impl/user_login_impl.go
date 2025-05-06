package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"Go/global"
	consts "Go/internal/const"
	"Go/internal/database"
	"Go/internal/model"
	"Go/internal/utils"
	"Go/internal/utils/auth"
	"Go/internal/utils/crypto"
	"Go/internal/utils/random"
	"Go/internal/utils/sendto"
	"Go/pkg/response"
)

type sUserLogin struct {
	// implement the IUserLogin interface here
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin{

	return &sUserLogin{r: r}

}

func (s *sUserLogin) Login(ctx context.Context, in *model.LoginInput )  ( codeResult int, out model.LoginOutput ,err error) {
	// logic login
	userBase, err := s.r.GetOneUserInfo(ctx, in.UserAccount)
	if err != nil {
        return response.ErrCodeAuthFailed, out, err
    }

	// 2. check password
	if !crypto.MatchingPassword(userBase.UserPassword, in.UserPassword, userBase.UserSalt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("does not match password")
	}
	// 3. Check two-factor authentication

	// 4. update password time
	go s.r.LoginUserBase(ctx, database.LoginUserBaseParams{
		UserLoginIp: sql.NullString{String: "127.0.0.1", Valid: true},
        UserAccount: in.UserAccount,
	})

	// 5. Create UUUID user
	subToken := utils.GenerateCliTokenUUID(int(userBase.UserID))
	log.Println("subToken:", subToken)

	// 6. get user_info table
	infoUser, err := s.r.GetUser(ctx, uint64(userBase.UserID));

	if err != nil {
        return response.ErrCodeAuthFailed, out, err
    }

	// convert to json
	infoUserJson, err := json.Marshal(infoUser)
    if err != nil {
        return response.ErrCodeAuthFailed, out, fmt.Errorf("convert to json failed: %v", err)
    }
	// 7. give infoUserJson to redis with key = subtoken
	err = global.Rdb.Set(ctx, subToken, infoUserJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()

	if err != nil {
        return response.ErrCodeAuthFailed, out, err
    }

	// 8. create token
    out.Token, err = auth.CreateToken(subToken)
    if err != nil {
        return response.ErrCodeAuthFailed, out, err
    }


	return 200, out, nil
}
func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput)  ( codeResult int, err error) {
	// logic
	// 1. hash email
	fmt.Printf("VerifyKey: %s\n", in.VerifyKey)
	fmt.Printf("VerifyKey: %d\n", in.VerifyType)
	hashkey :=crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("HashKey: %s\n", hashkey)
	// 2. check email existence in db
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)

	if err != nil {
        return response.ErrCodeUserHasExits, err
    }

	if userFound > 0 {
		return response.ErrCodeUserHasExits, fmt.Errorf("user has already registered")
	}

	// 3. create OTP
	userKey := utils.GetUserKey(hashkey)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()

	switch {
		case err == redis.Nil:
			fmt.Println("Key does not exist")
            
        case err != nil:
			fmt.Println("Get failed::", err)
            return response.ErrInvalidOTP, err
		case otpFound != "":
			return response.ErrCodeUserHasExits, fmt.Errorf("")
	}
	// 4. Generate OTP
	otpNew := random.GenerateSixDigitOtp()

	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("OTP is :::%d\n", otpNew)
	// 5. save OTP in redis with expiration time

	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)* time.Minute).Err()

	if err != nil {
        return response.ErrInvalidOTP ,err
    }
	// 6. send OTP to email
	switch in.VerifyType {
	case consts.EMAIL:
		err = sendto.SendTextEmailOtp([]string{in.VerifyKey}, consts.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
            return response.ErrSendEmailOtp, err
        }

		// 7. save OTP to db

		result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
            VerifyKey:     in.VerifyKey,
            VerifyKeyHash: hashkey ,
            VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
		})
		if err != nil {
            return response.ErrInvalidOTP, err
        }
		// 8. getLastID
		lastIdVerifyUser, err := result.LastInsertId()

		if err != nil {
            return response.ErrSendEmailOtp, err
        }
        log.Println("lastIdVerifyUser", lastIdVerifyUser)
       	return response.ErrCodeSuccess, nil

	case consts.MOBILE:
		return response.ErrCodeSuccess, nil
	}

	return response.ErrCodeSuccess, nil

}
func (s *sUserLogin) VerifyOTP(ctx context.Context, in *model.VerifyInput)  (out model.VerifyOTPOutput, err error) {
	// logic
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// get otp
	otpFound, err := global.Rdb.Get(ctx, utils.GetUserKey(hashKey)).Result()

	if err != nil {
		return out, err
	}

	if in.VerifyCode != otpFound {
		// nếu như sai 3 lần trong 1 phút


		return out, fmt.Errorf("OTP not match")
	}
	infoOTP, err := s.r.GetInfoOTP(ctx, hashKey)

	if err != nil {
        return out, err
    }

	// update status verified
	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)

    if err != nil {
        return out, err
    }

    // output
    out.Token = infoOTP.VerifyKeyHash
	out.Message = "success"
	return out, err
}
func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error) {
	// token is already verified
	infoOTP, err := s.r.GetInfoOTP(ctx, token)

	if err != nil {
        return response.ErrCodeUserOtpNotExists, err
    }

	// check isVerified OK

	if infoOTP.IsVerified.Int32 == 0 {
        return response.ErrCodeUserOtpNotExists, fmt.Errorf("user OTP not verified")
    }

	// update status
	userBase := database.AddUserBaseParams{}

	userBase.UserAccount = infoOTP.VerifyKey
	userSalt, err := crypto.GenetareSalt(16)

	if err != nil {
        return response.ErrCodeUserOtpNotExists, err
    }

	userBase.UserSalt = userSalt
	userBase.UserPassword = crypto.HashPassword(password, userSalt)
	// add userBase to user_base table
	newUserBase, err := s.r.AddUserBase(ctx, userBase)

	if err != nil {
        return response.ErrCodeUserOtpNotExists, err
    }

	user_id, err := newUserBase.LastInsertId()

	if err != nil {
        return response.ErrCodeUserOtpNotExists, err
    }

	// add user_id to user_info table

	newUserInfo, err := s.r.AddUserHaveUserId(ctx, database.AddUserHaveUserIdParams{
		UserID:   int64(user_id),
        UserAccount: infoOTP.VerifyKey,
		UserNickname: sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserAvatar: sql.NullString{String: "", Valid: true},
		UserState: 1,
		UserMobile: sql.NullString{String: "", Valid: true},
		UserGender: sql.NullInt16{Int16: 0, Valid: true},
		UserBirthday: sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail: sql.NullString{String: "", Valid: true},
		UserIsAuthentication: 1,
	}) 
	if err != nil {
        return response.ErrCodeUserOtpNotExists, err
    }

	user_id, err = newUserInfo.LastInsertId() 
	if err != nil {
        return response.ErrCodeUserOtpNotExists, err
    }


	return int(user_id), nil
}