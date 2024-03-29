package main

import "fmt"

// Methods are functions that are associated with a type.
// Methods are functions that have a receiver. Receive is a  struct type or interface.

// Syntax: func (receiverName receiverType) methodName() returnType {}

// Example:
type EmailClient struct {
	Sender    string
	VendorURL string
}

// SendEmail is a method of EmailClient type.
func (emailClient EmailClient) SendEmail(email string) bool {
	fmt.Println("Sending email to", email)
	// Add your logic here.
	return true
}

func TestingEmailClient() {
	emailClient := EmailClient{Sender: "michael.ngowi@gmail.com", VendorURL: "https://www.gmail.com"}
	response := emailClient.SendEmail("jerrylusato@gmail.com")

	if response {
		fmt.Println("Email was send successfully")
	} else {
		fmt.Println("Failed to send email")
	}
}
