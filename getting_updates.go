package tgbot

// https://core.telegram.org/bots/api#update
//
// This object represents an incoming update.At most one of the optional parameters can be present in any given update.
type Update struct {
    // The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This identifier becomes especially handy if you're using webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
    UpdateId int `json:"update_id,omitempty"`
    // Optional. New incoming message of any kind - text, photo, sticker, etc.
    Message Message `json:"message,omitempty"`
    // Optional. New version of a message that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
    EditedMessage Message `json:"edited_message,omitempty"`
    // Optional. New incoming channel post of any kind - text, photo, sticker, etc.
    ChannelPost Message `json:"channel_post,omitempty"`
    // Optional. New version of a channel post that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
    EditedChannelPost Message `json:"edited_channel_post,omitempty"`
    // Optional. The bot was connected to or disconnected from a business account, or a user edited an existing connection with the bot
    BusinessConnection BusinessConnection `json:"business_connection,omitempty"`
    // Optional. New message from a connected business account
    BusinessMessage Message `json:"business_message,omitempty"`
    // Optional. New version of a message from a connected business account
    EditedBusinessMessage Message `json:"edited_business_message,omitempty"`
    // Optional. Messages were deleted from a connected business account
    DeletedBusinessMessages BusinessMessagesDeleted `json:"deleted_business_messages,omitempty"`
    // Optional. A reaction to a message was changed by a user. The bot must be an administrator in the chat and must explicitly specify "message_reaction" in the list of allowed_updates to receive these updates. The update isn't received for reactions set by bots.
    MessageReaction MessageReactionUpdated `json:"message_reaction,omitempty"`
    // Optional. Reactions to a message with anonymous reactions were changed. The bot must be an administrator in the chat and must explicitly specify "message_reaction_count" in the list of allowed_updates to receive these updates. The updates are grouped and can be sent with delay up to a few minutes.
    MessageReactionCount MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
    // Optional. New incoming inline query
    InlineQuery InlineQuery `json:"inline_query,omitempty"`
    // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
    ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result,omitempty"`
    // Optional. New incoming callback query
    CallbackQuery CallbackQuery `json:"callback_query,omitempty"`
    // Optional. New incoming shipping query. Only for invoices with flexible price
    ShippingQuery ShippingQuery `json:"shipping_query,omitempty"`
    // Optional. New incoming pre-checkout query. Contains full information about checkout
    PreCheckoutQuery PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
    // Optional. New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot
    Poll Poll `json:"poll,omitempty"`
    // Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
    PollAnswer PollAnswer `json:"poll_answer,omitempty"`
    // Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
    MyChatMember ChatMemberUpdated `json:"my_chat_member,omitempty"`
    // Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
    ChatMember ChatMemberUpdated `json:"chat_member,omitempty"`
    // Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
    ChatJoinRequest ChatJoinRequest `json:"chat_join_request,omitempty"`
    // Optional. A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
    ChatBoost ChatBoostUpdated `json:"chat_boost,omitempty"`
    // Optional. A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
    RemovedChatBoost ChatBoostRemoved `json:"removed_chat_boost,omitempty"`
    
}


// https://core.telegram.org/bots/api#getupdates
//
// Use this method to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
//
// Notes1. This method will not work if an outgoing webhook is set up.2. In order to avoid getting duplicate updates, recalculate offset after each server response.
type GetUpdates struct {
    // Required: Optional
    //
    // Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will be forgotten.
    Offset int `json:"offset,omitempty"`
    // Required: Optional
    //
    // Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
    Limit int `json:"limit,omitempty"`
    // Required: Optional
    //
    // Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
    Timeout int `json:"timeout,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
    AllowedUpdates []string `json:"allowed_updates,omitempty"`
    
}

func (m *GetUpdates) Method() string {
	return "getUpdates"
}

// https://core.telegram.org/bots/api#setwebhook
//
// Use this method to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
//
// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as content.
//
// Notes1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.3. Ports currently supported for webhooks: 443, 80, 88, 8443.
//
// If you're having any trouble setting up webhooks, please check out this amazing guide to webhooks.
type SetWebhook struct {
    // Required: Yes
    //
    // HTTPS URL to send updates to. Use an empty string to remove webhook integration
    Url string `json:"url,omitempty"`
    // Required: Optional
    //
    // Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
    Certificate InputFile `json:"certificate,omitempty"`
    // Required: Optional
    //
    // The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
    IpAddress string `json:"ip_address,omitempty"`
    // Required: Optional
    //
    // The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
    MaxConnections int `json:"max_connections,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
    AllowedUpdates []string `json:"allowed_updates,omitempty"`
    // Required: Optional
    //
    // Pass True to drop all pending updates
    DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
    // Required: Optional
    //
    // A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
    SecretToken string `json:"secret_token,omitempty"`
    
}

func (m *SetWebhook) Method() string {
	return "setWebhook"
}

// https://core.telegram.org/bots/api#deletewebhook
//
// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
type DeleteWebhook struct {
    // Required: Optional
    //
    // Pass True to drop all pending updates
    DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
    
}

func (m *DeleteWebhook) Method() string {
	return "deleteWebhook"
}

// https://core.telegram.org/bots/api#getwebhookinfo
//
// Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
type GetWebhookInfo struct {
    
}

func (m *GetWebhookInfo) Method() string {
	return "getWebhookInfo"
}

// https://core.telegram.org/bots/api#webhookinfo
//
// Describes the current status of a webhook.
type WebhookInfo struct {
    // Webhook URL, may be empty if webhook is not set up
    Url string `json:"url,omitempty"`
    // True, if a custom certificate was provided for webhook certificate checks
    HasCustomCertificate bool `json:"has_custom_certificate,omitempty"`
    // Number of updates awaiting delivery
    PendingUpdateCount int `json:"pending_update_count,omitempty"`
    // Optional. Currently used webhook IP address
    IpAddress string `json:"ip_address,omitempty"`
    // Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
    LastErrorDate int `json:"last_error_date,omitempty"`
    // Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
    LastErrorMessage string `json:"last_error_message,omitempty"`
    // Optional. Unix time of the most recent error that happened when trying to synchronize available updates with Telegram datacenters
    LastSynchronizationErrorDate int `json:"last_synchronization_error_date,omitempty"`
    // Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
    MaxConnections int `json:"max_connections,omitempty"`
    // Optional. A list of update types the bot is subscribed to. Defaults to all update types except chat_member
    AllowedUpdates []string `json:"allowed_updates,omitempty"`
    
}

