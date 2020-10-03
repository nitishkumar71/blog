package pkg

import "fmt"

// FormatSize formats size of file and directory
func FormatSize(size int64) string {
	if size > (1024 * 1024) {
		return fmt.Sprintf("%d Mb", (size / (1024 * 1024)))
	} else if size > 1024 {
		return fmt.Sprintf("%d Kb", (size / 1024))
	}
	return fmt.Sprintf("%d bytes", size)
}
