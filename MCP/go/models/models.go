package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Networktransactionreferenceresponse represents the Networktransactionreferenceresponse schema from the OpenAPI specification
type Networktransactionreferenceresponse struct {
	Network string `json:"network,omitempty"` // The card network or brand. Applies to credit, debit, gift, and payment cards.
	Time string `json:"time,omitempty"` // The date and time, in [Internet date and time format](https://tools.ietf.org/html/rfc3339#section-5.6). Seconds are required while fractional seconds are optional.<blockquote><strong>Note:</strong> The regular expression provides guidance but does not reject all invalid dates.</blockquote>
	Date string `json:"date,omitempty"` // The date that the transaction was authorized by the scheme. This field may not be returned for all networks. MasterCard refers to this field as "BankNet reference date.
	Id string `json:"id"` // Transaction reference id returned by the scheme. For Visa and Amex, this is the "Tran id" field in response. For MasterCard, this is the "BankNet reference id" field in response. For Discover, this is the "NRID" field in response. The pattern we expect for this field from Visa/Amex/CB/Discover is numeric, Mastercard/BNPP is alphanumeric and Paysecure is alphanumeric with special character -.
}

// Setuptokenrequest represents the Setuptokenrequest schema from the OpenAPI specification
type Setuptokenrequest struct {
	Customer map[string]interface{} `json:"customer,omitempty"` // This object defines a customer in your system. Use it to manage customer profiles, save payment methods and contact details.
	Payment_source interface{} `json:"payment_source"` // The payment method to vault with the instrument details.
}

// Linkdescription represents the Linkdescription schema from the OpenAPI specification
type Linkdescription struct {
	Method string `json:"method,omitempty"` // The HTTP method required to make the related call.
	Rel string `json:"rel"` // The [link relation type](https://tools.ietf.org/html/rfc5988#section-4), which serves as an ID for a link that unambiguously describes the semantics of the link. See [Link Relations](https://www.iana.org/assignments/link-relations/link-relations.xhtml).
	Href string `json:"href"` // The complete target URL. To make the related call, combine the method with this [URI Template-formatted](https://tools.ietf.org/html/rfc6570) link. For pre-processing, include the `$`, `(`, and `)` characters. The `href` is the key HATEOAS component that links a completed call with a subsequent call.
}

// Error409 represents the Error409 schema from the OpenAPI specification
type Error409 struct {
	Details []Errordetails `json:"details,omitempty"`
	Links []Errorlinkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](https://en.wikipedia.org/wiki/HATEOAS).
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Debug_id string `json:"debug_id,omitempty"` // The PayPal internal ID. Used for correlation purposes.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Card map[string]interface{} `json:"card,omitempty"` // A Resource representing a request to vault a Card.
	Token Tokenidrequest `json:"token,omitempty"` // The Tokenized Payment Source representing a Request to Vault a Token.
}

// Paymenttokenresponse represents the Paymenttokenresponse schema from the OpenAPI specification
type Paymenttokenresponse struct {
	Id string `json:"id,omitempty"` // The PayPal-generated ID for the vault token.
	Links []Linkdescription `json:"links,omitempty"` // An array of related [HATEOAS links](/api/rest/responses/#hateoas).
	Payment_source Paymentsource `json:"payment_source,omitempty"` // The vaulted payment method details.
	Customer Customer `json:"customer,omitempty"` // This object defines a customer in your system. Use it to manage customer profiles, save payment methods and contact details.
}

// Paymentsource represents the Paymentsource schema from the OpenAPI specification
type Paymentsource struct {
	Apple_pay Applepaypaymenttokenresponse `json:"apple_pay,omitempty"` // A resource representing a response for Apple Pay.
	Card Cardresponse `json:"card,omitempty"` // Full representation of a Card Payment Token including network token.
	Paypal Paypalwalletresponse `json:"paypal,omitempty"`
	Venmo Venmoresponse `json:"venmo,omitempty"`
}

// Cardverificationdetails represents the Cardverificationdetails schema from the OpenAPI specification
type Cardverificationdetails struct {
	Date string `json:"date,omitempty"` // DEPRECATED. This field is DEPRECATED. Please find the date data in the 'date' field under the 'network_transaction_reference' object instead of the 'verification' object.
	Network string `json:"network,omitempty"` // The card network or brand. Applies to credit, debit, gift, and payment cards.
	Network_transaction_id string `json:"network_transaction_id,omitempty"` // DEPRECATED. This field is DEPRECATED. Please find the network transaction id data in the 'id' field under the 'network_transaction_reference' object instead of the 'verification' object.
	Processor_response map[string]interface{} `json:"processor_response,omitempty"` // The processor response information for payment requests, such as direct credit card transactions.
	Three_d_secure interface{} `json:"three_d_secure,omitempty"` // DEPRECATED. This field is DEPRECATED. Please find the 3D secure authentication data in the 'three_d_secure' object under the 'authentication_result' object instead of the 'verification' object.
	Time string `json:"time,omitempty"` // The date and time, in [Internet date and time format](https://tools.ietf.org/html/rfc3339#section-5.6). Seconds are required while fractional seconds are optional.<blockquote><strong>Note:</strong> The regular expression provides guidance but does not reject all invalid dates.</blockquote>
	Amount Money `json:"amount,omitempty"` // The currency and amount for a financial transaction, such as a balance or payment due.
}

// Pricingscheme represents the Pricingscheme schema from the OpenAPI specification
type Pricingscheme struct {
	Price Money `json:"price,omitempty"` // The currency and amount for a financial transaction, such as a balance or payment due.
	Pricing_model string `json:"pricing_model"` // The pricing model for the billing cycle.
	Reload_threshold_amount Money `json:"reload_threshold_amount,omitempty"` // The currency and amount for a financial transaction, such as a balance or payment due.
}

// Customervaultpaymenttokensresponse represents the Customervaultpaymenttokensresponse schema from the OpenAPI specification
type Customervaultpaymenttokensresponse struct {
	Total_items int `json:"total_items,omitempty"` // Total number of items.
	Total_pages int `json:"total_pages,omitempty"` // Total number of pages.
	Customer map[string]interface{} `json:"customer,omitempty"` // This object defines a customer in your system. Use it to manage customer profiles, save payment methods and contact details.
	Links []Linkdescription `json:"links,omitempty"` // An array of related [HATEOAS links](/api/rest/responses/#hateoas).
	Payment_tokens []Paymenttokenresponse `json:"payment_tokens,omitempty"`
}

// Error500 represents the Error500 schema from the OpenAPI specification
type Error500 struct {
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Debug_id string `json:"debug_id,omitempty"` // The PayPal internal ID. Used for correlation purposes.
	Links []Errorlinkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](https://en.wikipedia.org/wiki/HATEOAS).
}

// Payerbase represents the Payerbase schema from the OpenAPI specification
type Payerbase struct {
	Email_address string `json:"email_address,omitempty"` // The internationalized email address.<blockquote><strong>Note:</strong> Up to 64 characters are allowed before and 255 characters are allowed after the <code>@</code> sign. However, the generally accepted maximum length for an email address is 254 characters. The pattern verifies that an unquoted <code>@</code> sign exists.</blockquote>
	Payer_id string `json:"payer_id,omitempty"` // The account identifier for a PayPal account.
}

// Threedsecureauthenticationresponse represents the Threedsecureauthenticationresponse schema from the OpenAPI specification
type Threedsecureauthenticationresponse struct {
	Authentication_status string `json:"authentication_status,omitempty"` // Transactions status result identifier. The outcome of the issuer's authentication.
	Enrollment_status string `json:"enrollment_status,omitempty"` // Status of Authentication eligibility.
}

// Bindetails represents the Bindetails schema from the OpenAPI specification
type Bindetails struct {
	Bin string `json:"bin,omitempty"` // The Bank Identification Number (BIN) signifies the number that is being used to identify the granular level details (except the PII information) of the card.
	Bin_country_code string `json:"bin_country_code,omitempty"` // The [two-character ISO 3166-1 code](/api/rest/reference/country-codes/) that identifies the country or region.<blockquote><strong>Note:</strong> The country code for Great Britain is <code>GB</code> and not <code>UK</code> as used in the top-level domain names for that country. Use the `C2` country code for China worldwide for comparable uncontrolled price (CUP) method, bank card, and cross-border transactions.</blockquote>
	Issuing_bank string `json:"issuing_bank,omitempty"` // The issuer of the card instrument.
	Products []string `json:"products,omitempty"` // The type of card product assigned to the BIN by the issuer. These values are defined by the issuer and may change over time. Some examples include: PREPAID_GIFT, CONSUMER, CORPORATE.
}

// Paymenttokenrequest represents the Paymenttokenrequest schema from the OpenAPI specification
type Paymenttokenrequest struct {
	Payment_source GeneratedType `json:"payment_source"` // The payment method to vault with the instrument details.
	Customer map[string]interface{} `json:"customer,omitempty"` // This object defines a customer in your system. Use it to manage customer profiles, save payment methods and contact details.
}

// Walletbase represents the Walletbase schema from the OpenAPI specification
type Walletbase struct {
	Shipping map[string]interface{} `json:"shipping,omitempty"` // The shipping details.
	Usage_pattern string `json:"usage_pattern,omitempty"` // Expected business/charge model for the billing agreement.
	Usage_type string `json:"usage_type,omitempty"` // The usage type associated with a digital wallet payment token.
	Customer_type string `json:"customer_type,omitempty"` // The customer type associated with a digital wallet payment token. This is to indicate whether the customer acting on the merchant / platform is either a business or a consumer.
	Description string `json:"description,omitempty"` // The description displayed to the consumer on the approval flow for a digital wallet, as well as on the merchant view of the payment token management experience. exp: PayPal.com.
	Permit_multiple_payment_tokens bool `json:"permit_multiple_payment_tokens,omitempty"` // Create multiple payment tokens for the same payer, merchant/platform combination. Use this when the customer has not logged in at merchant/platform. The payment token thus generated, can then also be used to create the customer account at merchant/platform. Use this also when multiple payment tokens are required for the same payer, different customer at merchant/platform. This helps to identify customers distinctly even though they may share the same PayPal account. This only applies to PayPal payment source.
}

// Venmoresponse represents the Venmoresponse schema from the OpenAPI specification
type Venmoresponse struct {
	Shipping map[string]interface{} `json:"shipping,omitempty"` // The shipping details.
	Usage_pattern string `json:"usage_pattern,omitempty"` // Expected business/charge model for the billing agreement.
	Usage_type string `json:"usage_type,omitempty"` // The usage type associated with a digital wallet payment token.
	Customer_type string `json:"customer_type,omitempty"` // The customer type associated with a digital wallet payment token. This is to indicate whether the customer acting on the merchant / platform is either a business or a consumer.
	Description string `json:"description,omitempty"` // The description displayed to the consumer on the approval flow for a digital wallet, as well as on the merchant view of the payment token management experience. exp: PayPal.com.
	Permit_multiple_payment_tokens bool `json:"permit_multiple_payment_tokens,omitempty"` // Create multiple payment tokens for the same payer, merchant/platform combination. Use this when the customer has not logged in at merchant/platform. The payment token thus generated, can then also be used to create the customer account at merchant/platform. Use this also when multiple payment tokens are required for the same payer, different customer at merchant/platform. This helps to identify customers distinctly even though they may share the same PayPal account. This only applies to PayPal payment source.
	Email_address string `json:"email_address,omitempty"` // The internationalized email address.<blockquote><strong>Note:</strong> Up to 64 characters are allowed before and 255 characters are allowed after the <code>@</code> sign. However, the generally accepted maximum length for an email address is 254 characters. The pattern verifies that an unquoted <code>@</code> sign exists.</blockquote>
	Payer_id string `json:"payer_id,omitempty"` // The account identifier for a PayPal account.
	Address map[string]interface{} `json:"address,omitempty"` // The portable international postal address. Maps to [AddressValidationMetadata](https://github.com/googlei18n/libaddressinput/wiki/AddressValidationMetadata) and HTML 5.1 [Autofilling form controls: the autocomplete attribute](https://www.w3.org/TR/html51/sec-forms.html#autofilling-form-controls-the-autocomplete-attribute).
	Name map[string]interface{} `json:"name,omitempty"` // The name of the party.
	Phone Phonewithtype `json:"phone,omitempty"` // The phone information.
	User_name string `json:"user_name,omitempty"` // The Venmo username, as chosen by the user.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Debug_id string `json:"debug_id"` // The PayPal internal ID. Used for correlation purposes.
	Details []GeneratedType `json:"details,omitempty"` // An array of additional details about the error.
	Links []Linkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](/api/rest/responses/#hateoas-links).
	Message string `json:"message"` // The message that describes the error.
	Name string `json:"name"` // The human-readable, unique name of the error.
}

// Phonewithtype represents the Phonewithtype schema from the OpenAPI specification
type Phonewithtype struct {
	Phone_number map[string]interface{} `json:"phone_number"` // The phone number, in its canonical international [E.164 numbering plan format](https://www.itu.int/rec/T-REC-E.164/en).
	Phone_type string `json:"phone_type,omitempty"` // The phone type.
}

// Paypalwalletresponse represents the Paypalwalletresponse schema from the OpenAPI specification
type Paypalwalletresponse struct {
	Description string `json:"description,omitempty"` // The description displayed to the consumer on the approval flow for a digital wallet, as well as on the merchant view of the payment token management experience. exp: PayPal.com.
	Permit_multiple_payment_tokens bool `json:"permit_multiple_payment_tokens,omitempty"` // Create multiple payment tokens for the same payer, merchant/platform combination. Use this when the customer has not logged in at merchant/platform. The payment token thus generated, can then also be used to create the customer account at merchant/platform. Use this also when multiple payment tokens are required for the same payer, different customer at merchant/platform. This helps to identify customers distinctly even though they may share the same PayPal account. This only applies to PayPal payment source.
	Shipping map[string]interface{} `json:"shipping,omitempty"` // The shipping details.
	Usage_pattern string `json:"usage_pattern,omitempty"` // Expected business/charge model for the billing agreement.
	Usage_type string `json:"usage_type,omitempty"` // The usage type associated with a digital wallet payment token.
	Customer_type string `json:"customer_type,omitempty"` // The customer type associated with a digital wallet payment token. This is to indicate whether the customer acting on the merchant / platform is either a business or a consumer.
	Email_address string `json:"email_address,omitempty"` // The internationalized email address.<blockquote><strong>Note:</strong> Up to 64 characters are allowed before and 255 characters are allowed after the <code>@</code> sign. However, the generally accepted maximum length for an email address is 254 characters. The pattern verifies that an unquoted <code>@</code> sign exists.</blockquote>
	Payer_id string `json:"payer_id,omitempty"` // The account identifier for a PayPal account.
	Name map[string]interface{} `json:"name,omitempty"` // The name of the party.
	Phone Phonewithtype `json:"phone,omitempty"` // The phone information.
	Address map[string]interface{} `json:"address,omitempty"` // The portable international postal address. Maps to [AddressValidationMetadata](https://github.com/googlei18n/libaddressinput/wiki/AddressValidationMetadata) and HTML 5.1 [Autofilling form controls: the autocomplete attribute](https://www.w3.org/TR/html51/sec-forms.html#autofilling-form-controls-the-autocomplete-attribute).
	Phone_number Phone `json:"phone_number,omitempty"` // The phone number, in its canonical international [E.164 numbering plan format](https://www.itu.int/rec/T-REC-E.164/en).
	Account_id string `json:"account_id,omitempty"` // The account identifier for a PayPal account.
}

// Networktransactionreferenceentity represents the Networktransactionreferenceentity schema from the OpenAPI specification
type Networktransactionreferenceentity struct {
	Network string `json:"network,omitempty"` // The card network or brand. Applies to credit, debit, gift, and payment cards.
	Time string `json:"time,omitempty"` // The date and time, in [Internet date and time format](https://tools.ietf.org/html/rfc3339#section-5.6). Seconds are required while fractional seconds are optional.<blockquote><strong>Note:</strong> The regular expression provides guidance but does not reject all invalid dates.</blockquote>
	Date string `json:"date,omitempty"` // The date that the transaction was authorized by the scheme. This field may not be returned for all networks. MasterCard refers to this field as "BankNet reference date.
	Id string `json:"id"` // Transaction reference id returned by the scheme. For Visa and Amex, this is the "Tran id" field in response. For MasterCard, this is the "BankNet reference id" field in response. For Discover, this is the "NRID" field in response. The pattern we expect for this field from Visa/Amex/CB/Discover is numeric, Mastercard/BNPP is alphanumeric and Paysecure is alphanumeric with special character -.
}

// Setuptokenresponse represents the Setuptokenresponse schema from the OpenAPI specification
type Setuptokenresponse struct {
	Links []Linkdescription `json:"links,omitempty"` // An array of related [HATEOAS links](/api/rest/responses/#hateoas).
	Payment_source interface{} `json:"payment_source,omitempty"` // The setup payment method details.
	Status string `json:"status,omitempty"` // The status of the payment token.
	Customer map[string]interface{} `json:"customer,omitempty"` // This object defines a customer in your system. Use it to manage customer profiles, save payment methods and contact details.
	Id string `json:"id,omitempty"` // The PayPal-generated ID for the vault token.
}

// Phone represents the Phone schema from the OpenAPI specification
type Phone struct {
	Country_code string `json:"country_code"` // The country calling code (CC), in its canonical international [E.164 numbering plan format](https://www.itu.int/rec/T-REC-E.164/en). The combined length of the CC and the national number must not be greater than 15 digits. The national number consists of a national destination code (NDC) and subscriber number (SN).
	Extension_number string `json:"extension_number,omitempty"` // The extension number.
	National_number string `json:"national_number"` // The national number, in its canonical international [E.164 numbering plan format](https://www.itu.int/rec/T-REC-E.164/en). The combined length of the country calling code (CC) and the national number must not be greater than 15 digits. The national number consists of a national destination code (NDC) and subscriber number (SN).
}

// Error400 represents the Error400 schema from the OpenAPI specification
type Error400 struct {
	Links []Errorlinkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](https://en.wikipedia.org/wiki/HATEOAS).
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Debug_id string `json:"debug_id,omitempty"` // The PayPal internal ID. Used for correlation purposes.
	Details []Errordetails `json:"details,omitempty"`
}

// Tokenidrequest represents the Tokenidrequest schema from the OpenAPI specification
type Tokenidrequest struct {
	Id string `json:"id"` // The PayPal-generated ID for the token.
	TypeField string `json:"type"` // The tokenization method that generated the ID.
}

// Errorlinkdescription represents the Errorlinkdescription schema from the OpenAPI specification
type Errorlinkdescription struct {
	Href string `json:"href"` // The complete target URL. To make the related call, combine the method with this [URI Template-formatted](https://tools.ietf.org/html/rfc6570) link. For pre-processing, include the `$`, `(`, and `)` characters. The `href` is the key HATEOAS component that links a completed call with a subsequent call.
	Method string `json:"method,omitempty"` // The HTTP method required to make the related call.
	Rel string `json:"rel"` // The [link relation type](https://tools.ietf.org/html/rfc5988#section-4), which serves as an ID for a link that unambiguously describes the semantics of the link. See [Link Relations](https://www.iana.org/assignments/link-relations/link-relations.xhtml).
}

// Error415 represents the Error415 schema from the OpenAPI specification
type Error415 struct {
	Links []Errorlinkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](https://en.wikipedia.org/wiki/HATEOAS).
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Debug_id string `json:"debug_id,omitempty"` // The PayPal internal ID. Used for correlation purposes.
	Details []Errordetails `json:"details,omitempty"`
}

// Cardresponseentity represents the Cardresponseentity schema from the OpenAPI specification
type Cardresponseentity struct {
	TypeField string `json:"type,omitempty"` // Type of card. i.e Credit, Debit and so on.
	Verification Cardverificationdetails `json:"verification,omitempty"` // Card Verification details including the authorization details and 3D SECURE details.
	Authentication_result map[string]interface{} `json:"authentication_result,omitempty"` // Results of Authentication such as 3D Secure.
	Billing_address map[string]interface{} `json:"billing_address,omitempty"` // Address request details.
	Expiry string `json:"expiry,omitempty"` // The year and month, in ISO-8601 `YYYY-MM` date format. See [Internet date and time format](https://tools.ietf.org/html/rfc3339#section-5.6).
	Verification_status string `json:"verification_status,omitempty"` // Verification status of Card.
	Brand string `json:"brand,omitempty"` // The card network or brand. Applies to credit, debit, gift, and payment cards.
	Name string `json:"name,omitempty"` // The card holder's name as it appears on the card.
	Bin_details Bindetails `json:"bin_details,omitempty"` // Bank Identification Number (BIN) details used to fund a payment.
	Last_digits string `json:"last_digits,omitempty"` // The last digits of the payment card.
	Network_transaction_reference Networktransactionreferenceresponse `json:"network_transaction_reference,omitempty"` // Previous network transaction reference including id in response.
}

// Error422 represents the Error422 schema from the OpenAPI specification
type Error422 struct {
	Links []Errorlinkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](https://en.wikipedia.org/wiki/HATEOAS).
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Debug_id string `json:"debug_id,omitempty"` // The PayPal internal ID. Used for correlation purposes.
	Details []Errordetails `json:"details,omitempty"`
}

// Venmorequest represents the Venmorequest schema from the OpenAPI specification
type Venmorequest struct {
	Permit_multiple_payment_tokens bool `json:"permit_multiple_payment_tokens,omitempty"` // Create multiple payment tokens for the same payer, merchant/platform combination. Use this when the customer has not logged in at merchant/platform. The payment token thus generated, can then also be used to create the customer account at merchant/platform. Use this also when multiple payment tokens are required for the same payer, different customer at merchant/platform. This helps to identify customers distinctly even though they may share the same PayPal account. This only applies to PayPal payment source.
	Shipping map[string]interface{} `json:"shipping,omitempty"` // The shipping details.
	Usage_pattern string `json:"usage_pattern,omitempty"` // Expected business/charge model for the billing agreement.
	Usage_type string `json:"usage_type,omitempty"` // The usage type associated with a digital wallet payment token.
	Customer_type string `json:"customer_type,omitempty"` // The customer type associated with a digital wallet payment token. This is to indicate whether the customer acting on the merchant / platform is either a business or a consumer.
	Description string `json:"description,omitempty"` // The description displayed to the consumer on the approval flow for a digital wallet, as well as on the merchant view of the payment token management experience. exp: PayPal.com.
	Experience_context map[string]interface{} `json:"experience_context,omitempty"` // Customizes the Vault creation flow experience for your customers.
}

// Customer represents the Customer schema from the OpenAPI specification
type Customer struct {
	Merchant_customer_id string `json:"merchant_customer_id,omitempty"` // Merchants and partners may already have a data-store where their customer information is persisted. Use merchant_customer_id to associate the PayPal-generated customer.id to your representation of a customer.
	Id string `json:"id,omitempty"` // The unique ID for a customer generated by PayPal.
}

// Error503 represents the Error503 schema from the OpenAPI specification
type Error503 struct {
	Debug_id string `json:"debug_id,omitempty"` // The PayPal internal ID. Used for correlation purposes.
	Links []Errorlinkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](https://en.wikipedia.org/wiki/HATEOAS).
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
}

// Error401 represents the Error401 schema from the OpenAPI specification
type Error401 struct {
	Debug_id string `json:"debug_id,omitempty"` // The PayPal internal ID. Used for correlation purposes.
	Details []Errordetails `json:"details,omitempty"`
	Links []Errorlinkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](https://en.wikipedia.org/wiki/HATEOAS).
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
}

// Money represents the Money schema from the OpenAPI specification
type Money struct {
	Currency_code string `json:"currency_code"` // The [three-character ISO-4217 currency code](/api/rest/reference/currency-codes/) that identifies the currency.
	Value string `json:"value"` // The value, which might be:<ul><li>An integer for currencies like `JPY` that are not typically fractional.</li><li>A decimal fraction for currencies like `TND` that are subdivided into thousandths.</li></ul>For the required number of decimal places for a currency code, see [Currency Codes](/api/rest/reference/currency-codes/).
}

// Plan represents the Plan schema from the OpenAPI specification
type Plan struct {
	Name string `json:"name,omitempty"` // Name of the recurring plan.
	One_time_charges Onetimecharges `json:"one_time_charges"` // The one-time charge info at the time of checkout.
	Product Productoverride `json:"product,omitempty"`
	Billing_cycles []Billingcycle `json:"billing_cycles"` // An array of billing cycles for trial billing and regular billing. A plan can have at most two trial cycles and only one regular cycle.
}

// Cardresponse represents the Cardresponse schema from the OpenAPI specification
type Cardresponse struct {
	Bin_details Bindetails `json:"bin_details,omitempty"` // Bank Identification Number (BIN) details used to fund a payment.
	Last_digits string `json:"last_digits,omitempty"` // The last digits of the payment card.
	Network_transaction_reference Networktransactionreferenceresponse `json:"network_transaction_reference,omitempty"` // Previous network transaction reference including id in response.
	TypeField string `json:"type,omitempty"` // Type of card. i.e Credit, Debit and so on.
	Verification Cardverificationdetails `json:"verification,omitempty"` // Card Verification details including the authorization details and 3D SECURE details.
	Authentication_result map[string]interface{} `json:"authentication_result,omitempty"` // Results of Authentication such as 3D Secure.
	Billing_address map[string]interface{} `json:"billing_address,omitempty"` // Address request details.
	Expiry string `json:"expiry,omitempty"` // The year and month, in ISO-8601 `YYYY-MM` date format. See [Internet date and time format](https://tools.ietf.org/html/rfc3339#section-5.6).
	Verification_status string `json:"verification_status,omitempty"` // Verification status of Card.
	Brand string `json:"brand,omitempty"` // The card network or brand. Applies to credit, debit, gift, and payment cards.
	Name string `json:"name,omitempty"` // The card holder's name as it appears on the card.
	Network_token Networktokenentity `json:"network_token,omitempty"`
}

// Error404 represents the Error404 schema from the OpenAPI specification
type Error404 struct {
	Details []Errordetails `json:"details,omitempty"`
	Links []Errorlinkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](https://en.wikipedia.org/wiki/HATEOAS).
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Debug_id string `json:"debug_id,omitempty"` // The PayPal internal ID. Used for correlation purposes.
}

// Error403 represents the Error403 schema from the OpenAPI specification
type Error403 struct {
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Debug_id string `json:"debug_id,omitempty"` // The PayPal internal ID. Used for correlation purposes.
	Details []Errordetails `json:"details,omitempty"`
	Links []Errorlinkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](https://en.wikipedia.org/wiki/HATEOAS).
}

// Experiencecontext represents the Experiencecontext schema from the OpenAPI specification
type Experiencecontext struct {
	Brand_name string `json:"brand_name,omitempty"` // The label that overrides the business name in the PayPal account on the PayPal site. The pattern is defined by an external party and supports Unicode.
	Cancel_url string `json:"cancel_url,omitempty"` // The URL where the customer is redirected after customer cancels or leaves the flow. It is a required field for contingency flows like PayPal wallet, 3DS.
	Locale string `json:"locale,omitempty"` // The [language tag](https://tools.ietf.org/html/bcp47#section-2) for the language in which to localize the error-related strings, such as messages, issues, and suggested actions. The tag is made up of the [ISO 639-2 language code](https://www.loc.gov/standards/iso639-2/php/code_list.php), the optional [ISO-15924 script tag](https://www.unicode.org/iso15924/codelists.html), and the [ISO-3166 alpha-2 country code](/api/rest/reference/country-codes/) or [M49 region code](https://unstats.un.org/unsd/methodology/m49/).
	Return_url string `json:"return_url,omitempty"` // The URL where the customer is redirected after customer approves leaves the flow. It is a required field for contingency flows like PayPal wallet, 3DS.
	Shipping_preference string `json:"shipping_preference,omitempty"` // The shipping preference. This only applies to PayPal payment source.
	Vault_instruction string `json:"vault_instruction,omitempty"` // Vault Instruction on action to be performed after a successful payer approval.
}

// Applepaypaymenttokenresponse represents the Applepaypaymenttokenresponse schema from the OpenAPI specification
type Applepaypaymenttokenresponse struct {
	Card map[string]interface{} `json:"card,omitempty"` // The payment card to be used to fund a payment. Can be a credit or debit card.
}

// Paypalwalletrequest represents the Paypalwalletrequest schema from the OpenAPI specification
type Paypalwalletrequest struct {
	Usage_type string `json:"usage_type,omitempty"` // The usage type associated with a digital wallet payment token.
	Customer_type string `json:"customer_type,omitempty"` // The customer type associated with a digital wallet payment token. This is to indicate whether the customer acting on the merchant / platform is either a business or a consumer.
	Description string `json:"description,omitempty"` // The description displayed to the consumer on the approval flow for a digital wallet, as well as on the merchant view of the payment token management experience. exp: PayPal.com.
	Permit_multiple_payment_tokens bool `json:"permit_multiple_payment_tokens,omitempty"` // Create multiple payment tokens for the same payer, merchant/platform combination. Use this when the customer has not logged in at merchant/platform. The payment token thus generated, can then also be used to create the customer account at merchant/platform. Use this also when multiple payment tokens are required for the same payer, different customer at merchant/platform. This helps to identify customers distinctly even though they may share the same PayPal account. This only applies to PayPal payment source.
	Shipping map[string]interface{} `json:"shipping,omitempty"` // The shipping details.
	Usage_pattern string `json:"usage_pattern,omitempty"` // Expected business/charge model for the billing agreement.
	Billing_plan Plan `json:"billing_plan,omitempty"` // The merchant level Recurring Billing plan metadata for the Billing Agreement.
	Experience_context Experiencecontext `json:"experience_context,omitempty"` // Customizes the Vault creation flow experience for your customers.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Links []Linkdescription `json:"links,omitempty"` // An array of request-related [HATEOAS links](/api/rest/responses/#hateoas-links) that are either relevant to the issue by providing additional information or offering potential resolutions.
	Location string `json:"location,omitempty"` // The location of the field that caused the error. Value is `body`, `path`, or `query`.
	Value string `json:"value,omitempty"` // The value of the field that caused the error.
	Description string `json:"description,omitempty"` // The human-readable description for an issue. The description can change over the lifetime of an API, so clients must not depend on this value.
	Field string `json:"field,omitempty"` // The field that caused the error. If this field is in the body, set this value to the field's JSON pointer value. Required for client-side errors.
	Issue string `json:"issue"` // The unique, fine-grained application-level error code.
}

// Errordetails represents the Errordetails schema from the OpenAPI specification
type Errordetails struct {
	Description string `json:"description,omitempty"` // The human-readable description for an issue. The description can change over the lifetime of an API, so clients must not depend on this value.
	Field string `json:"field,omitempty"` // The field that caused the error. If this field is in the body, set this value to the field's JSON pointer value. Required for client-side errors.
	Issue string `json:"issue"` // The unique, fine-grained application-level error code.
	Location string `json:"location,omitempty"` // The location of the field that caused the error. Value is `body`, `path`, or `query`.
	Value string `json:"value,omitempty"` // The value of the field that caused the error.
}

// Billingcycle represents the Billingcycle schema from the OpenAPI specification
type Billingcycle struct {
	Total_cycles int `json:"total_cycles,omitempty"` // The number of times this billing cycle gets executed. Trial billing cycles can only be executed a finite number of times (value between <code>1</code> and <code>999</code> for <code>total_cycles</code>). Regular billing cycles can be executed infinite times (value of <code>0</code> for <code>total_cycles</code>) or a finite number of times (value between <code>1</code> and <code>999</code> for <code>total_cycles</code>).
	Frequency Frequency `json:"frequency,omitempty"`
	Pricing_scheme Pricingscheme `json:"pricing_scheme,omitempty"` // The pricing scheme details.
	Sequence int `json:"sequence,omitempty"` // The order in which this cycle is to run among other billing cycles. For example, a trial billing cycle has a `sequence` of `1` while a regular billing cycle has a `sequence` of `2`, so that trial cycle runs before the regular cycle.
	Start_date string `json:"start_date,omitempty"` // The stand-alone date, in [Internet date and time format](https://tools.ietf.org/html/rfc3339#section-5.6). To represent special legal values, such as a date of birth, you should use dates with no associated time or time-zone data. Whenever possible, use the standard `date_time` type. This regular expression does not validate all dates. For example, February 31 is valid and nothing is known about leap years.
	Tenure_type string `json:"tenure_type"` // The tenure type of the billing cycle identifies if the billing cycle is a trial(free or discounted) or regular billing cycle.
}

// Onetimecharges represents the Onetimecharges schema from the OpenAPI specification
type Onetimecharges struct {
	Subtotal Money `json:"subtotal,omitempty"` // The currency and amount for a financial transaction, such as a balance or payment due.
	Taxes Money `json:"taxes,omitempty"` // The currency and amount for a financial transaction, such as a balance or payment due.
	Total_amount Money `json:"total_amount"` // The currency and amount for a financial transaction, such as a balance or payment due.
	Product_price Money `json:"product_price,omitempty"` // The currency and amount for a financial transaction, such as a balance or payment due.
	Setup_fee Money `json:"setup_fee,omitempty"` // The currency and amount for a financial transaction, such as a balance or payment due.
	Shipping_amount Money `json:"shipping_amount,omitempty"` // The currency and amount for a financial transaction, such as a balance or payment due.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enrollment_status string `json:"enrollment_status,omitempty"` // Status of Authentication eligibility.
	Authentication_status string `json:"authentication_status,omitempty"` // Transactions status result identifier. The outcome of the issuer's authentication.
	Authentication_id string `json:"authentication_id,omitempty"` // The externally received 3ds authentication id, to be returned in card detokenization response.
}
