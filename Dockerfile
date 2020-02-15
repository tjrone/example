FROM busybox
ADD example /
ENV Num 10
EXPOSE 9090
CMD ["/example"]