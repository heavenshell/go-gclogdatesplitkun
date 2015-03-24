package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/codegangsta/cli"
)

var logs = map[string][]string{}

const datetimeFormat = "2006-01-02T15:04:05.000+0900"

func parse(line string) {
	words := strings.Split(line, " ")
	t, err := time.Parse(datetimeFormat, strings.Trim(words[0], ":"))
	if err != nil {
		return
	}
	date := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
	logs[date] = append(logs[date], line)
}

func read(targetFileName string) {
	_, err := os.Stat(targetFileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var fp *os.File
	fp, err = os.Open(targetFileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		parse(string(line))
	}
}

func write(path string, fileName string, log []string) {
	fp, err := os.Create(fmt.Sprintf("%s/%s", path, fileName))
	if err != nil {
		return
	}
	defer fp.Close()
	w := bufio.NewWriterSize(fp, 4098)
	for i := range log {
		w.WriteString(fmt.Sprintf("%s\n", log[i]))
		w.Flush()
	}
}

func do(c *cli.Context) {
	if len(c.Args()) == 0 {
		log.Fatal("gc.log file path required.")
		os.Exit(0)
	}
	log.Println("Start.")
	path := c.Args()[0]
	targetFileName, _ := filepath.Abs(path)
	read(targetFileName)
	fileName := filepath.Base(targetFileName)

	dir := filepath.Dir(targetFileName)
	var wg sync.WaitGroup
	for k := range logs {
		wg.Add(1)
		go func(dir, fileName string, lines []string) {
			defer wg.Done()
			write(dir, fileName, lines)
			log.Println(fmt.Sprintf("Generate %s", fileName))
		}(dir, fmt.Sprintf("%s_%s", k, fileName), logs[k])
	}
	wg.Wait()
	log.Println("Done.")
}

func main() {
	app := cli.NewApp()
	app.Name = "gclogdatesplitchan"
	app.Usage = "Split gc.log by date."
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:   "read",
			Usage:  "Read gc.log",
			Action: do,
		},
	}
	app.Run(os.Args)
}
