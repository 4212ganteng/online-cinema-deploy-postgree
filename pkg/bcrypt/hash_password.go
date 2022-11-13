package bcryptpkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashing, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "failed hash", err
	}
	return string(hashing), nil
}

func CheckPasswordHash(password, hashedpass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(password))
	return err == nil
}
