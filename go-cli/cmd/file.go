package cmd

import (
	"fmt"
	"os"

	"github.com/nitishkumar71/blog/go-cli/pkg"
	"github.com/spf13/cobra"
)

var fSize bool

// CreateNewFileCommand : create new directory command
func createNewFileCommand() *cobra.Command {
	var fileCmd = &cobra.Command{
		Use:   "file",
		Short: "Perform operations on Files",
		Long:  `This will allow users to perform operations on files`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			processFileCommand(args[0])
		},
	}

	fileCmd.Flags().BoolVarP(&fSize, "size", "s", false, "Size of the File")

	return fileCmd
}

func processFileCommand(fName string) {
	if fName == "" {
		fmt.Printf("File name is not provided\n")
		return
	}

	file, err := os.OpenFile(fName, os.O_RDONLY, 0444)
	defer file.Close()

	if err != nil {
		fmt.Printf("Issue faced while accessing file %v\n", err)
		return
	}

	if file != nil {
		fStat, _ := file.Stat()

		if fStat.IsDir() {
			fmt.Printf("Please use dir command for Diretory\n")
			return
		}

		if fSize == true {
			fmt.Printf("File Size : %s\n", pkg.FormatSize(fStat.Size()))
			return
		}

		fmt.Printf("File Name %s\n", file.Name())
		fmt.Printf("Size of the File %s\n", pkg.FormatSize(fStat.Size()))
		fmt.Printf("File Modified Time: %v\n", fStat.ModTime())
	}

}
