# syntax=docker/dockerfile:1
FROM golang:1.18-bullseye
WORKDIR /forum
COPY . .
RUN go mod download
RUN go build
EXPOSE 8080
RUN ls
RUN chmod +x forum
CMD [ "./forum" ]