package hatchbuckapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateTag requires a contactID or email
func (api HatchbuckClient) CreateTag(target string, tag []Tag) error {
	endpoint := fmt.Sprintf("%v/contact/%v/tags?api_key=%v", api.baseURL, target, api.key)
	payload, _ := json.Marshal(tag)
	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	if res.StatusCode != 201 {
		body, _ := ioutil.ReadAll(res.Body)
		return errors.New(string(body))
	}
	return nil
}

// DeleteTag requires a contactID or email
func (api HatchbuckClient) DeleteTag(target string, tags []Tag) error {
	client := &http.Client{}
	endpoint := fmt.Sprintf("%v/contact/%v/tags?api_key=%v", api.baseURL, target, api.key)
	payload, _ := json.Marshal(tags)
	req, err := http.NewRequest("DELETE", endpoint, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 201 {
		body, _ := ioutil.ReadAll(res.Body)
		return errors.New(string(body))
	}
	return nil
}