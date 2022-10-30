FROM ubuntu:18.04
ARG DEBIAN_FRONTEND=noninteractive
RUN apt update
RUN apt install wget -y
RUN wget https://dl.google.com/go/go1.19.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

