# GoLang Setup

## install and hello world

* `brew update`
* `brew install go --with-cc-all`
* in `.bashrc`, put `export GOPATH=/Users/pivotal/workspace/go`
* actual git clone of a go project is under `/Users/pivotal/workspace/go/src/github.com/thewoolleyman/gitrflow`
* make hello world:
```
package main
import "fmt"

func main() {
  fmt.Println("hello go")
}
```
* `go build`
* `./hellogo`
* export GOOS=linux
* export GOARCH=amd_64
* `go build -o bin/linux-amd_64/hellogo`

## structuring a project

* http://blog.parse.com/2014/05/13/dependency-injection-with-go/

## dependencies

* godep

## Testing

* https://github.com/stretchr/testify
* https://github.com/shageman/gotestit 