FROM golang:1.22-alpine AS go_build

RUN apk --update --no-cache add git build-base

ENV TZ Asia/Tokyo

WORKDIR /himo-outgame

COPY go.mod go.sum /himo-outgame/
RUN go mod download

COPY . /himo-outgame
# GOARCH はビルド環境に合わせる(Apple Silicon では arm64 になる)
RUN CGO_ENABLED=0 make

FROM alpine:3.20

ENV TZ Asia/Tokyo

RUN apk --update --no-cache add curl tzdata bash ca-certificates jq && \
  update-ca-certificates

WORKDIR /himo-outgame
COPY ./entry-point.sh ./entry-point.sh
COPY --from=go_build /himo-outgame/bin/outgame ./bin/

RUN chmod 755 ./entry-point.sh
ENTRYPOINT [ "./entry-point.sh" ]
