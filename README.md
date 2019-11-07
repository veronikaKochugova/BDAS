# BDAS
BDAS = Big Data Analysis Systems. Repository for practical tasks on the subject BDAS.

## Project structure

The main program code is in the `/cmd` directory. Each laboratory is separate application (app) and has its own directory with an appropriate name `/cmd/lab[n]`. Inside the appl there is a `main.go` file, which is the main endpoint of program.

All examples are located in the `examles/` directory.

All unit tests are located in the same directory as the source code and end with `_test.go`. 

In `/build` directory [executable files](#run-builds) are located. 

In `/build/docker` Dockerfiles are located. 

> TODO

## Technology stack

* GoLang go1.12.7
* Docker 19.03.2
* OS Ubuntu 18.04
> TODO

## Usage 

The following commands are described for run/build `lab1`, but are valid for all laboratory (lab1..4).

### Run from src

```
go run ./cmd/lab1/main.go [OPTIONS]
```

### Run builds

For Windows:
```
./build/lab1.exe [OPTIONS]
```
For Linux:
```
./build/lab1 [OPTIONS]
```

### Build from src

This command create executable file with <custom name>.
```
go build -o ./build/<custom name> ./cmd/lab1/
```

### Docker

> You can use the docker container so as not to download the go language and all the dependencies.

Image pushed on [DockerHub repo](https://hub.docker.com/repository/docker/veronikakochugova/bdas).

To build docker image from src:
```
docker build --rm -f "build/docker/Dockerfile" -t veronikakochugova/bdas:latest .
```

> You can skip previous step and use image from DockerHub like in next step.

Run app in docker container:
```
docker run -it veronikakochugova/bdas:latest /bin/bash
```
Inside the container, you can run the application (lab1..4) from both [source](#run-from-src) and [build](#run-builds).

### CI

> TODO

## Lab overview

### Lab1. Data obfuscation and deobfuscation.

#### Usage

```
usage: main.go [-h|--help] -i|--input "<value>" [-o|--output "<value>"]
               [-d|--deobfuscate]

               Obfuscate/Deobfuscate text from input file

Arguments:

  -h  --help         Print help information
  -i  --input        Path to input file (required)
  -o  --output       Path to output file (by default: [input file].result)
  -d  --deobfuscate  Deobfuscation flag
```

#### Examples

To obfuscate file:
`go run ./cmd/lab1/main.go -i ./examples/example.xml -o ./examples/obfuscated-example.xml`

To deobfuscate file:
`go run ./cmd/lab1/main.go -i ./examples/obfuscated-example.xml -o ./examples/example.xml -d`

To obfuscate file without specifying a output file:
`go run ./cmd/lab1/main.go -i ./examples/example.xml`

To deobfuscate file without specifying a output file:
`go run ./cmd/lab1/main.go -i ./examples/obfuscated-example.xml -d`

#### Tests

To run unit tests manualy with coverage stats:
```
go test -cover .\cmd\lab1\obfuscation\
```

*100% coverage*

### Lab2. Encryption scenarios using standart libraries

Creating program for demonstration of testing encryption scenarios using standart libraries: 

- BouncyCastle
- SafeNet
- KeySecure
- Gemalto
- SunJCE

> TODO

### Lab3. SSL/TLS interacting

Design and implementation of client-server application interacting via HTTPS protocol using encryption keys for SSL/TLS (Two-Way TLS)

> TODO

### Lab4. Gateway

Creating router for client-server application working through Gateway using Netflix Zuul library.

> TODO


