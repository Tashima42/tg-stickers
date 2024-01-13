package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"strconv"

	"github.com/sunshineplan/imgconv"
)

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	file, _, err := req.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	defer req.Body.Close()

	buffer := new(bytes.Buffer)
	if err := convert(buffer, file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)

	s := base64.URLEncoding.EncodeToString(buffer.Bytes())
	sb := []byte(s)
	w.Write(sb)
	w.Header().Set("Content-Length", strconv.FormatInt(int64(len(sb)), 10))
}

func convert(w io.Writer, r io.Reader) error {
	srcImg, err := imgconv.Decode(r)
	if err != nil {
		return err
	}
	img := imgconv.Resize(srcImg, &imgconv.ResizeOption{Width: 500})
	return imgconv.Write(w, img, &imgconv.FormatOption{Format: imgconv.PNG})
}
