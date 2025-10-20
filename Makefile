# Makefile for solace-dsemp-agent

.PHONY: build clean swagger_semp posttargetstatejson

build:
	go mod tidy
	go build -o solace-dsemp-agent main.go

clean:
	rm -rf build

swagger_semp:
	mkdir -p spec
	curl -X GET -u admin:admin -o spec/semp-v2-swagger-config.json host.docker.internal:8080/SEMP/v2/config/spec
	mkdir -p semp_swagger
	docker run --rm --user $(id -u):$(id -g) -v "$(WORKSPACE_LOCALDIR)/spec:/spec" -v "$(WORKSPACE_LOCALDIR)/semp_swagger:/semp_swagger" swaggerapi/swagger-codegen-cli-v3:3.0.71 generate -l go -i /spec/semp-v2-swagger-config.json -o /semp_swagger/config --type-mappings boolean=*bool
	sudo chown -R $(id -u):$(id -g) semp_swagger/

posttargetstatejson:
	curl -X POST -d @targetstate.json -u another-user:solace-dsemp-agent http://host.docker.internal:9000/TOPIC/config/my-broker/target-state