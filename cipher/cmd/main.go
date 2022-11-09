package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"service/cmd/root"
)

func main() {
	var action = flag.String("action", "", "the action you need to perform (encrypting or decrypting)")
	var text = flag.String("text", "", "text to be enciphered or deciphered")
	var key = flag.Int("key", 0, "value that specifies the alphabet rotation factor")
	flag.Usage = func() {
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()

	if *action == "" {
		log.Fatalln("unknown action requested")
	}

	if *key <= 0 {
		log.Fatalln("key must be greater than 0")
	}

	result, err := root.Run(*action, *text, *key)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}
