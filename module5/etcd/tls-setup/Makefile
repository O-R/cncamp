.PHONY: cfssl ca req clean

CFSSL	= @env PATH=$(GOPATH)/bin:$(PATH) cfssl
JSON	= env PATH=$(GOPATH)/bin:$(PATH) cfssljson

all:  ca req

cfssl:
	HTTPS_PROXY=127.0.0.1:12639 go get -u -tags nopkcs11 github.com/cloudflare/cfssl/cmd/cfssl
	HTTPS_PROXY=127.0.0.1:12639 go get -u github.com/cloudflare/cfssl/cmd/cfssljson
	HTTPS_PROXY=127.0.0.1:12639 go get -u github.com/mattn/goreman

ca:
	mkdir -p ${certdir}
	$(CFSSL) gencert -initca config/ca-csr.json | $(JSON) -bare ${certdir}/ca

req:
	$(CFSSL) gencert \
	  -ca ${certdir}/ca.pem \
	  -ca-key ${certdir}/ca-key.pem \
	  -config config/ca-config.json \
	  config/req-csr.json | $(JSON) -bare ${certdir}/${infra0}
	$(CFSSL) gencert \
	  -ca ${certdir}/ca.pem \
	  -ca-key ${certdir}/ca-key.pem \
	  -config config/ca-config.json \
	  config/req-csr.json | $(JSON) -bare ${certdir}/${infra1}
	$(CFSSL) gencert \
	  -ca ${certdir}/ca.pem \
	  -ca-key ${certdir}/ca-key.pem \
	  -config config/ca-config.json \
	  config/req-csr.json | $(JSON) -bare ${certdir}/${infra2}
	$(CFSSL) gencert \
	  -ca ${certdir}/ca.pem \
	  -ca-key ${certdir}/ca-key.pem \
	  -config config/ca-config.json \
	  config/req-csr.json | $(JSON) -bare ${certdir}/peer-${infra0}
	$(CFSSL) gencert \
	  -ca ${certdir}/ca.pem \
	  -ca-key ${certdir}/ca-key.pem \
	  -config config/ca-config.json \
	  config/req-csr.json | $(JSON) -bare ${certdir}/peer-${infra1}
	$(CFSSL) gencert \
	  -ca ${certdir}/ca.pem \
	  -ca-key ${certdir}/ca-key.pem \
	  -config config/ca-config.json \
	  config/req-csr.json | $(JSON) -bare ${certdir}/peer-${infra2}

clean:
	rm -rf ${certdir}