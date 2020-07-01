FROM golang:1.14.4-alpine AS base
ENV CGO_ENABLED 0
COPY ./ /src
WORKDIR /src
RUN go mod download

FROM base AS build
RUN go build -o /out/app ./cmd/server/main.go

FROM base AS test
RUN go test -v ./...

FROM scratch AS bin
COPY --from=build /out/app /
RUN /app

FROM base AS watch
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build ./cmd/server/main.go" --command=./main