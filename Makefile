compile = env GOOS=linux  GOARCH=arm64  go build -v -ldflags '-s -w -v' -o
upx_url = https://github.com/upx/upx/releases/download/v3.96
upx_file = upx-3.96-arm64_linux
percent = go tool cover -func=coverage.out | sed '$d' | sort -g -r -k 3
test_to_file = go test -coverprofile=coverage.out
mock_to_file = mockery
current_path = $(shell pwd)
current_dir = $(shell basename "$(PWD)")
gopath = $(GOPATH)
mock_to_file = mockery --dir

test:
	go test -count=2 ./service/ ./config/aws/

coverage:
	$(test_to_file) ./config/aws/ ./service/
	go tool cover -html=coverage.out

validate:
	make imports
	make gofmt
	$(test_to_file) ./config/aws/  && $(percent)
	$(test_to_file) ./service/  && $(percent)
	golangci-lint run --fast

percent:
	$(percent)

docs:
	rm -rf $(gopath)/src/godoc/src
	mkdir -p $(gopath)/src/godoc/src
	rsync -a $(current_path) $(gopath)/src/godoc/src --exclude "node_modules" --exclude ".git"
	(cd $(gopath)/src/godoc; godoc -goroot .)
	cd $(current_path)

mock:
	mockery --dir ./gateway --output ./mocks/gateway --all
	mockery --dir ./service --output ./mocks/service --all
	mockery --dir ./repository --output ./mocks/repository --all

mock-clear:
	find mocks -type f ! -name "*main.go" -delete

imports:
	find . -name \*.go ! -path "./mocks/*" ! -path "./node_modules/*" ! -path "./vendor/*" -exec goimports -w {} \;

gofmt:
	find . -name \*.go ! -path "./mocks/*" ! -path "./node_modules/*" ! -path "./vendor/*" -exec gofmt -s -w {} \;
