package viber_webhook_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	send_message_URL = "https://chatapi.viber.com/pa/send_message"
	senderName       = "БОТ"
	senderAvatar     = ""
)

type Sender struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	//min_api_version && tracking_data
}

func SendMessage(i interface{}, url string, token string) ([]byte, error) {
	data, err := json.Marshal(i)
	if err != nil {
		return data, err
	}
	client := http.Client{}

	buff := bytes.Buffer{}
	buff.Write(data)
	request, err := http.NewRequest("POST", url, &buff)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-Viber-Auth-Token", token)
	request.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type GeneralMessage struct {
	Receiver string `json:"receiver"`
	Type     string `json:"type"`
	Sender   `json:"sender"`
}

type TextMessage struct {
	GeneralMessage
	Text string `json:"text"`
}

func SendTextMessage(receiverID, text, token string) error {
	message := TextMessage{GeneralMessage{receiverID, "text", Sender{senderName, senderAvatar}}, text}
	resp, err := SendMessage(message, send_message_URL, token)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

type PictureMessage struct {
	GeneralMessage
	Text  string `json:"text"`
	Media string `json:"media"`
}

func SendPictureMessage(receiverID, text, media, token string) error {
	message := PictureMessage{GeneralMessage{receiverID, "picture", Sender{senderName, senderAvatar}}, text, media}
	resp, err := SendMessage(message, send_message_URL, token)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

type VideoMessage struct {
	GeneralMessage
	Text  string `json:"text"`
	Size  int64  `json:"size"`
	Media string `json:"media"`
}

func SendVideoMessage(receiverID, media, text, token string, size int64) error {
	message := VideoMessage{GeneralMessage{receiverID, "video", Sender{senderName, senderAvatar}}, text, size, media}
	resp, err := SendMessage(message, send_message_URL, token)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

type FileMessage struct {
	GeneralMessage
	Size     int64  `json:"size"`
	Media    string `json:"media"`
	FileName string `json:"file_name"`
}

func SendFileMessage(receiverID, media, fileName, token string, size int64) error {
	message := FileMessage{GeneralMessage{receiverID, "file", Sender{senderName, senderAvatar}}, size, media, fileName}
	resp, err := SendMessage(message, send_message_URL, token)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

type ContactMessage struct {
	GeneralMessage
	Contact Contact `json:"contact"`
}

func SendContactMessage(receiverID, name, phoneNumber, token string) error {
	message := ContactMessage{GeneralMessage{receiverID, "contact", Sender{senderName, senderAvatar}}, Contact{name, phoneNumber}}

	resp, err := SendMessage(message, send_message_URL, token)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

type LocationMessage struct {
	GeneralMessage
	Location Location `json:"location"`
}

func SendLocationMessage(receiverID string, lat float32, lon float32, token string) error {
	message := LocationMessage{GeneralMessage{receiverID, "location", Sender{senderName, senderAvatar}}, Location{lat, lon}}
	resp, err := SendMessage(message, send_message_URL, token)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

type URLMessage struct {
	GeneralMessage
	Media string `json:"media"`
}

func SendURLMessage(receiverID, media, token string) error {
	message := URLMessage{GeneralMessage{receiverID, "url", Sender{senderName, senderAvatar}}, media}
	resp, err := SendMessage(message, send_message_URL, token)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

type StickerMessage struct {
	GeneralMessage
	StickerID int32 `json:"sticker_id"`
}

func SendStickerMessage(receiverID string, sticker_id int32, token string) error {
	message := StickerMessage{GeneralMessage{receiverID, "sticker", Sender{senderName, senderAvatar}}, sticker_id}
	resp, err := SendMessage(message, send_message_URL, token)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
