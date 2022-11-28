package util

import (
	"io"
	"log"
	"net/http"
	"os"
)

func FileDownloadFromUrl(url, wantFileName string) bool {
	log.Printf("will download the file from the ur")
	file, err := os.Create(wantFileName)
	if err != nil {
		log.Println(err)
		return false
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)

	defer file.Close()
	log.Printf("Downloaded a file %s with size %d", wantFileName, size)
	return true
}
