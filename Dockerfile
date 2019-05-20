FROM alpine:3.2
ADD location-srv /
ENTRYPOINT [ "/location-srv" ]
