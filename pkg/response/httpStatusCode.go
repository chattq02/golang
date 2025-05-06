package response

const (
	ErrCodeSuccess      = 20001 // success
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken = 30001 // Invalid token
	ErrInvalidOTP   = 30002 // Invalid token type
	ErrSendEmailOtp = 30003 // Send email
	// User Authentication fail
	ErrCodeAuthFailed = 40005

	// Register code
	ErrCodeUserHasExits = 50001 // user has already registered

	// err login
	ErrCodeOtpHasExits      = 60009
	ErrCodeUserOtpNotExists = 60008
)

// message
var msg = map[int]string{
	ErrCodeSuccess:          "Success",
	ErrCodeParamInvalid:     "Email is invalid",
	ErrInvalidOTP:           "OTP error",
	ErrInvalidToken:         "Invalid token",
	ErrSendEmailOtp:         "Fail to send otp",
	ErrCodeUserHasExits:     "User has already registered",
	ErrCodeOtpHasExits:      "OTP exists but not registered",
	ErrCodeAuthFailed:       "Authentication failed",
	ErrCodeUserOtpNotExists: "OTP exists but not registered",
}