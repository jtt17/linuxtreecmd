package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func tree(pre, path string, level int) {
	id := strings.LastIndex(path, "/")
	if id == -1 {
		id = 0
	}
	tmppath := path[id:]
	fmt.Println(pre + "|---" + tmppath)
	pre += "|   "
	if level <= 0 {
		return
	}
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Printf("read path: %s  failed : %v\n", path, err)
		return
	}

	for _, x := range files {

		if x.IsDir() {
			tree(pre, filepath.Join(path, x.Name()), level-1)
		} else {
			fmt.Println(pre + "|--- " + x.Name())
		}
	}
}

func main() {
	var path string
	var level int
	var err error
	if len(os.Args) == 1 {
		path, err = os.Getwd()
		if err != nil {
			panic(err)
		}
		level = 999999999
	} else {
		path = os.Args[1]
		level, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}
	tree("", path, level)
}
