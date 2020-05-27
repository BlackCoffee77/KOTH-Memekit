package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/BlackCoffee77/KOTH-Memekit/fakeflags"
	"github.com/BlackCoffee77/KOTH-Memekit/signatures"
)

func signatureHandler(sigPointer *string) {
	switch *sigPointer {
	case "example":
		signatures.ExampleSignature()

	}
}
func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fakeFlagsPtr := flag.Int("fakeflags", 0, "Do you want to add fake flags?")

	signaturePtr := flag.String("signature", "", "Whose signature would you like to use")
	flag.Parse()

	if *fakeFlagsPtr > 0 {
		fakeflags.GenFakeFlags(*fakeFlagsPtr, "/")
	}

	signatureHandler(signaturePtr)

}
