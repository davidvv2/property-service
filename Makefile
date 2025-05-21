.PHONY: all build_grpc push_grpc build_gateway push_gateway deploy restart delete start_server

# Set your image name and other variables
GRPC_IMAGE       = gcr.io/property-project-459018/property-service-server:latest
GATEWAY_IMAGE    = gcr.io/property-project-459018/property-service-gateway:latest
GRPC_DOCKER      = docker/Dockerfile.server
GATEWAY_DOCKER   = docker/Dockerfile.gateway
DEPLOYMENTS      = deployments

all: build_grpc build_gateway push_grpc push_gateway deploy

build_grpc:
	docker build -t $(GRPC_IMAGE) -f $(GRPC_DOCKER) .

push_grpc: build_grpc
	docker push $(GRPC_IMAGE)

build_gateway:
	docker build -t $(GATEWAY_IMAGE) -f $(GATEWAY_DOCKER) .

push_gateway: build_gateway
	docker push $(GATEWAY_IMAGE)

deploy:
	kubectl apply -f $(DEPLOYMENTS)/configmap.yaml
	kubectl apply -f $(DEPLOYMENTS)/deployment.yaml
	kubectl apply -f $(DEPLOYMENTS)/service-grpc.yaml
	kubectl apply -f $(DEPLOYMENTS)/gateway-deployment.yaml
	kubectl apply -f $(DEPLOYMENTS)/service-gateway.yaml
	kubectl apply -f $(DEPLOYMENTS)/redis.yaml

restart:
	kubectl rollout restart deployment/property-service
	kubectl rollout restart deployment/property-service-gateway

delete:
	kubectl delete -f $(DEPLOYMENTS)/deployment.yaml
	kubectl delete -f $(DEPLOYMENTS)/service-grpc.yaml
	kubectl delete -f $(DEPLOYMENTS)/gateway-deployment.yaml
	kubectl delete -f $(DEPLOYMENTS)/service-gateway.yaml
	kubectl delete -f $(DEPLOYMENTS)/redis.yaml

start_server:
	kubectl port-forward deployment/property-service-gateway 8888:8080

gen_proto:
	@mkdir -p ./api/proto/gen
	@for f in $(shell find ./api/proto -maxdepth 1 -name '*.proto'); do \
		echo "Generating $$f"; \
		protoc -I=./api/proto -I=./api/proto/third_party/googleapis \
			--go_out=paths=source_relative:./api/proto/gen \
			--go-grpc_out=paths=source_relative:./api/proto/gen \
			--grpc-gateway_out=logtostderr=true,paths=source_relative:./api/proto/gen $$f; \
	done