package main

import (
	"fmt"
	"github.com/gomiran/escav"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Please provide a filename")
	}
	path := strings.Join(args[1:], " ")
	log.Println(path)

	fmt.Println(escav.File(path))
}
