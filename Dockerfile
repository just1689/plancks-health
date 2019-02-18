FROM golang:1.10.1 as pcdeperhealth
WORKDIR /go/src/github.com/plancks-cloud/plancks-health
COPY Gopkg.toml Gopkg.lock vendor ./
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v -vendor-only

FROM pcdeperhealth as builderhealth
WORKDIR /go/src/github.com/plancks-cloud/plancks-health
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pc-health .

FROM scratch
WORKDIR /
COPY --from=builderhealth /go/src/github.com/plancks-cloud/plancks-health/pc-health .
COPY --from=builderhealth /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/pc-health"]