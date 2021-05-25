package sub

import (
	"cyracs/data"
	"encoding/json"
	"io/ioutil"
	"log"
)

type Range data.Range

func Road(file string) Range {
	fl := "config/" + file
	rdata, err := ioutil.ReadFile(fl)
	if err != nil {
		log.Fatal(err)
	}
	var rangeinf Range
	err = json.Unmarshal(rdata, &rangeinf)
	if err != nil {
		log.Fatal(err)
	}
	return rangeinf
}
