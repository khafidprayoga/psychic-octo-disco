## About
Build Todo APP with Golang as backend services challenge by [https://devcode.gethired.id/](https://devcode.gethired.id/)

## For Future
- Add abstraction on use case layer return function with go generic
- Add error (zap2) with its context for consumer, and server log (debug and secret not to be public exposed)

## How to use this repo
1. Start `builder/docker-compose-dev` to running mysql services
2. Run `go play/play.go` to auto migrate tables and seed new dummy data
3. Run backend services `go run main.go`, dont forget to set the env

## Footnote For Reviewer
Just run `docker-compose -f test-deploy up` it will start my latest docker container, rather than using too long CLI for `docker run`
