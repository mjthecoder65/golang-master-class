package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

func (user User) Print() {
	fmt.Println("ID: ", user.ID)
	fmt.Println("Name: ", user.Name)
	fmt.Println("Username: ", user.Username)
	fmt.Println("Email: ", user.Email)
	fmt.Println("Phone: ", user.Phone)
	fmt.Println("Website: ", user.Website)
	fmt.Println("Street: ", user.Address.Street)
	fmt.Println("Suite: ", user.Address.Suite)
	fmt.Println("City: ", user.Address.City)
	fmt.Println("Zipcode: ", user.Address.Zipcode)
	fmt.Println("Company Name: ", user.Company.Name)
	fmt.Println("Catch Phrase: ", user.Company.CatchPhrase)
	fmt.Println("Bs: ", user.Company.Bs)
	fmt.Println()
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (post Post) PublishPost(userId int, title string, body string) Post {
	URL := "https://jsonplaceholder.typicode.com/posts"
	request, err := http.NewRequest("POST", URL, nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer {token}")

	payload := Post{
		UserID: userId,
		Title:  title,
		Body:   body,
	}

	postDataBytes, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	request.Body = io.NopCloser(bytes.NewReader(postDataBytes))

	// Send post request
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer response.Body.Close() // Close response body when function ends

	var newPost Post
	err = json.NewDecoder(response.Body).Decode(&newPost)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return newPost
}

// Decode response
func Decode(response *http.Response, users *[]User) error {
	err := json.NewDecoder(response.Body).Decode(&users)
	if err != nil {
		return err
	}
	return nil
}

func HTTPRequests() {
	const URL string = "https://jsonplaceholder.typicode.com/users"

	// Get request
	response, err := http.Get(URL)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer response.Body.Close() // Close response body when function ends

	// Decode response
	var users []User
	err = Decode(response, &users)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Print users
	for _, user := range users {
		user.Print()
	}
}
