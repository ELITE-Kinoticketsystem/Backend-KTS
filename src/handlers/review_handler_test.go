package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateReview(t *testing.T) {
	testCases := []struct {
		name            string
		requestBody     gin.H
		setExpectations func(mockCtrl *mocks.MockReviewControllerI, body gin.H)
		expectedStatus  int
	}{
		{
			name: "Success",
			requestBody: gin.H{
				"rating":    5,
				"comment":   "Comment",
				"datetime":  "2006-01-02T15:04:05Z",
				"isSpoiler": false,
				"userId":    "d2781ace-4d6e-4cc7-9285-bd310c5c6d25",
				"movieId":   "f3c4e8f8-1769-4029-a6b4-fd36b91918a9",
			},
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, body gin.H) {
				datetime, _ := time.Parse(time.RFC3339, body["datetime"].(string))
				reviewData := models.CreateReviewRequest{
					Rating:    int32(body["rating"].(int)),
					Comment:   body["comment"].(string),
					Datetime:  datetime,
					IsSpoiler: body["isSpoiler"].(bool),
					UserID:    body["userId"].(string),
					MovieID:   body["movieId"].(string),
				}
				id := uuid.New()
				mockCtrl.EXPECT().CreateReview(reviewData).Return(&id, nil)
			},
			expectedStatus: 200,
		},
		{
			name: "Internal error",
			requestBody: gin.H{
				"rating":    5,
				"comment":   "Comment",
				"datetime":  "2006-01-02T15:04:05Z",
				"isSpoiler": false,
				"userId":    "d2781ace-4d6e-4cc7-9285-bd310c5c6d25",
				"movieId":   "f3c4e8f8-1769-4029-a6b4-fd36b91918a9",
			},
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, body gin.H) {
				datetime, _ := time.Parse(time.RFC3339, body["datetime"].(string))
				reviewData := models.CreateReviewRequest{
					Rating:    int32(body["rating"].(int)),
					Comment:   body["comment"].(string),
					Datetime:  datetime,
					IsSpoiler: body["isSpoiler"].(bool),
					UserID:    body["userId"].(string),
					MovieID:   body["movieId"].(string),
				}
				mockCtrl.EXPECT().CreateReview(reviewData).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: 500,
		},
		{
			name:            "Invalid json",
			requestBody:     nil,
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, body gin.H) {},
			expectedStatus:  400,
		},
		{
			name: "Empty field",
			requestBody: gin.H{
				"rating":    5,
				"comment":   "",
				"datetime":  "2006-01-02T15:04:05Z",
				"isSpoiler": false,
				"userId":    "d2781ace-4d6e-4cc7-9285-bd310c5c6d25",
				"movieId":   "f3c4e8f8-1769-4029-a6b4-fd36b91918a9",
			},
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, body gin.H) {},
			expectedStatus:  400,
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
			c.Request = req

			// define expectations
			tc.setExpectations(reviewController, tc.requestBody)

			// WHEN
			// call CreateReviewHandler with mock context
			CreateReviewHandler(reviewController)(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			if tc.expectedStatus == 200 {
				_, err := uuid.Parse(w.Body.String())
				assert.True(t, err == nil)
			}
		})
	}

}

func TestDeleteReview(t *testing.T) {
	testCases := []struct {
		name                 string
		id                   string
		setExpectations      func(mockCtrl *mocks.MockReviewControllerI, id uuid.UUID)
		expectedStatus       int
		expectedResponseBody string
	}{
		{
			name: "Success",
			id:   uuid.NewString(),
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, id uuid.UUID) {
				mockCtrl.EXPECT().DeleteReview(&id).Return(nil)
			},
			expectedStatus:       200,
			expectedResponseBody: "",
		},
		{
			name: "Internal error",
			id:   uuid.NewString(),
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, id uuid.UUID) {
				mockCtrl.EXPECT().DeleteReview(&id).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: 500,
			expectedResponseBody: func() string {
				expectedResponseBody, _ := json.Marshal(gin.H{
					"errorMessage": "INTERNAL_ERROR",
				})
				return bytes.NewBuffer(expectedResponseBody).String()
			}(),
		},
		{
			name:            "Invalid id",
			id:              "invalid id",
			setExpectations: func(mockCtrl *mocks.MockReviewControllerI, id uuid.UUID) {},
			expectedStatus:  400,
			expectedResponseBody: func() string {
				expectedResponseBody, _ := json.Marshal(gin.H{
					"errorMessage": "BAD_REQUEST",
				})
				return bytes.NewBuffer(expectedResponseBody).String()
			}(),
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
			req := httptest.NewRequest("DELETE", "/reviews/:id", nil)
			c.Request = req
			c.AddParam("id", tc.id)

			// define expectations
			id, _ := uuid.Parse(tc.id)
			tc.setExpectations(reviewController, id)

			// WHEN
			// call DeleteReviewHandler with mock context
			DeleteReviewHandler(reviewController)(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			assert.Equal(t, tc.expectedResponseBody, w.Body.String(), "wrong response body")
		})
	}
}
