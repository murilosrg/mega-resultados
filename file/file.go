package file

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
)

//Download source file with all results
func Download(src, dest string) error {
	if err := downloadFile(dest, src); err != nil {
		return err
	}

	fmt.Println("Downloaded file")

	if err := unzip(dest, "tmp"); err != nil {
		return err
	}

	fmt.Println("unzip file")

	return nil
}

func downloadFile(path string, url string) error {
	if _, err := os.Stat(path); err == nil {
		os.Remove(path)
	}

	out, err := os.Create(path)

	if err != nil {
		return err
	}

	defer out.Close()

	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	res, err := client.Get(url)
	fmt.Println(res)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if _, err = io.Copy(out, res.Body); err != nil {
		return err
	}

	return nil
}

func unzip(src, dest string) error {
	if err := os.Mkdir(dest, 0755); err != nil {
		return err
	}

	zipReader, _ := zip.OpenReader(src)
	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		extractedFilePath := filepath.Join(
			dest,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			fmt.Println("Directory Created:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			fmt.Println("File extracted:", file.Name)

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				return err
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
