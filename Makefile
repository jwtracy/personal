RELEASE_TAG ?= latest

all: src deploy

deploy: generate-apiserver
	kustomize build deploy/ > deploy/manifest.yaml
	kubectl apply -f deploy/manifest.yaml

generate-%:
	kustomize build deploy/$*/ > deploy/$*/manifest.yaml

src: apiserver

apiserver: proto-apiserver build-apiserver push-apiserver 

proto-%:
	protoc -Isrc/apiserver/$* --go_opt=paths=source_relative --go_out=plugins=grpc:. pb/$*.proto

build-%: 
	docker build -t registry.digitalocean.com/jwtracy-personal/app-$*:$(RELEASE_TAG) \
			src/$*/

push-%: 
	docker push registry.digitalocean.com/jwtracy-personal/app-$*:$(RELEASE_TAG)
	
deploy-%: 
	kubectl apply -k deploy/$*/$(ENV)

clean:
	find deploy -name manifest.yaml -delete
	find -wholename "*/pb/*.pb.go" -delete
