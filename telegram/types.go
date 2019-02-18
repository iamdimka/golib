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

// Response is telegram api response
type Response struct {
	Ok               bool            `json:"ok"`
	Result           json.RawMessage `json:"result,omitempty"`
	ErrorCode        int             `json:"error_code,omitempty"`
	ErrorDescription string          `json:"description,omitempty"`
}

func (e *Response) Error() string {
	return e.ErrorDescription
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
	MessageID             int64            `json:"message_id"`
	From                  *User            `json:"from,omitempty"`
	Date                  uint64           `json:"date"`
	Chat                  *Chat            `json:"chat"`
	ForwardFrom           *User            `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat            `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  uint64           `json:"forward_from_message_id,omitempty"`
	ForwardSignature      string           `json:"forward_signature,omitempty"`
	ForwardDate           uint64           `json:"forward_date,omitempty"`
	ReplyToMessage        *Message         `json:"reply_to_message,omitempty"`
	EditDate              uint64           `json:"edit_date,omitempty"`
	MediaGroupID          string           `json:"media_group_id,omitempty"`
	AuthorSignature       string           `json:"author_signature,omitempty"`
	Text                  string           `json:"text,omitempty"`
	Entities              []*MessageEntity `json:"enitites,omitempty"`
	CaptionENtities       []*MessageEntity `json:"caption_entities,omitempty"`
	Audio                 interface{}      `json:"audio,omitempty"`
	Document              interface{}      `json:"document,omitempty"`
	Game                  interface{}      `json:"game,omitempty"`
	Photo                 []*PhotoSize     `json:"photo,omitempty"`
	Sticker               interface{}      `json:"sticker,omitempty"`
	Video                 interface{}      `json:"video,omitempty"`
	Voice                 interface{}      `json:"voice,omitempty"`
	VideoNote             interface{}      `json:"video_note,omitempty"`
	Caption               string           `json:"caption,omitempty"`
	Contact               interface{}      `json:"contact,omitempty"`
	Location              *Location        `json:"location,omitempty"`
	Venue                 interface{}      `json:"venue,omitempty"`
	NewChatMembers        []*User          `json:"new_chat_members,omitempty"`
	LeftChatMember        *User            `json:"left_chat_member,omitempty"`
	NewChatTitle          string           `json:"new_chat_title,omitempty"`
	NewChatPhoto          []*PhotoSize     `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool             `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool             `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool             `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool             `json:"channelchat_created,omitempty"`
	MigrateToChatID       int64            `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     int64            `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message         `json:"pinned_message,omitempty"`
	Invoice               interface{}      `json:"invoice,omitempty"`
	SuccessfulPayment     interface{}      `json:"successful_payment,omitempty"`
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
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
}

// CallbackQuery https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
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

// ReplyMarkup is part of https://core.telegram.org/bots/api#sendmessage
type ReplyMarkup struct {
	// InlineKeyboardMarkup https://core.telegram.org/bots/api#inlinekeyboardmarkup
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard,omitempty"`

	// ReplyKeyboardMarkup https://core.telegram.org/bots/api#replykeyboardmarkup
	Keyboard        [][]*KeyboardButton `json:"keyboard,omitempty"`
	ResizeKeyboard  bool                `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool                `json:"one_time_keyboard,omitempty"`
	Selective       bool                `json:"selective,omitempty"`

	// ReplyKeyboardRemove https://core.telegram.org/bots/api#replykeyboardremove
	RemoveKeyboard bool `json:"remove_keyboard,omitempty"`
	//Selective      bool `json:"selective,omitempty"`

	// ForceReply https://core.telegram.org/bots/api#forcereply
	ForceReply bool `json:"force_reply,omitempy"`
	//Selective      bool `json:"selective,omitempty"`
}

// InlineKeyboardButton https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string `json:"text"`
	URL                          string `json:"url,omitempty"`
	CallbackData                 string `json:"callback_data,omitempty"`
	SwitchInlineQuery            string `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
}

// KeyboardButton https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text            string `json:"text,omitempty"`
	RequestContact  bool   `json:"request_contact,omitempty"`
	RequestLocation bool   `json:"request_location,omitempty"`
}

// ChosenInlineResult https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id,omitempty"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query,omitempty"`
}

// PhotoSize https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FieldID  string `json:"field_id"`
	Width    uint   `json:"width"`
	Height   uint   `json:"height"`
	FileSize uint   `json:"file_size,omitempty"`
}
