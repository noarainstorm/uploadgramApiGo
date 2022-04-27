package uploadgramApiGo

import "os"

type Config struct {
	Api       string
	UserAgent string
}

type Response struct {
	Ok    bool
	Url   string
	Token string `json:"delete"`

	DeleteStat int
}

type DownloadInfo struct {
	Ok       bool
	Filename string
	Size     int
}

type All struct {
	File *os.File

	Out []byte

	Config
	Response
	DownloadInfo
}
