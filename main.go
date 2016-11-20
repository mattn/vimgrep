package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	grepcmd := `silent vimgrep`
	for _, arg := range os.Args[1:] {
		grepcmd += ` ` + strings.Replace(arg, ` `, `\ `, -1)
	}
	errcmd := `if v:errmsg != '' | cquit! | endif`
	outcmd := `echo join(map(getqflist(), 'printf("%s:%d:%s", bufname(v:val.bufnr), v:val.lnum, v:val.text)'), "\n")`
	cmd := exec.Command("vim", "--not-a-term", "--cmd", grepcmd, "--cmd", errcmd, "--cmd", outcmd, "--cmd", "qall")
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], string(b))
		os.Exit(1)
	}
	fmt.Println(strings.Replace(string(b), "\r", "", -1))
}
