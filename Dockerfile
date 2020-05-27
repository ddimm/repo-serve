FROM balenalib/raspberrypi4-64-golang:latest as build

WORKDIR /go/src/github.com/ddimm/repo-serve

COPY . . 
RUN go build

FROM balenalib/raspberrypi4-64-debian:stretch
EXPOSE 8080
COPY --from=build /go/src/github.com/ddimm/repo-serve .
CMD [ "./repo-serve" ]