// Git pull tool for auto pull all the list repositories in local machine
// All list repositories should in .target-git-pull
// The script will help to create a config file under user home directory if the config file not exists

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const CONFIGFILE = ".target-git-pull"

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func ExecCommand(dir string, command string, args []string) ([]string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	return lines, nil
}

func main() {

	var issueFilePath []string
	var pullFailFilePath []string
	var filesPath []string

	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Println("Internal Error.")
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(home + "/" + CONFIGFILE); err != nil {
		os.Create(home + "/" + CONFIGFILE)
		fmt.Println("No config file, create " + CONFIGFILE + " in home directory. Please update config file.")
	}

	file, err := os.Open(home + "/" + CONFIGFILE)
	CheckErr(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCounter := 0
	for scanner.Scan() {

		if _, err := os.Stat(scanner.Text()); err != nil {
			// collect filepath that cannot open
			issueFilePath = append(issueFilePath, scanner.Text())
		} else {
			lineCounter++
			filesPath = append(filesPath, scanner.Text())
		}
	}

	if len(issueFilePath) > 0 {
		fmt.Println("Find directory not exists")
		for _, path := range issueFilePath {
			fmt.Println("- " + path)
		}
	}

	if lineCounter < 1 {
		fmt.Println("No directory path list in config file")
		return
	}

	fmt.Printf("Find %d directory need to git pull action \n", lineCounter)

	for _, path := range filesPath {
		result, err := ExecCommand(path, "git", []string{"pull"})
		if err != nil {
			pullFailFilePath = append(pullFailFilePath, path)
		} else {
			fmt.Printf("Success Pull %s : %s \n", path, result)
		}

	}

	if len(pullFailFilePath) > 0 {
		fmt.Println("============================")
		fmt.Println("List fail pull directory: ")
		for _, path := range pullFailFilePath {
			fmt.Printf("- %s \n", path)
		}
	}

}
