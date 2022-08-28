package atack

import (
	"bytes"
	"fmt"
	zzz "github.com/yeka/zip"
	"io"
	"os"
	"sync"
)

type zipArchive struct {
	fileName string
	content  []byte
	size     int64
	r        *zzz.Reader
	mx       *sync.Mutex
}

func NewZipArchive(fileName string) *zipArchive {
	var err error
	z := &zipArchive{fileName: fileName}

	z.content, err = os.ReadFile(fileName)
	if err != nil {
		panic("Error open file!!!")
		return nil
	}
	z.size = int64(len(z.content))
	z.mx = &sync.Mutex{}
	return z
}

func (a *zipArchive) check(pass string) bool {
	r, _ := zzz.NewReader(bytes.NewReader(a.content), a.size)

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(pass)
		}
		r1, err1 := f.Open()
		if err1 != nil {
			return false
		}
		defer r1.Close()
		buf, err := io.ReadAll(r1)
		if err != nil {
			return false
		}
		fmt.Printf("Size of %v: %v byte(s)\n", f.Name, len(buf))
	}
	return true
}
