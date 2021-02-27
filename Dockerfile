FROM alpine
ADD go-user /go-user
ENTRYPOINT [ "/go-user" ]
