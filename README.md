# go-trex-runner
This repository is to show that you can also develop games in Go and also how one can enchanced thier skills in Go. 
I tried writing Google Chrome's offline T-Rex game in Go by using veandco/sdl2 package, please note that this is NOT a 100% Clone of Chrome's 
it's missing lot of features, hence it is not meant to be the ultimate Go-Code, but it is meant to be **simple**, **easy to understand** and yes its **playable**
This project is distributed under the [DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE](https://en.wikipedia.org/wiki/WTFPL).

### Requirements for Running it on Mac OS (Tested & Works on Catalina)
* Go lang - `version go1.13+ darwin/amd64` (https://golang.org/dl/)
* Docker for Mac or Docker Desktop (https://docs.docker.com/docker-for-mac/install)
* SDL2 - ` brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config`
* Go Package Dependencies 
   * `go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}`

### Steps to Run 
1. Clone the repo 
2. cd to `/your/directory/` and `go build -i && ./go-trex-runner` in your terminal
Includes the binary file, use that instead if it doesn't compile on your version of MacOS

### Cross Platform Compile
##### compile for linux on mac
* Use the provided dockerfile to compile it for `linux amd64` 
##### compile for windows on mac
* 
```
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -tags static -ldflags "-s -w"
```

### Controls 
* SPACE BAR to jump
* Game auto restarts after 1 second of game over

### Useful links for references
1. https://github.com/veandco/sdl2
