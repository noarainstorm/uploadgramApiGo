# uploadgramApiGo
Easy to use uploadgram.me api
### Already done:

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

### Read Docs **nothing**
