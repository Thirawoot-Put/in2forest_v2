package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPwd(pwd string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	return string(byte), err
}

func CheckPwd(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
