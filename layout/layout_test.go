package layout

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	tele "gopkg.in/tucnak/telebot.v3"
)

func TestLayout(t *testing.T) {
	os.Setenv("TOKEN", "TEST")

	lt, err := New("example.yml")
	if err != nil {
		t.Fatal(err)
	}

	pref := lt.Settings()
	assert.Equal(t, "TEST", pref.Token)
	assert.Equal(t, "html", pref.ParseMode)
	assert.Equal(t, &tele.LongPoller{}, pref.Poller)

	assert.Equal(t, "string", lt.Get("str"))
	assert.Equal(t, 123, lt.Int("num"))
	assert.Equal(t, int64(123), lt.Int64("num"))
	assert.Equal(t, float64(123), lt.Float("num"))
	assert.Equal(t, 10*time.Minute, lt.Duration("dur"))

	assert.Equal(t, &tele.ReplyMarkup{
		ReplyKeyboard:       [][]tele.ReplyButton{{{Text: "Send a contact", Contact: true}}},
		ResizeReplyKeyboard: true,
		OneTimeKeyboard:     true,
	}, lt.Markup("reply"))

	assert.Equal(t, &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{{{Unique: "inline"}}},
	}, lt.Markup("inline"))

	assert.Equal(t, &tele.ReplyMarkup{
		ReplyKeyboard: [][]tele.ReplyButton{
			{{Text: "Help"}},
			{{Text: "Settings"}},
		},
		ResizeReplyKeyboard: true,
	}, lt.Markup("embedded"))

	assert.Equal(t, &tele.ReplyMarkup{
		InlineKeyboard: [][]tele.InlineButton{{{Unique: "anchored"}}},
	}, lt.Markup("anchored"))
}
