package globals

import (
	"os"
	"strings"
)

/*Function to split files with newlines.
It returns an array where each member is the newline content.
Made this for AOC to help users*/
func SplitFile(fileName string) []string {
	file, _ := os.Open(fileName)
	stats, _ := os.Stat(fileName)
	size := stats.Size()
	fileByte := make([]byte, size)
	fileSlice := fileByte[:]
	file.Read(fileSlice)
	arr := strings.Split(string(fileSlice), "\n")
	return arr
}
