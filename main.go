package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func removeEmptyNewFolders() {

}

func main() {

	root := os.Args
	var directoryPathToWalk string = root[1]
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
					fmt.Println("something went wrong with finding the abspath...")
					//continue;
				} else {
					fmt.Println(absPath)
					os.Remove(absPath)
					fmt.Println("removed %s", absPath)
				}
			} else {
				//fmt.Println("something went wrong with finding the abspath...")
				//fmt.Println(path, info.Size())
				files, _ := os.ReadDir(absPath)

				fmt.Println("%s : %s", absPath, len(files))
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}

}
