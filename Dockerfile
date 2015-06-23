FROM google/golang

RUN go get github.com/tools/godep

RUN mkdir -p /gopath/src/github.com/lavab/onboarding-webhook
ADD . /gopath/src/github.com/lavab/onboarding-webhook
RUN cd /gopath/src/github.com/lavab/onboarding-webhook && godep go install

CMD []
ENTRYPOINT ["/gopath/bin/onboarding-webhook"]
