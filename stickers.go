package tgbot

// https://core.telegram.org/bots/api#sticker
//
// This object represents a sticker.
type Sticker struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id,omitempty"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id,omitempty"`
    // Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of the sticker is independent from its format, which is determined by the fields is_animated and is_video.
    Type string `json:"type,omitempty"`
    // Sticker width
    Width int `json:"width,omitempty"`
    // Sticker height
    Height int `json:"height,omitempty"`
    // True, if the sticker is animated
    IsAnimated bool `json:"is_animated,omitempty"`
    // True, if the sticker is a video sticker
    IsVideo bool `json:"is_video,omitempty"`
    // Optional. Sticker thumbnail in the .WEBP or .JPG format
    Thumbnail PhotoSize `json:"thumbnail,omitempty"`
    // Optional. Emoji associated with the sticker
    Emoji string `json:"emoji,omitempty"`
    // Optional. Name of the sticker set to which the sticker belongs
    SetName string `json:"set_name,omitempty"`
    // Optional. For premium regular stickers, premium animation for the sticker
    PremiumAnimation File `json:"premium_animation,omitempty"`
    // Optional. For mask stickers, the position where the mask should be placed
    MaskPosition MaskPosition `json:"mask_position,omitempty"`
    // Optional. For custom emoji stickers, unique identifier of the custom emoji
    CustomEmojiId string `json:"custom_emoji_id,omitempty"`
    // Optional. True, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
    NeedsRepainting bool `json:"needs_repainting,omitempty"`
    // Optional. File size in bytes
    FileSize int `json:"file_size,omitempty"`
    
}


// https://core.telegram.org/bots/api#stickerset
//
// This object represents a sticker set.
type StickerSet struct {
    // Sticker set name
    Name string `json:"name,omitempty"`
    // Sticker set title
    Title string `json:"title,omitempty"`
    // Type of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”
    StickerType string `json:"sticker_type,omitempty"`
    // List of all set stickers
    Stickers []Sticker `json:"stickers,omitempty"`
    // Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
    Thumbnail PhotoSize `json:"thumbnail,omitempty"`
    
}


// https://core.telegram.org/bots/api#maskposition
//
// This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
    // The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
    Point string `json:"point,omitempty"`
    // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
    XShift float64 `json:"x_shift,omitempty"`
    // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
    YShift float64 `json:"y_shift,omitempty"`
    // Mask scaling coefficient. For example, 2.0 means double size.
    Scale float64 `json:"scale,omitempty"`
    
}


// https://core.telegram.org/bots/api#inputsticker
//
// This object describes a sticker to be added to a sticker set.
type InputSticker struct {
    // The added sticker. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, upload a new one using multipart/form-data, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. Animated and video stickers can't be uploaded via HTTP URL. More information on Sending Files »
    Sticker string `json:"sticker,omitempty"`
    // Format of the added sticker, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, “video” for a WEBM video
    Format string `json:"format,omitempty"`
    // List of 1-20 emoji associated with the sticker
    EmojiList []string `json:"emoji_list,omitempty"`
    // Optional. Position where the mask should be placed on faces. For “mask” stickers only.
    MaskPosition MaskPosition `json:"mask_position,omitempty"`
    // Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For “regular” and “custom_emoji” stickers only.
    Keywords []string `json:"keywords,omitempty"`
    
}


// https://core.telegram.org/bots/api#sendsticker
//
// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
type SendSticker struct {
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
    // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data. More information on Sending Files ». Video and animated stickers can't be sent via an HTTP URL.
    Sticker string `json:"sticker,omitempty"`
    // Required: Optional
    //
    // Emoji associated with the sticker; only for just uploaded stickers
    Emoji string `json:"emoji,omitempty"`
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

func (m *SendSticker) Method() string {
	return "sendSticker"
}

// https://core.telegram.org/bots/api#getstickerset
//
// Use this method to get a sticker set. On success, a StickerSet object is returned.
type GetStickerSet struct {
    // Required: Yes
    //
    // Name of the sticker set
    Name string `json:"name,omitempty"`
    
}

func (m *GetStickerSet) Method() string {
	return "getStickerSet"
}

// https://core.telegram.org/bots/api#getcustomemojistickers
//
// Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
type GetCustomEmojiStickers struct {
    // Required: Yes
    //
    // A JSON-serialized list of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
    CustomEmojiIds []string `json:"custom_emoji_ids,omitempty"`
    
}

func (m *GetCustomEmojiStickers) Method() string {
	return "getCustomEmojiStickers"
}

// https://core.telegram.org/bots/api#uploadstickerfile
//
// Use this method to upload a file with a sticker for later use in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods (the file can be used multiple times). Returns the uploaded File on success.
type UploadStickerFile struct {
    // Required: Yes
    //
    // User identifier of sticker file owner
    UserId int `json:"user_id,omitempty"`
    // Required: Yes
    //
    // A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format. See https://core.telegram.org/stickers for technical requirements. More information on Sending Files »
    Sticker InputFile `json:"sticker,omitempty"`
    // Required: Yes
    //
    // Format of the sticker, must be one of “static”, “animated”, “video”
    StickerFormat string `json:"sticker_format,omitempty"`
    
}

func (m *UploadStickerFile) Method() string {
	return "uploadStickerFile"
}

// https://core.telegram.org/bots/api#createnewstickerset
//
// Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
type CreateNewStickerSet struct {
    // Required: Yes
    //
    // User identifier of created sticker set owner
    UserId int `json:"user_id,omitempty"`
    // Required: Yes
    //
    // Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>". <bot_username> is case insensitive. 1-64 characters.
    Name string `json:"name,omitempty"`
    // Required: Yes
    //
    // Sticker set title, 1-64 characters
    Title string `json:"title,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
    Stickers []InputSticker `json:"stickers,omitempty"`
    // Required: Optional
    //
    // Type of stickers in the set, pass “regular”, “mask”, or “custom_emoji”. By default, a regular sticker set is created.
    StickerType string `json:"sticker_type,omitempty"`
    // Required: Optional
    //
    // Pass True if stickers in the sticker set must be repainted to the color of text when used in messages, the accent color if used as emoji status, white on chat photos, or another appropriate color based on context; for custom emoji sticker sets only
    NeedsRepainting bool `json:"needs_repainting,omitempty"`
    
}

func (m *CreateNewStickerSet) Method() string {
	return "createNewStickerSet"
}

// https://core.telegram.org/bots/api#addstickertoset
//
// Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.
type AddStickerToSet struct {
    // Required: Yes
    //
    // User identifier of sticker set owner
    UserId int `json:"user_id,omitempty"`
    // Required: Yes
    //
    // Sticker set name
    Name string `json:"name,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set isn't changed.
    Sticker InputSticker `json:"sticker,omitempty"`
    
}

func (m *AddStickerToSet) Method() string {
	return "addStickerToSet"
}

// https://core.telegram.org/bots/api#setstickerpositioninset
//
// Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
type SetStickerPositionInSet struct {
    // Required: Yes
    //
    // File identifier of the sticker
    Sticker string `json:"sticker,omitempty"`
    // Required: Yes
    //
    // New sticker position in the set, zero-based
    Position int `json:"position,omitempty"`
    
}

func (m *SetStickerPositionInSet) Method() string {
	return "setStickerPositionInSet"
}

// https://core.telegram.org/bots/api#deletestickerfromset
//
// Use this method to delete a sticker from a set created by the bot. Returns True on success.
type DeleteStickerFromSet struct {
    // Required: Yes
    //
    // File identifier of the sticker
    Sticker string `json:"sticker,omitempty"`
    
}

func (m *DeleteStickerFromSet) Method() string {
	return "deleteStickerFromSet"
}

// https://core.telegram.org/bots/api#replacestickerinset
//
// Use this method to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet. Returns True on success.
type ReplaceStickerInSet struct {
    // Required: Yes
    //
    // User identifier of the sticker set owner
    UserId int `json:"user_id,omitempty"`
    // Required: Yes
    //
    // Sticker set name
    Name string `json:"name,omitempty"`
    // Required: Yes
    //
    // File identifier of the replaced sticker
    OldSticker string `json:"old_sticker,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set remains unchanged.
    Sticker InputSticker `json:"sticker,omitempty"`
    
}

func (m *ReplaceStickerInSet) Method() string {
	return "replaceStickerInSet"
}

// https://core.telegram.org/bots/api#setstickeremojilist
//
// Use this method to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
type SetStickerEmojiList struct {
    // Required: Yes
    //
    // File identifier of the sticker
    Sticker string `json:"sticker,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized list of 1-20 emoji associated with the sticker
    EmojiList []string `json:"emoji_list,omitempty"`
    
}

func (m *SetStickerEmojiList) Method() string {
	return "setStickerEmojiList"
}

// https://core.telegram.org/bots/api#setstickerkeywords
//
// Use this method to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
type SetStickerKeywords struct {
    // Required: Yes
    //
    // File identifier of the sticker
    Sticker string `json:"sticker,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64 characters
    Keywords []string `json:"keywords,omitempty"`
    
}

func (m *SetStickerKeywords) Method() string {
	return "setStickerKeywords"
}

// https://core.telegram.org/bots/api#setstickermaskposition
//
// Use this method to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.
type SetStickerMaskPosition struct {
    // Required: Yes
    //
    // File identifier of the sticker
    Sticker string `json:"sticker,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized object with the position where the mask should be placed on faces. Omit the parameter to remove the mask position.
    MaskPosition MaskPosition `json:"mask_position,omitempty"`
    
}

func (m *SetStickerMaskPosition) Method() string {
	return "setStickerMaskPosition"
}

// https://core.telegram.org/bots/api#setstickersettitle
//
// Use this method to set the title of a created sticker set. Returns True on success.
type SetStickerSetTitle struct {
    // Required: Yes
    //
    // Sticker set name
    Name string `json:"name,omitempty"`
    // Required: Yes
    //
    // Sticker set title, 1-64 characters
    Title string `json:"title,omitempty"`
    
}

func (m *SetStickerSetTitle) Method() string {
	return "setStickerSetTitle"
}

// https://core.telegram.org/bots/api#setstickersetthumbnail
//
// Use this method to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.
type SetStickerSetThumbnail struct {
    // Required: Yes
    //
    // Sticker set name
    Name string `json:"name,omitempty"`
    // Required: Yes
    //
    // User identifier of the sticker set owner
    UserId int `json:"user_id,omitempty"`
    // Required: Optional
    //
    // A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size (see https://core.telegram.org/stickers#animated-sticker-requirements for animated sticker technical requirements), or a WEBM video with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#video-sticker-requirements for video sticker technical requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files ». Animated and video sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
    Thumbnail string `json:"thumbnail,omitempty"`
    // Required: Yes
    //
    // Format of the thumbnail, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, or “video” for a WEBM video
    Format string `json:"format,omitempty"`
    
}

func (m *SetStickerSetThumbnail) Method() string {
	return "setStickerSetThumbnail"
}

// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
//
// Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.
type SetCustomEmojiStickerSetThumbnail struct {
    // Required: Yes
    //
    // Sticker set name
    Name string `json:"name,omitempty"`
    // Required: Optional
    //
    // Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail.
    CustomEmojiId string `json:"custom_emoji_id,omitempty"`
    
}

func (m *SetCustomEmojiStickerSetThumbnail) Method() string {
	return "setCustomEmojiStickerSetThumbnail"
}

// https://core.telegram.org/bots/api#deletestickerset
//
// Use this method to delete a sticker set that was created by the bot. Returns True on success.
type DeleteStickerSet struct {
    // Required: Yes
    //
    // Sticker set name
    Name string `json:"name,omitempty"`
    
}

func (m *DeleteStickerSet) Method() string {
	return "deleteStickerSet"
}
