FROM golang:1.10.1 as pcdeperhealth
WORKDIR /go/src/git.amabanana.com/plancks-cloud/pc-go-brutus
COPY Gopkg.toml Gopkg.lock vendor ./
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v -vendor-only

FROM pcdeperhealth as builderhealth
WORKDIR /go/src/git.amabanana.com/plancks-cloud/pc-go-brutus
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pc-health .

FROM scratch
WORKDIR /
COPY --from=builderhealth /go/src/git.amabanana.com/plancks-cloud/pc-go-brutus/pc-health .
COPY --from=builderhealth /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/pc-health"]