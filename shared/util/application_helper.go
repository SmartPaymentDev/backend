package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	cfg "smart-payment-dev-be/config"
	"time"
)

type JwtClaims struct {
	CUSTID int    `json:"custid"`
	NOCUST string `json:"nocust"`
	NMCUST string `json:"nmcust"`
	NO_WA  string `json:"no_wa"`
	jwt.StandardClaims
}

func ParseJwtToken(token string) (string, error) {
	tokenString := token
	cfg := cfg.GetConfig()

	secretKey := []byte(cfg.SECREET_KEY)

	respToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token's signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the key for validation
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	// Check if the token is valid and extract claims
	if claims, ok := respToken.Claims.(jwt.MapClaims); ok && respToken.Valid {
		if value, exists := claims["custid"]; exists {
			return fmt.Sprintf("%v", value), nil
		}
		return "", errors.New("CAN'T FIND USER")
	} else {
		return "", errors.New("CAN'T FIND USER")
	}

}

func CreateJwtToken(custId int, noCust, nmCust, noWa string) (string, error) {
	cfg := cfg.GetConfig()

	claimss := JwtClaims{
		custId,
		noCust,
		nmCust,
		noWa,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 320).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claimss)
	token, err := rawToken.SignedString([]byte(cfg.SECREET_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ComparePassword(passInput, originalHashed string) bool {
	inputHashed := EncrypCode(passInput)

	// Compare the hashed passwords
	if originalHashed == inputHashed {
		return true
	}
	return false
}

func EncrypCode(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func TransactonNo(count int) string {
	date := GetCurrentDate()
	return fmt.Sprintf(`MOBILE%d%02d%02d%05d`, date.Year(), date.Month(), date.Day(), count)
}
