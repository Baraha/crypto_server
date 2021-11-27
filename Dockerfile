FROM golang:latest
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
EXPOSE 8080
CMD ./main