package crypto

import (
	"strings"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

//Hash implements root.Hash
type Hash struct{}

var deliminator = "||"

//Generate a salted hash for the input string
func (c *Hash) Generate(s string) (string, error) {
	salt := uuid.New().String()
	saltedBytes := []byte(s + salt)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash + deliminator + salt, nil
}

//Compare string to generated hash
func (c *Hash) Compare(hash string, s string) error {
	parts := strings.Split(hash, deliminator)
	if len(parts) != 2 {
		return errors.New("Invalid hash, must have 2 parts")
	}

	incoming := []byte(s + parts[1])
	existing := []byte(parts[0])
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

