package http

import (
	"os"
	"strings"

	"github.com/elmsec/telegram-file-backend/pkg/crypto"
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

	return botToken, fileId, nil
}
