package mailer

import (
	"fmt"

	"github.com/mattbaird/gochimp"
)

func SendMail() {
	apiKey := "OMOxFhoklMo7hPjkmxUJxg"
	mandrillAPI, err := gochimp.NewMandrill(apiKey)

	if err != nil {
		fmt.Println("Error instantiating client")
	}

	templateName := "simple"
	contentVar := gochimp.Var{Name: "jamName", Content: "testing from go"}
	content := []gochimp.Var{contentVar}

	//_, err = mandrillAPI.TemplateAdd(templateName, fmt.Sprintf("%s", contentVar.Content), true)
	//if err != nil {
	//fmt.Println("Error adding template")
	//return
	//}
	//defer mandrillAPI.TemplateDelete(templateName)

	renderedTemplate, err := mandrillAPI.TemplateRender(templateName, content, nil)

	if err != nil {
		fmt.Println("Error rendering template")
		return
	}

	recipients := []gochimp.Recipient{
		gochimp.Recipient{Email: "marlon@monroy.io"},
		gochimp.Recipient{Email: "david.j.strom@gmail.com"},
	}

	message := gochimp.Message{
		Html:      renderedTemplate,
		Subject:   "Texting from our mailer",
		FromEmail: "acounts@draglabs.com",
		FromName:  "Drag Labs",
		To:        recipients,
	}

	_, err = mandrillAPI.MessageSend(message, false)

	if err != nil {
		fmt.Println("Error sending message")
	}
}
