package constant

import "errors"

var (
	// OAuth / Auth errors
	ErrInvalidRequest                 = errors.New("invalid_request")
	ErrUnauthorizedClient             = errors.New("unauthorized_client")
	ErrAccessDenied                   = errors.New("access_denied")
	ErrUnsupportedResponseType        = errors.New("unsupported_response_type")
	ErrInvalidScope                   = errors.New("invalid_scope")
	ErrServerError                    = errors.New("server_error")
	ErrTemporarilyUnavailable         = errors.New("temporarily_unavailable")
	ErrInvalidClient                  = errors.New("invalid_client")
	ErrInvalidGrant                   = errors.New("invalid_grant")
	ErrUnsupportedGrantType           = errors.New("unsupported_grant_type")
	ErrCodeChallengeRquired           = errors.New("code_challenge_required")
	ErrUnsupportedCodeChallengeMethod = errors.New("unsupported_code_challenge_method")
	ErrInvalidCodeChallengeLen        = errors.New("invalid_code_challenge_length")
	ErrNotFoundItem                   = errors.New("not_found_item")
	ErrDatabaseAccess                 = errors.New("database_access_error")
	ErrEmailAlreadyExists             = errors.New("email is existed")

	// Judge / Coding platform errors
	ErrCompileFailed = errors.New("compile_failed")
	ErrRuntimeError  = errors.New("runtime_error")
	ErrWrongAnswer   = errors.New("wrong_answer")
	ErrTimeLimit     = errors.New("time_limit_exceeded")
	ErrMemoryLimit   = errors.New("memory_limit_exceeded")
	ErrOutputLimit   = errors.New("output_limit_exceeded")
)

var Descriptions = map[error]string{
	ErrInvalidRequest:                 "Request is missing or has invalid parameters",
	ErrUnauthorizedClient:             "Client is not authorized for this request",
	ErrAccessDenied:                   "Request was denied by user or server",
	ErrUnsupportedResponseType:        "Response type is not supported",
	ErrInvalidScope:                   "Requested scope is invalid or unknown",
	ErrServerError:                    "Unexpected server error occurred",
	ErrTemporarilyUnavailable:         "Service temporarily unavailable",
	ErrInvalidClient:                  "Client authentication failed",
	ErrInvalidGrant:                   "Authorization code or token is invalid or expired",
	ErrUnsupportedGrantType:           "Grant type is not supported",
	ErrCodeChallengeRquired:           "PKCE code_challenge is required",
	ErrUnsupportedCodeChallengeMethod: "Unsupported code_challenge_method",
	ErrInvalidCodeChallengeLen:        "Code challenge length must be 43–128 characters",
	ErrNotFoundItem:                   "The requested item was not found",
	ErrDatabaseAccess:                 "Database access error",
	ErrEmailAlreadyExists:             "Email duplicated",
	// Judge errors
	ErrCompileFailed: "Compilation failed",
	ErrRuntimeError:  "Runtime error occurred while executing",
	ErrWrongAnswer:   "Output does not match expected result",
	ErrTimeLimit:     "Execution exceeded time limit",
	ErrMemoryLimit:   "Execution exceeded memory limit",
	ErrOutputLimit:   "Output exceeded allowed limit",
}

var StatusCodes = map[error]int{
	// OAuth
	ErrInvalidRequest:                 400,
	ErrUnauthorizedClient:             401,
	ErrAccessDenied:                   403,
	ErrUnsupportedResponseType:        400,
	ErrInvalidScope:                   400,
	ErrServerError:                    500,
	ErrTemporarilyUnavailable:         503,
	ErrInvalidClient:                  401,
	ErrInvalidGrant:                   400,
	ErrUnsupportedGrantType:           400,
	ErrCodeChallengeRquired:           400,
	ErrUnsupportedCodeChallengeMethod: 400,
	ErrInvalidCodeChallengeLen:        400,
	ErrNotFoundItem:                   404,
	ErrDatabaseAccess:                 500,
	ErrEmailAlreadyExists:             419,

	// Judge
	ErrCompileFailed: 211,
	ErrRuntimeError:  212,
	ErrWrongAnswer:   213,
	ErrTimeLimit:     214,
	ErrMemoryLimit:   215,
	ErrOutputLimit:   216,
}
