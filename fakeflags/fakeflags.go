package fakeflags

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func genFlagName() string {
	names := []string{"user", "flag", "king", "root", "congrats", "your_flag"}
	ext := []string{".txt", ".md", "", ".lol"}

	return names[rand.Intn(len(names))] + ext[rand.Intn(len(ext))]
}

// GenFakeFlags fake flag generator
func GenFakeFlags(num int, root string) {

	fmt.Println("Placing fake flags at: ")
	for i := 0; i < num; i++ {

		depth := rand.Intn(10) + 1
		curDir := root

		for j := 0; j < depth; j++ {

			files, err := ioutil.ReadDir(curDir)

			if err != nil {
				log.Fatal(err)
			}

			dirs := []string{}

			for _, file := range files {
				if file.IsDir() {
					dirs = append(dirs, file.Name()+"/")
				}
			}

			if len(dirs) == 0 {
				//break in order to write file
				break
			}

			curDir = curDir + dirs[rand.Intn(len(dirs))]
		}

		// create file at current location
		loc := curDir + genFlagName()
		fmt.Println(loc)
		f, _ := os.Create(loc)
		w := bufio.NewWriter(f)
		bytes := md5.Sum([]byte(time.Now().String()))
		finalHash := hex.EncodeToString(bytes[:])

		w.WriteString(finalHash + "\n")
		w.Flush()
	}
}
