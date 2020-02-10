package viber_webhook_server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	Token = ""
	Host  = ""
)

func StartServer() {

	//	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("./pictures"))))
	//http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	//	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("./videos"))))
	http.HandleFunc("/", GetCallbacks)
	http.ListenAndServe("localhost:2517", nil)
}

func GetCallbacks(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	resp := GeneralResponse{}
	err = json.Unmarshal(body, &resp)
	var response Callback
	switch resp.Event {
	case "message":
		response = new(CallbackReceiveMessage)
	case "subscribed":
		response = new(CallbackSubscribed)
	case "unsubscribed":
		response = new(CallbackUnsubscribed)
	case "delivered", "seen":
		response = new(CallbackMessageReceipts)
	case "failed":
		response = new(CallbackFailedCallback)
	case "webhook":
		response = new(CallbackWebhook)
	default:
		fmt.Println(string(body))
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	response.PrintCallback()
	if resp.Event == "message" {
		message := CallbackReceiveMessage{}
		err = json.Unmarshal(body, &message)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = SendResultMessage(message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
func SaveMedia(url, typeMedia, fileName string) (string, int64, error) {
	var path string
	resp, err := http.Get(url)
	if err != nil {
		return path, 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return path, 0, err
	}

	switch typeMedia {
	case "picture":
		path = "pictures/"
	case "video":
		path = "videos/"
	case "file":
		path = "files/"
	default:
		err = errors.New("Не поддерживаемый тип сообщения")
		return path, 0, err
	}

	file, err := os.Create("./" + path + fileName)
	defer file.Close()
	if err != nil {
		return path, 0, err
	}

	if err != nil {
		return path, 0, err
	}

	_, err = file.Write(body)
	if err != nil {
		return path, 0, err
	}
	stat, err := file.Stat()
	return Host + path + fileName, stat.Size(), nil
}
func SendResultMessage(message CallbackReceiveMessage) error {

	switch message.Type {

	case "picture":
		err := SendPictureMessage(message.Id, message.Text, message.Media, Token)
		if err != nil {
			return err
		}

	case "video":
		err := SendVideoMessage(message.Id, message.Media, message.Text, Token, message.FileSize)
		if err != nil {
			return err
		}
	case "file":

		err := SendFileMessage(message.Id, message.Media, message.FileName, Token, message.FileSize)
		if err != nil {
			return err
		}
	case "sticker":
		err := SendStickerMessage(message.Id, message.StickerID, Token)
		if err != nil {
			return err
		}
	case "url":
		err := SendURLMessage(message.Id, message.Media, Token)
		if err != nil {
			return err
		}
	case "location":
		err := SendLocationMessage(message.Id, message.Lat, message.Lon, Token)
		if err != nil {
			return err
		}
	case "contact":
		err := SendContactMessage(message.Id, message.Contact.Name, message.Contact.PhoneNumber, Token)
		if err != nil {
			return err
		}
	case "text":
		err := SendTextMessage(message.Id, message.Text, Token)
		if err != nil {
			return err
		}
	default:
		err := errors.New("Не поддерживаемы тип сообщения")
		return err
	}

	return nil
}
