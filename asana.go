package asana

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	asana "bitbucket.org/mikehouston/asana-go"
)

func getTaskURL(c *asana.Client, taskid string) string {
	path := "/tasks/" + taskid
	return c.BaseURL.String() + path
}

func sendRequest(c *asana.Client, url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func parseResponse(res *http.Response) (*asana.Response, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Decode the response
	value := &asana.Response{}
	json.Unmarshal(body, value)

	switch res.StatusCode {
	case 200: // OK
	case 201: // Object created
	default:
		return nil, fmt.Errorf("statuscode %d: %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	if value.Data == nil {
		return nil, errors.New("not found: taskを取得できませんでした。")
	}

	return value, nil
}

func FetchTask(c *asana.Client, taskid string) (*asana.Task, error) {
	url := getTaskURL(c, taskid)
	res, err := sendRequest(c, url)
	if err != nil {
		return nil, err
	}

	value, err := parseResponse(res)
	if err != nil {
		return nil, err
	}

	task := &asana.Task{}
	if err := json.Unmarshal(value.Data, task); err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateTaskNotes(c *asana.Client, task *asana.Task, pr string) error {
	// NOTE: もとの説明が削除されて以下に置換される
	task.TaskBase.Notes = pr
	updatereq := &asana.UpdateTaskRequest{
		TaskBase: task.TaskBase,
	}

	if err := task.Update(c, updatereq); err != nil {
		return err
	}
	return nil
}
