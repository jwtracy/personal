RELEASE_TAG := $(shell ./release.sh 2> /dev/null)

all: src deploy

src: apiserver

protos: proto-greeter

apiserver: build-apiserver push-apiserver 

proto-%:
	protoc -Isrc/apiserver/$* \
			--go_opt=paths=source_relative \
			--go_out=plugins=grpc:src/apiserver/$* \
			--descriptor_set_out=src/apiserver/$*/pb/$*.pb \
			pb/$*.proto

deploy: generate-apiserver
	kustomize build deploy/ > deploy/manifest.yaml
	kubectl apply -f deploy/manifest.yaml

generate-%:
	cd deploy/$*/ && kustomize edit set image app-$*=registry.digitalocean.com/jwtracy-personal/app-$*:$(RELEASE_TAG)
	cd deploy/$*/ && kustomize build > manifest.yaml

build-%: generate-apiserver protos
	docker build -t registry.digitalocean.com/jwtracy-personal/app-$*:$(RELEASE_TAG) \
			src/$*/

push-%: 
	docker push registry.digitalocean.com/jwtracy-personal/app-$*:$(RELEASE_TAG)
	
clean:
	find deploy -name manifest.yaml -delete
	find \( -wholename "*/pb/*.pb.go" -o -wholename "*/pb/*.pb" \) -delete

.PHONY: deploy apiserver
