package bcryptclient

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func authenticate(userID string, password string) error {
	hashedPassword, err := getHashedPassword(userID)
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	}
	return err
}

func getHashedPassword(userID string) (string, error) {
	var hashedPassword string
	if userID == "0" {
		hashedPassword = "$2a$12$yq6qOGgCBzUFukaXIPuEUeBcaglmV35B0FKrps.fInksuzfLsF/iC"
	} else if userID == "1" {
		hashedPassword = "$2a$12$yq6qOGgCBzUFukaXIPuEUeBcaglmV35B0FKrps.fInksuzfLsF/iC"
	}
	if hashedPassword == "" {
		return hashedPassword, fmt.Errorf("User '%v' not found", userID)
	}
	return hashedPassword, nil
}
