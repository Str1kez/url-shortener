FROM golang:1.19.5-bullseye

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

COPY . .

RUN make

CMD ["make", "run-prod"]

