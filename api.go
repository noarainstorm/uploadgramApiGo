package uploadgramApiGo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func New(api string, userAgent string) All {
	all := All{}
	if api == "" {
		all.Api = "https://api.uploadgram.me/"
	} else {
		all.Api = api
	}
	if userAgent == "" {
		all.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36"
	} else {
		all.UserAgent = userAgent
	}
	return all
}

func (tool *All) Upload(filePath string) (err error) {
	err = tool.loadFile(filePath)
	defer tool.File.Close()

	if err != nil {
		return ErrOpenFile
	}
	_, tool.Out, err = tool.request("PUT", tool.Api+"upload", tool.File)
	if err != nil {
		return
	}

	err = tool.unjson()

	if err != nil {
		return ErrUnJson
	}

	if !tool.Response.Ok {
		return ErrServer
	}

	return
}

func (tool *All) UploadBytes(input []byte) (err error) {
	_, tool.Out, err = tool.request("PUT", tool.Api+"upload", input)
	if err != nil {
		return
	}
	err = tool.unjson()

	if err != nil {
		return ErrUnJson
	}

	if !tool.Response.Ok {
		return ErrServer
	}

	return
}

func (tool *All) loadFile(path string) (err error) {
	tool.File, err = os.Open(path)
	return err
}

func (tool *All) unjson() (err error) {
	err = json.Unmarshal(tool.Out, &tool.Response)
	return
}

func (tool *All) Download(link string) (data []byte, err error) {
	err = tool.GetInfo(link)

	if err != nil {
		return
	}

	_, data, err = tool.request("GET", link+"?raw", nil)

	return
}

func (tool *All) GetInfo(link string) (err error) {
	_, cuttedForInfo, isOk := strings.Cut(link, "https://dl.uploadgram.me/")
	if !isOk {
		return ErrLink
	}

	ans, out, err := tool.request("GET", tool.Api+"get/"+cuttedForInfo, nil)
	if ans.StatusCode == 403 || ans.StatusCode == 404 {
		return ErrNotFound
	}
	if err != nil {
		return
	}

	err = json.Unmarshal(out, &tool.DownloadInfo)
	if err != nil {
		err = ErrUnJson
	}

	return
}

func (tool *All) Delete(token string) (err error) {
	if len(token) != 49 {
		return ErrToken
	}

	ans, _, err := tool.request("GET", tool.Api+"delete/"+token, nil)
	if err != nil {
		return
	}
	tool.DeleteStat = ans.StatusCode
	return
}

func (tool All) Rename(token string, newName string) (err error) {
	if len(token) != 49 {
		return ErrToken
	}
	if len(newName) < 1 {
		return ErrFileName
	}
	inputJson := strings.NewReader(fmt.Sprintf(`{"new_filename": "%s"}`, newName))
	_, _, err = tool.request("POST", tool.Api+"rename/"+token, inputJson)
	return err
}

func (tool All) request(method, url string, data io.Reader) (ans *http.Response, output []byte, err error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return
	}
	req.Header.Add("User-Agent", tool.UserAgent)
	if method == "PUT" {
		info, err := tool.File.Stat()
		if err != nil {
			return ans, output, err
		}
		req.Header.Add("content-type", "application/octet-stream")
		req.Header.Add("upload-filename", info.Name())
	}
	ans, err = client.Do(req)
	if err != nil {
		return
	}
	output, err = io.ReadAll(ans.Body)

	if ans.StatusCode == 403 || ans.StatusCode == 404 {
		return ans, nil, ErrNotFound
	}

	return
}
