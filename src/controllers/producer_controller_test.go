package controllers

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

func TestGetProducers(t *testing.T) {

	sampleProducers := &[]models.GetProducersDTO{
		{
			Producers: model.Producers{
				ID:          utils.NewUUID(),
				Name:        "Test Producer",
				Description: "Test Description",
			},
			Pictures: nil,
		},
	}

	testCases := []struct {
		name                 string
		setExpectations      func(mockRepo mocks.MockProducerRepositoryI)
		expectedProducersDTO *[]models.GetProducersDTO
		expectedError        *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockProducerRepositoryI) {
				mockRepo.EXPECT().GetProducers().Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedProducersDTO: nil,
			expectedError:        kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mockRepo mocks.MockProducerRepositoryI) {
				mockRepo.EXPECT().GetProducers().Return(sampleProducers, nil)
			},
			expectedProducersDTO: sampleProducers,
			expectedError:        nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mockRepo mocks.MockProducerRepositoryI) {
				mockRepo.EXPECT().GetProducers().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedProducersDTO: nil,
			expectedError:        kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			producerRepoMock := mocks.NewMockProducerRepositoryI(mockCtrl)
			producerController := ProducerController{
				ProducerRepo: producerRepoMock,
			}

			tc.setExpectations(*producerRepoMock)

			// WHEN
			producers, kts_err := producerController.GetProducers()

			// THEN
			assert.Equal(t, tc.expectedProducersDTO, producers)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetProducerByID(t *testing.T) {
	samplpleProducer := &models.ProducerDTO{
		Producers: model.Producers{
			ID:          utils.NewUUID(),
			Name:        "Test Producer",
			Description: "Test Description",
		},
		Pictures: nil,
		Movies:   nil,
	}

	testCases := []struct {
		name             string
		setExpectations  func(mockRepo mocks.MockProducerRepositoryI, movieId *uuid.UUID)
		expectedProducer *models.ProducerDTO
		expectedError    *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockProducerRepositoryI, producerId *uuid.UUID) {
				mockRepo.EXPECT().GetProducerById(producerId).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedProducer: nil,
			expectedError:    kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mockRepo mocks.MockProducerRepositoryI, producerId *uuid.UUID) {
				mockRepo.EXPECT().GetProducerById(producerId).Return(samplpleProducer, nil)
			},
			expectedProducer: samplpleProducer,
			expectedError:    nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mockRepo mocks.MockProducerRepositoryI, producerId *uuid.UUID) {
				mockRepo.EXPECT().GetProducerById(producerId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedProducer: nil,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			producerRepoMock := mocks.NewMockProducerRepositoryI(mockCtrl)

			producerController := ProducerController{
				ProducerRepo: producerRepoMock,
			}

			// define expectations
			tc.setExpectations(*producerRepoMock, samplpleProducer.ID)

			// WHEN
			movies, kts_err := producerController.GetProducerById(samplpleProducer.ID)

			// THEN
			assert.Equal(t, tc.expectedProducer, movies)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestCreateProducer(t *testing.T) {

	sampleProducer := &models.CreateProducerDTO{
		Producers: model.Producers{
			ID:          utils.NewUUID(),
			Name:        "Test Producer",
			Description: "Test Description",
		},
		PicturesUrls: []string{
			"test1",
			"test2",
		},
	}

	samplePictureUrlId := utils.NewUUID()

	testCases := []struct {
		name               string
		producerDto        *models.CreateProducerDTO
		setExpectations    func(mockProducerRepo mocks.MockProducerRepositoryI, producer *models.CreateProducerDTO)
		expectedProducerID bool
		expectedError      *models.KTSError
	}{
		{
			name: "Create producer",
			producerDto: &models.CreateProducerDTO{
				Producers:    sampleProducer.Producers,
				PicturesUrls: sampleProducer.PicturesUrls,
			},
			setExpectations: func(mockProducerRepo mocks.MockProducerRepositoryI, producer *models.CreateProducerDTO) {
				mockProducerRepo.EXPECT().CreateProducer(&producer.Producers).Return(sampleProducer.ID, nil)
				mockProducerRepo.EXPECT().CreateProducerPicture(gomock.Any()).AnyTimes().Return(samplePictureUrlId, nil)

			},
			expectedProducerID: true,
			expectedError:      nil,
		},
		{
			name:        "Create producer fails",
			producerDto: &models.CreateProducerDTO{},
			setExpectations: func(mockProducerRepo mocks.MockProducerRepositoryI, producer *models.CreateProducerDTO) {

			},
			expectedProducerID: false,
			expectedError:      kts_errors.KTS_BAD_REQUEST,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			producerRepoMock := mocks.NewMockProducerRepositoryI(mockCtrl)
			producerController := ProducerController{
				ProducerRepo: producerRepoMock,
			}

			// define expectations
			tc.setExpectations(*producerRepoMock, tc.producerDto)

			// WHEN
			producerId, kts_err := producerController.CreateProducer(tc.producerDto)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)

			if tc.expectedProducerID && producerId == nil {
				t.Error("Expected producer ID, got nil")
			}
		})
	}
}
