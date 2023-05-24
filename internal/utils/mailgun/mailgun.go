package mailgun

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"strings"
)

type MailgunEventParsed struct {
	BodyPlain string   `json:"body-plain"`
	From      string   `json:"from"`
	Subject   string   `json:"subject"`
	To        []string `json:"to"`
}

type MailgunEventParser struct {
	Boundary string
}

func (m *MailgunEventParser) MailgunEventParser(body string) (*MailgunEventParsed, error) {
	reader := multipart.NewReader(strings.NewReader(body), m.Boundary)
	mailgunEventParsed := &MailgunEventParsed{To: []string{}}

	// Iterate over each part in the multipart form
	for {
		part, err := reader.NextPart()
		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				return mailgunEventParsed, nil
			}
			return mailgunEventParsed, err
		}

		partName := part.FormName()

		// Process the desired parts of the email
		switch strings.ToLower(partName) {
		case "subject":
			subjectBytes, _ := ioutil.ReadAll(part)
			log.Println("Subject:", string(subjectBytes))
			mailgunEventParsed.Subject = string(subjectBytes)
		case "from":
			fromBytes, _ := ioutil.ReadAll(part)
			log.Println("From:", string(fromBytes))
			mailgunEventParsed.From = string(fromBytes)
		case "to":
			toBytes, _ := ioutil.ReadAll(part)
			log.Println("To:", string(toBytes))
			mailgunEventParsed.To = append(mailgunEventParsed.To, string(toBytes))
		case "body-plain":
			bodyBytes, _ := ioutil.ReadAll(part)
			log.Println("Body:", string(bodyBytes))
			mailgunEventParsed.BodyPlain = string(bodyBytes)
		}

	}
}
