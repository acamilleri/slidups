# slidups

[![Go Report Card](https://goreportcard.com/badge/github.com/acamilleri/slidups)](https://goreportcard.com/report/github.com/acamilleri/slidups)

Mini go web server to easily upload my reveal-md presentations to my self-hosted instance of [webpro/reveal-mid](https://github.com/webpro/reveal-md).

## Installation

### Docker

Run with docker
```
$ docker run -p 80:8080 acamilleri/slidups:latest
```

### From source

clone the project
```
$ git clone git@github.com:acamilleri/slidups.git
```

run
```
make run
```

Note: with `make run` command, slidups will start with /tmp directory as upload destination. 

## Usage

Push file
```
$ curl -F file=@presentation.md http://localhost/upload
```

Push file to a custom directory
```
$ curl -F file=@presentation.md http://localhost/upload?destionation=/slides/custom
```

# Build from source

clone the project
```
$ git clone git@github.com:acamilleri/slidups.git
```

build
```
$ make build
```

# TODO:
- Improve code (tests!)
- Adding metrics
