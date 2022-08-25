package main

import (
	"flag"
	"fmt"
	"github.com/xpoh/bzip/pkg/atack"
)

func main() {
	charset := flag.String("c", "1234567890", "char set")
	fileName := flag.String("f", "input.zip", "zip file name")
	maxLength := flag.Int("l", 8, "max password lenght")
	flag.Parse()

	za := atack.NewZipArchive(*fileName)
	at := atack.NewAtack(za, *maxLength, []rune(*charset))

	pass, err := at.Brute()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Find password: ", pass)
	}
}
