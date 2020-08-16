RELEASE_TAG := $(shell ./release.sh 2> /dev/null)

all: src deploy

src: apiserver webapp

protos: proto-greeter

webapp: build-webapp push-webapp 

apiserver: build-apiserver push-apiserver 

#--go_opt=paths=source_relative
proto-%:
	cd src/apiserver/$*/pb && protoc --proto_path=$(GOPATH)/src:. \
			--twirp_out=. \
			--go_out=. \
			--descriptor_set_out=$*.pb \
			$*.proto

deploy: generate-apiserver generate-webapp
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
	find \( -wholename "*/pb/*.twirp.go" -o -wholename "*/pb/*.pb.go" -o -wholename "*/pb/*.pb" \) -delete

.PHONY: deploy apiserver
