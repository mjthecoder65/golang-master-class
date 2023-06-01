package main

import "fmt"

// Methods are functions that are associated with a type.
// Methods are functions that have a receiver.

// Syntax: func (receiverName receiverType) methodName() returnType {}

// Example:
type EmailClient struct {
	sender    string
	vendorURL string
}

// SendEmail is a method of EmailClient type.
func (emailClient EmailClient) SendEmail(email string) bool {
	fmt.Println("Sending email to", email)
	// Add your logic here.
	return true
}

func testingEmailClient() {
	emailClient := EmailClient{sender: "michael.ngowi@gmail.com", vendorURL: "https://www.gmail.com"}
	response := emailClient.SendEmail("jerrylusato@gmail.com")

	if response {
		fmt.Println("Email was send successfully")
	} else {
		fmt.Println("Failed to send email")
	}
}
