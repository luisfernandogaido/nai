package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("auto_increment=\\d+ ")
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
			if err := removeAutoIncrementHeader(filepath.Join("./", arq.Name())); err != nil {
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
	out := strings.ToLower(string(bytes))
	out = re.ReplaceAllString(out, "")
	out = reHeader.ReplaceAllString(out, "")
	return ioutil.WriteFile(filepath, []byte(out), 0664)
}
