package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func run() int {
	var kind string
	flag.StringVar(&kind, "k", "post", "")
	flag.Parse()

	if flag.NArg() <= 0 {
		fmt.Fprintln(os.Stderr, "{cmd} [opts] {title ...}")
		return 128
	}

	timestamp := time.Now().UTC()
	title := strings.Replace(strings.Join(flag.Args(), "-"), " ", "-", -1) + ".md"

	filename := filepath.Join("post", timestamp.Format("2006"), timestamp.Format("01"), timestamp.Format("02-")+title)

	fmt.Printf("hugo new --kind %s %s\n", kind, filename)
	cmd := exec.Command("hugo", "new", "--kind", kind, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(run())
}
