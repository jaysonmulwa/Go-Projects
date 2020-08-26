package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/joho/godotenv"
)

func SimpleMail(){

	//API
	err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	api_key := os.Getenv("SENDGRID_API_KEY")

	//Email
	from := mail.NewEmail("Sender Name", "jayson.mulwa@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Receiver Name", "jayson.mulwa@gmail.com")
	plainTextContent := "Eureka"
	htmlContent := "<strong>Eureka!</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	
	client := sendgrid.NewSendClient(api_key)

	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}


}

func main() {

	sendEmail()//has attachment

	//SimpleMail()
}