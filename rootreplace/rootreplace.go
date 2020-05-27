package rootreplace

import (
	"bufio"
	"os"
)

// RootReplace constantly deletes and creates the root.txt file
func RootReplace(arr []string) {
	f, _ := os.Create("/root/king.txt")
	w := bufio.NewWriter(f)
	for _, name := range arr {
		w.WriteString(name + "\n")
	}
	w.Flush()
}
