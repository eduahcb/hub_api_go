package security

import "golang.org/x/crypto/bcrypt"


func CreateHashPassword(password string) ([]byte, error) {
   return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) 
}

func ComparePasswords(hashPassword, password string) error {
  return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
