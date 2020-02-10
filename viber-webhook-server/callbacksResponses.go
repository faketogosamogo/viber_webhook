package viber_webhook_server

import "fmt"

type GeneralResponse struct {
	Event     string `json:"event"`
	Timestamp int64  `json:"timestamp"`
}

func (g GeneralResponse) PrintCallback() {
	fmt.Println(g.Event)
	fmt.Println("timestamp: ", g.Timestamp)
}

type Location struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}
type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Country    string `json:"country"`
	Language   string `json:"language"`
	ApiVersion int    `json:"api_version"`
}
type Contact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type Callback interface {
	PrintCallback()
}

func printRazdelitel() {
	fmt.Println("*************************************")
}

type CallbackWebhook struct {
	GeneralResponse
	ChatHostname string `json:"chat_hostname"`
	MessageToken int64  `json:"message_token"`
}

func (w CallbackWebhook) PrintCallback() {
	w.GeneralResponse.PrintCallback()
	fmt.Println("Chat Hostname: ", w.ChatHostname)
	fmt.Println("Message token: ", w.MessageToken)
	printRazdelitel()
}

type CallbackSubscribed struct {
	GeneralResponse
	User
	MessageToken int64 `json:"message_token"`
}

func (w CallbackSubscribed) PrintCallback() {
	w.GeneralResponse.PrintCallback()
	fmt.Println("User ID: ", w.User.Id)
	fmt.Println("User name: ", w.User.Name)
	fmt.Println("MessageToken: ", w.MessageToken)
	printRazdelitel()
}

type CallbackUnsubscribed struct {
	GeneralResponse
	UserId       string `json:"user_id"`
	MessageToken int64  `json:"message_token"`
}

func (c CallbackUnsubscribed) PrintCallback() {
	c.GeneralResponse.PrintCallback()
	fmt.Println("User ID: ", c.UserId)
	fmt.Println("Message token: ", c.MessageToken)
	printRazdelitel()
}

type CallbackConversationStarted struct {
	GeneralResponse
	MessageToken string `json:"message_token"`
	Type         string `json:"type"`
	Context      string `json:"context"`
	User
	Subscribed bool `json:"subscribed"`
}

func (c CallbackConversationStarted) PrintCallback() {
	c.GeneralResponse.PrintCallback()
	fmt.Println("Message token: ", c.MessageToken)
	fmt.Println("Type: ", c.MessageToken)
	fmt.Println("Context: ", c.Context)
	fmt.Println("User ID: ", c.User.Id)
	fmt.Println("User name: ", c.User.Name)
	fmt.Println("Subscribed: ", c.Subscribed)
	printRazdelitel()
}

type CallbackMessageReceipts struct {
	GeneralResponse
	MessageToken int64  `json:"message_token"`
	UserId       string `json:"user_id"`
}

func (c CallbackMessageReceipts) PrintCallback() {
	c.GeneralResponse.PrintCallback()
	fmt.Println("Message token: ", c.MessageToken)
	fmt.Println("User id: ", c.UserId)
	printRazdelitel()
}

type CallbackFailedCallback struct {
	GeneralResponse
	MessageToken string `json:"message_token"`
	UserId       string `json:"user_id"`
	Desc         string `json:"desc"`
}

func (c CallbackFailedCallback) PrintCallback() {
	c.GeneralResponse.PrintCallback()
	fmt.Println("Message token: ", c.MessageToken)
	fmt.Println("User id: ", c.UserId)
	fmt.Println("Desc: ", c.Desc)
	printRazdelitel()
}

type CallbackReceiveMessage struct {
	GeneralResponse
	MessageToken    int64 `json:"message_token"`
	User            `json:"sender"`
	CallbackMessage `json:"message"`
}

func (c CallbackReceiveMessage) PrintCallback() {
	c.GeneralResponse.PrintCallback()
	fmt.Println("Message token: ", c.MessageToken)
	fmt.Println("User ID", c.User.Id)
	fmt.Println("User name", c.User.Name)
	fmt.Println("Type: ", c.CallbackMessage.Type)
	switch c.CallbackMessage.Type {
	case "text":
		fmt.Println("Text: ", c.CallbackMessage.Text)
	case "picture":
		fmt.Println("Text: ", c.CallbackMessage.Text)
		fmt.Println("Media: ", c.CallbackMessage.Media)
	case "video", "file":
		fmt.Println("Media: ", c.CallbackMessage.Media)
	case "location":
		fmt.Println("Lat: ", c.CallbackMessage.Lat)
		fmt.Println("Lon: ", c.CallbackMessage.Lon)
	case "contact":
		fmt.Println("Name: ", c.Contact.Name)
		fmt.Println("Number: ", c.Contact.PhoneNumber)
	case "sticker":
		fmt.Println("Sticker ID", c.StickerID)
	case "url":
		fmt.Println("Media: ", c.Media)
	}

	printRazdelitel()
}

type CallbackMessage struct {
	Type      string `json:"type"`
	Text      string `json:"text"`
	Location  `json:"location"`
	Media     string `json:"media"`
	Thumbnail string `json:"thumbnail"`
	FileName  string `json:"file_name"`
	FileSize  int64 `json:"size"`
	StickerID int32  `json:"sticker_id"`
	Contact   `json:"contact"`
}
