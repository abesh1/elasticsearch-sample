package mecab

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var dicDir string

func init() {
	out, err := exec.Command("mecab-config", "--dicdir").Output()
	if err != nil {
		log.Fatalln(err)
	}
	dicDir = strings.TrimRight(string(out), "\n")
}

func YomiOption(dicName string) string {
	return fmt.Sprintf("-Oyomi -d%s/%s", dicDir, dicName)
}
