FROM golang
LABEL maintainer Sergey Nesterov

RUN mkdir /go/src/telehorn
COPY . /go/src/telehorn

WORKDIR /go/src/telehorn

COPY .env.docker .env
COPY glide* .

# Install glide #
RUN curl https://glide.sh/get | sh

# Install fresh (live reload) #
RUN go get github.com/pilu/fresh
ENV PATH="$GOPATH/bin:$PATH"

RUN glide up
