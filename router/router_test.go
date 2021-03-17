package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/varopxndx/bootcamp-workshop-w4/router/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//mockgen -source=router/router.go -destination=router/mocks/router.go -package=mocks

func Test_New(t *testing.T) {
	testCases := []struct {
		name           string
		endpoint       string
		handlerName    string
		status         int
		callController bool
	}{
		{
			name:           "OK, Get square",
			endpoint:       "/get/shapes/square",
			handlerName:    "GetShapeByName",
			status:         200,
			callController: true,
		},
		{
			name:           "404, Bad endpoint",
			endpoint:       "/badRequest",
			handlerName:    "GetShapeByName",
			status:         200,
			callController: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			c := mocks.NewMockController(mockCtrl)

			if tc.callController {
				c.EXPECT().GetShapeByName(gomock.Any(), gomock.Any()).Times(1)
			}

			r := New(c)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, tc.endpoint, nil)

			r.ServeHTTP(recorder, request)
			assert.Equal(t, tc.status, recorder.Code)
			assert.Nil(t, err)
		})
	}
}
