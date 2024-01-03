package controllers

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func TestGetProducers(t *testing.T) {

	samplePriceCategories := &[]model.PriceCategories{
		{
			ID:           utils.NewUUID(),
			CategoryName: "StudentDiscount",
			Price:        16,
		},
		{
			ID:           utils.NewUUID(),
			CategoryName: "ChildDiscount",
			Price:        16,
		},
	}

	testCases := []struct {
		name                    string
		setExpectations         func(mockRepo mocks.MockPriceCategoryRepositoryI)
		expectedPriceCategories *[]model.PriceCategories
		expectedError           *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockPriceCategoryRepositoryI) {
				mockRepo.EXPECT().GetPriceCategories().Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedPriceCategories: nil,
			expectedError:           kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple priceCategories",
			setExpectations: func(mockRepo mocks.MockPriceCategoryRepositoryI) {
				mockRepo.EXPECT().GetPriceCategories().Return(samplePriceCategories, nil)
			},
			expectedPriceCategories: samplePriceCategories,
			expectedError:           nil,
		},
		{
			name: "Error while querying priceCategories",
			setExpectations: func(mockRepo mocks.MockPriceCategoryRepositoryI) {
				mockRepo.EXPECT().GetPriceCategories().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedPriceCategories: nil,
			expectedError:           kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			priceCategoryRepoMock := mocks.NewMockPriceCategoryRepositoryI(mockCtrl)
			priceCategoryController := PriceCategoryController{
				PriceCategoryRepository: priceCategoryRepoMock,
			}

			tc.setExpectations(*priceCategoryRepoMock)

			// WHEN
			priceCategories, kts_err := priceCategoryController.GetPriceCategories()

			// THEN
			assert.Equal(t, tc.expectedPriceCategories, priceCategories)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetProducerByID(t *testing.T) {
	samplePriceCategory := &model.PriceCategories{
		ID:           utils.NewUUID(),
		CategoryName: "StudentDiscount",
		Price:        16,
	}

	testCases := []struct {
		name                  string
		setExpectations       func(mockRepo mocks.MockPriceCategoryRepositoryI, priceCategoryID *uuid.UUID)
		expectedPriceCategory *model.PriceCategories
		expectedError         *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockPriceCategoryRepositoryI, priceCategoryID *uuid.UUID) {
				mockRepo.EXPECT().GetPriceCategoryById(priceCategoryID).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedPriceCategory: nil,
			expectedError:         kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "One priceCategory",
			setExpectations: func(mockRepo mocks.MockPriceCategoryRepositoryI, priceCategoryID *uuid.UUID) {
				mockRepo.EXPECT().GetPriceCategoryById(priceCategoryID).Return(samplePriceCategory, nil)
			},
			expectedPriceCategory: samplePriceCategory,
			expectedError:         nil,
		},
		{
			name: "Error while querying for priceCategory",
			setExpectations: func(mockRepo mocks.MockPriceCategoryRepositoryI, priceCategoryID *uuid.UUID) {
				mockRepo.EXPECT().GetPriceCategoryById(priceCategoryID).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedPriceCategory: nil,
			expectedError:         kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			priceCategoryRepoMock := mocks.NewMockPriceCategoryRepositoryI(mockCtrl)
			priceCategoryController := PriceCategoryController{
				PriceCategoryRepository: priceCategoryRepoMock,
			}

			// define expectations
			tc.setExpectations(*priceCategoryRepoMock, samplePriceCategory.ID)

			// WHEN
			priceCategory, kts_err := priceCategoryController.GetPriceCategoryById(samplePriceCategory.ID)

			// THEN
			assert.Equal(t, tc.expectedPriceCategory, priceCategory)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestCreateProducer(t *testing.T) {

	samplePriceCategory := &model.PriceCategories{
		ID:           utils.NewUUID(),
		CategoryName: "StudentDiscount",
		Price:        16,
	}

	testCases := []struct {
		name                    string
		priceCategory           *model.PriceCategories
		setExpectations         func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategory *model.PriceCategories)
		expectedPriceCategoryID bool
		expectedError           *models.KTSError
	}{
		{
			name:          "priceCategory == nil",
			priceCategory: nil,
			setExpectations: func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategory *model.PriceCategories) {

			},
			expectedPriceCategoryID: false,
			expectedError:           kts_errors.KTS_BAD_REQUEST,
		},
		{
			name:          "Create priceCategory",
			priceCategory: samplePriceCategory,
			setExpectations: func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategory *model.PriceCategories) {
				mockProducerRepo.EXPECT().CreatePriceCategory(priceCategory).Return(samplePriceCategory.ID, nil)
			},
			expectedPriceCategoryID: true,
			expectedError:           nil,
		},
		{
			name:          "Create priceCategory fails",
			priceCategory: samplePriceCategory,
			setExpectations: func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategory *model.PriceCategories) {
				mockProducerRepo.EXPECT().CreatePriceCategory(priceCategory).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedPriceCategoryID: false,
			expectedError:           kts_errors.KTS_INTERNAL_ERROR,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			priceCategoryRepoMock := mocks.NewMockPriceCategoryRepositoryI(mockCtrl)
			priceCategoryController := PriceCategoryController{
				PriceCategoryRepository: priceCategoryRepoMock,
			}

			// define expectations
			tc.setExpectations(*priceCategoryRepoMock, tc.priceCategory)

			// WHEN
			producerId, kts_err := priceCategoryController.CreatePriceCategory(tc.priceCategory)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)

			if tc.expectedPriceCategoryID && producerId == nil {
				t.Error("Expected producer ID, got nil")
			}
		})
	}
}

func TestUpdateProducer(t *testing.T) {

	samplePriceCategory := &model.PriceCategories{
		ID:           utils.NewUUID(),
		CategoryName: "StudentDiscount",
		Price:        16,
	}

	testCases := []struct {
		name                    string
		priceCategory           *model.PriceCategories
		setExpectations         func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategory *model.PriceCategories)
		expectedPriceCategoryID bool
		expectedError           *models.KTSError
	}{
		{
			name:          "priceCategory == nil",
			priceCategory: nil,
			setExpectations: func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategory *model.PriceCategories) {

			},
			expectedPriceCategoryID: false,
			expectedError:           kts_errors.KTS_BAD_REQUEST,
		},
		{
			name:          "Create priceCategory",
			priceCategory: samplePriceCategory,
			setExpectations: func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategory *model.PriceCategories) {
				mockProducerRepo.EXPECT().UpdatePriceCategory(priceCategory).Return(samplePriceCategory.ID, nil)
			},
			expectedPriceCategoryID: true,
			expectedError:           nil,
		},
		{
			name:          "Create priceCategory fails",
			priceCategory: samplePriceCategory,
			setExpectations: func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategory *model.PriceCategories) {
				mockProducerRepo.EXPECT().UpdatePriceCategory(priceCategory).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedPriceCategoryID: false,
			expectedError:           kts_errors.KTS_INTERNAL_ERROR,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			priceCategoryRepoMock := mocks.NewMockPriceCategoryRepositoryI(mockCtrl)
			priceCategoryController := PriceCategoryController{
				PriceCategoryRepository: priceCategoryRepoMock,
			}

			// define expectations
			tc.setExpectations(*priceCategoryRepoMock, tc.priceCategory)

			// WHEN
			producerId, kts_err := priceCategoryController.UpdatePriceCategory(tc.priceCategory)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)

			if tc.expectedPriceCategoryID && producerId == nil {
				t.Error("Expected producer ID, got nil")
			}
		})
	}
}

func TestDeleteProducer(t *testing.T) {

	priceCategoryID := utils.NewUUID()

	testCases := []struct {
		name            string
		priceCategoryID *uuid.UUID
		setExpectations func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategoryID *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name:            "priceCategory == nil",
			priceCategoryID: nil,
			setExpectations: func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategoryID *uuid.UUID) {

			},
			expectedError: kts_errors.KTS_BAD_REQUEST,
		},
		{
			name:            "Create priceCategory",
			priceCategoryID: priceCategoryID,
			setExpectations: func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategoryID *uuid.UUID) {
				mockProducerRepo.EXPECT().DeletePriceCategory(priceCategoryID).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:            "Create priceCategory fails",
			priceCategoryID: priceCategoryID,
			setExpectations: func(mockProducerRepo mocks.MockPriceCategoryRepositoryI, priceCategoryID *uuid.UUID) {
				mockProducerRepo.EXPECT().DeletePriceCategory(priceCategoryID).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			priceCategoryRepoMock := mocks.NewMockPriceCategoryRepositoryI(mockCtrl)
			priceCategoryController := PriceCategoryController{
				PriceCategoryRepository: priceCategoryRepoMock,
			}

			// define expectations
			tc.setExpectations(*priceCategoryRepoMock, tc.priceCategoryID)

			// WHEN
			kts_err := priceCategoryController.DeletePriceCategory(tc.priceCategoryID)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}
