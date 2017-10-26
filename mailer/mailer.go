package mailer

import (
	"compressor/models"
	"fmt"

	"github.com/mattbaird/gochimp"
)

const (
	jamName  = "jamName"
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
	contentJamVar := gochimp.Var{Name: "jam_name", Content: jam.Name}
	contentCreatorVar := gochimp.Var{Name: "creator", Content: jam.Creator.Name}
	contentS3URLVar := gochimp.Var{Name: "s3_url", Content: s3url}
	content := []gochimp.Var{contentJamVar, contentCreatorVar, contentS3URLVar}

	renderedTemplate, err := mandrillAPI.TemplateRender(templateName, content, nil)

	if err != nil {
		fmt.Println("Error rendering template")
		return
	}

	recipients := extractRecipients(jam)

	message := gochimp.Message{
		Html:      renderedTemplate,
		Subject:   "Jam audio files from dSoundboy",
		FromEmail: "acounts@draglabs.com",
		FromName:  "Drag Labs, dSoundboy",
		To:        recipients,
	}

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
	if recepients != nil {
		return recepients
	}
	return nil
}
