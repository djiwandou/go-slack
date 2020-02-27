# go-slack

A simple, flexible Golang wrapper around the [Slack webhook API](https://api.slack.com). Makes it easy to send notifications to Slack from your application.

## Features
* Send plain text and message with attachments

## Installation
```sh
$ go get -u github.com/fadlymuhammadf/go-slack
```

## Usage
Basic send text message
```go
// You need get your slack WebhookURL from https://api.slack.com/incoming-webhooks
	webhookURL := "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
  
	attachments := []Attachment{
    		{
    			Text: "test text",
    			Color: "good",
    			Fields: []Field{
    				{
    					Title: "fieldTitle",
    					Value: "Low",
    					Short: true,
    				},
    				{
    					Title: "fieldTitle2",
    					Value: "Low",
    					Short: true,
    				},
    				{
    					Title: "fieldTitle3",
    					Value: "Low",
    					Short: true,
    				},
    				{
    					Title: "fieldTitle4",
    					Value: "Low",
    					Short: true,
    				},
    			},
    			Ts: 123456789,
    		},
    	}

	err := slack.Init().
        UserName("User Name").
        Level("success").
        Channel("#channel").
        Content("test content").
        Icon(":sunglasses:").
        WebHookURL(webhookURL).
        Attachments(attachments).
        SendMessage()
```

Level Option
```level
good, danger, warning
