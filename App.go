package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	initEnv()
	readFile("stackTrace.txt")
}

func initEnv() {
	executePath, _ := exec.LookPath(os.Args[0])
	dir, _ := path.Split(executePath)
	os.Chdir(dir)
}

func readFile(name string) {

	file, e := os.Open(name)

	if e != nil {
		fmt.Println(e)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	create, i := os.Create("newStack.txt")

	if i != nil {
		fmt.Println(i)
		return
	}

	create.WriteString("java.lang.Exception" + "\n")

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			return
		}

		javaStackTraceLine := string(line)

		split := strings.Split(javaStackTraceLine, " ")

		if len(split) != 3 {
			return
		}

		methodInfo := split[0]

		if strings.Contains(methodInfo, "<") ||
			strings.Contains(methodInfo, ">") {
			fmt.Println("can not parse this line,because it contain <> char:", javaStackTraceLine)
			continue
		}

		methodNameWithLineNumber := strings.Split(strings.TrimRight(methodInfo, ","), ":")
		methodName := methodNameWithLineNumber[0]
		lineNumber := methodNameWithLineNumber[1]

		className := split[1]
		qualified := split[2]

		stackTrace := "at " + strings.Trim(qualified, "()") + "." + methodName + "(" + className + ":" + lineNumber + ")"

		create.WriteString(stackTrace + "\n")
	}

}
