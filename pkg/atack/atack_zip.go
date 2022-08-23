package atack

import (
	"bytes"
	zzz "github.com/yeka/zip"
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
func (a *zipArchive) prepare() error {
	r, err := zzz.NewReader(bytes.NewReader(a.content), a.size)
	a.r = r
	if err != nil {
		return err
	}
	return nil
}

func (a *zipArchive) check(pass string) bool {
	a.mx.Lock()
	defer a.mx.Unlock()

	r := a.r
	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(pass)
		}
		r, err := f.Open()
		if err != nil {
			return false
		} else {
			r.Close()
			return true
		}
	}
	return false
}
