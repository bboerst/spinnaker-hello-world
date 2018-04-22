FROM scratch

EXPOSE 9090

COPY ./src/build/gowebapp /helloworld/

CMD ["/helloworld/gowebapp"]