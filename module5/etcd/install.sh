

if [ -z "$1" ]; then
        echo "require install location..."
	exit 1;
fi

ETCD_VER=v3.5.1

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}
INSTALL_ROOT=$1
INSTALL_SUBPATH=$2
INSTALL_PATH=${INSTALL_ROOT}/${INSTALL_SUBPATH}



rm -f /${INSTALL_ROOT}/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /${INSTALL_ROOT}/${INSTALL_SUBPATH} && mkdir -p /${INSTALL_ROOT}/${INSTALL_SUBPATH}

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /${INSTALL_ROOT}/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /${INSTALL_ROOT}/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /${INSTALL_ROOT}/${INSTALL_SUBPATH} --strip-components=1
rm -f /${INSTALL_ROOT}/etcd-${ETCD_VER}-linux-amd64.tar.gz

/${INSTALL_ROOT}/${INSTALL_SUBPATH}/etcd --version
/${INSTALL_ROOT}/${INSTALL_SUBPATH}/etcdctl version
/${INSTALL_ROOT}/${INSTALL_SUBPATH}/etcdutl version
