CC=go
OUTDIR=bin
TARGETARCH=amd64
CGO_ENABLED=0
 

all: setup build

build:
	CGO_ENABLED=$(CGO_ENABLED) GOARCH=$(TARGETARCH) go build -o $(OUTDIR)/dtf .

setup:
	go mod download -x
	mkdir -p $(OUTDIR)

run:
	cp ./api.yaml $(OUTDIR)/api.yaml 
	cd $(OUTDIR) && ./dtf
	cd .. 

clean: 
	rm -rf $(OUTDIR)
 
 


