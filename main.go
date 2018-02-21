package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile("AUTO_INCREMENT=\\d+ ")

func main() {
	arquivos, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, arq := range arquivos {
		if filepath.Ext(arq.Name()) == ".sql" {
			err = removeAutoIncrement(filepath.Join("./", arq.Name()))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func removeAutoIncrement(filepath string) error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	out := re.ReplaceAllString(string(bytes), "")
	return ioutil.WriteFile(filepath, []byte(out), 0664)
}
