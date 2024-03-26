FROM golang:1.19.2-bullseye
 
WORKDIR /app
 
COPY . .
 
RUN go mod download
 
RUN go build -o /server
 
EXPOSE 8080/tcp
 
CMD [ "/server", "8080" ]
