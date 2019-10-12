package hatchbuckapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// SearchCriteria is needed to search for a contactl emails do not require a type or typeID
type SearchCriteria struct {
	ID        string  `json:"contactId,omitempty"`
	FirstName string  `json:"firstName,omitempty"`
	LastName  string  `json:"lastName,omitempty"`
	Emails    []Email `json:"emails,omitempty"`
}

// Contact represents a contact entity in the Hatchbuck API
type Contact struct {
	ID               string            `json:"contactId,omitempty"`
	FirstName        string            `json:"firstName,omitempty"`
	LastName         string            `json:"lastName,omitempty"`
	Title            string            `json:"title,omitempty"`
	Company          string            `json:"company,omitempty"`
	CreatedDt        string            `json:"createdDt,omitempty"`
	ContactURL       string            `json:"contactUrl,omitempty"`
	Emails           []Email           `json:"emails,omitempty"`
	Phones           []Phone           `json:"phones,omitempty"`
	Tags             []Tag             `json:"tags,omitempty"`
	Campaigns        []Campaign        `json:"campaigns,omitempty"`
	Status           Status            `json:"status,omitempty"`
	SalesRep         SalesRep          `json:"salesRep,omitempty"`
	Addresses        []Address         `json:"addresses,omitempty"`
	Timezone         string            `json:"timezone,omitempty"`
	SocialNetworks   []SocialNetwork   `json:"socialNetworks,omitempty"`
	InstantMessaging []InstantMessager `json:"instantMessaging,omitempty"`
	Websites         []Website         `json:"website,omitempty"`
	ReferredBy       string            `json:"referredBy,omitempty"`
	Subscribed       bool              `json:"subscribed,omitempty"`
	CustomFields     []CustomField     `json:"customFields,omitempty"`
}

// Email is an element of Contact; type or typeID is required when creating (but not searching)
type Email struct {
	ID      string `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
	Type    string `json:"type,omitempty"`
	TypeID  string `json:"typeId,omitempty"`
}

// Phone is an element of Contact; type or typeID is required
type Phone struct {
	ID      string `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
	Type    string `json:"type,omitempty"`
	TypeID  string `json:"typeId,omitempty"`
}

// Tag is an element of the tags array present on contact records
type Tag struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Score int    `json:"score,omitempty"`
}

// Campaign is an element of the campaigns array present on contact records
type Campaign struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Step int    `json:"step,omitempty"`
}

// Status is a required field on contacts
type Status struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// SalesRep will fallback to the owner of this API key if left blank
type SalesRep struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

// Address is an element of Contact; type or typeID is required
type Address struct {
	ID        string `json:"id,omitempty"`
	Street    string `json:"street,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	CountryID string `json:"countryId,omitempty"`
	Type      string `json:"type,omitempty"`
	TypeID    string `json:"typeId,omitempty"`
	IsPrimary bool   `json:"isPrimary,omitempty"`
}

// SocialNetwork is an element of the socialNetworks array present on contact records
type SocialNetwork struct {
	Address string `json:"address,omitempty"`
	Type    string `json:"type,omitempty"`
}

// InstantMessager is an element of the instantMessaging array present on contact records
type InstantMessager struct {
	Address string `json:"address,omitempty"`
	Type    string `json:"type,omitempty"`
}

// Website is an element of the website (sic) array present on contact records
type Website struct {
	ID         string `json:"id,omitempty"`
	WebsiteURL string `json:"websiteUrl,omitempty"`
}

// CustomField values are always returned as strings; only include type if the field doesn't exist and you are creating a new one
type CustomField struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"` // API sends everything back as strings
}

// SearchContact returns a list of contacts based on email address or first + last name
func (api HatchbuckClient) SearchContact(criteria SearchCriteria) ([]Contact, error) {
	var c []Contact
	endpoint := fmt.Sprintf("%v/contact/search?api_key=%v", api.baseURL, api.key)
	payload, _ := json.Marshal(criteria)
	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// CreateContact pushes a single contact to the Hatchbuck API
func (api HatchbuckClient) CreateContact(contact Contact) (Contact, error) {
	var c Contact
	endpoint := fmt.Sprintf("%v/contact?api_key=%v", api.baseURL, api.key)
	payload, _ := json.Marshal(contact)
	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
	log.Println(res.StatusCode)
	if err != nil {
		return c, err
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&c)
	if err != nil {
		return c, err
	}
	return c, nil
}

// UpdateContact updates a single contact to the Hatchbuck API
func (api HatchbuckClient) UpdateContact(contact Contact) (Contact, error) {
	var c Contact
	client := &http.Client{}
	endpoint := fmt.Sprintf("%v/contact?api_key=%v", api.baseURL, api.key)
	payload, _ := json.Marshal(contact)
	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(payload))
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	log.Println(res.StatusCode)
	if err != nil {
		return c, err
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&c)
	if err != nil {
		return c, err
	}
	return c, nil
}
