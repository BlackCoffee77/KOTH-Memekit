package rootreplace

import (
	"bufio"
	"os"
	"time"
)

// RootReplace constantly deletes and creates the root.txt file and puts a provided list of users in it
// usage: ./main -rootreplace -users <username> -users <username>
// You can add as many users as you want
func RootReplace(arr []string) {
	for {
		f, _ := os.Create("/root/king.txt")
		w := bufio.NewWriter(f)
		for _, name := range arr {
			w.WriteString(name + "\n")
		}
		w.Flush()
		time.Sleep(time.Second)
	}
}
