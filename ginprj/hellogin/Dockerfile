FROM centos 

COPY ./go /apps/go
COPY ./go.sh /etc/profile.d/go.sh

RUN mkdir -p /go/{src,bin,pkg}

WORKDIR /go
