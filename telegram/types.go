package telegram

import (
	"encoding/json"
)

// Message Entity Types
const (
	MEMention     = "mention"
	MEHashtag     = "hashtag"
	MEBotCommand  = "bot_command"
	MEUrl         = "url"
	MEEmail       = "email"
	MEBold        = "bold"
	MEItalic      = "italic"
	MECode        = "code"
	MEPre         = "pre"
	METextLink    = "text_link"
	MeTextMention = "text_mention"
)

// Reponse is telegram api response
type Response struct {
	Ok               bool            `json:"ok"`
	Result           json.RawMessage `json:"result,omitempty"`
	ErrorCode        int             `json:"error_code,omitempty"`
	ErrorDescription string          `json:"description,omitempty"`
}

// User https://core.telegram.org/bots/api#user
type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot,omitempty"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// Chat https://core.telegram.org/bots/api#chat
type Chat struct {
	ID                          int64       `json:"id"`
	Type                        string      `json:"type"`
	Title                       string      `json:"title,omitempty"`
	Username                    string      `json:"username,omitempty"`
	FirstName                   string      `json:"first_name,omitempty"`
	LastName                    string      `json:"last_name,omitempty"`
	AllMembersAreAdministrators bool        `json:"all_members_are_administrators,omitempty"`
	Photo                       interface{} `json:"photo,omitempty"`
	Description                 string      `json:"description,omitempty"`
	InviteLink                  string      `json:"invite_link,omitempty"`
	PinnedMessage               interface{} `json:"pinned_message,omitempty"`
}

// Message https://core.telegram.org/bots/api#message
type Message struct {
	MessageID            int64            `json:"message_id"`
	From                 *User            `json:"from,omitempty"`
	Date                 uint64           `json:"date"`
	Chat                 *Chat            `json:"chat"`
	ForwardFrom          *User            `json:"forward_from,omitempty"`
	ForwardFromChat      *Chat            `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID uint64           `json:"forward_from_message_id,omitempty"`
	ForwardDate          uint64           `json:"forward_date,omitempty"`
	EditDate             uint64           `json:"edit_date,omitempty"`
	AuthorSignature      string           `json:"author_signature,omitempty"`
	Text                 string           `json:"text,omitempty"`
	Entities             []*MessageEntity `json:"enitites,omitempty"`
}

// MessageEntity https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url,omitempty"`
	User   *User  `json:"user,omitempty"`
}

// Update https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID          int          `json:"update_id"`
	Message           *Message     `json:"message,omitempty"`
	EditedMessage     *Message     `json:"edited_message,omitempty"`
	ChannelPost       *Message     `json:"channel_post,omitempty"`
	EditedChannelPost *Message     `json:"edited_channel_post,omitempty"`
	InlineQuery       *InlineQuery `json:"inline_query,omitempty"`
}

// InlineQuery https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query,omitempty"`
	Offset   string    `json:"offset,omitempty"`
}

// Location https://core.telegram.org/bots/api#location
type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type ReplyMarkup struct {
	InlineKeyboard *[][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

// InlineKeyboardButton https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string `json:"text"`
	URL                          string `json:"url,omitempty"`
	CallbackData                 string `json:"callback_data,omitempty"`
	SwitchInlineQuery            string `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
}
