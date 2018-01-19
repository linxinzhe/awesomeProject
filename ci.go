package main

import (
	"strings"
	"os/exec"
	"path/filepath"
	"runtime"
	"log"
	"fmt"
	"os"
	"flag"
)

func main() {
	if _, err := os.Stat(filepath.Join(".", "ci.go")); os.IsNotExist(err) {
		log.Fatal("this script must be run from the root of the repository")
	}

	if len(os.Args) < 2 {
		log.Fatal("need subcommand as first argument")
	}

	switch os.Args[1] {
	case "test":
		doTest(os.Args[2:])
	default:
		log.Fatal("unknown command ", os.Args[1])
	}

}
func doTest(cmdline []string) {
	flag.CommandLine.Parse(cmdline)

	packages := []string{"./..."}
	packages = ExpandPackagesNoVendor(packages)
	if flag.NArg() > 0 {
		packages = flag.Args()
	}

	gotest := GoTool("test", "-v")
	gotest.Args = append(gotest.Args, packages...)
	MustRun(gotest)
}

func GoTool(tool string, args ...string) *exec.Cmd {
	args = append([]string{tool}, args...)
	return exec.Command(filepath.Join(runtime.GOROOT(), "bin", "go"), args...)
}

func ExpandPackagesNoVendor(patterns []string) []string {
	expand := false
	for _, pkg := range patterns {
		if strings.Contains(pkg, "...") {
			expand = true
		}
	}
	if expand {
		cmd := GoTool("list", patterns...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("package listing failed: %v\n%s", err, string(out))
		}
		var packages []string
		for _, line := range strings.Split(string(out), "\n") {
			if !strings.Contains(line, "/vendor/") {
				packages = append(packages, strings.TrimSpace(line))
			}
		}
		return packages
	}
	return patterns
}

func MustRun(cmd *exec.Cmd) {
	fmt.Println(">>>", strings.Join(cmd.Args, " "))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
