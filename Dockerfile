FROM scratch
ADD geo-srv /
ENTRYPOINT [ "/geo-srv" ]
