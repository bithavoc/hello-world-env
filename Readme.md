## hello-world-env

A docker image for learning and testing with web apps and environment variables.

```
bithavoc/hello-world-env
```

### Using with local docker

```
docker run --publish=3000:3000 --rm bithavoc/hello-world-env
```

```
GET http://localhost:3000/this
```

Prints

```
Hi there, I love this!
```

### Using with command arguments

```
docker run --publish=3000:3000 --rm bithavoc/hello-world-env /root/app -message "cool by arg"
```

### Using with environment variables

```
docker run --publish=3000:3000 --rm -e "MESSAGE=cool by env" bithavoc/hello-world-env
```

```
GET http://localhost:3000/this
```

Prints

```
Hi there, I love this! cool
```
