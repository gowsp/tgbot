package tgbot

// https://core.telegram.org/bots/api#sendinvoice
//
// Use this method to send invoices. On success, the sent Message is returned.
type SendInvoice struct {
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
    // Product name, 1-32 characters
    Title string `json:"title,omitempty"`
    // Required: Yes
    //
    // Product description, 1-255 characters
    Description string `json:"description,omitempty"`
    // Required: Yes
    //
    // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
    Payload string `json:"payload,omitempty"`
    // Required: Yes
    //
    // Payment provider token, obtained via @BotFather
    ProviderToken string `json:"provider_token,omitempty"`
    // Required: Yes
    //
    // Three-letter ISO 4217 currency code, see more on currencies
    Currency string `json:"currency,omitempty"`
    // Required: Yes
    //
    // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
    Prices []LabeledPrice `json:"prices,omitempty"`
    // Required: Optional
    //
    // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
    MaxTipAmount int `json:"max_tip_amount,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
    SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`
    // Required: Optional
    //
    // Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button), with the value used as the start parameter
    StartParameter string `json:"start_parameter,omitempty"`
    // Required: Optional
    //
    // JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
    ProviderData string `json:"provider_data,omitempty"`
    // Required: Optional
    //
    // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
    PhotoUrl string `json:"photo_url,omitempty"`
    // Required: Optional
    //
    // Photo size in bytes
    PhotoSize int `json:"photo_size,omitempty"`
    // Required: Optional
    //
    // Photo width
    PhotoWidth int `json:"photo_width,omitempty"`
    // Required: Optional
    //
    // Photo height
    PhotoHeight int `json:"photo_height,omitempty"`
    // Required: Optional
    //
    // Pass True if you require the user's full name to complete the order
    NeedName bool `json:"need_name,omitempty"`
    // Required: Optional
    //
    // Pass True if you require the user's phone number to complete the order
    NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
    // Required: Optional
    //
    // Pass True if you require the user's email address to complete the order
    NeedEmail bool `json:"need_email,omitempty"`
    // Required: Optional
    //
    // Pass True if you require the user's shipping address to complete the order
    NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
    // Required: Optional
    //
    // Pass True if the user's phone number should be sent to provider
    SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
    // Required: Optional
    //
    // Pass True if the user's email address should be sent to provider
    SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
    // Required: Optional
    //
    // Pass True if the final price depends on the shipping method
    IsFlexible bool `json:"is_flexible,omitempty"`
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
    // A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
    ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    
}

func (m *SendInvoice) Method() string {
	return "sendInvoice"
}

// https://core.telegram.org/bots/api#createinvoicelink
//
// Use this method to create a link for an invoice. Returns the created invoice link as String on success.
type CreateInvoiceLink struct {
    // Required: Yes
    //
    // Product name, 1-32 characters
    Title string `json:"title,omitempty"`
    // Required: Yes
    //
    // Product description, 1-255 characters
    Description string `json:"description,omitempty"`
    // Required: Yes
    //
    // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
    Payload string `json:"payload,omitempty"`
    // Required: Yes
    //
    // Payment provider token, obtained via BotFather
    ProviderToken string `json:"provider_token,omitempty"`
    // Required: Yes
    //
    // Three-letter ISO 4217 currency code, see more on currencies
    Currency string `json:"currency,omitempty"`
    // Required: Yes
    //
    // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
    Prices []LabeledPrice `json:"prices,omitempty"`
    // Required: Optional
    //
    // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
    MaxTipAmount int `json:"max_tip_amount,omitempty"`
    // Required: Optional
    //
    // A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
    SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`
    // Required: Optional
    //
    // JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
    ProviderData string `json:"provider_data,omitempty"`
    // Required: Optional
    //
    // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
    PhotoUrl string `json:"photo_url,omitempty"`
    // Required: Optional
    //
    // Photo size in bytes
    PhotoSize int `json:"photo_size,omitempty"`
    // Required: Optional
    //
    // Photo width
    PhotoWidth int `json:"photo_width,omitempty"`
    // Required: Optional
    //
    // Photo height
    PhotoHeight int `json:"photo_height,omitempty"`
    // Required: Optional
    //
    // Pass True if you require the user's full name to complete the order
    NeedName bool `json:"need_name,omitempty"`
    // Required: Optional
    //
    // Pass True if you require the user's phone number to complete the order
    NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
    // Required: Optional
    //
    // Pass True if you require the user's email address to complete the order
    NeedEmail bool `json:"need_email,omitempty"`
    // Required: Optional
    //
    // Pass True if you require the user's shipping address to complete the order
    NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
    // Required: Optional
    //
    // Pass True if the user's phone number should be sent to the provider
    SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
    // Required: Optional
    //
    // Pass True if the user's email address should be sent to the provider
    SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
    // Required: Optional
    //
    // Pass True if the final price depends on the shipping method
    IsFlexible bool `json:"is_flexible,omitempty"`
    
}

func (m *CreateInvoiceLink) Method() string {
	return "createInvoiceLink"
}

// https://core.telegram.org/bots/api#answershippingquery
//
// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
type AnswerShippingQuery struct {
    // Required: Yes
    //
    // Unique identifier for the query to be answered
    ShippingQueryId string `json:"shipping_query_id,omitempty"`
    // Required: Yes
    //
    // Pass True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
    Ok bool `json:"ok,omitempty"`
    // Required: Optional
    //
    // Required if ok is True. A JSON-serialized array of available shipping options.
    ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
    // Required: Optional
    //
    // Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
    ErrorMessage string `json:"error_message,omitempty"`
    
}

func (m *AnswerShippingQuery) Method() string {
	return "answerShippingQuery"
}

// https://core.telegram.org/bots/api#answerprecheckoutquery
//
// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
type AnswerPreCheckoutQuery struct {
    // Required: Yes
    //
    // Unique identifier for the query to be answered
    PreCheckoutQueryId string `json:"pre_checkout_query_id,omitempty"`
    // Required: Yes
    //
    // Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
    Ok bool `json:"ok,omitempty"`
    // Required: Optional
    //
    // Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
    ErrorMessage string `json:"error_message,omitempty"`
    
}

func (m *AnswerPreCheckoutQuery) Method() string {
	return "answerPreCheckoutQuery"
}

// https://core.telegram.org/bots/api#labeledprice
//
// This object represents a portion of the price for goods or services.
type LabeledPrice struct {
    // Portion label
    Label string `json:"label,omitempty"`
    // Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    Amount int `json:"amount,omitempty"`
    
}


// https://core.telegram.org/bots/api#invoice
//
// This object contains basic information about an invoice.
type Invoice struct {
    // Product name
    Title string `json:"title,omitempty"`
    // Product description
    Description string `json:"description,omitempty"`
    // Unique bot deep-linking parameter that can be used to generate this invoice
    StartParameter string `json:"start_parameter,omitempty"`
    // Three-letter ISO 4217 currency code
    Currency string `json:"currency,omitempty"`
    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    TotalAmount int `json:"total_amount,omitempty"`
    
}


// https://core.telegram.org/bots/api#shippingaddress
//
// This object represents a shipping address.
type ShippingAddress struct {
    // Two-letter ISO 3166-1 alpha-2 country code
    CountryCode string `json:"country_code,omitempty"`
    // State, if applicable
    State string `json:"state,omitempty"`
    // City
    City string `json:"city,omitempty"`
    // First line for the address
    StreetLine1 string `json:"street_line1,omitempty"`
    // Second line for the address
    StreetLine2 string `json:"street_line2,omitempty"`
    // Address post code
    PostCode string `json:"post_code,omitempty"`
    
}


// https://core.telegram.org/bots/api#orderinfo
//
// This object represents information about an order.
type OrderInfo struct {
    // Optional. User name
    Name string `json:"name,omitempty"`
    // Optional. User's phone number
    PhoneNumber string `json:"phone_number,omitempty"`
    // Optional. User email
    Email string `json:"email,omitempty"`
    // Optional. User shipping address
    ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
    
}


// https://core.telegram.org/bots/api#shippingoption
//
// This object represents one shipping option.
type ShippingOption struct {
    // Shipping option identifier
    Id string `json:"id,omitempty"`
    // Option title
    Title string `json:"title,omitempty"`
    // List of price portions
    Prices []LabeledPrice `json:"prices,omitempty"`
    
}


// https://core.telegram.org/bots/api#successfulpayment
//
// This object contains basic information about a successful payment.
type SuccessfulPayment struct {
    // Three-letter ISO 4217 currency code
    Currency string `json:"currency,omitempty"`
    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    TotalAmount int `json:"total_amount,omitempty"`
    // Bot specified invoice payload
    InvoicePayload string `json:"invoice_payload,omitempty"`
    // Optional. Identifier of the shipping option chosen by the user
    ShippingOptionId string `json:"shipping_option_id,omitempty"`
    // Optional. Order information provided by the user
    OrderInfo OrderInfo `json:"order_info,omitempty"`
    // Telegram payment identifier
    TelegramPaymentChargeId string `json:"telegram_payment_charge_id,omitempty"`
    // Provider payment identifier
    ProviderPaymentChargeId string `json:"provider_payment_charge_id,omitempty"`
    
}


// https://core.telegram.org/bots/api#shippingquery
//
// This object contains information about an incoming shipping query.
type ShippingQuery struct {
    // Unique query identifier
    Id string `json:"id,omitempty"`
    // User who sent the query
    From User `json:"from,omitempty"`
    // Bot specified invoice payload
    InvoicePayload string `json:"invoice_payload,omitempty"`
    // User specified shipping address
    ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
    
}


// https://core.telegram.org/bots/api#precheckoutquery
//
// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
    // Unique query identifier
    Id string `json:"id,omitempty"`
    // User who sent the query
    From User `json:"from,omitempty"`
    // Three-letter ISO 4217 currency code
    Currency string `json:"currency,omitempty"`
    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    TotalAmount int `json:"total_amount,omitempty"`
    // Bot specified invoice payload
    InvoicePayload string `json:"invoice_payload,omitempty"`
    // Optional. Identifier of the shipping option chosen by the user
    ShippingOptionId string `json:"shipping_option_id,omitempty"`
    // Optional. Order information provided by the user
    OrderInfo OrderInfo `json:"order_info,omitempty"`
    
}

