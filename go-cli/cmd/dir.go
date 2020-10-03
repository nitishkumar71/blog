package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/nitishkumar71/blog/go-cli/pkg"

	"github.com/spf13/cobra"
)

var dSize bool

// CreateNewDirectoryCommand : create new directory command
func createNewDirectoryCommand() *cobra.Command {
	var dirCmd = &cobra.Command{
		Use:   "dir",
		Short: "Perform operations on Directory",
		Long:  `This will allow users to perform operations on Directory`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			processDirCommand(args[0])
		},
	}

	dirCmd.Flags().BoolVarP(&dSize, "size", "s", false, "Size of the directory")

	return dirCmd
}

func processDirCommand(dName string) {
	if dName == "" {
		fmt.Printf("File name is not provided\n")
		return
	}

	size, filesCount, dirCount, err := getDirInfo(dName)

	if err != nil {
		return
	}

	if dSize {
		fmt.Printf("Size of Directory: %s\n", pkg.FormatSize(size))
		return
	}

	fmt.Printf("Directory Name: %s\n", dName)
	fmt.Printf("Size of Directory: %s\n", pkg.FormatSize(size))
	fmt.Printf("Total %d files and %d directories found\n", filesCount, dirCount)

}

func getDirInfo(dName string) (int64, int64, int64, error) {
	files, err := ioutil.ReadDir(dName)

	var size, filesCount, dirCount int64
	size, filesCount, dirCount = 0, 0, 0

	if err != nil {
		fmt.Printf("Issue faced while accessing file %v\n", err)
		return 0, 0, 0, err
	}

	for _, file := range files {
		size += file.Size()
		if file.IsDir() {
			dirCount++
			tSize, tFilesCount, tdirCount, _ := getDirInfo(fmt.Sprintf("%s/%s", dName, file.Name()))
			size += tSize
			filesCount += tFilesCount
			dirCount += tdirCount
		} else {
			filesCount++
		}
	}

	return size, filesCount, dirCount, nil
}
