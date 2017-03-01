package main

import (
	"encoding/json"
)

// /v1/gateway/creditcard/verify

// VerifyRequest ...
type VerifyRequest struct {
	MsgID          string  `url:"msgID"`
	UserID         int64   `url:"userID"`
	UserType       string  `url:"userType"`
	PaymentTypeID  string  `url:"paymentTypeID"`
	FareUpperBound float64 `url:"fareUpperBound"`
	Currency       string  `url:"currency"`
}

// DeleteAllRequest ...
type DeleteAllRequest struct {
	MsgID    string `url:"msgID"`
	UserID   int64  `url:"userID"`
	UserType string `url:"userType"`
}

// VerifyResponse ...
type VerifyResponse struct {
	UserGroupID int64 `json:"userGroupID"`
}

// VerifyErrorResponse ...
type VerifyErrorResponse struct {
	Reason string `json:"reason"`
}

// /v1/booking/{:bookingCode}/cancellation/

// CancellationRequest ...
type CancellationRequest struct {
	MsgID    string  `url:"msgID"`
	UserID   int64   `url:"userID"`
	Amount   float64 `url:"amount"`
	Currency string  `url:"currency"`
}

// CancellationResponse ...
type CancellationResponse struct {
}

// PreAuthChargeRequest ...
type PreAuthChargeRequest struct {
	UserID        int64   `json:"userID"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentTypeID string  `json:"paymentTypeID"`
	BookingCode   string  `json:"bookingCode" `
}

// PreAuthChargeFailResponse ...
type PreAuthChargeFailResponse struct {
	Status string `json:"status"` // this will be a "success", "failed", "unknown", "retry"
}

// AuthRequest ...
type AuthRequest struct {
	UserID        int64   `json:"userID"`
	PaymentTypeID string  `json:"paymentTypeID"`
	BookingCode   string  `json:"bookingCode"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
}

// AuthResponse ...
type AuthResponse struct {
	AuthTxID string `json:"txID"`
	Status   string `json:"status"`
}

// AuthCancelRequest ...
type AuthCancelRequest struct {
	BookingCode string `json:"bookingCode"`
}

// GetCreditTopupMethodsRequest defines request to get all supported credit/topup methods
type GetCreditTopupMethodsRequest struct {
	MsgID       *string `url:",omitempty"`
	CountryCode *string `url:",omitempty"`
	IsSorted    *bool   `url:",omitempty"` //it is not omitempty in paysi dto
	//IsSorted    *bool   `url:"isSorted,omitempty"` //it is not omitempty in paysi dto
}

// GenerateCreditTopupFrmBrandRequest defines the request of getting topup page info
type GenerateCreditTopupFrmBrandRequest struct {
	MsgID       *string  `url:"msgID",omitempty`
	Amount      *float64 `url:"amount",omitempty`
	CountryCode *string  `url:"countryCode",omitempty`
	BrandCode   *string  `url:"brandCode",omitempty`
	RewardID    *int64   `url:"rewardID,omitempty"`
}

// GetCreditTopUpFrmBrandAdyenCallbackInfoRequest defines the callback struct for adyen brand
type GetCreditTopUpFrmBrandAdyenCallbackInfoRequest struct {
	TxID              *string `url:"txID"`
	ProviderInfo      *string `url:"providerInfo"`
	AuthResult        *string `url:"authResult",omitempty`
	PspReference      *string `url:"pspReference"`
	MerchantReference *string `url:"merchantReference",omitempty`
	MerchantSig       *string `url:"merchantSig",omitempty`
	SkinCode          *string `url:"skinCode",omitempty`
	PaymentMethod     *string `url:"paymentMethod"`
	ShopperLocale     *string `url:"shopperLocale"`
}

// GenerateCreditTopUpInfoFrmCardRequest defines the request of getting topup page info
type GenerateCreditTopUpInfoFrmCardRequest struct {
	MsgID         *string  `url:"msgID",omitempty`
	PaymentTypeID *string  `url:"paymentTypeID",omitempty`
	Amount        *float64 `url:"amount",omitempty`
	Currency      *string  `url:"currency",omitempty`
}

// TopUpFrmVoucherRequest defines request for topup/voucher
type TopUpFrmVoucherRequest struct {
	MsgID       *string `url:"msgID",omitempty`
	VoucherCode *string `url:"voucherCode",omitempty`
	CountryCode *string `url:"countryCode",omitempty`
}

//BasePaysiAuthRequest ...
type BasePaysiAuthRequest struct {
	MsgID *string `json:"msgID",omitempty`
}

//ECashBalanceRequest ...
type ECashBalanceRequest struct {
	BasePaysiAuthRequest
	UserType *string `json:"userType",omitempty`
}

//ECashBindCheckRequest defines request to get current balance of passenger's e-cash account
type ECashBindCheckRequest struct {
	BasePaysiAuthRequest
	UserType *string `json:"userType",omitempty`
	MobileNo *string `json:"mobileNo",omitempty`
}

//ECashBindApplyRequest defines request to get current balance of passenger's e-cash account
type ECashBindApplyRequest struct {
	BasePaysiAuthRequest
	UserType     *string `json:"userType",omitempty`
	MobileNo     *string `json:"mobileNo",omitempty`
	EncryptedPIN *string `json:"encryptedPIN",omitempty`
	EncryptedOTP *string `json:"encryptedOTP",omitempty`
}

// WalletLoadInfoRequest defines request to enquire about provider, wallet, cards for a pax
type WalletLoadInfoRequest struct {
	BasePaysiAuthRequest
	UserType    *string `json:"userType",omitempty`
	CountryCode *string `json:"countryCode",omitempty`
}

// WalletUpdateRequest ...
type WalletUpdateRequest struct {
	BasePaysiAuthRequest
	UserType       *string `json:"userType",omitempty`
	CardID         *string `json:"cardID",omitempty`
	PrimaryGroupID *int64  `json:"primaryGroupID"`     //both PrimaryGroupID and Nickname can't be nil
	Nickname       *string `json:"nickname",omitempty` //
}

// WalletDeleteRequest ...
type WalletDeleteRequest struct {
	BasePaysiAuthRequest
	UserType *string `json:"userType",omitempty`
	CardID   *string `json:"cardID",omitempty`
}

// WalletSetPrimaryRequest ...
type WalletSetPrimaryRequest struct {
	BasePaysiAuthRequest
	UserType *string `json:"userType",omitempty`
	CardID   *string `json:"cardID",omitempty`
}

// WalletGatewayAddCardRequest ...
type WalletGatewayAddCardRequest struct {
	BasePaysiAuthRequest
	UserType    *string `json:"userType",omitempty`
	CountryCode *string `json:"countryCode",omitempty`
}

// WalletGatewayCallbackRequest ...
type WalletGatewayCallbackRequest struct {
	BasePaysiAuthRequest
	UserType *string          `json:"userType",omitempty`
	Provider *string          `json:"provider",omitempty`
	Email    *string          `json:"email",omitempty` //no omitempty in paysi dto, but it can be empty or nil
	Payload  *json.RawMessage `json:"payload",omitempty`
}

// FailedPaymentRetryRequest dto for POST failedpayment/retry/
type FailedPaymentRetryRequest struct {
	MsgID         *string `json:"msgID",omitempty`
	UserType      *string `json:"userType",omitempty`
	PaymentTypeID *string `json:"paymentTypeID",omitempty`
}

// FailedPaymentSearchRequest dto for GET failedpayment/search/
type FailedPaymentSearchRequest struct {
	MsgID    *string `json:"msgID",omitempty`
	UserType *string `json:"userType",omitempty`
}

// AlipaySignRequestRequest dto for GET alipay/sign/request/
type AlipaySignRequestRequest struct {
	MsgID       *string `json:"msgID",omitempty`
	UserType    *string `json:"userType",omitempty`
	CountryCode *string `json:"countryCode",omitempty`
	ReturnURL   *string `json:"returnURL",omitempty`
}

// AlipaySignCallbackRequest dto for PUT alipay/sign/callback/
type AlipaySignCallbackRequest struct {
	MsgID          *string `json:"msgID",omitempty`
	UserType       *string `json:"userType",omitempty`
	IsSuccess      *string `json:"is_success",omitempty`
	SignType       *string `json:"sign_type",omitempty`
	Sign           *string `json:"sign",omitempty`
	InputCharacter *string `json:"_input_charset",omitempty`
	AgreementNo    *string `json:"agreement_no",omitempty`
	ProductCode    *string `json:"product_code",omitempty`
	Scene          *string `json:"scene",omitempty`
	Status         *string `json:"status",omitempty`
	SignTime       *string `json:"sign_time",omitempty`
	SignModifyTime *string `json:"sign_modify_time",omitempty`
	ValidTime      *string `json:"valid_time",omitempty`
	InValidTime    *string `json:"invalid_time",omitempty`
	AlipayUserID   *string `json:"alipay_user_id",omitempty`
	ExternalSignNo *string `json:"external_sign_no",omitempty`
	ZmOpenID       *string `json:"zm_open_id",omitempty`
	ErrorCode      *string `json:"error",omitempty`
}

// PostAdyenCallbackRequest ...
type PostAdyenCallbackRequest struct {
	MsgID  *string `url:"msgID",omitempty`
	CardID *string `url:"cardID",omitempty`
	MD     *string `url:"md",omitempty`
	PaRes  *string `url:"paRes",omitempty`
}

// GetBinInfoRequest ...
type GetBinInfoRequest struct {
	MsgID *string `url:"msgID",omitempty`
	Bin   *string `url:"bin",omitempty`
}

// MobilePayBindRequest dto for POST mobilepay/bind
type MobilePayBindRequest struct {
	MsgID      *string `json:"msgID",omitempty`
	UserType   *string `json:"userType",omitempty`
	MobileType *string `json:"mobileType",omitempty`
	CardType   *string `json:"cardType",omitempty`   //no omitempty in paysi dto, but it can be empty or nil
	CardNumber *string `json:"cardNumber",omitempty` //no omitempty in paysi dto, but it can be empty or nil
}

// MobilePayUpdateRequest defines ext. request to update mobile pay account
type MobilePayUpdateRequest struct {
	MobilePayBindRequest
	PaymentTypeID *string `json:"paymentTypeID"` //error, no contraints, from paysi codes, it is omitempty
}

// MobilePayPrepareChargeRequest defines ext. request to prepare mobile pay account for charge
type MobilePayPrepareChargeRequest struct {
	MsgID         *string         `json:"msgID",omitempty`
	UserType      *string         `json:"userType",omitempty`
	PaymentTypeID *string         `json:"paymentTypeID",omitempty`
	CountryCode   *string         `json:"countryCode",omitempty`
	ReferenceCode string          `json:"referenceCode",omitempty` // this is optional parameter, for manual retry declined booking
	Payload       json.RawMessage `json:"payload",omitempty`       //no omitempty in paysi dto, but it can be empty or nil
}
