package sdk

// EMV struct contains constants used for encoding and decoding QR codes
// for transactions supported by the Bakong app.
type EMV struct {
	// Default QR Code Types
	DefaultDynamicQR string
	DefaultStaticQR  string

	// Currency Codes
	TransactionCurrencyUSD string
	TransactionCurrencyKHR string
	TransactionCurrency    string

	// Payload and Point of Initiation
	PayloadFormatIndicator        string
	DefaultPayloadFormatIndicator string
	PointOfInitiationMethod       string

	// Merchant Information
	MerchantName                string
	MerchantCity                string
	DefaultMerchantCity         string
	MerchantCategoryCode        string
	DefaultMerchantCategoryCode string

	// QR Code Identifiers
	StaticQR                             string
	DynamicQR                            string
	MerchantAccountInformationIndividual string
	MerchantAccountInformationMerchant   string

	// Transaction Details
	TransactionAmount        string
	DefaultTransactionAmount string
	CountryCode              string
	DefaultCountryCode       string

	// Additional Data Tags
	AdditionalDataTag                   string
	BillNumberTag                       string
	AdditionDataFieldMobileNumber       string
	StoreLabel                          string
	TerminalLabel                       string
	PurposeOfTransaction                string
	TimestampTag                        string
	MerchantInformationLanguageTemplate string

	// Language Preferences
	LanguagePreference              string
	MerchantNameAlternativeLanguage string
	MerchantCityAlternativeLanguage string

	// UnionPay Specific
	UnionPayMerchantAccount string

	// CRC Tag
	CRC           string
	CRCLength     string
	DefaultCRCTag string

	// Invalid Length Constraints
	InvalidLengthKHQR                            int
	InvalidLengthMerchantName                    int
	InvalidLengthBakongAccount                   int
	InvalidLengthAmount                          int
	InvalidLengthCountryCode                     int
	InvalidLengthMerchantCategoryCode            int
	InvalidLengthMerchantCity                    int
	InvalidLengthTimestamp                       int
	InvalidLengthTransactionAmount               int
	InvalidLengthTransactionCurrency             int
	InvalidLengthBillNumber                      int
	InvalidLengthStoreLabel                      int
	InvalidLengthTerminalLabel                   int
	InvalidLengthPurposeOfTransaction            int
	InvalidLengthMerchantID                      int
	InvalidLengthAcquiringBank                   int
	InvalidLengthMobileNumber                    int
	InvalidLengthAccountInformation              int
	InvalidLengthMerchantNameLanguageTemplate    int
	InvalidLengthUPIMerchant                     int
	InvalidLengthLanguagePreference              int
	InvalidLengthMerchantNameAlternativeLanguage int
	InvalidLengthMerchantCityAlternativeLanguage int
}

// EMV creates and initializes a new EMV instance with default values.
func NewEMV() *EMV {
	return &EMV{
		// Default QR Code Types
		DefaultDynamicQR: "010212",
		DefaultStaticQR:  "010211",

		// Currency Codes
		TransactionCurrencyUSD: "840", // USD
		TransactionCurrencyKHR: "116", // KHR
		TransactionCurrency:    "53",

		// Payload and Point of Initiation
		PayloadFormatIndicator:        "00",
		DefaultPayloadFormatIndicator: "01",
		PointOfInitiationMethod:       "01",

		// Merchant Information
		MerchantName:                "59",
		MerchantCity:                "60",
		DefaultMerchantCity:         "Phnom Penh",
		MerchantCategoryCode:        "52",
		DefaultMerchantCategoryCode: "5999",

		// QR Code Identifiers
		StaticQR:                             "11",
		DynamicQR:                            "12",
		MerchantAccountInformationIndividual: "29",
		MerchantAccountInformationMerchant:   "30",

		// Transaction Details
		TransactionAmount:        "54",
		DefaultTransactionAmount: "0",
		CountryCode:              "58",
		DefaultCountryCode:       "KH",

		// Additional Data Tags
		AdditionalDataTag:                   "62",
		BillNumberTag:                       "01",
		AdditionDataFieldMobileNumber:       "02",
		StoreLabel:                          "03",
		TerminalLabel:                       "07",
		PurposeOfTransaction:                "08",
		TimestampTag:                        "99",
		MerchantInformationLanguageTemplate: "64",

		// Language Preferences
		LanguagePreference:              "00",
		MerchantNameAlternativeLanguage: "01",
		MerchantCityAlternativeLanguage: "02",

		// UnionPay Specific
		UnionPayMerchantAccount: "15",

		// CRC Tag
		CRC:           "63",
		CRCLength:     "04",
		DefaultCRCTag: "6304",

		// Invalid Length Constraints
		InvalidLengthKHQR:                            2,
		InvalidLengthMerchantName:                    25,
		InvalidLengthBakongAccount:                   32,
		InvalidLengthAmount:                          13,
		InvalidLengthCountryCode:                     3,
		InvalidLengthMerchantCategoryCode:            4,
		InvalidLengthMerchantCity:                    15,
		InvalidLengthTimestamp:                       13,
		InvalidLengthTransactionAmount:               14,
		InvalidLengthTransactionCurrency:             3,
		InvalidLengthBillNumber:                      25,
		InvalidLengthStoreLabel:                      25,
		InvalidLengthTerminalLabel:                   25,
		InvalidLengthPurposeOfTransaction:            25,
		InvalidLengthMerchantID:                      32,
		InvalidLengthAcquiringBank:                   32,
		InvalidLengthMobileNumber:                    25,
		InvalidLengthAccountInformation:              32,
		InvalidLengthMerchantNameLanguageTemplate:    99,
		InvalidLengthUPIMerchant:                     99,
		InvalidLengthLanguagePreference:              2,
		InvalidLengthMerchantNameAlternativeLanguage: 25,
		InvalidLengthMerchantCityAlternativeLanguage: 15,
	}
}
