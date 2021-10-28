.PHONY: cert install cluster clean

CERTROOT	= ./tls-setup
CERTSUBDIR	= certs
CERTDIR	= $(CERTROOT)/$(CERTSUBDIR)

INFRA0	= etcd0
INFRA1	= etcd1
INFRA2	= etcd2

etcd0_client_url	= https://localhost:2379
etcd0_peer_url	= https://localhost:2380
etcd1_client_url	= https://localhost:12379
etcd1_peer_url	= https://localhost:12380
etcd2_client_url	= https://localhost:22379
etcd2_peer_url	= https://localhost:22380

etcd_cluster	= $(INFRA0)=$($(INFRA0)_peer_url),$(INFRA1)=$($(INFRA1)_peer_url),$(INFRA2)=$($(INFRA2)_peer_url)
etcd_cluster_token	= etcd-cluster-1

INSTALLROOT	= /data/faror/etcd
INSTALLBIN	= bin
INSTALLPATH	= $(INSTALLROOT)/$(INSTALLBIN)

all: install cert cluster 

cert:
	make -C $(CERTROOT) all certdir=$(CERTSUBDIR) infra0=$(INFRA0) infra1=$(INFRA1) infra2=$(INFRA2)

install:
	./install.sh $(INSTALLROOT) $(INSTALLBIN)

cluster:$(INFRA0) $(INFRA1) $(INFRA2)

$(INFRA0) $(INFRA1) $(INFRA2):
	touch cluster.log && chmod +w cluster.log
	nohup $(INSTALLPATH)/etcd \
		--name $@ \
		--listen-client-urls $($@_client_url) \
		--advertise-client-urls $($@_client_url) \
		--listen-peer-urls $($@_peer_url) \
		--initial-advertise-peer-urls $($@_peer_url) \
		--initial-cluster-token $(etcd_cluster_token) \
		--initial-cluster '$(etcd_cluster)' \
		--initial-cluster-state new \
		--cert-file=$(CERTDIR)/$@.pem \
		--key-file=$(CERTDIR)/$@-key.pem \
		--peer-cert-file=$(CERTDIR)/peer-$@.pem \
		--peer-key-file=$(CERTDIR)/peer-$@-key.pem \
		--client-cert-auth --trusted-ca-file=$(CERTDIR)/ca.pem \
		--peer-client-cert-auth --peer-trusted-ca-file=$(CERTDIR)/ca.pem >cluster.log 2>&1 &
clean:
	kill -TERM $$(ps -aux | grep "/etcd" | grep -v grep | awk '{print $$2}') && rm -rf $(INSTALLBIN) $(CERTDIR)