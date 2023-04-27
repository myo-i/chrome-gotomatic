FROM ubuntu:latest

RUN apt-get update && apt-get install vim git software-properties-common -y

RUN add-apt-repository ppa:longsleep/golang-backports &&\
    apt install golang-go -y

RUN mkdir -p /opt/chrome-gotomatic

WORKDIR /opt/chrome-gotomatic

COPY . .

CMD ["/bin/bash"]
