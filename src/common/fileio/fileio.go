package fileio

import (
	"bufio"
	"flag"
	"os"

	assert "github.com/arl/assertgo"
)

var (
	seqFile = flag.String("seq_file", "./data/3_8.txt", "qaq")
)

// ReadFile for sth.
func ReadFile() []string {
	filePath := *seqFile
	file, err := os.Open(filePath)
	assert.True(err == nil, "读文件失败...[", filePath, "]")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ctx []string
	for scanner.Scan() {
		line := scanner.Text()
		ctx = append(ctx, line)
	}
	return ctx
}
