clean:
	@rm ./build/*

build_macos:
	@GOOS=darwin GOARCH=arm64 go build -o ./build/guessing-game main.go

build_windows:
	@GOOS=windows GOARCH=amd64 go build -o ./build/guessing-game.exe main.go

package:
	@zip -vr build/builds.zip build/* -x "*.DS_Store" -x "_MACOSX"

build: clean build_macos build_windows package