package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"encoding/base64"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/joho/godotenv"
)

func sendEmail(){
	//API
	err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	api_key := os.Getenv("SENDGRID_API_KEY")


	m := mail.NewV3Mail()

	
	from := mail.NewEmail("Me", "jayson.mulwa@gmail.com")
	content := mail.NewContent("text/html", "<p>Sending different attachments.</p>")
	to := mail.NewEmail("Receiver", "jayson.mulwa@gmail.com")

	m.SetFrom(from)
  	m.AddContent(content)

	// create new *Personalization
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.Subject = "Attachments with SendGrid"

	// add `personalization` to `m`
	m.AddPersonalizations(personalization)

	 // read/attach .pdf file
	 a_pdf := mail.NewAttachment()
	 dat, err := ioutil.ReadFile("p.pdf")
	 if err != nil {
	   fmt.Println(err)
	 }
	 encoded := base64.StdEncoding.EncodeToString([]byte(dat))
	 a_pdf.SetContent(encoded)
	 a_pdf.SetType("application/pdf")
	 a_pdf.SetFilename("payslip.pdf")
	 a_pdf.SetDisposition("attachment")
	 a_pdf.SetContentID("Test Attachment")

	 // add `a_pdf``m`
	 m.AddAttachment(a_pdf)
  
  
	 request := sendgrid.GetRequest(api_key, "/v3/mail/send", "https://api.sendgrid.com")
	 request.Method = "POST"
	 request.Body = mail.GetRequestBody(m)
	 response, err := sendgrid.API(request)
	 if err != nil {
	   log.Println(err)
	 } else {
	   fmt.Println(response.StatusCode)
	   fmt.Println(response.Body)
	   fmt.Println(response.Headers)
	 }



}


