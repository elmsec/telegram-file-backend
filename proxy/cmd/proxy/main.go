package main

import (
	"fmt"
	"github.com/elmsec/telegram-file-backend/pkg/config"
	pkg_http "github.com/elmsec/telegram-file-backend/pkg/http"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	appConfig := config.InitConfig()

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
	})

	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		director(req)
		req.Host = req.URL.Host

		// parse url and decrypt the bot token and the file id
		botToken, fileId, err := pkg_http.ParseUrl(req.URL.RequestURI()[1:])
		if err != nil {
			log.Println("Cannot parse URL: ", err)
			return
		}
		// get file
		telegramResponse, err := pkg_http.GetFile(botToken, fileId)
		if err != nil {
			log.Println("Cannot get file: ", err)
			return
		}

		// proxy to the telegram servers
		newPath := fmt.Sprintf("/file/bot%s/%s", botToken, telegramResponse.Result.FilePath)
		req.URL.Path = newPath
	}
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Del("Content-Disposition")
		return nil
	}

	url := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)
	log.Fatal(http.ListenAndServe(url, proxy))
}
