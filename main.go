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
var reDefiner = regexp.MustCompile(
	"definer=`root`@`[^`]+` ",
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
	out := strings.ToLower(string(bytes))
	out = re.ReplaceAllString(out, "")
	out = reHeader.ReplaceAllString(out, "")
	out = reDefiner.ReplaceAllString(out, "")
	out = strings.ReplaceAll(
		out,
		"no_auto_value_on_zero",
		"strict_trans_tables,no_zero_in_date,no_zero_date,error_for_division_by_zero,no_engine_substitution",
	)
	return ioutil.WriteFile(filepath, []byte(out), 0664)
}
