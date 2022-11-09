package root

import (
	"appital-service/business/core/models"
	"appital-service/caesar"
	"errors"
	"fmt"
)

func Run(action, text string, key int) (string, error) {
	result := ""
	switch action {
	case "encode":
		input := models.NewEncoding{
			Text: text,
			Key:  key,
		}

		encoded, err := caesar.Encrypter(input)
		if err != nil {
			return "", errors.New(fmt.Sprintf("encoding text: %v", err))
		}

		result = encoded
	case "decode":
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
