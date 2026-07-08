package aibot

import (
	"encoding/json"
	"testing"
)

func TestNewSendMarkdownMsgBody(t *testing.T) {
	body := NewSendMarkdownMsgBody("## hello")

	if body.MsgType != string(MessageTypeMarkdown) || body.Markdown.Content != "## hello" {
		t.Fatalf("body = %+v", body)
	}
	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	if string(raw) != `{"msgtype":"markdown","markdown":{"content":"## hello"},"chatid":""}` {
		t.Fatalf("json = %s", raw)
	}
}

func TestNewSendTemplateCardMsgBody(t *testing.T) {
	body := NewSendTemplateCardMsgBody(TemplateCard{
		CardType: string(TemplateCardTypeTextNotice),
		MainTitle: &TemplateCardMainTitle{
			Title: "通知",
			Desc:  "今晚收档确认",
		},
	})

	if body.MsgType != string(MessageTypeTemplateCard) ||
		body.TemplateCard.CardType != string(TemplateCardTypeTextNotice) ||
		body.TemplateCard.MainTitle.Title != "通知" {
		t.Fatalf("body = %+v", body)
	}
}

func TestNewSendMediaMsgBody(t *testing.T) {
	fileBody := NewSendMediaMsgBody(WeComMediaTypeFile, "media-file", nil)
	if fileBody.MsgType != WeComMediaTypeFile || fileBody.File == nil || fileBody.File.MediaID != "media-file" {
		t.Fatalf("fileBody = %+v", fileBody)
	}

	videoBody := NewSendMediaMsgBody(WeComMediaTypeVideo, "media-video", &VideoMediaContent{
		Title:       "视频标题",
		Description: "视频说明",
	})
	if videoBody.MsgType != WeComMediaTypeVideo ||
		videoBody.Video == nil ||
		videoBody.Video.MediaID != "media-video" ||
		videoBody.Video.Title != "视频标题" ||
		videoBody.Video.Description != "视频说明" {
		t.Fatalf("videoBody = %+v", videoBody)
	}
}
