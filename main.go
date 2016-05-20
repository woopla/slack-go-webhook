/*
   Copyright 2016 Ashwanth Kumar & Clement Moyroud

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package slack

import (
	"encoding/json"
	"strings"

	"github.com/parnurzeal/gorequest"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachment struct {
	Fallback   *string  `json:"fallback"`
	Color      *string  `json:"color"`
	PreText    *string  `json:"pretext"`
	AuthorName *string  `json:"author_name"`
	AuthorLink *string  `json:"author_link"`
	AuthorIcon *string  `json:"author_icon"`
	Title      *string  `json:"title"`
	TitleLink  *string  `json:"title_link"`
	Text       *string  `json:"text"`
	ImageUrl   *string  `json:"image_url"`
	Fields     []*Field `json:"fields"`
}

func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	return attachment
}

func Payload(text, username, imageOrIcon, channel string, attachments []Attachment) map[string]interface{} {
	payload := make(map[string]interface{})
	payload["parse"] = "full"
	if username != "" {
		payload["username"] = username
	}

	if strings.HasPrefix("http", imageOrIcon) {
		payload["icon_url"] = imageOrIcon
	} else if imageOrIcon != "" {
		payload["icon_emoji"] = imageOrIcon
	}

	if channel != "" {
		payload["channel"] = channel
	}

	if text != "" {
		payload["text"] = text
	}

	if len(attachments) > 0 {
		payload["attachments"] = attachments
	}

	return payload
}

func Send(webhookUrl string, proxy string, payload map[string]interface{}) []error {
	data, _ := json.Marshal(payload)
	request := gorequest.New().Proxy(proxy)
	_, _, err := request.
		Post(webhookUrl).
		Send(string(data)).
		End()

	if err != nil {
		return err
	} else {
		return nil
	}
}
