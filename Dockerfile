  FROM ubuntu

  ENV DEBIAN_FRONTEND=noninteractive

  RUN apt-get update && \
      apt-get install -y libsdl2-image-dev libsdl2-mixer-dev libsdl2-ttf-dev libsdl2-gfx-dev wget git gcc 
  RUN wget -P /tmp https://dl.google.com/go/go1.13.6.linux-amd64.tar.gz  && \
      tar -C /usr/local -xzf /tmp/go1.13.6.linux-amd64.tar.gz && \
      rm /tmp/go1.13.6.linux-amd64.tar.gz

  ENV GOPATH /go
  ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

  RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

  WORKDIR $GOPATH/src

  COPY . .

  RUN go get -v github.com/veandco/go-sdl2/sdl \
      && go get -v github.com/veandco/go-sdl2/img \
      && go get -v github.com/veandco/go-sdl2/mix \
      && go get -v github.com/veandco/go-sdl2/ttf \
      && go get -v github.com/veandco/go-sdl2/gfx \
      && CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags static -ldflags "-s -w"
