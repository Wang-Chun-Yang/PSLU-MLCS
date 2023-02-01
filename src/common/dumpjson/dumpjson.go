package dumpjson

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	dump     = flag.Bool("dump", false, "whether to dump")
	jsonFile = flag.String("json_file", "../resource/mlcs.json", "input json file")
)

// Dump for sth.
func Dump(content interface{}) {
	if !*dump {
		return
	}
	b, _ := json.Marshal(content)
	data := []byte(fmt.Sprintf("%v", string(b)))
	ioutil.WriteFile(*jsonFile, data, 0644)
}
