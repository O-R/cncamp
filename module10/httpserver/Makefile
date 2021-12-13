build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: 
	echo "building httpserver container: huangzhihong/cncamp:${tag}"
	docker build -t huangzhihong/cncamp:${tag} .

push: release
	echo "pushing huangzhihong/cncamp:${tag}"
	docker push huangzhihong/cncamp:${tag}