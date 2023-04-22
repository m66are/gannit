package directory_pkg

import (
	"fmt"
	"os"
	"os/user"
)

func GetCurrentDirectory() string {
	// Get the current directory path
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("👉 Current directory:", dir)
	return dir

}
func GetHomeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	fmt.Println("🏠 Home directory:", currentUser.HomeDir)
	return currentUser.HomeDir, nil
}

func GetFilesInDir(dir string) ([]os.DirEntry, error) {
	// read directory contents
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// filter out directories
	var files []os.DirEntry
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry)
		}
	}

	return files, nil
}
func ChangeDir(dir string) (string, error) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return "", err

	}
	return GetCurrentDirectory(), nil
}
