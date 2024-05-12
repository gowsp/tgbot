package tgbot

// https://core.telegram.org/bots/api#getme
//
// A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
type GetMe struct {
    
}

func (m *GetMe) Method() string {
	return "getMe"
}

// https://core.telegram.org/bots/api#logout
//
// Use this method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
type LogOut struct {
    
}

func (m *LogOut) Method() string {
	return "logOut"
}

// https://core.telegram.org/bots/api#close
//
// Use this method to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
type Close struct {
    
}

func (m *Close) Method() string {
	return "close"
}

// https://core.telegram.org/bots/api#sendmessage
//
// Use this method to send text messages. On success, the sent Message is returned.
type SendMessage struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Text of the message to be sent, 1-4096 characters after entities parsing
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
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendMessage) Method() string {
	return "sendMessage"
}

// https://core.telegram.org/bots/api#formatting-options
//
// The Bot API supports basic formatting for messages. You can use bold, italic, underlined, strikethrough, spoiler text, block quotations as well as inline links and pre-formatted code in your bots' messages. Telegram clients will render them accordingly. You can specify text entities directly, or use markdown-style or HTML-style formatting.
//
// Note that Telegram clients will display an alert to the user before opening an inline link ('Open this link?' together with the full URL).
//
// Message entities can be nested, providing following restrictions are met:- If two entities have common characters, then one of them is fully contained inside another.- bold, italic, underline, strikethrough, and spoiler entities can contain and can be part of any other entities, except pre and code.- blockquote entities can't be nested.- All other entities can't contain each other.
//
// Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username. Please note:
//
//   - These links will work only if they are used inside an inline link or in an inline keyboard button. For example, they will not work, when used in a message text.
//
//   - Unless the user is a member of the chat where they were mentioned, these mentions are only guaranteed to work if the user has contacted the bot in private in the past or has sent a callback query to the bot via an inline button and doesn't have Forwarded Messages privacy enabled for the bot.
//
// You can find the list of programming and markup languages for which syntax highlighting is supported at libprisma#supported-languages.
//
// To use this mode, pass MarkdownV2 in the parse_mode field. Use the following syntax in your message:
//
// Please note:
//
//   - Any character with code between 1 and 126 inclusively can be escaped anywhere with a preceding '\' character, in which case it is treated as an ordinary character and not a part of the markup. This implies that '\' character usually must be escaped with a preceding '\' character.
//
//   - Inside pre and code entities, all '`' and '\' characters must be escaped with a preceding '\' character.
//
//   - Inside the (...) part of the inline link and custom emoji definition, all ')' and '\' must be escaped with a preceding '\' character.
//
//   - In all other places characters '_', '*', '[', ']', '(', ')', '~', '`', '>', '#', '+', '-', '=', '|', '{', '}', '.', '!' must be escaped with the preceding character '\'.
//
//   - In case of ambiguity between italic and underline entities __ is always greadily treated from left to right as beginning or end of underline entity, so instead of ___italic underline___ use ___italic underline_\r__, where \r is a character with code 13, which will be ignored.
//
//   - A valid emoji must be provided as an alternative value for the custom emoji. The emoji will be shown instead of the custom emoji in places where a custom emoji cannot be displayed (e.g., system notifications) or if the message is forwarded by a non-premium user. It is recommended to use the emoji from the emoji field of the custom emoji sticker.
//
//   - Custom emoji entities can only be used by bots that purchased additional usernames on Fragment.
//
// To use this mode, pass HTML in the parse_mode field. The following tags are currently supported:
//
// Please note:
//
//   - Only the tags mentioned above are currently supported.
//
//   - All <, > and & symbols that are not a part of a tag or an HTML entity must be replaced with the corresponding HTML entities (< with &lt;, > with &gt; and & with &amp;).
//
//   - All numerical HTML entities are supported.
//
//   - The API currently supports only the following named HTML entities: &lt;, &gt;, &amp; and &quot;.
//
//   - Use nested pre and code tags, to define programming language for pre entity.
//
//   - Programming language can't be specified for standalone code tags.
//
//   - A valid emoji must be used as the content of the tg-emoji tag. The emoji will be shown instead of the custom emoji in places where a custom emoji cannot be displayed (e.g., system notifications) or if the message is forwarded by a non-premium user. It is recommended to use the emoji from the emoji field of the custom emoji sticker.
//
//   - Custom emoji entities can only be used by bots that purchased additional usernames on Fragment.
//
// This is a legacy mode, retained for backward compatibility. To use this mode, pass Markdown in the parse_mode field. Use the following syntax in your message:
//
// Please note:
//
//   - Entities must not be nested, use parse mode MarkdownV2 instead.
//
//   - There is no way to specify “underline”, “strikethrough”, “spoiler”, “blockquote” and “custom_emoji” entities, use parse mode MarkdownV2 instead.
//
//   - To escape characters '_', '*', '`', '[' outside of an entity, prepend the characters '\' before them.
//
//   - Escaping inside entities is not allowed, so entity must be closed first and reopened again: use _snake_\__case_ for italic snake_case and *2*\**2=4* for bold 2*2=4.
type Formatting_options struct {
    
}


// https://core.telegram.org/bots/api#forwardmessage
//
// Use this method to forward messages of any kind. Service messages and messages with protected content can't be forwarded. On success, the sent Message is returned.
type ForwardMessage struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
    FromChatId string `json:"from_chat_id,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the forwarded message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Yes
    //
    // Message identifier in the chat specified in from_chat_id
    MessageId int `json:"message_id,omitempty"`
    
}

func (m *ForwardMessage) Method() string {
	return "forwardMessage"
}

// https://core.telegram.org/bots/api#forwardmessages
//
// Use this method to forward multiple messages of any kind. If some of the specified messages can't be found or forwarded, they are skipped. Service messages and messages with protected content can't be forwarded. Album grouping is kept for forwarded messages. On success, an array of MessageId of the sent messages is returned.
type ForwardMessages struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
    FromChatId string `json:"from_chat_id,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to forward. The identifiers must be specified in a strictly increasing order.
    MessageIds []int `json:"message_ids,omitempty"`
    // Required: Optional
    //
    // Sends the messages silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the forwarded messages from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    
}

func (m *ForwardMessages) Method() string {
	return "forwardMessages"
}

// https://core.telegram.org/bots/api#copymessage
//
// Use this method to copy messages of any kind. Service messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
type CopyMessage struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
    FromChatId string `json:"from_chat_id,omitempty"`
    // Required: Yes
    //
    // Message identifier in the chat specified in from_chat_id
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
    Caption string `json:"caption,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the new caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *CopyMessage) Method() string {
	return "copyMessage"
}

// https://core.telegram.org/bots/api#copymessages
//
// Use this method to copy messages of any kind. If some of the specified messages can't be found or copied, they are skipped. Service messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message. Album grouping is kept for copied messages. On success, an array of MessageId of the sent messages is returned.
type CopyMessages struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
    FromChatId string `json:"from_chat_id,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to copy. The identifiers must be specified in a strictly increasing order.
    MessageIds []int `json:"message_ids,omitempty"`
    // Required: Optional
    //
    // Sends the messages silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent messages from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Pass True to copy the messages without their captions
    RemoveCaption bool `json:"remove_caption,omitempty"`
    
}

func (m *CopyMessages) Method() string {
	return "copyMessages"
}

// https://core.telegram.org/bots/api#sendphoto
//
// Use this method to send photos. On success, the sent Message is returned.
type SendPhoto struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More information on Sending Files »
    Photo string `json:"photo,omitempty"`
    // Required: Optional
    //
    // Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the photo caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Required: Optional
    //
    // Pass True if the photo needs to be covered with a spoiler animation
    HasSpoiler bool `json:"has_spoiler,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendPhoto) Method() string {
	return "sendPhoto"
}

// https://core.telegram.org/bots/api#sendaudio
//
// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
//
// For sending voice messages, use the sendVoice method instead.
type SendAudio struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
    Audio string `json:"audio,omitempty"`
    // Required: Optional
    //
    // Audio caption, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the audio caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Required: Optional
    //
    // Duration of the audio in seconds
    Duration int `json:"duration,omitempty"`
    // Required: Optional
    //
    // Performer
    Performer string `json:"performer,omitempty"`
    // Required: Optional
    //
    // Track name
    Title string `json:"title,omitempty"`
    // Required: Optional
    //
    // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
    Thumbnail string `json:"thumbnail,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendAudio) Method() string {
	return "sendAudio"
}

// https://core.telegram.org/bots/api#senddocument
//
// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocument struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
    Document string `json:"document,omitempty"`
    // Required: Optional
    //
    // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
    Thumbnail string `json:"thumbnail,omitempty"`
    // Required: Optional
    //
    // Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the document caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Required: Optional
    //
    // Disables automatic server-side content type detection for files uploaded using multipart/form-data
    DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendDocument) Method() string {
	return "sendDocument"
}

// https://core.telegram.org/bots/api#sendvideo
//
// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideo struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More information on Sending Files »
    Video string `json:"video,omitempty"`
    // Required: Optional
    //
    // Duration of sent video in seconds
    Duration int `json:"duration,omitempty"`
    // Required: Optional
    //
    // Video width
    Width int `json:"width,omitempty"`
    // Required: Optional
    //
    // Video height
    Height int `json:"height,omitempty"`
    // Required: Optional
    //
    // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
    Thumbnail string `json:"thumbnail,omitempty"`
    // Required: Optional
    //
    // Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the video caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Required: Optional
    //
    // Pass True if the video needs to be covered with a spoiler animation
    HasSpoiler bool `json:"has_spoiler,omitempty"`
    // Required: Optional
    //
    // Pass True if the uploaded video is suitable for streaming
    SupportsStreaming bool `json:"supports_streaming,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendVideo) Method() string {
	return "sendVideo"
}

// https://core.telegram.org/bots/api#sendanimation
//
// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type SendAnimation struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More information on Sending Files »
    Animation string `json:"animation,omitempty"`
    // Required: Optional
    //
    // Duration of sent animation in seconds
    Duration int `json:"duration,omitempty"`
    // Required: Optional
    //
    // Animation width
    Width int `json:"width,omitempty"`
    // Required: Optional
    //
    // Animation height
    Height int `json:"height,omitempty"`
    // Required: Optional
    //
    // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
    Thumbnail string `json:"thumbnail,omitempty"`
    // Required: Optional
    //
    // Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the animation caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Required: Optional
    //
    // Pass True if the animation needs to be covered with a spoiler animation
    HasSpoiler bool `json:"has_spoiler,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendAnimation) Method() string {
	return "sendAnimation"
}

// https://core.telegram.org/bots/api#sendvoice
//
// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoice struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files »
    Voice string `json:"voice,omitempty"`
    // Required: Optional
    //
    // Voice message caption, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the voice message caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Required: Optional
    //
    // Duration of the voice message in seconds
    Duration int `json:"duration,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendVoice) Method() string {
	return "sendVoice"
}

// https://core.telegram.org/bots/api#sendvideonote
//
// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
type SendVideoNote struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files ». Sending video notes by a URL is currently unsupported
    VideoNote string `json:"video_note,omitempty"`
    // Required: Optional
    //
    // Duration of sent video in seconds
    Duration int `json:"duration,omitempty"`
    // Required: Optional
    //
    // Video width and height, i.e. diameter of the video message
    Length int `json:"length,omitempty"`
    // Required: Optional
    //
    // Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files »
    Thumbnail string `json:"thumbnail,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendVideoNote) Method() string {
	return "sendVideoNote"
}

// https://core.telegram.org/bots/api#sendmediagroup
//
// Use this method to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Messages that were sent is returned.
type SendMediaGroup struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized array describing messages to be sent, must include 2-10 items
    Media InputMediaAudio `json:"media,omitempty"`
    // Required: Optional
    //
    // Sends messages silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent messages from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    
}

func (m *SendMediaGroup) Method() string {
	return "sendMediaGroup"
}

// https://core.telegram.org/bots/api#sendlocation
//
// Use this method to send point on the map. On success, the sent Message is returned.
type SendLocation struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Latitude of the location
    Latitude float64 `json:"latitude,omitempty"`
    // Required: Yes
    //
    // Longitude of the location
    Longitude float64 `json:"longitude,omitempty"`
    // Required: Optional
    //
    // The radius of uncertainty for the location, measured in meters; 0-1500
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    // Required: Optional
    //
    // Period in seconds during which the location will be updated (see Live Locations, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
    LivePeriod int `json:"live_period,omitempty"`
    // Required: Optional
    //
    // For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
    Heading int `json:"heading,omitempty"`
    // Required: Optional
    //
    // For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
    ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendLocation) Method() string {
	return "sendLocation"
}

// https://core.telegram.org/bots/api#sendvenue
//
// Use this method to send information about a venue. On success, the sent Message is returned.
type SendVenue struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Latitude of the venue
    Latitude float64 `json:"latitude,omitempty"`
    // Required: Yes
    //
    // Longitude of the venue
    Longitude float64 `json:"longitude,omitempty"`
    // Required: Yes
    //
    // Name of the venue
    Title string `json:"title,omitempty"`
    // Required: Yes
    //
    // Address of the venue
    Address string `json:"address,omitempty"`
    // Required: Optional
    //
    // Foursquare identifier of the venue
    FoursquareId string `json:"foursquare_id,omitempty"`
    // Required: Optional
    //
    // Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
    FoursquareType string `json:"foursquare_type,omitempty"`
    // Required: Optional
    //
    // Google Places identifier of the venue
    GooglePlaceId string `json:"google_place_id,omitempty"`
    // Required: Optional
    //
    // Google Places type of the venue. (See supported types.)
    GooglePlaceType string `json:"google_place_type,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendVenue) Method() string {
	return "sendVenue"
}

// https://core.telegram.org/bots/api#sendcontact
//
// Use this method to send phone contacts. On success, the sent Message is returned.
type SendContact struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Contact's phone number
    PhoneNumber string `json:"phone_number,omitempty"`
    // Required: Yes
    //
    // Contact's first name
    FirstName string `json:"first_name,omitempty"`
    // Required: Optional
    //
    // Contact's last name
    LastName string `json:"last_name,omitempty"`
    // Required: Optional
    //
    // Additional data about the contact in the form of a vCard, 0-2048 bytes
    Vcard string `json:"vcard,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendContact) Method() string {
	return "sendContact"
}

// https://core.telegram.org/bots/api#sendpoll
//
// Use this method to send a native poll. On success, the sent Message is returned.
type SendPoll struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Poll question, 1-300 characters
    Question string `json:"question,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the question. See formatting options for more details. Currently, only custom emoji entities are allowed
    QuestionParseMode string `json:"question_parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in the poll question. It can be specified instead of question_parse_mode
    QuestionEntities []MessageEntity `json:"question_entities,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized list of 2-10 answer options
    Options []InputPollOption `json:"options,omitempty"`
    // Required: Optional
    //
    // True, if the poll needs to be anonymous, defaults to True
    IsAnonymous bool `json:"is_anonymous,omitempty"`
    // Required: Optional
    //
    // Poll type, “quiz” or “regular”, defaults to “regular”
    Type string `json:"type,omitempty"`
    // Required: Optional
    //
    // True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
    AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`
    // Required: Optional
    //
    // 0-based identifier of the correct answer option, required for polls in quiz mode
    CorrectOptionId int `json:"correct_option_id,omitempty"`
    // Required: Optional
    //
    // Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
    Explanation string `json:"explanation,omitempty"`
    // Required: Optional
    //
    // Mode for parsing entities in the explanation. See formatting options for more details.
    ExplanationParseMode string `json:"explanation_parse_mode,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of special entities that appear in the poll explanation. It can be specified instead of explanation_parse_mode
    ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
    // Required: Optional
    //
    // Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
    OpenPeriod int `json:"open_period,omitempty"`
    // Required: Optional
    //
    // Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
    CloseDate int `json:"close_date,omitempty"`
    // Required: Optional
    //
    // Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
    IsClosed bool `json:"is_closed,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding and saving
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendPoll) Method() string {
	return "sendPoll"
}

// https://core.telegram.org/bots/api#senddice
//
// Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
type SendDice struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Optional
    //
    // Emoji on which the dice throw animation is based. Currently, must be one of “”, “”, “”, “”, “”, or “”. Dice can have values 1-6 for “”, “” and “”, values 1-5 for “” and “”, and values 1-64 for “”. Defaults to “”
    Emoji string `json:"emoji,omitempty"`
    // Required: Optional
    //
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification bool `json:"disable_notification,omitempty"`
    // Required: Optional
    //
    // Protects the contents of the sent message from forwarding
    ProtectContent bool `json:"protect_content,omitempty"`
    // Required: Optional
    //
    // Description of the message to reply to
    ReplyParameters ReplyParameters `json:"reply_parameters,omitempty"`
    // Required: Optional
    //
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendDice) Method() string {
	return "sendDice"
}

// https://core.telegram.org/bots/api#sendchataction
//
// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
//
// Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.
//
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
type SendChatAction struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the action will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread; for supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.
    Action string `json:"action,omitempty"`
    
}

func (m *SendChatAction) Method() string {
	return "sendChatAction"
}

// https://core.telegram.org/bots/api#setmessagereaction
//
// Use this method to change the chosen reactions on a message. Service messages can't be reacted to. Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel. Returns True on success.
type SetMessageReaction struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Identifier of the target message. If the message belongs to a media group, the reaction is set to the first non-deleted message in the group instead.
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of reaction types to set on the message. Currently, as non-premium users, bots can set up to one reaction per message. A custom emoji reaction can be used if it is either already present on the message or explicitly allowed by chat administrators.
    Reaction []ReactionType `json:"reaction,omitempty"`
    // Required: Optional
    //
    // Pass True to set the reaction with a big animation
    IsBig bool `json:"is_big,omitempty"`
    
}

func (m *SetMessageReaction) Method() string {
	return "setMessageReaction"
}

// https://core.telegram.org/bots/api#getuserprofilephotos
//
// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
type GetUserProfilePhotos struct {
    // Required: Yes
    //
    // Unique identifier of the target user
    UserId int `json:"user_id,omitempty"`
    // Required: Optional
    //
    // Sequential number of the first photo to be returned. By default, all photos are returned.
    Offset int `json:"offset,omitempty"`
    // Required: Optional
    //
    // Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
    Limit int `json:"limit,omitempty"`
    
}

func (m *GetUserProfilePhotos) Method() string {
	return "getUserProfilePhotos"
}

// https://core.telegram.org/bots/api#getfile
//
// Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
//
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
type GetFile struct {
    // Required: Yes
    //
    // File identifier to get information about
    FileId string `json:"file_id,omitempty"`
    
}

func (m *GetFile) Method() string {
	return "getFile"
}

// https://core.telegram.org/bots/api#banchatmember
//
// Use this method to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type BanChatMember struct {
    // Required: Yes
    //
    // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target user
    UserId int `json:"user_id,omitempty"`
    // Required: Optional
    //
    // Date when the user will be unbanned; Unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
    UntilDate int `json:"until_date,omitempty"`
    // Required: Optional
    //
    // Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.
    RevokeMessages bool `json:"revoke_messages,omitempty"`
    
}

func (m *BanChatMember) Method() string {
	return "banChatMember"
}

// https://core.telegram.org/bots/api#unbanchatmember
//
// Use this method to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
type UnbanChatMember struct {
    // Required: Yes
    //
    // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target user
    UserId int `json:"user_id,omitempty"`
    // Required: Optional
    //
    // Do nothing if the user is not banned
    OnlyIfBanned bool `json:"only_if_banned,omitempty"`
    
}

func (m *UnbanChatMember) Method() string {
	return "unbanChatMember"
}

// https://core.telegram.org/bots/api#restrictchatmember
//
// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
type RestrictChatMember struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target user
    UserId int `json:"user_id,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized object for new user permissions
    Permissions ChatPermissions `json:"permissions,omitempty"`
    // Required: Optional
    //
    // Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
    UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`
    // Required: Optional
    //
    // Date when restrictions will be lifted for the user; Unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
    UntilDate int `json:"until_date,omitempty"`
    
}

func (m *RestrictChatMember) Method() string {
	return "restrictChatMember"
}

// https://core.telegram.org/bots/api#promotechatmember
//
// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
type PromoteChatMember struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target user
    UserId int `json:"user_id,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator's presence in the chat is hidden
    IsAnonymous bool `json:"is_anonymous,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
    CanManageChat bool `json:"can_manage_chat,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can delete messages of other users
    CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can manage video chats
    CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can restrict, ban or unban chat members, or access supergroup statistics
    CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
    CanPromoteMembers bool `json:"can_promote_members,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can change chat title, photo and other settings
    CanChangeInfo bool `json:"can_change_info,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can invite new users to the chat
    CanInviteUsers bool `json:"can_invite_users,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can post stories to the chat
    CanPostStories bool `json:"can_post_stories,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
    CanEditStories bool `json:"can_edit_stories,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can delete stories posted by other users
    CanDeleteStories bool `json:"can_delete_stories,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can post messages in the channel, or access channel statistics; for channels only
    CanPostMessages bool `json:"can_post_messages,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can edit messages of other users and can pin messages; for channels only
    CanEditMessages bool `json:"can_edit_messages,omitempty"`
    // Required: Optional
    //
    // Pass True if the administrator can pin messages; for supergroups only
    CanPinMessages bool `json:"can_pin_messages,omitempty"`
    // Required: Optional
    //
    // Pass True if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
    CanManageTopics bool `json:"can_manage_topics,omitempty"`
    
}

func (m *PromoteChatMember) Method() string {
	return "promoteChatMember"
}

// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
//
// Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
type SetChatAdministratorCustomTitle struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target user
    UserId int `json:"user_id,omitempty"`
    // Required: Yes
    //
    // New custom title for the administrator; 0-16 characters, emoji are not allowed
    CustomTitle string `json:"custom_title,omitempty"`
    
}

func (m *SetChatAdministratorCustomTitle) Method() string {
	return "setChatAdministratorCustomTitle"
}

// https://core.telegram.org/bots/api#banchatsenderchat
//
// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
type BanChatSenderChat struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target sender chat
    SenderChatId int `json:"sender_chat_id,omitempty"`
    
}

func (m *BanChatSenderChat) Method() string {
	return "banChatSenderChat"
}

// https://core.telegram.org/bots/api#unbanchatsenderchat
//
// Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
type UnbanChatSenderChat struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target sender chat
    SenderChatId int `json:"sender_chat_id,omitempty"`
    
}

func (m *UnbanChatSenderChat) Method() string {
	return "unbanChatSenderChat"
}

// https://core.telegram.org/bots/api#setchatpermissions
//
// Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
type SetChatPermissions struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized object for new default chat permissions
    Permissions ChatPermissions `json:"permissions,omitempty"`
    // Required: Optional
    //
    // Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
    UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`
    
}

func (m *SetChatPermissions) Method() string {
	return "setChatPermissions"
}

// https://core.telegram.org/bots/api#exportchatinvitelink
//
// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other administrators. If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink or by calling the getChat method. If your bot needs to generate a new primary invite link replacing its previous one, use exportChatInviteLink again.
type ExportChatInviteLink struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *ExportChatInviteLink) Method() string {
	return "exportChatInviteLink"
}

// https://core.telegram.org/bots/api#createchatinvitelink
//
// Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
type CreateChatInviteLink struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Invite link name; 0-32 characters
    Name string `json:"name,omitempty"`
    // Required: Optional
    //
    // Point in time (Unix timestamp) when the link will expire
    ExpireDate int `json:"expire_date,omitempty"`
    // Required: Optional
    //
    // The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
    MemberLimit int `json:"member_limit,omitempty"`
    // Required: Optional
    //
    // True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
    CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
    
}

func (m *CreateChatInviteLink) Method() string {
	return "createChatInviteLink"
}

// https://core.telegram.org/bots/api#editchatinvitelink
//
// Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
type EditChatInviteLink struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // The invite link to edit
    InviteLink string `json:"invite_link,omitempty"`
    // Required: Optional
    //
    // Invite link name; 0-32 characters
    Name string `json:"name,omitempty"`
    // Required: Optional
    //
    // Point in time (Unix timestamp) when the link will expire
    ExpireDate int `json:"expire_date,omitempty"`
    // Required: Optional
    //
    // The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
    MemberLimit int `json:"member_limit,omitempty"`
    // Required: Optional
    //
    // True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
    CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
    
}

func (m *EditChatInviteLink) Method() string {
	return "editChatInviteLink"
}

// https://core.telegram.org/bots/api#revokechatinvitelink
//
// Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
type RevokeChatInviteLink struct {
    // Required: Yes
    //
    // Unique identifier of the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // The invite link to revoke
    InviteLink string `json:"invite_link,omitempty"`
    
}

func (m *RevokeChatInviteLink) Method() string {
	return "revokeChatInviteLink"
}

// https://core.telegram.org/bots/api#approvechatjoinrequest
//
// Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
type ApproveChatJoinRequest struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target user
    UserId int `json:"user_id,omitempty"`
    
}

func (m *ApproveChatJoinRequest) Method() string {
	return "approveChatJoinRequest"
}

// https://core.telegram.org/bots/api#declinechatjoinrequest
//
// Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
type DeclineChatJoinRequest struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target user
    UserId int `json:"user_id,omitempty"`
    
}

func (m *DeclineChatJoinRequest) Method() string {
	return "declineChatJoinRequest"
}

// https://core.telegram.org/bots/api#setchatphoto
//
// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatPhoto struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // New chat photo, uploaded using multipart/form-data
    Photo InputFile `json:"photo,omitempty"`
    
}

func (m *SetChatPhoto) Method() string {
	return "setChatPhoto"
}

// https://core.telegram.org/bots/api#deletechatphoto
//
// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type DeleteChatPhoto struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *DeleteChatPhoto) Method() string {
	return "deleteChatPhoto"
}

// https://core.telegram.org/bots/api#setchattitle
//
// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatTitle struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // New chat title, 1-128 characters
    Title string `json:"title,omitempty"`
    
}

func (m *SetChatTitle) Method() string {
	return "setChatTitle"
}

// https://core.telegram.org/bots/api#setchatdescription
//
// Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatDescription struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // New chat description, 0-255 characters
    Description string `json:"description,omitempty"`
    
}

func (m *SetChatDescription) Method() string {
	return "setChatDescription"
}

// https://core.telegram.org/bots/api#pinchatmessage
//
// Use this method to add a message to the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
type PinChatMessage struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Identifier of a message to pin
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // Pass True if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
    DisableNotification bool `json:"disable_notification,omitempty"`
    
}

func (m *PinChatMessage) Method() string {
	return "pinChatMessage"
}

// https://core.telegram.org/bots/api#unpinchatmessage
//
// Use this method to remove a message from the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
type UnpinChatMessage struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Identifier of a message to unpin. If not specified, the most recent pinned message (by sending date) will be unpinned.
    MessageId int `json:"message_id,omitempty"`
    
}

func (m *UnpinChatMessage) Method() string {
	return "unpinChatMessage"
}

// https://core.telegram.org/bots/api#unpinallchatmessages
//
// Use this method to clear the list of pinned messages in a chat. If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel. Returns True on success.
type UnpinAllChatMessages struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *UnpinAllChatMessages) Method() string {
	return "unpinAllChatMessages"
}

// https://core.telegram.org/bots/api#leavechat
//
// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
type LeaveChat struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *LeaveChat) Method() string {
	return "leaveChat"
}

// https://core.telegram.org/bots/api#getchat
//
// Use this method to get up-to-date information about the chat. Returns a ChatFullInfo object on success.
type GetChat struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *GetChat) Method() string {
	return "getChat"
}

// https://core.telegram.org/bots/api#getchatadministrators
//
// Use this method to get a list of administrators in a chat, which aren't bots. Returns an Array of ChatMember objects.
type GetChatAdministrators struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *GetChatAdministrators) Method() string {
	return "getChatAdministrators"
}

// https://core.telegram.org/bots/api#getchatmembercount
//
// Use this method to get the number of members in a chat. Returns Int on success.
type GetChatMemberCount struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *GetChatMemberCount) Method() string {
	return "getChatMemberCount"
}

// https://core.telegram.org/bots/api#getchatmember
//
// Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
type GetChatMember struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target user
    UserId int `json:"user_id,omitempty"`
    
}

func (m *GetChatMember) Method() string {
	return "getChatMember"
}

// https://core.telegram.org/bots/api#setchatstickerset
//
// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type SetChatStickerSet struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Name of the sticker set to be set as the group sticker set
    StickerSetName string `json:"sticker_set_name,omitempty"`
    
}

func (m *SetChatStickerSet) Method() string {
	return "setChatStickerSet"
}

// https://core.telegram.org/bots/api#deletechatstickerset
//
// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type DeleteChatStickerSet struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *DeleteChatStickerSet) Method() string {
	return "deleteChatStickerSet"
}

// https://core.telegram.org/bots/api#getforumtopiciconstickers
//
// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of Sticker objects.
type GetForumTopicIconStickers struct {
    
}

func (m *GetForumTopicIconStickers) Method() string {
	return "getForumTopicIconStickers"
}

// https://core.telegram.org/bots/api#createforumtopic
//
// Use this method to create a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns information about the created topic as a ForumTopic object.
type CreateForumTopic struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Topic name, 1-128 characters
    Name string `json:"name,omitempty"`
    // Required: Optional
    //
    // Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)
    IconColor int `json:"icon_color,omitempty"`
    // Required: Optional
    //
    // Unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
    IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
    
}

func (m *CreateForumTopic) Method() string {
	return "createForumTopic"
}

// https://core.telegram.org/bots/api#editforumtopic
//
// Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
type EditForumTopic struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target message thread of the forum topic
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Optional
    //
    // New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept
    Name string `json:"name,omitempty"`
    // Required: Optional
    //
    // New unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be kept
    IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
    
}

func (m *EditForumTopic) Method() string {
	return "editForumTopic"
}

// https://core.telegram.org/bots/api#closeforumtopic
//
// Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
type CloseForumTopic struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target message thread of the forum topic
    MessageThreadId int `json:"message_thread_id,omitempty"`
    
}

func (m *CloseForumTopic) Method() string {
	return "closeForumTopic"
}

// https://core.telegram.org/bots/api#reopenforumtopic
//
// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
type ReopenForumTopic struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target message thread of the forum topic
    MessageThreadId int `json:"message_thread_id,omitempty"`
    
}

func (m *ReopenForumTopic) Method() string {
	return "reopenForumTopic"
}

// https://core.telegram.org/bots/api#deleteforumtopic
//
// Use this method to delete a forum topic along with all its messages in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
type DeleteForumTopic struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target message thread of the forum topic
    MessageThreadId int `json:"message_thread_id,omitempty"`
    
}

func (m *DeleteForumTopic) Method() string {
	return "deleteForumTopic"
}

// https://core.telegram.org/bots/api#unpinallforumtopicmessages
//
// Use this method to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
type UnpinAllForumTopicMessages struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target message thread of the forum topic
    MessageThreadId int `json:"message_thread_id,omitempty"`
    
}

func (m *UnpinAllForumTopicMessages) Method() string {
	return "unpinAllForumTopicMessages"
}

// https://core.telegram.org/bots/api#editgeneralforumtopic
//
// Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have can_manage_topics administrator rights. Returns True on success.
type EditGeneralForumTopic struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // New topic name, 1-128 characters
    Name string `json:"name,omitempty"`
    
}

func (m *EditGeneralForumTopic) Method() string {
	return "editGeneralForumTopic"
}

// https://core.telegram.org/bots/api#closegeneralforumtopic
//
// Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
type CloseGeneralForumTopic struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *CloseGeneralForumTopic) Method() string {
	return "closeGeneralForumTopic"
}

// https://core.telegram.org/bots/api#reopengeneralforumtopic
//
// Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
type ReopenGeneralForumTopic struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *ReopenGeneralForumTopic) Method() string {
	return "reopenGeneralForumTopic"
}

// https://core.telegram.org/bots/api#hidegeneralforumtopic
//
// Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
type HideGeneralForumTopic struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *HideGeneralForumTopic) Method() string {
	return "hideGeneralForumTopic"
}

// https://core.telegram.org/bots/api#unhidegeneralforumtopic
//
// Use this method to unhide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
type UnhideGeneralForumTopic struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *UnhideGeneralForumTopic) Method() string {
	return "unhideGeneralForumTopic"
}

// https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
//
// Use this method to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
type UnpinAllGeneralForumTopicMessages struct {
    // Required: Yes
    //
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId string `json:"chat_id,omitempty"`
    
}

func (m *UnpinAllGeneralForumTopicMessages) Method() string {
	return "unpinAllGeneralForumTopicMessages"
}

// https://core.telegram.org/bots/api#answercallbackquery
//
// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
type AnswerCallbackQuery struct {
    // Required: Yes
    //
    // Unique identifier for the query to be answered
    CallbackQueryId string `json:"callback_query_id,omitempty"`
    // Required: Optional
    //
    // Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
    Text string `json:"text,omitempty"`
    // Required: Optional
    //
    // If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
    ShowAlert bool `json:"show_alert,omitempty"`
    // Required: Optional
    //
    // URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @BotFather, specify the URL that opens your game - note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
    Url string `json:"url,omitempty"`
    // Required: Optional
    //
    // The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
    CacheTime int `json:"cache_time,omitempty"`
    
}

func (m *AnswerCallbackQuery) Method() string {
	return "answerCallbackQuery"
}

// https://core.telegram.org/bots/api#getuserchatboosts
//
// Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
type GetUserChatBoosts struct {
    // Required: Yes
    //
    // Unique identifier for the chat or username of the channel (in the format @channelusername)
    ChatId string `json:"chat_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier of the target user
    UserId int `json:"user_id,omitempty"`
    
}

func (m *GetUserChatBoosts) Method() string {
	return "getUserChatBoosts"
}

// https://core.telegram.org/bots/api#getbusinessconnection
//
// Use this method to get information about the connection of the bot with a business account. Returns a BusinessConnection object on success.
type GetBusinessConnection struct {
    // Required: Yes
    //
    // Unique identifier of the business connection
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    
}

func (m *GetBusinessConnection) Method() string {
	return "getBusinessConnection"
}

// https://core.telegram.org/bots/api#setmycommands
//
// Use this method to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
type SetMyCommands struct {
    // Required: Yes
    //
    // A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
    Commands []BotCommand `json:"commands,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
    Scope BotCommandScope `json:"scope,omitempty"`
    // Required: Optional
    //
    // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
    LanguageCode string `json:"language_code,omitempty"`
    
}

func (m *SetMyCommands) Method() string {
	return "setMyCommands"
}

// https://core.telegram.org/bots/api#deletemycommands
//
// Use this method to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
type DeleteMyCommands struct {
    // Required: Optional
    //
    // A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
    Scope BotCommandScope `json:"scope,omitempty"`
    // Required: Optional
    //
    // A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
    LanguageCode string `json:"language_code,omitempty"`
    
}

func (m *DeleteMyCommands) Method() string {
	return "deleteMyCommands"
}

// https://core.telegram.org/bots/api#getmycommands
//
// Use this method to get the current list of the bot's commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
type GetMyCommands struct {
    // Required: Optional
    //
    // A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
    Scope BotCommandScope `json:"scope,omitempty"`
    // Required: Optional
    //
    // A two-letter ISO 639-1 language code or an empty string
    LanguageCode string `json:"language_code,omitempty"`
    
}

func (m *GetMyCommands) Method() string {
	return "getMyCommands"
}

// https://core.telegram.org/bots/api#setmyname
//
// Use this method to change the bot's name. Returns True on success.
type SetMyName struct {
    // Required: Optional
    //
    // New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
    Name string `json:"name,omitempty"`
    // Required: Optional
    //
    // A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name.
    LanguageCode string `json:"language_code,omitempty"`
    
}

func (m *SetMyName) Method() string {
	return "setMyName"
}

// https://core.telegram.org/bots/api#getmyname
//
// Use this method to get the current bot name for the given user language. Returns BotName on success.
type GetMyName struct {
    // Required: Optional
    //
    // A two-letter ISO 639-1 language code or an empty string
    LanguageCode string `json:"language_code,omitempty"`
    
}

func (m *GetMyName) Method() string {
	return "getMyName"
}

// https://core.telegram.org/bots/api#setmydescription
//
// Use this method to change the bot's description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
type SetMyDescription struct {
    // Required: Optional
    //
    // New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
    Description string `json:"description,omitempty"`
    // Required: Optional
    //
    // A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description.
    LanguageCode string `json:"language_code,omitempty"`
    
}

func (m *SetMyDescription) Method() string {
	return "setMyDescription"
}

// https://core.telegram.org/bots/api#getmydescription
//
// Use this method to get the current bot description for the given user language. Returns BotDescription on success.
type GetMyDescription struct {
    // Required: Optional
    //
    // A two-letter ISO 639-1 language code or an empty string
    LanguageCode string `json:"language_code,omitempty"`
    
}

func (m *GetMyDescription) Method() string {
	return "getMyDescription"
}

// https://core.telegram.org/bots/api#setmyshortdescription
//
// Use this method to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot. Returns True on success.
type SetMyShortDescription struct {
    // Required: Optional
    //
    // New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
    ShortDescription string `json:"short_description,omitempty"`
    // Required: Optional
    //
    // A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description.
    LanguageCode string `json:"language_code,omitempty"`
    
}

func (m *SetMyShortDescription) Method() string {
	return "setMyShortDescription"
}

// https://core.telegram.org/bots/api#getmyshortdescription
//
// Use this method to get the current bot short description for the given user language. Returns BotShortDescription on success.
type GetMyShortDescription struct {
    // Required: Optional
    //
    // A two-letter ISO 639-1 language code or an empty string
    LanguageCode string `json:"language_code,omitempty"`
    
}

func (m *GetMyShortDescription) Method() string {
	return "getMyShortDescription"
}

// https://core.telegram.org/bots/api#setchatmenubutton
//
// Use this method to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
type SetChatMenuButton struct {
    // Required: Optional
    //
    // Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
    ChatId int `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
    MenuButton MenuButton `json:"menu_button,omitempty"`
    
}

func (m *SetChatMenuButton) Method() string {
	return "setChatMenuButton"
}

// https://core.telegram.org/bots/api#getchatmenubutton
//
// Use this method to get the current value of the bot's menu button in a private chat, or the default menu button. Returns MenuButton on success.
type GetChatMenuButton struct {
    // Required: Optional
    //
    // Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
    ChatId int `json:"chat_id,omitempty"`
    
}

func (m *GetChatMenuButton) Method() string {
	return "getChatMenuButton"
}

// https://core.telegram.org/bots/api#setmydefaultadministratorrights
//
// Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
type SetMyDefaultAdministratorRights struct {
    // Required: Optional
    //
    // A JSON-serialized object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
    Rights ChatAdministratorRights `json:"rights,omitempty"`
    // Required: Optional
    //
    // Pass True to change the default administrator rights of the bot in channels. Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
    ForChannels bool `json:"for_channels,omitempty"`
    
}

func (m *SetMyDefaultAdministratorRights) Method() string {
	return "setMyDefaultAdministratorRights"
}

// https://core.telegram.org/bots/api#getmydefaultadministratorrights
//
// Use this method to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
type GetMyDefaultAdministratorRights struct {
    // Required: Optional
    //
    // Pass True to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
    ForChannels bool `json:"for_channels,omitempty"`
    
}

func (m *GetMyDefaultAdministratorRights) Method() string {
	return "getMyDefaultAdministratorRights"
}

// https://core.telegram.org/bots/api#inline-mode-methods
//
// Methods and objects used in the inline mode are described in the Inline mode section.
type Inline_mode_methods struct {
    
}

