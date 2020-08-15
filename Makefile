RELEASE_TAG ?= latest

all: src deploy

deploy: generate-apiserver
	kustomize build deploy/ > deploy/manifest.yaml
	kubectl apply -f deploy/manifest.yaml

generate-%:
	kustomize build deploy/$*/ > deploy/$*/manifest.yaml

src: build-apiserver push-apiserver 

build-%: 
	docker build -t registry.digitalocean.com/jwtracy-personal/app-$*:$(RELEASE_TAG) \
			src/$*/

push-%: 
	docker push registry.digitalocean.com/jwtracy-personal/app-$*:$(RELEASE_TAG)
	
deploy-%: 
	kubectl apply -k deploy/$*/$(ENV)

clean:
	find deploy -name manifest.yaml -delete
