package main

import (
	"fmt"
	"os"
	"parex/internal/lib"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Use: parex <filepath> <offset> <level>")
		fmt.Println("Level 0-: FLS")
		fmt.Println("Level 1: Recursive FLS")
		fmt.Println("Level 2+: Recursive FCAT")
	}

	imagefile, err := os.Open(os.Args[1])
	handle(err)
	offset, err := strconv.ParseInt(os.Args[2], 10, 64)
	handle(err)
	level, err := strconv.Atoi(os.Args[3])
	handle(err)

	err = lib.Explore(imagefile, uint64(offset), level)
	handle(err)
}

func handle(err error) {
	if err != nil {
		fmt.Printf("n\n%v\n\n", err)
		os.Exit(1)
	}
}
