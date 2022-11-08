FROM ubuntu:20.04
RUN apt-get update 
RUN apt install docker.io -y
RUN apt update
RUN apt install python3 python3-pip -y
RUN pip3 install docker docker-compose
RUN docker --version

RUN apt install ansible -y
RUN ansible --version

WORKDIR /app

RUN chmod 755 /app

CMD [""]
