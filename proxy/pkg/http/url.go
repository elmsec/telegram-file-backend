package http

import (
	"github.com/elmsec/telegram-file-backend/pkg/crypto"
	"os"
	"strings"
)

var key = []byte(os.Getenv("PROXY_AES_KEY"))
var iv = []byte(os.Getenv("PROXY_AES_IV"))

func ParseUrl(path string) (botToken string, fileId string, err error) {
	information := strings.SplitN(path, "/", 2)
	fileId = information[1]

	encryptedString := information[0]
	botToken, err = crypto.DecryptPayload(key, iv, encryptedString)
	if err != nil {
		return "", "", err
	}

	if err != nil {
		return "", "", err
	}

	return botToken, fileId, nil
}
