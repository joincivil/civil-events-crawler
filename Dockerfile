#FROM alpine:3.7
FROM golang:1.12.7
ADD build build
ADD build/crawler /crawler
RUN chmod u+x /crawler

EXPOSE 9000
CMD ["/crawler", "-logtostderr=true", "-stderrthreshold=INFO", "-v=2"]
