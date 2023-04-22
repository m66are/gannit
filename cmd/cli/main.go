package main

import (
	"fmt"
	"gannit/pkg/directory_pkg"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "my-cli",
		Short: "A brief description of my CLI",
		Long:  "A longer description of my CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, world!")
		},
	}

	// Add commands defined in commands.go
	rootCmd.AddCommand(helloCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// homeDr, err := GetHomeDir()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Get a list of all files in the homeDr directory Downloads subfolder and handle errors

	files, err := directory_pkg.GetFilesInDir(directory_pkg.GetCurrentDirectory())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("📁 Files in current directory:", files)

	// // Define the extensions for the files to be grouped
	// imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	// musicExtensions := []string{".mp3", ".wav", ".flac"}
	// videoExtensions := []string{".mp4", ".avi", ".mov"}
	// documentExtensions := []string{".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx"}
	// archiveExtensions := []string{".zip", ".rar", ".7z"}
	// programExtensions := []string{".exe", ".dmg", ".app"}

	// // Group the files based on their types
	// imageFiles := groupFilesByExtensions(files, imageExtensions)
	// musicFiles := groupFilesByExtensions(files, musicExtensions)
	// videoFiles := groupFilesByExtensions(files, videoExtensions)
	// documentFiles := groupFilesByExtensions(files, documentExtensions)
	// archiveFiles := groupFilesByExtensions(files, archiveExtensions)
	// programFiles := groupFilesByExtensions(files, programExtensions)
	// otherFiles := findUncategorized(files, imageFiles, musicFiles, videoFiles, documentFiles, archiveFiles, programFiles)
	// // call moveFilesToDirectory for all above lists with the proper name as directory name
	// err = moveFilesToDirectory(imageFiles, "Images")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// err = moveFilesToDirectory(musicFiles, "Music")
	// if err != nil {
	// 	fmt.Println(err)

	// 	return
	// }
	// err = moveFilesToDirectory(videoFiles, "Videos")
	// if err != nil {
	// 	fmt.Println(err)

	// 	return
	// }
	// err = moveFilesToDirectory(documentFiles, "Documents")
	// if err != nil {
	// 	fmt.Println(err)

	// 	return
	// }
	// err = moveFilesToDirectory(archiveFiles, "Archives")
	// if err != nil {
	// 	fmt.Println(err)

	// 	return
	// }
	// err = moveFilesToDirectory(programFiles, "Programs")
	// if err != nil {
	// 	fmt.Println(err)

	// 	return
	// }

	// err = moveFilesToDirectory(otherFiles, "Other")
	// if err != nil {
	// 	fmt.Println(err)

	// 	return
	// }
	// fmt.Println("✅ Done moving files to their respective folders!")

	// // Print the lists of files
	// fmt.Println(" 📸 Image files:", len(imageFiles))
	// fmt.Println(" 🎧 Music files:", len(musicFiles))
	// fmt.Println(" 🎥 Video files:", len(videoFiles))
	// fmt.Println(" 📄 Document files:", len(documentFiles))
	// fmt.Println(" 📦 Archive files:", len(archiveFiles))
	// fmt.Println(" 💻 Program files:", len(programFiles))
	// fmt.Println(" 🗑️ Other files:", len(otherFiles))
	// // dispose all arrays and files
	// imageFiles = nil
	// musicFiles = nil
	// videoFiles = nil
	// documentFiles = nil
	// archiveFiles = nil
	// programFiles = nil

	// files = nil

}
