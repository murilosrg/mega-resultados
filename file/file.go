package file

import (
	"archive/zip"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
)

//Download source file with all results
func Download(src, dest string) ([]string, error) {
	if err := downloadFile(dest, src); err != nil {
		return nil, err
	}

	files, err := unzip(dest, "tmp")

	if err != nil {
		return nil, err
	}

	return files, nil
}

//Remove all files after processing
func Remove(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	if err := os.RemoveAll("tmp"); err != nil {
		return err
	}

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

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if _, err = io.Copy(out, res.Body); err != nil {
		return err
	}

	return nil
}

func unzip(src, dest string) ([]string, error) {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := os.Mkdir(dest, 0755); err != nil {
			return nil, err
		}
	}

	var files []string
	zipReader, _ := zip.OpenReader(src)
	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer zippedFile.Close()

		extractedFilePath := filepath.Join(
			dest,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				return nil, err
			}

			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				return nil, err
			}

			files = append(files, outputFile.Name())
		}
	}

	return files, nil
}
