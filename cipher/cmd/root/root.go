package root

import (
	"errors"
	"fmt"
	"service/business/core/models"
	"service/caesar"
)

func Run(action, text string, key int) (string, error) {
	result := ""
	switch action {
	case "encrypt":
		input := models.NewEncoding{
			Text: text,
			Key:  key,
		}

		encoded, err := caesar.Encrypter(input)
		if err != nil {
			return "", errors.New(fmt.Sprintf("encoding text: %v", err))
		}

		result = encoded
	case "decrypt":
		input := models.NewDecoding{
			Text: text,
			Key:  key,
		}
		decoded, err := caesar.Decrypter(input)
		if err != nil {
			return "", errors.New(fmt.Sprintf("decoding text: %v", err))
		}

		result = decoded
	default:
		return "", errors.New("unknown action requested")
	}

	return result, nil
}
