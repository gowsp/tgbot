package tgbot

// https://core.telegram.org/bots/api#passportdata
//
// Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
    // Array with information about documents and other Telegram Passport elements that was shared with the bot
    Data []EncryptedPassportElement `json:"data,omitempty"`
    // Encrypted credentials required to decrypt the data
    Credentials EncryptedCredentials `json:"credentials,omitempty"`
    
}


// https://core.telegram.org/bots/api#passportfile
//
// This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id,omitempty"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id,omitempty"`
    // File size in bytes
    FileSize int `json:"file_size,omitempty"`
    // Unix time when the file was uploaded
    FileDate int `json:"file_date,omitempty"`
    
}


// https://core.telegram.org/bots/api#encryptedpassportelement
//
// Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
    // Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
    Type string `json:"type,omitempty"`
    // Optional. Base64-encoded encrypted Telegram Passport element data provided by the user; available only for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials.
    Data string `json:"data,omitempty"`
    // Optional. User's verified phone number; available only for “phone_number” type
    PhoneNumber string `json:"phone_number,omitempty"`
    // Optional. User's verified email address; available only for “email” type
    Email string `json:"email,omitempty"`
    // Optional. Array of encrypted files with documents provided by the user; available only for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
    Files []PassportFile `json:"files,omitempty"`
    // Optional. Encrypted file with the front side of the document, provided by the user; available only for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
    FrontSide PassportFile `json:"front_side,omitempty"`
    // Optional. Encrypted file with the reverse side of the document, provided by the user; available only for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
    ReverseSide PassportFile `json:"reverse_side,omitempty"`
    // Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available if requested for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
    Selfie PassportFile `json:"selfie,omitempty"`
    // Optional. Array of encrypted files with translated versions of documents provided by the user; available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
    Translation []PassportFile `json:"translation,omitempty"`
    // Base64-encoded element hash for using in PassportElementErrorUnspecified
    Hash string `json:"hash,omitempty"`
    
}


// https://core.telegram.org/bots/api#encryptedcredentials
//
// Describes data required for decrypting and authenticating EncryptedPassportElement. See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
    // Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
    Data string `json:"data,omitempty"`
    // Base64-encoded data hash for data authentication
    Hash string `json:"hash,omitempty"`
    // Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
    Secret string `json:"secret,omitempty"`
    
}


// https://core.telegram.org/bots/api#setpassportdataerrors
//
// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
//
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
type SetPassportDataErrors struct {
    // Required: Yes
    //
    // User identifier
    UserId int `json:"user_id,omitempty"`
    // Required: Yes
    //
    // A JSON-serialized array describing the errors
    Errors []PassportElementError `json:"errors,omitempty"`
    
}

func (m *SetPassportDataErrors) Method() string {
	return "setPassportDataErrors"
}

// https://core.telegram.org/bots/api#passportelementerror
//
// This object represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of:
//
//   - PassportElementErrorDataField
//
//   - PassportElementErrorFrontSide
//
//   - PassportElementErrorReverseSide
//
//   - PassportElementErrorSelfie
//
//   - PassportElementErrorFile
//
//   - PassportElementErrorFiles
//
//   - PassportElementErrorTranslationFile
//
//   - PassportElementErrorTranslationFiles
//
//   - PassportElementErrorUnspecified
type PassportElementError struct {
    
}


// https://core.telegram.org/bots/api#passportelementerrordatafield
//
// Represents an issue in one of the data fields that was provided by the user. The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
    // Error source, must be data
    Source string `json:"source,omitempty"`
    // The section of the user's Telegram Passport which has the error, one of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
    Type string `json:"type,omitempty"`
    // Name of the data field which has the error
    FieldName string `json:"field_name,omitempty"`
    // Base64-encoded data hash
    DataHash string `json:"data_hash,omitempty"`
    // Error message
    Message string `json:"message,omitempty"`
    
}


// https://core.telegram.org/bots/api#passportelementerrorfrontside
//
// Represents an issue with the front side of a document. The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
    // Error source, must be front_side
    Source string `json:"source,omitempty"`
    // The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
    Type string `json:"type,omitempty"`
    // Base64-encoded hash of the file with the front side of the document
    FileHash string `json:"file_hash,omitempty"`
    // Error message
    Message string `json:"message,omitempty"`
    
}


// https://core.telegram.org/bots/api#passportelementerrorreverseside
//
// Represents an issue with the reverse side of a document. The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
    // Error source, must be reverse_side
    Source string `json:"source,omitempty"`
    // The section of the user's Telegram Passport which has the issue, one of “driver_license”, “identity_card”
    Type string `json:"type,omitempty"`
    // Base64-encoded hash of the file with the reverse side of the document
    FileHash string `json:"file_hash,omitempty"`
    // Error message
    Message string `json:"message,omitempty"`
    
}


// https://core.telegram.org/bots/api#passportelementerrorselfie
//
// Represents an issue with the selfie with a document. The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
    // Error source, must be selfie
    Source string `json:"source,omitempty"`
    // The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
    Type string `json:"type,omitempty"`
    // Base64-encoded hash of the file with the selfie
    FileHash string `json:"file_hash,omitempty"`
    // Error message
    Message string `json:"message,omitempty"`
    
}


// https://core.telegram.org/bots/api#passportelementerrorfile
//
// Represents an issue with a document scan. The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
    // Error source, must be file
    Source string `json:"source,omitempty"`
    // The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
    Type string `json:"type,omitempty"`
    // Base64-encoded file hash
    FileHash string `json:"file_hash,omitempty"`
    // Error message
    Message string `json:"message,omitempty"`
    
}


// https://core.telegram.org/bots/api#passportelementerrorfiles
//
// Represents an issue with a list of scans. The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
    // Error source, must be files
    Source string `json:"source,omitempty"`
    // The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
    Type string `json:"type,omitempty"`
    // List of base64-encoded file hashes
    FileHashes []string `json:"file_hashes,omitempty"`
    // Error message
    Message string `json:"message,omitempty"`
    
}


// https://core.telegram.org/bots/api#passportelementerrortranslationfile
//
// Represents an issue with one of the files that constitute the translation of a document. The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
    // Error source, must be translation_file
    Source string `json:"source,omitempty"`
    // Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
    Type string `json:"type,omitempty"`
    // Base64-encoded file hash
    FileHash string `json:"file_hash,omitempty"`
    // Error message
    Message string `json:"message,omitempty"`
    
}


// https://core.telegram.org/bots/api#passportelementerrortranslationfiles
//
// Represents an issue with the translated version of a document. The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
    // Error source, must be translation_files
    Source string `json:"source,omitempty"`
    // Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
    Type string `json:"type,omitempty"`
    // List of base64-encoded file hashes
    FileHashes []string `json:"file_hashes,omitempty"`
    // Error message
    Message string `json:"message,omitempty"`
    
}


// https://core.telegram.org/bots/api#passportelementerrorunspecified
//
// Represents an issue in an unspecified place. The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
    // Error source, must be unspecified
    Source string `json:"source,omitempty"`
    // Type of element of the user's Telegram Passport which has the issue
    Type string `json:"type,omitempty"`
    // Base64-encoded element hash
    ElementHash string `json:"element_hash,omitempty"`
    // Error message
    Message string `json:"message,omitempty"`
    
}

