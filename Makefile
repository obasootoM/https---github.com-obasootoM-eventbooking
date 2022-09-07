main:
	go run main.go

certificate:
	go run $GOROOT/usr/local/go/src/crypto/tls/generate_cert.go --host=localhost    
	
.PHONY: main certificate