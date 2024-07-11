package errorcode

const (
	SUCCESS        = "0000"
	PartialSuccess = "0001"
	EXCEPTION      = "3000"

	ParamsError     = "3001"
	CheckValueError = "3002"
	VersionError    = "3003"
	ExpiredError    = "3004"
	MinValue        = "3005"
	MaxValue        = "3006"
	SignatureError  = "3007"

	Unauthorized   = "3210"
	BadCredentials = "3211"
	AccessDenied   = "3212"

	ResourceExist    = "3011"
	ResourceNotExist = "3012"
)
