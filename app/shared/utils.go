package shared

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateAccessToken(s string) []byte {
	str := s + fmt.Sprint(time.Now().UnixNano())
	token, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return token
}
