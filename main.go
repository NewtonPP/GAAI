package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func TabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}

	return ret
}

func CalculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}

	return w
}

func NormalizeStringsLength(lines []string, maxWidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxWidth-utf8.RuneCountInString((l)))
		ret = append(ret, s)
	}
	return ret
}

func BuildBalloon(lines []string, maxWidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}
	top := " " + strings.Repeat("_", maxWidth+2)
	bottom := " " + strings.Repeat("-", maxWidth+2)
	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], lines[i], borders[4])
			ret = append(ret, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[2], lines[i], borders[3])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}
func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("You can only use this program for inputs with pipes")
		fmt.Println(("Usage: echo | go run main.go"))
	}

	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		lines = append(lines, string(line))

	}

	var cow = `         \  ^__^
          \ (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
            `

	lines = TabsToSpaces(lines)
	maxWidth := CalculateMaxWidth(lines)
	messages := NormalizeStringsLength(lines, maxWidth)
	balloon := BuildBalloon(messages, maxWidth)
	fmt.Println(balloon)
	fmt.Println(cow)
	fmt.Println()
}
