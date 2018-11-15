package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	readFile("stackTrace.txt")

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

		className := split[1]
		qualified := split[2]

		stackTrace := "at " + strings.Trim(qualified, "()") + "." + className + "(" + strings.TrimRight(methodInfo, ",") + ")"

		create.WriteString(stackTrace + "\n")
	}

}
