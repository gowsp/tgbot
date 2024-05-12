package tgbot

// https://core.telegram.org/bots/api#sendgame
//
// Use this method to send a game. On success, the sent Message is returned.
type SendGame struct {
    // Required: Optional
    //
    // Unique identifier of the business connection on behalf of which the message will be sent
    BusinessConnectionId string `json:"business_connection_id,omitempty"`
    // Required: Yes
    //
    // Unique identifier for the target chat
    ChatId int `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
    MessageThreadId int `json:"message_thread_id,omitempty"`
    // Required: Yes
    //
    // Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
    GameShortName string `json:"game_short_name,omitempty"`
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
    // A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown. If not empty, the first button must launch the game.
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendGame) Method() string {
	return "sendGame"
}

// https://core.telegram.org/bots/api#game
//
// This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
    // Title of the game
    Title string `json:"title,omitempty"`
    // Description of the game
    Description string `json:"description,omitempty"`
    // Photo that will be displayed in the game message in chats.
    Photo []PhotoSize `json:"photo,omitempty"`
    // Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
    Text string `json:"text,omitempty"`
    // Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
    TextEntities []MessageEntity `json:"text_entities,omitempty"`
    // Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
    Animation Animation `json:"animation,omitempty"`
    
}


// https://core.telegram.org/bots/api#callbackgame
//
// A placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame struct {
    
}


// https://core.telegram.org/bots/api#setgamescore
//
// Use this method to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
type SetGameScore struct {
    // Required: Yes
    //
    // User identifier
    UserId int `json:"user_id,omitempty"`
    // Required: Yes
    //
    // New score, must be non-negative
    Score int `json:"score,omitempty"`
    // Required: Optional
    //
    // Pass True if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
    Force bool `json:"force,omitempty"`
    // Required: Optional
    //
    // Pass True if the game message should not be automatically edited to include the current scoreboard
    DisableEditMessage bool `json:"disable_edit_message,omitempty"`
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Unique identifier for the target chat
    ChatId int `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Identifier of the sent message
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId string `json:"inline_message_id,omitempty"`
    
}

func (m *SetGameScore) Method() string {
	return "setGameScore"
}

// https://core.telegram.org/bots/api#getgamehighscores
//
// Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
//
// This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.
type GetGameHighScores struct {
    // Required: Yes
    //
    // Target user id
    UserId int `json:"user_id,omitempty"`
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Unique identifier for the target chat
    ChatId int `json:"chat_id,omitempty"`
    // Required: Optional
    //
    // Required if inline_message_id is not specified. Identifier of the sent message
    MessageId int `json:"message_id,omitempty"`
    // Required: Optional
    //
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId string `json:"inline_message_id,omitempty"`
    
}

func (m *GetGameHighScores) Method() string {
	return "getGameHighScores"
}

// https://core.telegram.org/bots/api#gamehighscore
//
// This object represents one row of the high scores table for a game.
//
// And that's about all we've got for now.If you've got any questions, please check out our Bot FAQ »
//
//   - FAQ
//
//   - Privacy
//
//   - Press
//
//   - iPhone/iPad
//
//   - Android
//
//   - Mobile Web
//
//   - PC/Mac/Linux
//
//   - macOS
//
//   - Web-browser
//
//   - API
//
//   - Translations
//
//   - Instant View
type GameHighScore struct {
    // Position in high score table for the game
    Position int `json:"position,omitempty"`
    // User
    User User `json:"user,omitempty"`
    // Score
    Score int `json:"score,omitempty"`
    
}

