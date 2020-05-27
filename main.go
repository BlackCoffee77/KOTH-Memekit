package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/BlackCoffee77/KOTH-Memekit/fakeflags"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fakeFlagsPtr := flag.Int("fakeflags", 0, "Do you want to add fake flags?")

	flag.Parse()

	fakeflags.GenFakeFlags(*fakeFlagsPtr, "/")
}
