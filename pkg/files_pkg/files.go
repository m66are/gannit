package files_pkg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func GroupFilesByExtensions(files []os.DirEntry, extensions []string) []os.DirEntry {
	// Create a slice to store the files with the specified extensions
	matchingFiles := make([]os.DirEntry, 0)

	// Iterate over the files in the directory
	for _, file := range files {
		if file.IsDir() {
			// Skip directories
			continue
		}

		// Get the file extension
		ext := filepath.Ext(file.Name())

		// Check if the file has one of the specified extensions
		for _, e := range extensions {
			if strings.EqualFold(ext, e) {
				matchingFiles = append(matchingFiles, file)
				break
			}
		}
	}

	return matchingFiles
}

func FindUncategorized(bigList []os.DirEntry, smallLists ...[]os.DirEntry) []os.DirEntry {
	uncategorized := []os.DirEntry{}

	for _, bigItem := range bigList {
		found := false

		for _, smallList := range smallLists {
			if Contains(smallList, bigItem) {

				found = true
				break
			}
		}

		if !found {
			fmt.Println("❔ uncategorized file found:", bigItem.Name())
			uncategorized = append(uncategorized, bigItem)
		}
	}

	return uncategorized
}

func Contains(list []os.DirEntry, item os.DirEntry) bool {
	for _, x := range list {
		if x.Name() == item.Name() {
			return true
		}
	}

	return false
}

func MoveFilesToDirectory(files []os.DirEntry, directoryName string) error {
	// create the destination directory if it doesn't exist
	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		err := os.Mkdir(directoryName, 0755)
		if err != nil {
			return fmt.Errorf("error creating destination directory: %s", err)
		}
	}

	// move each file to the destination directory
	for _, file := range files {
		src := file.Name()
		dest := filepath.Join(directoryName, file.Name())

		// open source file
		srcFile, err := os.Open(src)
		if err != nil {
			return fmt.Errorf("error opening source file: %s", err)
		}
		defer srcFile.Close()

		// create destination file
		destFile, err := os.Create(dest)
		if err != nil {
			return fmt.Errorf("error creating destination file: %s", err)
		}
		defer destFile.Close()

		// copy contents of source file to destination file
		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return fmt.Errorf("error copying file contents: %s", err)
		}

		// remove the original file
		err = os.Remove(src)
		if err != nil {
			return fmt.Errorf("error removing original file: %s", err)
		}
	}

	return nil
}
