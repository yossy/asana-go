package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

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
	flag.StringVar(&body, "body", "", "Set the body to write the task")
	flag.Parse()
}

func run() error {
	if pat == "" {
		return errors.New("please set your asana personal access token.")
	}
	if pr == "" {
		return errors.New("please set pullrequest url")
	}

	c := asana.NewClient(pat)
	taskid, err := asana.PickUpTaskID(body)
	if taskid == "" {
		return nil
	}
	if err != nil {
		return err
	}

	task, err := asana.FetchTask(c, taskid)
	if err != nil {
		return err
	}
	if err := asana.UpdateTaskNotes(c, task, pr); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
