run: build
	@echo "Running anagram-solver..."
	@${GOPATH}/bin/anagram-solver

build: generate dep
	@echo "Building application..."
	@CGO_ENABLED=0 go install -a

dep:
	@echo "Resolving dependencies..."
	@go mod tidy

generate: clean-generated generate-server

generate-server:
	@echo "Generating server..."
	@mkdir -p ./generated
	@${GOPATH}/bin/swagger generate server -f ./swagger.yml -t ./generated --exclude-main --regenerate-configureapi

clean-generated:
	@echo "Removing generated files..."
	@rm -rf ./app/generated