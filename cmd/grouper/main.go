package main

import (
	"os"

	"io/ioutil"

	"github.com/athom/grouper/web"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	confPath := dir + `/config.json`
	argv := os.Args
	if len(argv) > 1 {
		confPath = argv[1]
	}

	fd, err := os.OpenFile(confPath, os.O_RDONLY, 0644)
	defer fd.Close()
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}

	// Example conf
	conf := string(b)
	web.Run(conf)
}
