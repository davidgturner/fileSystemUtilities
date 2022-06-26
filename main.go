package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func yesNo() bool {
	prompt := promptui.Select{
		Label: "Select[Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

func main() {
	removeEmptyNewFolders()
}

func removeEmptyNewFolders() {

	var directoryPathToWalk string = os.Args[1]
	fmt.Println(directoryPathToWalk)

	fileSystem := os.DirFS(directoryPathToWalk)
	fmt.Println(fileSystem)

	fmt.Println(fmt.Sprintf("Walking the directory of %v , searching for \"New folder\" to delete...", fileSystem))

	err := filepath.Walk(directoryPathToWalk,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Size() > 0 {
				return nil
			}

			absPath, err := filepath.Abs(path)
			if info.IsDir() && strings.Contains(string(info.Name()), "New folder") && info.Size() == 0 {
				fmt.Println("Found an empty new folder = " + string(info.Name()))
				fmt.Println("we will need to delete this.")
				if err != nil {
					fmt.Println("something went wrong with finding the abspath...skipping over.")
				} else {
					fmt.Println(fmt.Sprintf("Pending Removal: %v. Remove?: ", absPath))
					removeNewFolderFlag := yesNo()
					if removeNewFolderFlag {
						// os.Remove(absPath)
						fmt.Println(fmt.Sprintf("Removed: %v", absPath))
					}
				}
			} else {
				files, _ := os.ReadDir(absPath)
				fmt.Println(fmt.Sprintf("File Path: %v # of Files: %v", absPath, len(files)))
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

}
