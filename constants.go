package gopyload

import (
	"fmt"
	"strings"
)

type endpoint string

const (
	loginEndpoint       endpoint = "/api/login"
	logoutEndpoint      endpoint = "/api/logout"
	fileDataEndpoint    endpoint = "/api/getFileData"
	getUserDataEndpoint endpoint = "/api/getUserData"

	getServicesEndpoint endpoint = "/api/getServices" // something here

	deleteFinishedEndpoint      endpoint = "/api/deleteFinished"
	freeSpaceEndpoint           endpoint = "/api/freeSpace"
	getAccountTypesEndpoint     endpoint = "/api/getAccountTypes"
	getAllInfoEndpoint          endpoint = "/api/getAllInfo" // DIDN'T USE DIDN'T TEST
	getAllUserDataEndpoint      endpoint = "/api/getAllUserData"
	getCollectorEndpoint        endpoint = "/api/getCollector"     // DIDN'T USE DIDN'T TEST
	getCollectorDataEndpoint    endpoint = "/api/getCollectorData" // DIDN'T USE DIDN'T TEST
	getConfigEndpoint           endpoint = "/api/getConfig"
	getConfigDictEndpoint       endpoint = "/api/getConfigDict"
	getPluginConfigEndpoint     endpoint = "/api/getPluginConfig"
	getPluginConfigDictEndpoint endpoint = "/api/getPluginConfigDict"
	getQueueEndpoint            endpoint = "/api/getQueue"
	getQueueDataEndpoint        endpoint = "/api/getQueueData"
	getServerVersionEndpoint    endpoint = "/api/getServerVersion"
	isCaptchaWaitingEndpoint endpoint = "/api/isCaptchaWaiting"
	isTimeDownloadEndpoint endpoint = "/api/isTimeDownload"
	isTimeReconnectEndpoint endpoint = "/api/isTimeReconnect"
	killEndpoint endpoint = "/api/kill"
	pauseServerEndpoint endpoint = "/api/pauseServer"
	unpauseServerEndpoint endpoint = "/api/unpauseServer" 

	restartEndpoint endpoint = "/api/restart"

	restartFailedEndpoint endpoint = "/api/restartFailed"

	statusDownloadsEndpoint endpoint = "/api/statusDownloads"

	statusServerEndpoint endpoint = "/api/statusServer"

	stopAllDownloadsEndpoint endpoint = "/api/stopAllDownloads"

	togglePauseEndpoint endpoint = "/api/togglePause"

	toggleReconnectEndpoint endpoint = "/api/toggleReconnect"

)

func (e endpoint) GetFuncName() string {
	uri := string(e)
	spl := strings.Split(uri, "/")
	if len(spl) != 3 {
		return ""
	}
	return spl[len(spl)-1]
}

var HttpMapToStatusCode = map[int]error{
	200: nil,
	400: fmt.Errorf("bad request"),
	401: fmt.Errorf("unauthorized"),
	403: fmt.Errorf("forbidden"),
	404: fmt.Errorf("not found"),
	500: fmt.Errorf("internal server error"),
	503: fmt.Errorf("service unavailable"),
}

var (
	ErrCookieNotValid = fmt.Errorf("cookie is not valid")
	ErrNoCookie       = fmt.Errorf("cookie does not exist")
	ErrNotBoolean     = fmt.Errorf("not a boolean")
)

var defaultPyloadHeaders = map[string]string{
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"Accept-Language":           "en-US,en;q=0.9,ar;q=0.8",
	"Cache-Control":             "max-age=0",
	"Connection":                "keep-alive",
	"Content-Type":              "application/x-www-form-urlencoded",
	"DNT":                       "1",
	"Origin":                    "http://pyload.lxc.1jz.cc:8000",
	"Referer":                   "http://pyload.lxc.1jz.cc:8000/login?next=dashboard",
	"Upgrade-Insecure-Requests": "1",
	"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
}
