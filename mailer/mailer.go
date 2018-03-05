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

// Each recipient is a collaborator in the jam
//
// Param: jam, a jam model object
//
// Returns: an array of gochimp Recipients, where each one is a collaborator from the jam
func extractRecipients(jam models.Jam) []gochimp.Recipient {
	// (Bug): in theory, this is supposed to send an email to each collaborator, because it goes through all of the collaborators
	// in the jam, but the issue is that the api doesn't initialize the collaborators, so only the jam creator gets the emails
	// This can be seen in the loop below.
	var recipients []gochimp.Recipient
	for _, user := range jam.Collaborators {

		recipients = append(recipients, gochimp.Recipient{Email: user.Email})

	}
	fmt.Println("email", jam.Creator.FBEmail)
	recipients = append(recipients, gochimp.Recipient{Email: jam.Creator.FBEmail})
	if recipients != nil {
		return recipients
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

// Composes the message, with the body. Uses the email fields of HTML body style, subject, from email, from name, and to email. Emails all collaborators.
//
// Params: recipients, an array of gochimp Recipient objects; template, a string HTML template of the message
//
// Returns: a gochimp Message object
func composeMessage(recipients []gochimp.Recipient, template string) gochimp.Message {
	return gochimp.Message {
		Html:      template,
		Subject:   "Jam audio files from dSoundboy",
		FromEmail: "acounts@draglabs.com",
		FromName:  "Drag Labs, dSoundboy",
		To:        recipients,
	}
}
