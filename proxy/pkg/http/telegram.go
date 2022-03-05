package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elmsec/telegram-file-backend/pkg/database"
	"github.com/elmsec/telegram-file-backend/pkg/entities"
	"log"
	"net/http"
	"time"
)

var redisClient = database.NewRedisClient()
var ctx = context.Background()
var BASE_URL = "https://api.telegram.org/bot"

func GetFile(botToken, fileId string) (*entities.TelegramResponse, error) {
	cachedFile, _ := GetCachedFile(fileId)
	if cachedFile != nil {
		return cachedFile, nil
	}

	return GetRemoteFile(botToken, fileId)
}

func GetCachedFile(fileId string) (*entities.TelegramResponse, error) {
	var fileObj entities.TelegramResponse
	fileStr, err := redisClient.Get(ctx, fileId).Result()

	if err != nil {
		return nil, err
	}

	// parse cached file
	err = json.Unmarshal([]byte(fileStr), &fileObj)
	if err != nil {
		return nil, err
	}
	return &fileObj, nil
}

func GetRemoteFile(botToken, fileId string) (*entities.TelegramResponse, error) {
	var fileObj entities.TelegramResponse

	// get remote data
	url := fmt.Sprintf("%s%s/getFile?file_id=%s", BASE_URL, botToken, fileId)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	// check status code
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Cannot fetch data: " + resp.Status)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&fileObj)
	if err != nil {
		return nil, err
	}

	err = AddFileToCache(fileId, fileObj)
	if err != nil {
		// just log it
		log.Println(err)
	}

	return &fileObj, nil
}

func AddFileToCache(fileId string, fileObj entities.TelegramResponse) error {
	fileObjJson, err := json.Marshal(fileObj)
	if err != nil {
		return err
	}
	// Telegram guarantees at least one hour of validity
	cachePeriod := time.Hour
	err = redisClient.Set(ctx, fileId, fileObjJson, time.Duration(cachePeriod)).Err()
	if err != nil {
		return err
	}
	return nil
}
