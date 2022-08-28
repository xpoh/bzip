package atack

import (
	"bytes"
	"github.com/yeka/zip"
	"io"
	"log"
	"os"
	"testing"
)

func createTestArchive(fileName string, pass string) {
	contents := []byte("Hello World")
	fzip, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	zipw := zip.NewWriter(fzip)
	defer func(zipw *zip.Writer) {
		err := zipw.Close()
		if err != nil {

		}
	}(zipw)
	w, err := zipw.Encrypt(`test.txt`, pass, zip.AES256Encryption)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	err = zipw.Flush()
	if err != nil {
		return
	}
}

func Benchmark_azip_check(t *testing.B) {
	t.StopTimer()
	path := "./test.zip"
	createTestArchive(path, "test123")
	a := NewZipArchive(path)
	t.StartTimer()
	for n := 0; n < t.N; n++ {
		a.check("12345")
	}
}

func Test_azip_check(t *testing.T) {
	createTestArchive("./test.zip", "test123")

	tests := []struct {
		name string
		path string
		pass string
		want bool
	}{
		{
			name: "Pass correct",
			path: "./test.zip",
			pass: "test123",
			want: true,
		},
		{
			name: "Pass uncorrect",
			path: "./test.zip",
			pass: "test321",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewZipArchive(tt.path)
			if got := a.check(tt.pass); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtack_brute_zip(t *testing.T) {

	createTestArchive("./test.zip", "1234")

	tests := []struct {
		name      string
		path      string
		wantPass  string
		wantErr   bool
		maxLength int
		chars     []rune
	}{
		{
			name:      "Pass = 1234",
			path:      "./test.zip",
			wantPass:  "1234",
			wantErr:   false,
			maxLength: 4,
			chars:     []rune("0123456789"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewZipArchive(tt.path)
			a := NewAtack(b,
				tt.maxLength,
				tt.chars)

			gotPass, err := a.Brute()
			if (err != nil) != tt.wantErr {
				t.Errorf("brute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPass != tt.wantPass {
				t.Errorf("brute() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
