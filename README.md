# psychic-guide

Physic-guide is a NSQ consummer to the topic: `effective_bassoon`

## How to use

Start nsqd:

``` shell
nsqlookupd & 
nsqd --lookupd-tcp-address=127.0.0.1:4160 &
```

You will need a producer available [here](https://github.com/MaxencePPierre/effective-bassoon)

Then run the consumer with the following command:

```shell
go run consumer.go
```

## Build and run in docker

To build the application:

```shell
docker build --tag psychic-guide .
```

In order to run this newly created image, use docker run command:

```shell
docker run -p 4161:4161 -it psychic-guide
```
