package flag

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

type flagFileSet map[string]struct{}

func (a *flagFileSet) Set(value string) error {
	if _, ok := (*a)[value]; !ok {
		(*a)[value] = struct{}{}
	}
	return nil
}

func (a *flagFileSet) String() string {
	var flagFileList []string
	for key := range *a {
		flagFileList = append(flagFileList, key)
	}
	return strings.Join(flagFileList, ",")
}

var (
	flagFiles flagFileSet = make(map[string]struct{})
)

func init() {
	flag.Var(&flagFiles, "flagfile", "")
}

func visitFlag(f *flag.Flag) {
	fmt.Printf("%s=%s\n", f.Name, f.Value)
}

// Parse 自带的flag库不支持传入文件，这个接口支持传入"-flagfile"方便传入配置文件
func Parse() error {
	flag.Parse()

	var validFlagLines []string
	for f := range flagFiles {
		flagContents, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Printf("err when read content from flagfile[%s], err:%s", f, err)
			return err
		}
		flagLines := strings.Split(string(flagContents), "\n")
		for _, line := range flagLines {
			if len([]rune(line)) != 0 && string([]rune(line)[0]) != "#" {
				validFlagLines = append(validFlagLines, line)
				//fmt.Println(line)
			}
		}
	}
	err := flag.CommandLine.Parse(validFlagLines)
	if err != nil {
		fmt.Printf("failed to parse flagfile[%s], err:%s", flagFiles.String(), err)
		return err
	}
	// 主动在命令行设置的参数具有更高的优先级
	flag.Parse()
	// flag.VisitAll(visitFlag)

	return nil
}
