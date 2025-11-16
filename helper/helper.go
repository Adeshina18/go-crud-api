package helper

import (
	"errors"
	"github/AdeleyeShina/models"
	"regexp"

	"github.com/google/uuid"
)

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func ValidateUserInput(user models.User) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	emailRegex := regexp.MustCompile(regex)

	if !emailRegex.MatchString(user.Email) {
		return errors.New("invalid email address")
	}

	if len(user.Password) < 6 {
		return errors.New("password cannot be less than 6 character")
	}
	return nil
}
