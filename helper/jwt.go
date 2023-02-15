package helper

import (
	"ActiveCitizenRESTAPI/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"
	"time"
)

var privateKey = []byte("secretkey")
var tokenKey string = "2000"

func GenerateJWT(user models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(tokenKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID,
		"iat":    time.Now().Unix(),
		"eat":    time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
		"roleId": user.Roleid,
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(context *gin.Context) error {
	token, err := GetToken(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func CurrentUser(context *gin.Context) models.User {
	err := ValidateJWT(context)
	if err != nil {
		return models.User{}
	}
	token, _ := GetToken(context)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	user, err := models.FindUserById(userId)
	if err != nil {
		return models.User{}
	}
	return user
}

func GetToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := GetTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func GetTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
