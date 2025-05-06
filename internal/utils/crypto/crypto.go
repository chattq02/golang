package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes)
}

// GenetareSalt a random salt
func GenetareSalt(length int) (string, error) { 
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil  // return salt as a hex string
}

// hash password 
func HashPassword(password string, salt string) string {
	// concatenate password and salt
	saltedPassword := password + salt

    // hash the combined string
    hashedPassword := sha256.Sum256([]byte(saltedPassword))

    return hex.EncodeToString(hashedPassword[:])
}

func MatchingPassword(storeHash string, password string ,salt string) bool {
	hashpassword := HashPassword(password,salt)
	return storeHash == hashpassword
}