package asana

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	asana "bitbucket.org/mikehouston/asana-go"
)

func getTaskURL(c *asana.Client, taskid string) string {
	path := "/tasks/" + taskid
	return c.BaseURL.String() + path
}

func sendRequest(c *asana.Client, url string) *http.Response {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.HTTPClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func parseResponse(res *http.Response) *asana.Response {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the response
	value := &asana.Response{}
	if err := json.Unmarshal(body, value); err != nil {
		value.Errors = []*asana.Error{{
			StatusCode: res.StatusCode,
			Type:       "unknown",
			Message:    http.StatusText(res.StatusCode),
		}}
	}
	if value.Data == nil {
		log.Fatal("Taskを取得できませんでした")
	}
	return value
}

func FetchTask(c *asana.Client, taskid string) *asana.Task {
	url := getTaskURL(c, taskid)
	res := sendRequest(c, url)

	value := parseResponse(res)

	task := &asana.Task{}
	if err := json.Unmarshal(value.Data, task); err != nil {
		log.Fatal(err)
	}
	return task
}

func UpdateTask(c *asana.Client, task *asana.Task, pr string) {
	// TODO: PRのURLとNotesの内容をaggregateして書き込む
	// NOTE: もとの説明が削除されて以下に置換される
	task.TaskBase.Notes = pr
	updatereq := &asana.UpdateTaskRequest{
		TaskBase: task.TaskBase,
	}

	task.Update(c, updatereq)
}
