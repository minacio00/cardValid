FROM golang:1.21-alpine
WORKDIR /app-docker
EXPOSE 8080
COPY . ./
RUN go build -o . .
CMD [ "./cardValid" ]