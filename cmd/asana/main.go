package main

import (
	"flag"
	"log"

	asana "github.com/yossy/asana-go"
)

var (
	pat  string
	pr   string
	body string
)

func init() {
	flag.StringVar(&pat, "pat", "", "Your asana personal access token")
	flag.StringVar(&pr, "pr", "", "URL of the PullRequest to write to the task")
	// TODO: ここの仕様考える
	flag.StringVar(&body, "body", "", "Set the body to write the task")
	flag.Parse()
}

func main() {
	if pat == "" {
		log.Fatal("access tokenを設定して下さい。")
	}
	c := asana.NewClient(pat)
	// TODO: githubのDesciptionからTaskのIDを抽出する
	taskid := body
	task := asana.FetchTask(c, taskid)
	asana.UpdateTask(c, task, pr)
}
