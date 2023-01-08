FROM golang:alpine

ARG EXE=users

COPY . /usr/local
COPY ../proto /usr/local/proto
WORKDIR /usr/local
RUN ./build.sh

# # CMD ["./${EXE}"]