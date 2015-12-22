FROM alpine:3.2
ADD geo-srv /
ENTRYPOINT [ "/geo-srv" ]
