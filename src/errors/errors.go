package kts_errors

import "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

var (
	// KTS_BAD_REQUEST is used to indicate that the request was malformed
	KTS_BAD_REQUEST = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "BAD_REQUEST"}, Status: 400}
	// KTS_UNAUTHORIZED is used to indicate that the request was unauthorized
	KTS_UNAUTHORIZED = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "UNAUTHORIZED"}, Status: 401}
	// KTS_CREDENTIALS_INVALID is used to indicate that the login credentials were invalid
	KTS_CREDENTIALS_INVALID = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "CREDENTIALS_INVALID"}, Status: 401}
	// KTS_FORBIDDEN is used to indicate that the request was forbidden due to insufficient permissions
	KTS_FORBIDDEN = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "FORBIDDEN"}, Status: 403}
	// KTS_USER_NOT_FOUND is used to indicate that the requested user was not found
	KTS_USER_NOT_FOUND = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "USER_NOT_FOUND"}, Status: 404}
	// KTS_NOT_FOUND is used to indicate that the requested resource was not found
	KTS_NOT_FOUND = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "NOT_FOUND"}, Status: 404}
	// KTS_USER_EXISTS is used to indicate that the creation of a user failed because the user already exists
	KTS_USER_EXISTS = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "USER_EXISTS"}, Status: 409}
	// KTS_EMAIL_EXISTS is used to indicate that the email already exists
	KTS_EMAIL_EXISTS = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "EMAIL_EXISTS"}, Status: 409}
	// KTS_USERNAME_EXISTS is used to indicate that the username already exists
	KTS_USERNAME_EXISTS = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "USERNAME_EXISTS"}, Status: 409}
	// KTS_CONFLICT is used to indicate that the request could not be processed due to a conflict
	KTS_CONFLICT = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "CONFLICT"}, Status: 409}
	// KTS_UPSTREAM_ERROR is used to indicate an error in 3rd party services
	KTS_UPSTREAM_ERROR = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "UPSTREAM_ERROR"}, Status: 500}
	// KTS_INTERNAL_ERROR is used to indicate an internal, unclassified error
	KTS_INTERNAL_ERROR = &models.KTSError{KTSErrorMessage: models.KTSErrorMessage{ErrorMessage: "INTERNAL_ERROR"}, Status: 500}
)
