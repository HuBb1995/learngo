package filelist

import (
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (u userError) Error() string {
	return u.Message()
}

func (u userError) Message() string {
	return string(u)
}

func Handler(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path
	if strings.Index(path, prefix) != 0 {
		return userError("path must start with " + prefix)
	}
	fileName := path[len(prefix):]
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
