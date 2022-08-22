package atack

import (
	"archive/zip"
	"bytes"
	"github.com/alexmullins/zip"
	"io"
	"log"
	"os"
)

type azip struct {
	fileName string
}

func New(fileName string) *azip {
	contents := []byte("Hello World")
	fzip, err := os.Create(`./test.zip`)
	if err != nil {
		log.Fatalln(err)
	}
	zipw := zip.NewWriter(fzip)
	defer zipw.Close()
	w, err := zipw.Encrypt(`test.txt`, `golang`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	zipw.Flush()
	return &azip{fileName: fileName}
}

func (a azip) check(pass string) bool {
	//TODO implement me
	panic("implement me")
}
