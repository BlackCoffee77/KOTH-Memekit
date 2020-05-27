package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/BlackCoffee77/KOTH-Memekit/fakeflags"
	"github.com/BlackCoffee77/KOTH-Memekit/signatures"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fakeFlagsPtr := flag.Int("fakeflags", 0, "Do you want to add fake flags?")

	signaturePtr := flag.String("signature", "", "Whose signature would you like to use")
	flag.Parse()

	if *fakeFlagsPtr > 0 {
		fakeflags.GenFakeFlags(*fakeFlagsPtr, "/")
	}

	switch *signaturePtr {
	case "example":
		signatures.ExampleSignature()
	}

}
