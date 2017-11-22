package mailer

import (
	"compressor/models"
	"fmt"

	"github.com/mattbaird/gochimp"
)

const (
	jamName    = "jam_name"
	jamNotes   = "jam_notes"
	jamCreator = "jam_creator"
	s3URL      = "s3_url"

	template = "simple"
	apiKey   = "OMOxFhoklMo7hPjkmxUJxg"
)

//SendMail send a email to the user
func SendMail(jam models.Jam, s3url string) {

	mandrillAPI, err := gochimp.NewMandrill(apiKey)

	if err != nil {
		fmt.Println("Error instantiating client")
	}

	templateName := template

	renderedTemplate, err := mandrillAPI.TemplateRender(templateName, composeContent(jam, s3url), nil)

	if err != nil {
		fmt.Println("Error rendering template")
		return
	}

	recipients := extractRecipients(jam)

	message := composeMessage(recipients, renderedTemplate)

	_, err = mandrillAPI.MessageSend(message, false)

	if err != nil {
		fmt.Println("Error sending message")
	}
}

func extractRecipients(jam models.Jam) []gochimp.Recipient {

	var recepients []gochimp.Recipient
	for _, r := range jam.Collaborators {

		recepients = append(recepients, gochimp.Recipient{Email: r.Email})

	}
	recepients = append(recepients, gochimp.Recipient{Email: jam.Creator.FBEmail})
	if recepients != nil {
		return recepients
	}
	return nil
}
func composeContent(jam models.Jam, url string) []gochimp.Var {
	contentJamVar := gochimp.Var{Name: jamName, Content: jam.Name}
	contentCreatorVar := gochimp.Var{Name: jamCreator, Content: jam.Creator.FirstName}
	contentS3URLVar := gochimp.Var{Name: s3URL, Content: url}
	contentJamNotes := gochimp.Var{Name: jamNotes, Content: jam.Notes}
	content := []gochimp.Var{contentJamVar, contentCreatorVar, contentS3URLVar, contentJamNotes}
	return content
}
func composeMessage(recipients []gochimp.Recipient, template string) gochimp.Message {
	return gochimp.Message{
		Html:      template,
		Subject:   "Jam audio files from dSoundboy",
		FromEmail: "acounts@draglabs.com",
		FromName:  "Drag Labs, dSoundboy",
		To:        recipients,
	}
}
