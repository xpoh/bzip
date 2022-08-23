package atack

import (
	"bytes"
	zzz "github.com/yeka/zip"
	"log"
	"os"
)

type azip struct {
	fileName string
	content  []byte
	size     int64
}

func New(fileName string) *azip {
	var err error
	z := &azip{fileName: fileName}

	z.content, err = os.ReadFile(fileName)
	if err != nil {
		panic("Error open file!!!")
		return nil
	}
	z.size = int64(len(z.content))
	return z
}

func (a *azip) check(pass string) bool {
	r, err := zzz.NewReader(bytes.NewReader(a.content), a.size)
	if err != nil {
		return false
	}

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(pass)
		}

		r, err := f.Open()
		if err != nil {
			log.Println(err)
			return false
		} else {
			r.Close()
			return true
		}
	}
	return false
}
