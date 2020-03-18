package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./gokyoto [url]")
		os.Exit(0)
	}
	url := os.Args[1]
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	inputList := []string{}
	answerList := []string{}

	doc.Find("section").Each(func(i int, s *goquery.Selection) {
		head := s.Find("h3").First().Text()

		if strings.Contains(head, "入力例") {
			input := s.Find("pre").First().Text()
			inputList = append(inputList, input)
		}
		if strings.Contains(head, "出力例") {
			answer := s.Find("pre").First().Text()
			answerList = append(answerList, answer)
		}
	})

	for i := 0; i < len(inputList); i++ {
		fmt.Println("* sample", i+1)

		cmd := exec.Command("./a.out")
		stdin, _ := cmd.StdinPipe()
		io.WriteString(stdin, inputList[i])
		stdin.Close()

		output, _ := cmd.Output()
		outArr := strings.Fields(string(output))

		ansArr := strings.Fields(answerList[i])

		if reflect.DeepEqual(outArr, ansArr) {
			color.Green("Accepted")
		} else {
			fmt.Println("input")
			fmt.Println(inputList[i])
			fmt.Println("\noutput")
			fmt.Println(string(output))
			fmt.Println("\nanswer")
			fmt.Println(inputList[i])
		}

		fmt.Println()
	}

}
