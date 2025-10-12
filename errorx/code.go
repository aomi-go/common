package errorx

const (
	SUCCESS        = "0000"
	PartialSuccess = "0001"
	EXCEPTION      = "3000"

	// ParamsError is used when the request parameters are invalid or missing.
	ParamsError     = "3001"
	CheckValueError = "3002"
	VersionError    = "3003"
	ExpiredError    = "3004"
	MinValue        = "3005"
	MaxValue        = "3006"
	// SignatureError is used when the signature verification fails.
	SignatureError = "3007"
	// ConfigError is used when the configuration is invalid or missing.
	ConfigError = "3008"
	MinSize     = "3009"
	MaxSize     = "3010"

	// ResourceExist is used when the requested resource already exists.
	ResourceExist = "3011"
	// ResourceNotExist is used when the requested resource does not exist.
	ResourceNotExist = "3012"

	// Unauthorized is used when the user is not authenticated.
	Unauthorized   = "3210"
	BadCredentials = "3211"
	AccessDenied   = "3212"
)
