# Build Monzo
#FROM golang:1.13 AS build-env-monzo
#RUN git clone https://github.com/monzo/envoy-preflight.git /src
##disable crosscompiling
#ENV CGO_ENABLED=0
##compile linux only
#ENV GOOS=linux
#RUN cd /src && go get -v -d && go build -ldflags '-w -s' -a -installsuffix cgo -o envoy-preflight 


# build stage
FROM golang:1.13 AS build-env
ADD . /src
#disable crosscompiling
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux
RUN go get -v -d && go build -ldflags '-w -s' -a -installsuffix cgo -o postgres_load

# final stage
FROM scratch
COPY --from=build-env /postgres_load /app/
WORKDIR /app
