package asana

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	asana "bitbucket.org/mikehouston/asana-go"
)

func UpdateTask() {
	// TODO: githubactionsから渡された値を参照する
	pat := ""
	taskid := ""
	pr := ""

	c := asana.NewClientWithAccessToken(pat)
	url := c.BaseURL.String() + "/tasks/" + taskid

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.HTTPClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
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

	// Decode the data field
	if value.Data == nil {
		log.Fatal("取得したTaskにdataがありません")
	}

	task := &asana.Task{}
	if err := json.Unmarshal(value.Data, task); err != nil {
		log.Fatal(err)
	}
	task.TaskBase.Notes = pr
	updatereq := &asana.UpdateTaskRequest{
		TaskBase: task.TaskBase,
	}

	task.Update(c, updatereq)
}
