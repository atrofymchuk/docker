FROM ubuntu:18.04

ENV DEBIAN_FRONTEND noninteractive

# Install basics

RUN apt-get update

RUN apt-get install -y software-properties-common && \
    add-apt-repository ppa:ondrej/php && \
    apt-get update

RUN apt-get install -y apache2

RUN apt-get install -y php7.0

RUN apt-get install -y mysql-server mysql-client mysql-common

RUN apt-get install -y curl wget vim

# Install PHP 7.0

RUN apt-get install -y libapache2-mod-php7.0 php7.0-mysql php7.0-mcrypt php7.0-cli php7.0-gd php7.0-curl

# Expose apache.

EXPOSE 80
