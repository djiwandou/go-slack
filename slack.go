package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Slack struct {
	webHookURL  string
	level       string
	userName    string
	icon        string
	channel     string
	content     string
	attachments []Attachment
}

type Attachment struct {
	Title    string   `json:"title"`
	Color    string   `json:"color"`
	Text     string   `json:"text"`
	MrkdwnIn []string `json:"mrkdwn_in"`
	Fields   []Field  `json:"fields"`
	Ts       int      `json:"ts"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func Init() *Slack {
	return &Slack{}
}

func (s *Slack) WebHookURL(webhookURL string) *Slack {
	s.webHookURL = webhookURL
	return s
}

func (s *Slack) UserName(username string) *Slack {
	s.userName = username
	return s
}

func (s *Slack) Level(level string) *Slack {
	s.level = level
	return s
}

func (s *Slack) Icon(icon string) *Slack {
	s.icon = icon
	return s
}

func (s *Slack) Channel(channel string) *Slack {
	s.channel = channel
	return s
}

func (s *Slack) Content(content string) *Slack {
	s.content = content
	return s
}

func (s *Slack) Attachments(attachments []Attachment) *Slack {
	s.attachments = attachments
	return s
}

func (s *Slack) SendMessage() error {
	var err error

	message := map[string]interface{}{
		"username": s.userName,
		"text": s.content,
		"channel": s.channel,
		"icon_emoji": s.icon,
		"attachments": s.attachments,
	}
	requestBody, _ := json.Marshal(message)
	request, err := http.NewRequest("POST", s.webHookURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return errors.New("Success")
}

func main() {
	//_ := []Attachment{
	//	{
	//		Text: "testtesttext",
	//		Color: "good",
	//		Fields: []Field{
	//			{
	//				Title: "fieldTitle",
	//				Value: "Low",
	//				Short: true,
	//			},
	//			{
	//				Title: "fieldTitle2",
	//				Value: "Low",
	//				Short: true,
	//			},
	//			{
	//				Title: "fieldTitle3",
	//				Value: "Low",
	//				Short: true,
	//			},
	//			{
	//				Title: "fieldTitle4",
	//				Value: "Low",
	//				Short: true,
	//			},
	//		},
	//		Ts: 123456789,
	//	},
	//}

	_ = Init().
		UserName("Fikray Brew").
		Level("success").
		Channel("#transactionservice").
		Content("test emsasdassage brewwwwww").
		Icon(":sunglasses:").
		WebHookURL("https://hooks.slack.com/services/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").
		//Attachments(attachments).
		SendMessage()

}