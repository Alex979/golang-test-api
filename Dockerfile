FROM golang:1.15.2-alpine3.12

WORKDIR /usr/src/app
COPY . .

# Install required packages
RUN apk add --no-cache gcc g++ sqlite

# Download dependencies and build app
RUN go mod download
RUN go build

EXPOSE 8080
CMD [ "./golang-test" ]