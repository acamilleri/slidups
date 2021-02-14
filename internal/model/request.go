package model

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Request - Wrap http request
type Request struct {
	*http.Request
}

// GetMediaType - Retrieve media type from content-type header
func (r *Request) GetMediaType() string {
	return r.Header.Get("Content-Type")
}

// GetFileSize - Retrieve size of file from content-length header
func (r *Request) GetFileSize() int {
	filesize := r.Header.Get("Content-Length")
	size, _ := strconv.Atoi(filesize)
	return size
}

// GetFile - Retrieve file from http request
func (r *Request) GetFile() (File, error) {
	var (
		filename    string
		filesize    int
		content     []byte
		destination string
		err         error
	)

	mediaType := r.GetMediaType()
	filesize = r.GetFileSize()
	destination = r.GetFileDestination()

	if strings.HasPrefix(mediaType, "multipart/") {
		filename, content, err = r.multipartFile()
		if err != nil {
			return File{}, fmt.Errorf("failed to get file from multipart: %v", err)
		}

		return NewFile(filename, filesize, content, destination), nil
	}

	return File{}, fmt.Errorf("failed to get file other than multipart request: %v", err)
}

// GetFileDestination - Retrieve destination from url queries
func (r *Request) GetFileDestination() string {
	destination := r.URL.Query().Get("destination")
	if destination != "" {
		return destination
	}

	return ""
}

func (r *Request) multipartFile() (string, []byte, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return "", nil, err
	}

	var (
		filename string
		content  []byte
	)

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		if err != nil {
			return "", nil, err
		}

		if part.FileName() != "" {
			filename = part.FileName()
		}

		data, err := ioutil.ReadAll(part)
		if err != nil {
			return "", nil, err
		}
		content = data

		_ = part.Close()
	}

	return filename, content, nil
}
