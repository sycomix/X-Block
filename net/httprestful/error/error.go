package error

const (
	SUCCESS            int64 = 0
	SESSION_EXPIRED    int64 = 41001
	SERVICE_CEILING    int64 = 41002
	ILLEGAL_DATAFORMAT int64 = 41003

	INVALID_METHOD int64 = 42001
	INVALID_PARAMS int64 = 42002
	INVALID_TOKEN  int64 = 42003

	INVALID_TRANSACTION int64 = 43001
	INVALID_ASSET       int64 = 43002
	INVALID_BLOCK       int64 = 43003

	UNKNOWN_TRANSACTION int64 = 44001
	UNKNOWN_ASSET       int64 = 44002
	UNKNOWN_BLOCK       int64 = 44003

	INVALID_VERSION int64 = 45001
	INTERNAL_ERROR  int64 = 45002
)

