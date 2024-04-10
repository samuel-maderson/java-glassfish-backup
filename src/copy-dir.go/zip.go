package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Dir(appPath string, fileName string, destination string) {

	zipFile, err := os.Create(destination + "/" + fileName)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(appPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func File(dumpPath string, fileName string, destination string) {

	fileToZip := dumpPath

	zipFile, err := os.Create(destination + "/" + fileName)
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return
	}
	defer zipFile.Close()

	// Create a new writer to write to the zip file
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	file, err := os.Open(fileToZip)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		fmt.Println("Error creating file header:", err)
		return
	}

	header.Name = filepath.Base(fileToZip)

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		fmt.Println("Error creating zip header:", err)
		return
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		fmt.Println("Error copying file data to zip archive:", err)
		return
	}

}
