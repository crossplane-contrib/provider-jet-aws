FROM bitnami/kubectl:1.20.12

COPY package/crds /Users/monus/go/src/github.com/crossplane/provider-tf-aws/package/crds

# Build locally and copy. It's too late already.
# RUN GOOS=linux GOARCH=amd64 go build -o cmd/test/main cmd/test/main.go && chmod +x cmd/test/main

COPY cmd/test/main /main

ENTRYPOINT ["./main"]