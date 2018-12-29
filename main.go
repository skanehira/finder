package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	exclude = flag.String("e", "node_modules", "execlude target")
	dir     = "."
)

func parseArgs() {
	flag.Parse()
	target := flag.Arg(0)
	if target != "" {
		dir = target
	}

	sep := string(os.PathSeparator)
	if !strings.HasSuffix(dir, "/") {
		dir += sep
	}
}

func finder(dir string) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		name := info.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}

		if strings.Index(name, *exclude) != -1 {
			continue
		}

		file := filepath.Join(dir, name)

		if info.IsDir() {
			finder(file)
		} else {
			fmt.Println(file)
		}
	}
}

func main() {
	parseArgs()
	finder(dir)
}
