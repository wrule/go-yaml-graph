package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	bytes, err := ioutil.ReadFile("mylua.yaml")
	if err != nil {
		panic("读取文件出错")
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(bytes, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
		panic("yaml解析出错")
	}
	fmt.Println(m)
}
