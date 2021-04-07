package main

import (
	"flag"
	"log"

	asana "github.com/yossy/asana-go"
)

var (
	pat   string
	pr    string
	notes string
)

func init() {
	flag.StringVar(&pat, "pat", "", "your asana personal access token")
	flag.StringVar(&pr, "pr", "", "URL of the PullRequest to write to the task.")
	flag.StringVar(&notes, "notes", "", "Sets the contents to be written to the task.")
	flag.Parse()
}

func main() {
	if pat == "" {
		log.Fatal("access tokenを設定して下さい。")
	}
	c := asana.NewClient(pat)
	// TODO: notesからTaskのIDを抽出する
	taskid := notes
	task := asana.FetchTask(c, taskid)
	asana.UpdateTask(c, task, pr)
}
