FROM golang:1.14-alpine as builder

WORKDIR /go/src/squareroot
COPY ./src/squareroot/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o squareroot
RUN ls

FROM scratch
COPY --from=builder  /go/src/squareroot .

CMD [ "./squareroot" ]

