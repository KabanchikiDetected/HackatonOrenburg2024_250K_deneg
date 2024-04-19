FROM golang:1.22

WORKDIR /users

# Download all dependencies.
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

EXPOSE 27017
COPY . .
# RUN ls
RUN go build -o bin ./cmd/main.go

CMD [ "./bin" ]