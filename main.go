package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
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

func genFakeFlags(num int, root string) {
	for i := 0; i < num; i++ {

		depth := rand.Intn(10) + 1
		curDir := root
		fmt.Println("depth", depth)

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
			fmt.Println(dirs)

			if len(dirs) == 0 {
				//break in order to write file
				break
			}

			curDir = curDir + dirs[rand.Intn(len(dirs))]

			fmt.Println(curDir)
		}

		// create file at current location
		fmt.Println("ree")
		f, _ := os.Create(curDir + genFlagName())
		w := bufio.NewWriter(f)
		bytes := md5.Sum([]byte(time.Now().String()))
		finalHash := hex.EncodeToString(bytes[:])

		w.WriteString(finalHash)
		w.Flush()
	}
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fakeFlagsPtr := flag.Int("fakeflags", 0, "Do you want to add fake flags?")

	flag.Parse()
	fmt.Println(*fakeFlagsPtr)

	genFakeFlags(*fakeFlagsPtr, "/")
}
