package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numberPtr := flag.Int("number", 45, "a number")
	forkPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	//标志声明后，解析标志
	flag.Parse()
	fmt.Println("word: ", *wordPtr)
	fmt.Println("number: ", *numberPtr)
	fmt.Println("fork: ", *forkPtr)
	fmt.Println("svar: ", svar)
	fmt.Println("flag arguments: ", flag.Args())
}
