package models

type QueueModel struct {
	Pid        int         `json:"pid"`
	Name       string      `json:"name"`
	Folder     string      `json:"folder"`
	Site       string      `json:"site"`
	Password   string      `json:"password"`
	Dest       int         `json:"dest"`
	Order      int         `json:"order"`
	Linksdone  int         `json:"linksdone"`
	Sizedone   int64       `json:"sizedone"`
	Sizetotal  int64       `json:"sizetotal"`
	Linkstotal interface{} `json:"linkstotal"`
	Links      []struct {
		Fid        int    `json:"fid"`
		URL        string `json:"url"`
		Name       string `json:"name"`
		Plugin     string `json:"plugin"`
		Size       int64  `json:"size"`
		FormatSize string `json:"format_size"`
		Status     int    `json:"status"`
		Statusmsg  string `json:"statusmsg"`
		PackageID  int    `json:"package_id"`
		Error      string `json:"error"`
		Order      int    `json:"order"`
	} `json:"links"`
	Fids interface{} `json:"fids"`
}

type QueueDataModel struct {
	Dest       int         `json:"dest"`
	Fids       interface{} `json:"fids"`
	Folder     string      `json:"folder"`
	Links      interface{} `json:"links"`
	Linksdone  int         `json:"linksdone"`
	Linkstotal int         `json:"linkstotal"`
	Name       string      `json:"name"`
	Order      int         `json:"order"`
	Password   string      `json:"password"`
	Pid        int         `json:"pid"`
	Site       string      `json:"site"`
	Sizedone   int         `json:"sizedone"`
	Sizetotal  int         `json:"sizetotal"`
}