package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateReview(t *testing.T) {
	user := samples.GetSampleUser()
	review := samples.GetSampleReview()
	testCases := []struct {
		name            string
		setContextUser  bool
		requestBody     gin.H
		setExpectations func(mockCtrl *mocks.MockReviewControllerI, body gin.H)
		expectedBody    gin.H
		expectedStatus  int
	}{
		{
			name:           "Success",
			setContextUser: true,
			requestBody: gin.H{
				"rating":    5,
				"comment":   "Comment",
				"datetime":  "2006-01-02T15:04:05Z",
				"isSpoiler": false,
				"movieId":   "f3c4e8f8-1769-4029-a6b4-fd36b91918a9",
			},
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, body gin.H) {
				datetime, _ := time.Parse(time.RFC3339, body["datetime"].(string))
				reviewData := models.CreateReviewRequest{
					Rating:    int32(body["rating"].(int)),
					Comment:   body["comment"].(string),
					Datetime:  datetime,
					IsSpoiler: body["isSpoiler"].(bool),
					MovieID:   body["movieId"].(string),
				}
				mockCtrl.EXPECT().CreateReview(reviewData, user.ID).Return(&review, *user.Username, nil)
			},
			expectedBody:   gin.H{"review": review, "username": user.Username},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "Internal error",
			setContextUser: true,
			requestBody: gin.H{
				"rating":    5,
				"comment":   "Comment",
				"datetime":  "2006-01-02T15:04:05Z",
				"isSpoiler": false,
				"movieId":   "f3c4e8f8-1769-4029-a6b4-fd36b91918a9",
			},
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, body gin.H) {
				datetime, _ := time.Parse(time.RFC3339, body["datetime"].(string))
				reviewData := models.CreateReviewRequest{
					Rating:    int32(body["rating"].(int)),
					Comment:   body["comment"].(string),
					Datetime:  datetime,
					IsSpoiler: body["isSpoiler"].(bool),
					MovieID:   body["movieId"].(string),
				}
				mockCtrl.EXPECT().CreateReview(reviewData, user.ID).Return(nil, "", kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:            "Invalid json",
			setContextUser:  true,
			requestBody:     nil,
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, body gin.H) {},
			expectedBody:    gin.H{"errorMessage": "BAD_REQUEST"},
			expectedStatus:  http.StatusBadRequest,
		},
		{
			name:           "Empty field",
			setContextUser: true,
			requestBody: gin.H{
				"rating":    5,
				"comment":   "",
				"datetime":  "2006-01-02T15:04:05Z",
				"isSpoiler": false,
				"userId":    "d2781ace-4d6e-4cc7-9285-bd310c5c6d25",
				"movieId":   "f3c4e8f8-1769-4029-a6b4-fd36b91918a9",
			},
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, body gin.H) {},
			expectedBody:    gin.H{"errorMessage": "BAD_REQUEST"},
			expectedStatus:  http.StatusBadRequest,
		},
		{
			name:           "Failed - no context user",
			setContextUser: false,
			requestBody: gin.H{
				"rating":    5,
				"comment":   "Comment",
				"datetime":  "2006-01-02T15:04:05Z",
				"isSpoiler": false,
				"movieId":   "f3c4e8f8-1769-4029-a6b4-fd36b91918a9",
			},
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, body gin.H) {

			},
			expectedBody:   gin.H{"errorMessage": "UNAUTHORIZED"},
			expectedStatus: http.StatusUnauthorized,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock context
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			// create mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			reviewController := mocks.NewMockReviewControllerI(mockCtrl)

			// create mock request
			jsonData, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest("POST", "/reviews", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			if tc.setContextUser {
				ctx := context.WithValue(req.Context(), models.ContextKeyUserID, user.ID)
				c.Request = req.WithContext(ctx)
			} else {
				c.Request = req
			}

			// define expectations
			tc.setExpectations(reviewController, tc.requestBody)

			// WHEN
			// call CreateReviewHandler with mock context
			CreateReviewHandler(reviewController)(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}

}

func TestDeleteReview(t *testing.T) {
	testCases := []struct {
		name                 string
		setContextUser       bool
		id                   string
		userId               string
		setExpectations      func(mockCtrl *mocks.MockReviewControllerI, id uuid.UUID, userId uuid.UUID)
		expectedStatus       int
		expectedResponseBody interface{}
	}{
		{
			name:           "Success",
			setContextUser: true,
			id:             uuid.NewString(),
			userId:         uuid.NewString(),
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, id uuid.UUID, userId uuid.UUID) {
				mockCtrl.EXPECT().DeleteReview(&id, &userId).Return(nil)
			},
			expectedStatus:       200,
			expectedResponseBody: "",
		},
		{
			name:           "Unauthorized",
			setContextUser: true,
			id:             uuid.NewString(),
			userId:         uuid.NewString(),
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, id uuid.UUID, userId uuid.UUID) {
				mockCtrl.EXPECT().DeleteReview(&id, &userId).Return(kts_errors.KTS_FORBIDDEN)
			},
			expectedStatus:       403,
			expectedResponseBody: gin.H{"errorMessage": "FORBIDDEN"},
		},
		{
			name:           "Internal error",
			setContextUser: true,
			id:             uuid.NewString(),
			userId:         uuid.NewString(),
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, id uuid.UUID, userId uuid.UUID) {
				mockCtrl.EXPECT().DeleteReview(&id, &userId).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus:       500,
			expectedResponseBody: gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:                 "Invalid id",
			setContextUser:       true,
			id:                   "invalid id",
			setExpectations:      func(mockCtrl *mocks.MockReviewControllerI, id uuid.UUID, userId uuid.UUID) {},
			expectedStatus:       400,
			expectedResponseBody: gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:           "Failed - no context user",
			setContextUser: false,
			id:             utils.NewUUID().String(),
			userId:         utils.NewUUID().String(),
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, id uuid.UUID, userId uuid.UUID) {

			},
			expectedResponseBody: gin.H{"errorMessage": "UNAUTHORIZED"},
			expectedStatus:       http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock context
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			// create mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			reviewController := mocks.NewMockReviewControllerI(mockCtrl)

			// create mock request
			id, _ := uuid.Parse(tc.id)
			userId, _ := uuid.Parse(tc.userId)

			req := httptest.NewRequest("DELETE", "/reviews/:id", nil)

			if tc.setContextUser {
				ctx := context.WithValue(req.Context(), models.ContextKeyUserID, &userId)
				c.Request = req.WithContext(ctx)
			} else {
				c.Request = req
			}
			c.AddParam("id", tc.id)

			// define expectations
			tc.setExpectations(reviewController, id, userId)

			// WHEN
			// call DeleteReviewHandler with mock context
			DeleteReviewHandler(reviewController)(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedResponseBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
