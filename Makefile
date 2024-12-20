build_macos:
	@GOOS=darwin GOARCH=arm64 go build -o ./build/guessing-game main.go

build_windows:
	@GOOS=darwin GOARCH=amd64 go build -o ./build/guessing-game.exe main.go