package hatchbuckapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Note is added to a contact by either email address or contactID
type Note struct {
	Subject               string   `json:"subject,omitempty"`
	Body                  string   `json:"body,omitempty"`
	CreatedDateTime       string   `json:"createdDateTime,omitempty"`
	SalesRep              SalesRep `json:"salesRep,omitempty"`
	CopyToCompany         bool     `json:"copyToCompany,omitempty"`
	UpdateLastContactDate bool     `json:"updateLastContactDate,omitempty"`
}

// CreateNote requires a contactID or email; salesRep will be author of the new note
func (api HatchbuckClient) CreateNote(target string, note Note) error {
	endpoint := fmt.Sprintf("%v/contact/%v/notes?api_key=%v", api.baseURL, target, api.key)
	payload, _ := json.Marshal(note)
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
