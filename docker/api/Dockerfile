FROM golang:1.12-alpine AS build_img


WORKDIR /tmp/build-app

COPY go.mod .
RUN apk add git \
	&& go mod download

COPY . .

RUN go build -o waco .

FROM alpine:3.9

COPY --from=build_img /tmp/build-app/waco /app/waco

CMD ["/app/waco"]

