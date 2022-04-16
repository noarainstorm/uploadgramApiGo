package uploadgramApiGo

import "errors"

var (
	ErrOpenFile = errors.New("file opening error")
	ErrFileName = errors.New("invalid file name")
	ErrToken    = errors.New("invalid token")
	ErrServer   = errors.New("error on server")
	ErrUnJson   = errors.New("parsing json was interrupted")
	ErrNotFound = errors.New("not found")
	ErrLink     = errors.New("invalid link")
)
