package tools

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}