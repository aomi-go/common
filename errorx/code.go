package errorx

const (
	SUCCESS        = "0000"
	PartialSuccess = "0001"
	EXCEPTION      = "3000"

	// ParamsError is used when the request parameters are invalid or missing.
	ParamsError = "3001"
	// ConfigError is used when the configuration is invalid or missing.
	ConfigError  = "3002"
	VersionError = "3003"
	ExpiredError = "3004"
	MinValue     = "3005"
	MaxValue     = "3006"
	MinSize      = "3007"
	MaxSize      = "3008"
	MinLimit     = "3009"
	MaxLimit     = "3010"

	FileTypeNotSupport = "3013"
	FileCorrupted      = "3014"

	// 服务端请求限制相关 start
	ToManyRequest = "3100"
	// 服务端请求限制相关 end

	// 安全相关权限 start
	// Unauthorized is used when the user is not authenticated.
	Unauthorized   = "3200"
	BadCredentials = "3201"
	AccessDenied   = "3202"
	// SignatureError is used when the signature verification fails.
	SignatureError = "3207"
	// 安全相关权限 end

	// ResourceExist is used when the requested resource already exists.
	ResourceExist = "3300"
	// ResourceNotExist is used when the requested resource does not exist.
	ResourceNotExist = "3301"
)
