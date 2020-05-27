package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/BlackCoffee77/KOTH-Memekit/fakeflags"
	"github.com/BlackCoffee77/KOTH-Memekit/rootreplace"
	"github.com/BlackCoffee77/KOTH-Memekit/signatures"
)

func signatureHandler(sigPointer *string) {
	switch *sigPointer {
	case "example":
		signatures.ExampleSignature()

	default:
		fmt.Println("There is no signature for that user")
	}

}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var users arrayFlags

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fakeFlagsPtr := flag.Int("fakeflags", 0, "Do you want to add fake flags?")

	signaturePtr := flag.String("signature", "", "Whose signature would you like to use")

	rootReplacePtr := flag.Bool("rootreplace", false, "If everyone is king, no one is king - Syndrome")

	flag.Var(&users, "users", "A user in the KOTH game")
	flag.Parse()

	if *fakeFlagsPtr > 0 {
		fakeflags.GenFakeFlags(*fakeFlagsPtr, "/")
	}

	if *signaturePtr != "" {
		signatureHandler(signaturePtr)
	}

	if *rootReplacePtr {
		rootreplace.RootReplace(users)
	}

}
