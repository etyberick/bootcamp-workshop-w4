package controller

import (
	"errors"
	"testing"

	"github.com/varopxndx/bootcamp-workshop-w4/controller/mocks"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/unrolled/render"
)

//mockgen -source=controller/controller.go -destination=controller/mocks/controller.go -package=mocks
func Test_Controller(t *testing.T) {
	l := &logrus.Logger{}
	r := render.New()

	tests := []struct {
		name                    string
		expectedParams          string
		expectedUsecaseResponse string
		expectUsecaseCall       bool
		expectedError           error
		wantError               bool
	}{
		{
			name:                    "OK, get square",
			expectedParams:          "square",
			expectedUsecaseResponse: "hi",
			expectUsecaseCall:       true,
			wantError:               false,
			expectedError:           nil,
		},
		//Create scenario for 404
		{
			name:                    "404 - Not Found",
			expectedParams:          "squares",
			expectedUsecaseResponse: "Not Found",
			expectUsecaseCall:       true,
			wantError:               true,
			expectedError:           errors.New("Not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			u := mocks.NewMockUseCase(mockCtrl)

			if tt.expectUsecaseCall {
				u.EXPECT().GetShapeByName(tt.expectedParams).Return(tt.expectedUsecaseResponse, tt.expectedError)
			}

			c := New(u, l, r)

			response, err := c.useCase.GetShapeByName(tt.expectedParams)
			assert.Equal(t, response, tt.expectedUsecaseResponse)

			// Pro gamer tip: Use wantError to check if the error should be nil
			if tt.wantError {
				assert.NotNil(t, err)
				assert.Equal(t, err, tt.expectedError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
