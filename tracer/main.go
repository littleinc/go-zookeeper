package main

import (
	"github.com/samuel/go-zookeeper"
)

func main() {
	zk.StartTracer("127.0.0.1:2182", "127.0.0.1:2181")
}
