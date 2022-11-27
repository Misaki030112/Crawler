package uti

import (
	"io"
	"net/http"
	"os"
)

// File remote download
func FileDownloadFromUrl(url, wantFileName string) bool {
	// Log.Printf("will download the file from the ur")
	fileURL, err := url.Parse(url)
	if err != nil {
		// Log.Fatal(er)
	}
	file, err := os.Create(wantFileName)
	if err != nil {
		// Log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fileURL)
	if err != nil {
		// Log.Fatal(er)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)

	defer file.Close()
	// Log.Printf("Downloaded a file %s with size %d", wantFileName, sie)
	return true
}
