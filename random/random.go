package random

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/pkg/errors"
)

var ErrorGeneratingRandomSeed = errors.New("error generating random seed")

func GenerateRandomString(lenght int) (string, error) {
	randomBytes := make([]byte, lenght)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", errors.Wrap(err, ErrorGeneratingRandomSeed.Error())
	}
	encoded := base64.URLEncoding.EncodeToString(randomBytes[:lenght])[:lenght]
	return encoded, nil
}
