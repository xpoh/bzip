GOOS=windows GOARCH=amd64 go build -o ./build/bzip-win64.exe ./cmd/bzip.go
GOOS=windows GOARCH=386 go build -o ./build/bzip-win32.exe ./cmd/bzip.go
GOOS=linux GOARCH=amd64 go build -o ./build/bzip-linux-amd64 ./cmd/bzip.go
GOOS=linux GOARCH=386 go build -o ./build/bzip-linux-386 ./cmd/bzip.go
