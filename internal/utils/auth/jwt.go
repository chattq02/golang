package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"Go/global"

)

type PayloadClaims struct { 
	jwt.StandardClaims
}

func GenTokenJWT(payload jwt.Claims) (string, error)  {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(global.Config.JWT.API_SECRET_KEY))
}

func CreateToken(uuidToken string) (string, error) {
	// 1. set thời gian hết hạn
	timeEx := global.Config.JWT.JWT_EXPIRATION

	if timeEx == "" {
		timeEx = "1h"
	}
	expiration, err := time.ParseDuration(timeEx)

	if err != nil {
        return "", err
    }
	now := time.Now()
	expiresAt := now.Add(expiration)

	return GenTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id: uuid.New().String(), //ID duy nhất của token.
			ExpiresAt: expiresAt.Unix(), //Thời gian hết hạn.
			IssuedAt: now.Unix(), //Thời gian phát hành.
			Issuer: "shopdevgo", //Tổ chức phát hành token.
			Subject: uuidToken, //Chủ thể của token (thường là ID người dùng hoặc thông tin liên quan).
		},
	})

}


func ParseJwtTokenSubject(token string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
        return []byte(global.Config.JWT.API_SECRET_KEY), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims ,nil
		}
	}
	return nil, err
}

// validateToken

func VerifyTokenSubject(token string) (*jwt.StandardClaims, error) {
    claims, err := ParseJwtTokenSubject(token) 
	if err != nil {
        return nil, err
    }

	if err = claims.Valid(); err != nil { 
		return nil, err
	}
	return claims, nil
}