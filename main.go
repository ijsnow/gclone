package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print("Usage:\n  gclone <uri>\n")
		os.Exit(1)
		return
	}

	var host, path string

	uri := os.Args[1]
	gpath := os.ExpandEnv("$GOPATH")

	u, err := url.Parse(uri)
	if err == nil {
		// Handle https cloning
		host = u.Host
		path = u.Path
	} else if strings.Contains(uri, "@") {
		// Handle ssh cloning
		parts := strings.Split(uri, ":")
		path = parts[1]

		parts = strings.Split(parts[0], "@")
		host = parts[1]
	}

	path = path[:len(path)-len(".git")]
	path = filepath.Join(gpath, "src", host, path)

	clone(uri, path)
}

func clone(uri, path string) {
	cmd := exec.Command("git", "clone", uri, path)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("error cloning %s: %v", uri, path)
	}

	fmt.Printf("Cloned: cd %s", path)
}
