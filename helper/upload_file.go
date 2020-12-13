package helper

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//// Create folder path
	// Split /
	filepathArray := strings.Split(filepath, "/")
	//  Remove last element
	filepathArray = filepathArray[:len(filepathArray)-1]
	// Convert array to string
	filepathAfter := strings.Join(filepathArray, "/")
	// Create folder
	os.MkdirAll(filepathAfter, os.ModePerm)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
