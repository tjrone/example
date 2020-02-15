FROM ubuntu
ADD example /
ENV Num 10
EXPOSE 9090
ENTRYPOINT ["./example"]