# BDAS
BDAS = Big Data Analysis Systems. Repository for practical tasks on the subject BDAS.

## Overview

### Lab1. Data obfuscation and deobfuscation.

## Usage

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

## Examples

To obfuscate file:
`go run ./cmd/lab1/main.go -i ./examples/example.xml -o ./examples/obfuscated-example.xml`

To deobfuscate file:
`go run ./cmd/lab1/main.go -i ./examples/obfuscated-example.xml -o ./examples/example.xml -d`

To obfuscate file without specifying a output file:
`go run ./cmd/lab1/main.go -i ./examples/example.xml`

To deobfuscate file without specifying a output file:
`go run ./cmd/lab1/main.go -i ./examples/obfuscated-example.xml -d`


## DevKit

> TODO

## Tests


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

## Project structure

The main program code is in the `/cmd` directory. Each laboratory has its own directory with an appropriate name `/cmd/lab<n>`. Inside the application there is a `main.go` file, which is the main endpoint of program.

All examples are located in the `examles` directory.

All tests are ...

All build files ...

ci ...

> TODO

### Stack

* GoLang go1.12.7
* Docker 19.03.2
* OS Ubuntu 18.04
> TODO

## Usage

In the project directory, you can run:

> TODO

### Build

> TODO

### Deploy

> TODO

### Tests

> TODO
