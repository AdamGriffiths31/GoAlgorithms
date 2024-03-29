package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	force := flag.Bool("force", false, "delete existing folder")
	flag.Parse()
	folderName := time.Now().Format("2006-01-02")
	if *force {
		err := deleteFolder(folderName)
		if err != nil {
			panic(fmt.Errorf("error deleting folder %s: %w", folderName, err))
		}
	}
	if folderExists(folderName) {
		panic(fmt.Errorf("folder %s already exists", folderName))
	}

	copyAlgorithms(folderName)
}

func copyAlgorithms(folderName string) {
	filePath := "algorithms.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var jsonArray []interface{}

	err = json.Unmarshal(data, &jsonArray)
	if err != nil {
		panic(err)
	}

	for _, element := range jsonArray {
		str, ok := element.(string)
		if !ok {
			str = fmt.Sprintf("%v", element)
		}
		copyFolder(str, folderName+"/"+str)
	}
}

func copyFolder(source, destination string) error {
	err := os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(source, entry.Name())
		dstPath := filepath.Join(destination, entry.Name())

		if entry.IsDir() {
			err = copyFolder(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(source, destination string) error {
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func folderExists(folderPath string) bool {
	_, err := os.Stat(folderPath)

	if os.IsNotExist(err) {
		return false
	}

	return err == nil
}

func deleteFolder(folderPath string) error {
	return os.RemoveAll(folderPath)
}
