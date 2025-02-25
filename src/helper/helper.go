package helper

import(
	"golang.org/x/crypto/bcrypt"
)
func HashingPassword(pass string) (string, error)  {
	pwhash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	
	return string(pwhash), err
}