package service

import (
	"context"

	"Go/internal/model"
)

type (
	IUserLogin interface {
		Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error)
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error )
		UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error)
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserAdmin IUserAdmin
	localUserInfo IUserInfo
	localUserLogin IUserLogin
)


// admin
func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("Implement localUserAdmin not found for interface IUserAdmin")
	}

	return localUserAdmin

}

func InitUserAdmin(i IUserAdmin)  {
	localUserAdmin= i
}


// info
func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("Implement localUserInfo not found for interface IUserInfo")
	}

	return localUserInfo

}

func InitUserInfo(i IUserInfo)  {
	localUserInfo= i
}



// Login
func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("Implement localUserAdmin not found for interface IUserLogin")
	}

	return localUserLogin

}

func InitUserLogin(i IUserLogin)  {
	localUserLogin = i
}
