
FROM ubuntu:latest


ENV TZ=Europe/Madrid
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone


RUN apt-get update 

RUN apt-get -y install curl
RUN apt-get -y install git
RUN apt-get -y install golang

# Install app dependencies
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"


# Bundle app source
COPY postter /opt/postter


EXPOSE 8082

CMD ["/opt/postter"]