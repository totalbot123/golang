FROM golang:alpine as builder

ARG EXE=location-processing
ENV GOOS=linux
ENV GOARCH=arm64

COPY . /usr/local
COPY ../proto /usr/local/proto
WORKDIR /usr/local
RUN ./build.sh
RUN chmod +x bin/${EXE}

# FROM alpine

# COPY /user/local/src/bin/${EXE} /bin/${EXE}

# CMD ["/bin/${EXE}"]