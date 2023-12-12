package kts_errors

import "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

var (
	// KTS_BAD_REQUEST is used to indicate that the request was malformed
	KTS_BAD_REQUEST = &models.KTSError{ErrorMessage: "BAD_REQUEST", Status: 400}
	// KTS_UNAUTHORIZED is used to indicate that the request was unauthorized
	KTS_UNAUTHORIZED = &models.KTSError{ErrorMessage: "UNAUTHORIZED", Status: 401}
	// KTS_CREDENTIALS_INVALID is used to indicate that the login credentials were invalid
	KTS_CREDENTIALS_INVALID = &models.KTSError{ErrorMessage: "CREDENTIALS_INVALID", Status: 401}
	// KTS_FORBIDDEN is used to indicate that the request was forbidden due to insufficient permissions
	KTS_FORBIDDEN = &models.KTSError{ErrorMessage: "FORBIDDEN", Status: 403}
	// KTS_USER_NOT_FOUND is used to indicate that the requested user was not found
	KTS_USER_NOT_FOUND = &models.KTSError{ErrorMessage: "USER_NOT_FOUND", Status: 404}
	// KTS_NOT_FOUND is used to indicate that the requested resource was not found
	KTS_MOVIE_NOT_FOUND = &models.KTSError{ErrorMessage: "MOVIE_NOT_FOUND", Status: 404}
	// KTS_NOT_FOUND is used to indicate that the requested resource was not found
	KTS_NOT_FOUND = &models.KTSError{ErrorMessage: "NOT_FOUND", Status: 404}
	// KTS_USER_EXISTS is used to indicate that the creation of a user failed because the user already exists
	KTS_USER_EXISTS = &models.KTSError{ErrorMessage: "USER_EXISTS", Status: 409}
	// KTS_EMAIL_EXISTS is used to indicate that the email already exists
	KTS_EMAIL_EXISTS = &models.KTSError{ErrorMessage: "EMAIL_EXISTS", Status: 409}
	// KTS_USERNAME_EXISTS is used to indicate that the username already exists
	KTS_USERNAME_EXISTS = &models.KTSError{ErrorMessage: "USERNAME_EXISTS", Status: 409}
	// KTS_CONFLICT is used to indicate that the request could not be processed due to a conflict
	KTS_CONFLICT = &models.KTSError{ErrorMessage: "CONFLICT", Status: 409}
	// KTS_UPSTREAM_ERROR is used to indicate an error in 3rd party services
	KTS_UPSTREAM_ERROR = &models.KTSError{ErrorMessage: "UPSTREAM_ERROR", Status: 500}
	// KTS_INTERNAL_ERROR is used to indicate an internal, unclassified error
	KTS_INTERNAL_ERROR = &models.KTSError{ErrorMessage: "INTERNAL_ERROR", Status: 500}
)
