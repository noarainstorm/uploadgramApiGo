# uploadgramApiGo
Easy to use uploadgram.me api

[![Go Reference](https://pkg.go.dev/badge/github.com/noarainstorm/uploadgramApiGo)](https://pkg.go.dev/github.com/noarainstorm/uploadgramApiGo)

# Already done:

+ Upload Files

+ Download Files + Files info

+ Delete files

+ Rename files

## How install it?

```
$ go get github.com/noarainstorm/uploadgramApiGo
```

## How use it?

### Init the  module

```
package main
import (
    "fmt"
    uploadgram "github.com/noarainstorm/uploadgramApiGo"
)

func main() {
    api := uploadgram.New("", "")
...
```

### Upload file

```
...
    err := api.Upload("file name")
    if err != nil {
        panic(err)
    }
    link, token := api.Response.Url, api.Response.Token
...
```

### Download file

```
...
    file, err := api.Download("url here")
    if err != nil {
        panic(err)
    }
   // you can save "file" to file with os
...

```

### Delete file

```
...
    err := api.Delete(token)
    if err != nil {
        panic(err)
    }
...
```

### Rename file

```
...
    err := api.Rename(token, newName)
    if err != nil {
        panic(err)
...
```

## Simple example

```
package main
import (
	ug "github.com/noarainstorm/uploadgramApiGo"
	"fmt"
	"os"
)

func main() {
	api := ug.New("","")
	err := api.Upload("Downloads/hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File uploaded! %s :: link, %s token\n", api.Response.Url, api.Response.Token)
	fileData, err := api.Download(api.Response.Url)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(api.DownloadInfo.Filename)
	if err != nil {
		panic(err)
	}
	file.Write(fileData)
	fmt.Println("File downloaded!")
	err = api.Rename(api.Response.Token, "WOW")
	if err != nil {
		panic(err)
	}
	fmt.Println("File renamed!")
	err = api.Delete(api.Response.Token)
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted from servers!")
}
```

## *Need more docs?*

### Read Docs [here](https://pkg.go.dev/github.com/noarainstorm/uploadgramApiGo "Docs here")
