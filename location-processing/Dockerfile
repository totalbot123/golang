FROM alpine

ARG EXE
ENV EXE=${EXE}

COPY bin/${EXE} /bin/${EXE}

CMD ["sh", "-c", "${EXE}"]