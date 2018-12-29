package main

import (
	"flag"
	"fmt"
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
}

func finder(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if strings.HasPrefix(path, ".") || strings.Contains(path, *exclude) || info.IsDir() {
			return nil
		}

		fmt.Println(path)

		return nil
	})
}

func main() {
	parseArgs()
	finder(dir)
}
