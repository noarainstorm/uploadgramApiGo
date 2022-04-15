# uploadgram-api-go
Easy to use uploadgram.me api
### Already done:

+ Upload Files

+ Download Files + Files info

+ Delete files

+ Rename files

## How use it?

### Init the  module

```
package main
import (
    "fmt"
    uploadgram "github.com/noarainstorm/uploadgram-go"
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
    link, token := api.Response.Url, api.Response.Delete
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
