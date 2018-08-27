package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func tree(pre, path string, level int) {
	tmppath := path[strings.LastIndex(path, "\\"):]
	fmt.Println(pre + "|---" + tmppath)
	pre += "|   "
	if level <= 0 {
		return
	}
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println("read path failed ", err)
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
	path := filepath.Join("D:\\")

	tree("", path, 3)
}
