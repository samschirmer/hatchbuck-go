package main

import (
	hb "hb_api/pkg/hatchbuckapi"
	"log"
	"os"
)

func main() {
	api := hb.Init(os.Getenv("HB_API_KEY"))

	var criteria hb.SearchCriteria
	email := hb.Email{
		Address: "whatever@yopmail.com",
		TypeID:  os.Getenv("EMAIL_WORK"),
	}
	criteria.Emails = append(criteria.Emails, email)

	contacts, err := api.SearchContact(criteria)
	if err != nil {
		log.Println(err)
	}

	var c hb.Contact
	log.Println("no errors")
	if len(contacts) > 0 {
		c = contacts[0]
		c.FirstName = "Alec"
		c.LastName = "Holland"
		c.Status.ID = os.Getenv("STATUS_PROSPECT")
		c.Emails = append(c.Emails, email)
		contact, err := api.UpdateContact(c)
		if err != nil {
			log.Println(err)
		}
		log.Println(contact)
	} else {
		var c hb.Contact
		c.FirstName = "Swamp"
		c.LastName = "Thing"
		c.Status.ID = os.Getenv("STATUS_PROSPECT")
		c.Emails = append(c.Emails, email)
		contact, err := api.CreateContact(c)
		if err != nil {
			log.Println(err)
		}
		log.Println(contact)
	}

	sr := hb.SalesRep{ID: os.Getenv("SALES_REP")}
	n := hb.Note{
		Subject:               "This is another note subject",
		Body:                  "This is another note body",
		CreatedDateTime:       "2019-10-02 12:00:00",
		SalesRep:              sr,
		CopyToCompany:         true,
		UpdateLastContactDate: true,
	}

	err = api.CreateNote(c.ID, n)
	if err != nil {
		log.Println(err)
	}

	var tags []hb.Tag
	tags = append(tags, hb.Tag{Name: "wonder if this works"})
	err = api.DeleteTag(c.ID, tags)
	if err != nil {
		log.Println(err)
	}
}
