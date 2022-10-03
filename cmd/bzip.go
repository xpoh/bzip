package main

import (
	"flag"
	"fmt"
	"github.com/xpoh/bzip/pkg/atack"
)

const version = "1.0.1"

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("bzip -f=ZipfileName -c=\"CharSet\" -l=MaxPasswordLenght")
}

func main() {
	fmt.Println("Zip password recovery utility. Version ", version)

	charset := flag.String("c", "1234567890", "char set")
	fileName := flag.String("f", "input.zip", "zip file name")
	maxLength := flag.Int("l", 8, "max password lenght")
	flag.Parse()

	if flag.NFlag() == 0 {
		PrintUsage()
		return
	}

	za := atack.NewZipArchive(*fileName)
	at := atack.NewAtack(za, *maxLength, []rune(*charset))

	pass, err := at.Brute()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Find password: ", pass)
	}
}
