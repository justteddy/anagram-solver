run: build
	@echo "Running anagram-solver..."
	@${GOPATH}/bin/anagram-solver

build: generate
	@echo "Building application..."
	@CGO_ENABLED=0 go install -a

generate: clean-generated go-swagger generate-server

go-swagger:
	@echo "Installing Go-Swagger code generator..."
	@go get -u github.com/go-swagger/go-swagger/cmd/swagger

generate-server:
	@echo "Generating server..."
	@mkdir -p ./generated
	@${GOPATH}/bin/swagger generate server -f ./swagger.yml -t ./generated --exclude-main --regenerate-configureapi

clean-generated:
	@echo "Removing generated files..."
	@rm -rf ./app/generated