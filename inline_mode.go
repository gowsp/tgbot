package tgbot

// https://core.telegram.org/bots/api#inlinequery
//
// This object represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
    // Unique identifier for this query
    Id string `json:"id,omitempty"`
    // Sender
    From User `json:"from,omitempty"`
    // Text of the query (up to 256 characters)
    Query string `json:"query,omitempty"`
    // Offset of the results to be returned, can be controlled by the bot
    Offset string `json:"offset,omitempty"`
    // Optional. Type of the chat from which the inline query was sent. Can be either “sender” for a private chat with the inline query sender, “private”, “group”, “supergroup”, or “channel”. The chat type should be always known for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat
    ChatType string `json:"chat_type,omitempty"`
    // Optional. Sender location, only for bots that request user location
    Location Location `json:"location,omitempty"`
    
}


// https://core.telegram.org/bots/api#answerinlinequery
//
// Use this method to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
type AnswerInlineQuery struct {
    // Required: Yes
    //
    // Unique identifier for the answered query
    InlineQueryId string `json:"inline_query_id,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized array of results for the inline query
    Results []InlineQueryResult `json:"results,omitempty"`
    // Required: Optional
    //
    // The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
    CacheTime int `json:"cache_time,omitempty"`
    // Required: Optional
    //
    // Pass True if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query.
    IsPersonal bool `json:"is_personal,omitempty"`
    // Required: Optional
    //
    // Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don't support pagination. Offset length can't exceed 64 bytes.
    NextOffset string `json:"next_offset,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object describing a button to be shown above inline query results
    Button InlineQueryResultsButton `json:"button,omitempty"`
    
}

func (m *AnswerInlineQuery) Method() string {
	return "answerInlineQuery"
}

// https://core.telegram.org/bots/api#inlinequeryresultsbutton
//
// This object represents a button to be shown above inline query results. You must use exactly one of the optional fields.
type InlineQueryResultsButton struct {
    // Label text on the button
    Text string `json:"text,omitempty"`
    // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to switch back to the inline mode using the method switchInlineQuery inside the Web App.
    WebApp WebAppInfo `json:"web_app,omitempty"`
    // Optional. Deep-linking parameter for the /start message sent to the bot when a user presses the button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account' button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
    StartParameter string `json:"start_parameter,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresult
//
// This object represents one result of an inline query. Telegram clients currently support results of the following 20 types:
//
//   - InlineQueryResultCachedAudio
//
//   - InlineQueryResultCachedDocument
//
//   - InlineQueryResultCachedGif
//
//   - InlineQueryResultCachedMpeg4Gif
//
//   - InlineQueryResultCachedPhoto
//
//   - InlineQueryResultCachedSticker
//
//   - InlineQueryResultCachedVideo
//
//   - InlineQueryResultCachedVoice
//
//   - InlineQueryResultArticle
//
//   - InlineQueryResultAudio
//
//   - InlineQueryResultContact
//
//   - InlineQueryResultGame
//
//   - InlineQueryResultDocument
//
//   - InlineQueryResultGif
//
//   - InlineQueryResultLocation
//
//   - InlineQueryResultMpeg4Gif
//
//   - InlineQueryResultPhoto
//
//   - InlineQueryResultVenue
//
//   - InlineQueryResultVideo
//
//   - InlineQueryResultVoice
//
// Note: All URLs passed in inline query results will be available to end users and therefore must be assumed to be public.
type InlineQueryResult struct {
    
}


// https://core.telegram.org/bots/api#inlinequeryresultarticle
//
// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
    // Type of the result, must be article
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id,omitempty"`
    // Title of the result
    Title string `json:"title,omitempty"`
    // Content of the message to be sent
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. URL of the result
    Url string `json:"url,omitempty"`
    // Optional. Pass True if you don't want the URL to be shown in the message
    HideUrl bool `json:"hide_url,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Url of the thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Thumbnail width
    ThumbnailWidth int `json:"thumbnail_width,omitempty"`
    // Optional. Thumbnail height
    ThumbnailHeight int `json:"thumbnail_height,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultphoto
//
// Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
    // Type of the result, must be photo
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
    PhotoUrl string `json:"photo_url,omitempty"`
    // URL of the thumbnail for the photo
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Width of the photo
    PhotoWidth int `json:"photo_width,omitempty"`
    // Optional. Height of the photo
    PhotoHeight int `json:"photo_height,omitempty"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the photo
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultgif
//
// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
    // Type of the result, must be gif
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid URL for the GIF file. File size must not exceed 1MB
    GifUrl string `json:"gif_url,omitempty"`
    // Optional. Width of the GIF
    GifWidth int `json:"gif_width,omitempty"`
    // Optional. Height of the GIF
    GifHeight int `json:"gif_height,omitempty"`
    // Optional. Duration of the GIF in seconds
    GifDuration int `json:"gif_duration,omitempty"`
    // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
    ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the GIF animation
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
//
// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
    // Type of the result, must be mpeg4_gif
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid URL for the MPEG4 file. File size must not exceed 1MB
    Mpeg4Url string `json:"mpeg4_url,omitempty"`
    // Optional. Video width
    Mpeg4Width int `json:"mpeg4_width,omitempty"`
    // Optional. Video height
    Mpeg4Height int `json:"mpeg4_height,omitempty"`
    // Optional. Video duration in seconds
    Mpeg4Duration int `json:"mpeg4_duration,omitempty"`
    // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”
    ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the video animation
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultvideo
//
// Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
//
// If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you must replace its content using input_message_content.
type InlineQueryResultVideo struct {
    // Type of the result, must be video
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid URL for the embedded video player or video file
    VideoUrl string `json:"video_url,omitempty"`
    // MIME type of the content of the video URL, “text/html” or “video/mp4”
    MimeType string `json:"mime_type,omitempty"`
    // URL of the thumbnail (JPEG only) for the video
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Video width
    VideoWidth int `json:"video_width,omitempty"`
    // Optional. Video height
    VideoHeight int `json:"video_height,omitempty"`
    // Optional. Video duration in seconds
    VideoDuration int `json:"video_duration,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultaudio
//
// Represents a link to an MP3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
    // Type of the result, must be audio
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid URL for the audio file
    AudioUrl string `json:"audio_url,omitempty"`
    // Title
    Title string `json:"title,omitempty"`
    // Optional. Caption, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Performer
    Performer string `json:"performer,omitempty"`
    // Optional. Audio duration in seconds
    AudioDuration int `json:"audio_duration,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the audio
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultvoice
//
// Represents a link to a voice recording in an .OGG container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
    // Type of the result, must be voice
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid URL for the voice recording
    VoiceUrl string `json:"voice_url,omitempty"`
    // Recording title
    Title string `json:"title,omitempty"`
    // Optional. Caption, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Recording duration in seconds
    VoiceDuration int `json:"voice_duration,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the voice recording
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultdocument
//
// Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
    // Type of the result, must be document
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // A valid URL for the file
    DocumentUrl string `json:"document_url,omitempty"`
    // MIME type of the content of the file, either “application/pdf” or “application/zip”
    MimeType string `json:"mime_type,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the file
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    // Optional. URL of the thumbnail (JPEG only) for the file
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Thumbnail width
    ThumbnailWidth int `json:"thumbnail_width,omitempty"`
    // Optional. Thumbnail height
    ThumbnailHeight int `json:"thumbnail_height,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultlocation
//
// Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
    // Type of the result, must be location
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id,omitempty"`
    // Location latitude in degrees
    Latitude float64 `json:"latitude,omitempty"`
    // Location longitude in degrees
    Longitude float64 `json:"longitude,omitempty"`
    // Location title
    Title string `json:"title,omitempty"`
    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    // Optional. Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
    LivePeriod int `json:"live_period,omitempty"`
    // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
    Heading int `json:"heading,omitempty"`
    // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
    ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the location
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    // Optional. Url of the thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Thumbnail width
    ThumbnailWidth int `json:"thumbnail_width,omitempty"`
    // Optional. Thumbnail height
    ThumbnailHeight int `json:"thumbnail_height,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultvenue
//
// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
    // Type of the result, must be venue
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id,omitempty"`
    // Latitude of the venue location in degrees
    Latitude float64 `json:"latitude,omitempty"`
    // Longitude of the venue location in degrees
    Longitude float64 `json:"longitude,omitempty"`
    // Title of the venue
    Title string `json:"title,omitempty"`
    // Address of the venue
    Address string `json:"address,omitempty"`
    // Optional. Foursquare identifier of the venue if known
    FoursquareId string `json:"foursquare_id,omitempty"`
    // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
    FoursquareType string `json:"foursquare_type,omitempty"`
    // Optional. Google Places identifier of the venue
    GooglePlaceId string `json:"google_place_id,omitempty"`
    // Optional. Google Places type of the venue. (See supported types.)
    GooglePlaceType string `json:"google_place_type,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the venue
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    // Optional. Url of the thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Thumbnail width
    ThumbnailWidth int `json:"thumbnail_width,omitempty"`
    // Optional. Thumbnail height
    ThumbnailHeight int `json:"thumbnail_height,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultcontact
//
// Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
    // Type of the result, must be contact
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id,omitempty"`
    // Contact's phone number
    PhoneNumber string `json:"phone_number,omitempty"`
    // Contact's first name
    FirstName string `json:"first_name,omitempty"`
    // Optional. Contact's last name
    LastName string `json:"last_name,omitempty"`
    // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
    Vcard string `json:"vcard,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the contact
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    // Optional. Url of the thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Thumbnail width
    ThumbnailWidth int `json:"thumbnail_width,omitempty"`
    // Optional. Thumbnail height
    ThumbnailHeight int `json:"thumbnail_height,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultgame
//
// Represents a Game.
type InlineQueryResultGame struct {
    // Type of the result, must be game
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // Short name of the game
    GameShortName string `json:"game_short_name,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
//
// Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
    // Type of the result, must be photo
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid file identifier of the photo
    PhotoFileId string `json:"photo_file_id,omitempty"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the photo
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
//
// Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
    // Type of the result, must be gif
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid file identifier for the GIF file
    GifFileId string `json:"gif_file_id,omitempty"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the GIF animation
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
//
// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
    // Type of the result, must be mpeg4_gif
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid file identifier for the MPEG4 file
    Mpeg4FileId string `json:"mpeg4_file_id,omitempty"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the video animation
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
//
// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
    // Type of the result, must be sticker
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid file identifier of the sticker
    StickerFileId string `json:"sticker_file_id,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the sticker
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
//
// Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
    // Type of the result, must be document
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // Title for the result
    Title string `json:"title,omitempty"`
    // A valid file identifier for the file
    DocumentFileId string `json:"document_file_id,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the file
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
//
// Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
    // Type of the result, must be video
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid file identifier for the video file
    VideoFileId string `json:"video_file_id,omitempty"`
    // Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the video
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
//
// Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
    // Type of the result, must be voice
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid file identifier for the voice message
    VoiceFileId string `json:"voice_file_id,omitempty"`
    // Voice message title
    Title string `json:"title,omitempty"`
    // Optional. Caption, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the voice message
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
//
// Represents a link to an MP3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
    // Type of the result, must be audio
    Type string `json:"type,omitempty"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id,omitempty"`
    // A valid file identifier for the audio file
    AudioFileId string `json:"audio_file_id,omitempty"`
    // Optional. Caption, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the audio
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    
}


// https://core.telegram.org/bots/api#inputmessagecontent
//
// This object represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 5 types:
//
//   - InputTextMessageContent
//
//   - InputLocationMessageContent
//
//   - InputVenueMessageContent
//
//   - InputContactMessageContent
//
//   - InputInvoiceMessageContent
type InputMessageContent struct {
    
}


// https://core.telegram.org/bots/api#inputtextmessagecontent
//
// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
    // Text of the message to be sent, 1-4096 characters
    MessageText string `json:"message_text,omitempty"`
    // Optional. Mode for parsing entities in the message text. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in message text, which can be specified instead of parse_mode
    Entities []MessageEntity `json:"entities,omitempty"`
    // Optional. Link preview generation options for the message
    LinkPreviewOptions LinkPreviewOptions `json:"link_preview_options,omitempty"`
    
}


// https://core.telegram.org/bots/api#inputlocationmessagecontent
//
// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
    // Latitude of the location in degrees
    Latitude float64 `json:"latitude,omitempty"`
    // Longitude of the location in degrees
    Longitude float64 `json:"longitude,omitempty"`
    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    // Optional. Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
    LivePeriod int `json:"live_period,omitempty"`
    // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
    Heading int `json:"heading,omitempty"`
    // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
    ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
    
}


// https://core.telegram.org/bots/api#inputvenuemessagecontent
//
// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
    // Latitude of the venue in degrees
    Latitude float64 `json:"latitude,omitempty"`
    // Longitude of the venue in degrees
    Longitude float64 `json:"longitude,omitempty"`
    // Name of the venue
    Title string `json:"title,omitempty"`
    // Address of the venue
    Address string `json:"address,omitempty"`
    // Optional. Foursquare identifier of the venue, if known
    FoursquareId string `json:"foursquare_id,omitempty"`
    // Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
    FoursquareType string `json:"foursquare_type,omitempty"`
    // Optional. Google Places identifier of the venue
    GooglePlaceId string `json:"google_place_id,omitempty"`
    // Optional. Google Places type of the venue. (See supported types.)
    GooglePlaceType string `json:"google_place_type,omitempty"`
    
}


// https://core.telegram.org/bots/api#inputcontactmessagecontent
//
// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
    // Contact's phone number
    PhoneNumber string `json:"phone_number,omitempty"`
    // Contact's first name
    FirstName string `json:"first_name,omitempty"`
    // Optional. Contact's last name
    LastName string `json:"last_name,omitempty"`
    // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
    Vcard string `json:"vcard,omitempty"`
    
}


// https://core.telegram.org/bots/api#inputinvoicemessagecontent
//
// Represents the content of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {
    // Product name, 1-32 characters
    Title string `json:"title,omitempty"`
    // Product description, 1-255 characters
    Description string `json:"description,omitempty"`
    // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
    Payload string `json:"payload,omitempty"`
    // Payment provider token, obtained via @BotFather
    ProviderToken string `json:"provider_token,omitempty"`
    // Three-letter ISO 4217 currency code, see more on currencies
    Currency string `json:"currency,omitempty"`
    // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
    Prices []LabeledPrice `json:"prices,omitempty"`
    // Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
    MaxTipAmount int `json:"max_tip_amount,omitempty"`
    // Optional. A JSON-serialized array of suggested amounts of tip in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
    SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`
    // Optional. A JSON-serialized object for data about the invoice, which will be shared with the payment provider. A detailed description of the required fields should be provided by the payment provider.
    ProviderData string `json:"provider_data,omitempty"`
    // Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
    PhotoUrl string `json:"photo_url,omitempty"`
    // Optional. Photo size in bytes
    PhotoSize int `json:"photo_size,omitempty"`
    // Optional. Photo width
    PhotoWidth int `json:"photo_width,omitempty"`
    // Optional. Photo height
    PhotoHeight int `json:"photo_height,omitempty"`
    // Optional. Pass True if you require the user's full name to complete the order
    NeedName bool `json:"need_name,omitempty"`
    // Optional. Pass True if you require the user's phone number to complete the order
    NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
    // Optional. Pass True if you require the user's email address to complete the order
    NeedEmail bool `json:"need_email,omitempty"`
    // Optional. Pass True if you require the user's shipping address to complete the order
    NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
    // Optional. Pass True if the user's phone number should be sent to provider
    SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
    // Optional. Pass True if the user's email address should be sent to provider
    SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
    // Optional. Pass True if the final price depends on the shipping method
    IsFlexible bool `json:"is_flexible,omitempty"`
    
}


// https://core.telegram.org/bots/api#choseninlineresult
//
// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
//
// Note: It is necessary to enable inline feedback via @BotFather in order to receive these objects in updates.
type ChosenInlineResult struct {
    // The unique identifier for the result that was chosen
    ResultId string `json:"result_id,omitempty"`
    // The user that chose the result
    From User `json:"from,omitempty"`
    // Optional. Sender location, only for bots that require user location
    Location Location `json:"location,omitempty"`
    // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
    InlineMessageId string `json:"inline_message_id,omitempty"`
    // The query that was used to obtain the result
    Query string `json:"query,omitempty"`
    
}


// https://core.telegram.org/bots/api#answerwebappquery
//
// Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
type AnswerWebAppQuery struct {
    // Required: Yes
    //
    // Unique identifier for the query to be answered
    WebAppQueryId string `json:"web_app_query_id,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized object describing the message to be sent
    Result InlineQueryResult `json:"result,omitempty"`
    
}

func (m *AnswerWebAppQuery) Method() string {
	return "answerWebAppQuery"
}

// https://core.telegram.org/bots/api#sentwebappmessage
//
// Describes an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
    // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message.
    InlineMessageId string `json:"inline_message_id,omitempty"`
    
}

