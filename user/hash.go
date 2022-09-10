package user

import "golang.org/x/crypto/bcrypt"

func hash(pswd []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pswd, bcrypt.DefaultCost)
}

func isAuthorized(pswd, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, pswd) == nil
}
