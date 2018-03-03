package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile("AUTO_INCREMENT=\\d+ ")
var reHeader = regexp.MustCompile(
	`/\*\s+SQLyog Community\s+MySQL - [\w\d.-]+ : Database - [\w\d_]+\s+\*+\s+\*/\s+`,
)

func main() {
	arquivos, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, arq := range arquivos {
		if filepath.Ext(arq.Name()) == ".sql" {
			err = removeAutoIncrementHeader(filepath.Join("./", arq.Name()))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

//remove AUTO_INCREMENT=\d+ e header do SQLyog
func removeAutoIncrementHeader(filepath string) error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	out := re.ReplaceAllString(string(bytes), "")
	out = reHeader.ReplaceAllString(out, "")
	return ioutil.WriteFile(filepath, []byte(out), 0664)
}
