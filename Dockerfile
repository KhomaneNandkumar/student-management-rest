FROM golang:1.23

WORKDIR /student-mgmt-rest

COPY . .

COPY cmd//student-management/.env .env

RUN go build -o student-mgmt-rest ./cmd/student-management/main.go

CMD ["./student-mgmt-rest"]


