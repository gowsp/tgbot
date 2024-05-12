package tgbot

// https://core.telegram.org/bots/api#editmessagetext
//
// Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageText struct {
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Identifier of the message to edit
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId string `json:"inline_message_id,omitempty"`
    // Required: Yes
    //
    // New text of the message, 1-4096 characters after entities parsing
    Text string `json:"text,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the message text. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
    Entities []MessageEntity `json:"entities,omitempty"`
    // Required: Optional
    //
    // Link preview generation options for the message
    LinkPreviewOptions LinkPreviewOptions `json:"link_preview_options,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object for an inline keyboard.
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *EditMessageText) Method() string {
	return "editMessageText"
}

// https://core.telegram.org/bots/api#editmessagecaption
//
// Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageCaption struct {
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Identifier of the message to edit
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId string `json:"inline_message_id,omitempty"`
    // Required: Optional
    //
    // New caption of the message, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the message caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object for an inline keyboard.
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *EditMessageCaption) Method() string {
	return "editMessageCaption"
}

// https://core.telegram.org/bots/api#editmessagemedia
//
// Use this method to edit animation, audio, document, photo, or video messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageMedia struct {
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Identifier of the message to edit
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId string `json:"inline_message_id,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized object for a new media content of the message
    Media InputMedia `json:"media,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object for a new inline keyboard.
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *EditMessageMedia) Method() string {
	return "editMessageMedia"
}

// https://core.telegram.org/bots/api#editmessagelivelocation
//
// Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageLiveLocation struct {
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Identifier of the message to edit
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId string `json:"inline_message_id,omitempty"`
    // Required: Yes
    //
    // Latitude of new location
    Latitude float64 `json:"latitude,omitempty"`
    // Required: Yes
    //
    // Longitude of new location
    Longitude float64 `json:"longitude,omitempty"`
    // Required: Optional
    //
    // New period in seconds during which the location can be updated, starting from the message send date. If 0x7FFFFFFF is specified, then the location can be updated forever. Otherwise, the new value must not exceed the current live_period by more than a day, and the live location expiration date must remain within the next 90 days. If not specified, then live_period remains unchanged
    LivePeriod int `json:"live_period,omitempty"`
    // Required: Optional
    //
    // The radius of uncertainty for the location, measured in meters; 0-1500
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    // Required: Optional
    //
    // Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
    Heading int `json:"heading,omitempty"`
    // Required: Optional
    //
    // The maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
    ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object for a new inline keyboard.
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *EditMessageLiveLocation) Method() string {
	return "editMessageLiveLocation"
}

// https://core.telegram.org/bots/api#stopmessagelivelocation
//
// Use this method to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
type StopMessageLiveLocation struct {
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Identifier of the message with live location to stop
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId string `json:"inline_message_id,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object for a new inline keyboard.
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *StopMessageLiveLocation) Method() string {
	return "stopMessageLiveLocation"
}

// https://core.telegram.org/bots/api#editmessagereplymarkup
//
// Use this method to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageReplyMarkup struct {
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Identifier of the message to edit
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId string `json:"inline_message_id,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object for an inline keyboard.
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *EditMessageReplyMarkup) Method() string {
	return "editMessageReplyMarkup"
}

// https://core.telegram.org/bots/api#stoppoll
//
// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
type StopPoll struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Identifier of the original message with the poll
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object for a new message inline keyboard.
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *StopPoll) Method() string {
	return "stopPoll"
}

// https://core.telegram.org/bots/api#deletemessage
//
// Use this method to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Service messages about a supergroup, channel, or forum topic creation can't be deleted.- A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.- Bots can delete outgoing messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.Returns True on success.
type DeleteMessage struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Identifier of the message to delete
    MessageId int `json:"message_id,omitempty"`
    
}

func (m *DeleteMessage) Method() string {
	return "deleteMessage"
}

// https://core.telegram.org/bots/api#deletemessages
//
// Use this method to delete multiple messages simultaneously. If some of the specified messages can't be found, they are skipped. Returns True on success.
type DeleteMessages struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized list of 1-100 identifiers of messages to delete. See deleteMessage for limitations on which messages can be deleted
    MessageIds []int `json:"message_ids,omitempty"`
    
}

func (m *DeleteMessages) Method() string {
	return "deleteMessages"
}
