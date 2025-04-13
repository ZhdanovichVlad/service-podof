package jwttoken

import (
	"fmt"
	"time"

	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey            = "jwtsecret"
	tokenLifetimeInHours = 12
)

type JwtTokenGenerator struct {
}

func NewJwtTokenGenerator() *JwtTokenGenerator {
	return &JwtTokenGenerator{}
}

type Claims struct {
    UserID string `json:"user_id"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

func (j *JwtTokenGenerator) GenerateToken(userUUID string, role string) (string, error) {
	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"userUUID": userUUID,
		"role":     role,
		"exp":      now.Add(tokenLifetimeInHours * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accesstoken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return accesstoken, nil
}



func (j *JwtTokenGenerator) ParseToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errorsx.ErrUnexpSignedMetod
        }
        return []byte(secretKey), nil
    })
}


func (j *JwtTokenGenerator) ValidateExpiration(claims jwt.MapClaims) error {
    exp, ok := claims["exp"].(float64)
    if !ok {
        fmt.Println("err 1")
        return errorsx.ErrInvalidToken
    }
    
    if time.Now().Unix() > int64(exp) {
        return errorsx.ErrTokenExpired
    }
    
    return nil
}


func (j *JwtTokenGenerator) GetClaims(token *jwt.Token) (jwt.MapClaims, error) {
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        fmt.Println("err 2")
        return nil, errorsx.ErrInvalidToken
    }
    return claims, nil
}


func (j *JwtTokenGenerator) ValidateToken(tokenString string) (*Claims, error) {
 
    token, err := j.ParseToken(tokenString)
    if err != nil {
        fmt.Println("err 3", err)
        return nil, errorsx.ErrInvalidToken
    }

   
    claims, err := j.GetClaims(token)
    if err != nil {
        return nil, err
    }


    if err := j.ValidateExpiration(claims); err != nil {
        return nil, err
    }

    userID, ok := claims["userUUID"].(string)
    if !ok {
        fmt.Println("err 4")
        return nil, errorsx.ErrInvalidToken
    }

    role, ok := claims["role"].(string)
    if !ok {
        fmt.Println("err 5")
        return nil, errorsx.ErrInvalidToken
    }

    return &Claims{
        UserID: userID,
        Role:   role,
    }, nil
}