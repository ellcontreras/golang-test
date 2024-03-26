# Stratplus Challenge

## There is a file called Stratplus.postman_collection.json with Postman collection to test
About time I did not write unit tests :c

### Commands
Since I run and build on win11 I had issues with docker, but this works
``` shell
    docker build --no-cache -t stratplus -f Dockerfile .
    docker run -p 127.0.0.1:8097:8080 stratplus
```

This commands will download our dependencies

```shell
$ npm i
$ go mod download
$ go mod tidy
$ go mod vendor
```
## Running the tests
Use the following command to run your tests

Using the Makefile:

```shell
$ make test
$ make validate
$ make coverage
```

Using bash:

```shell
$ bash scripts/test.sh
$ bash scripts/validate.sh
$ bash scripts/coverage.sh
```

If you want to check which lines are not being covered you can use this command

```shell
$ make coverage
$ bash scripts/coverage.sh
```