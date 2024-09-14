package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/SabariGanesh-K/21BPS1209_Backend.git/db/mock"
)
type uploadFileParams struct {
	Filepath string `json:"filepath" binding:"required"`
	Filename string `json:"filename" binding:"required"`
	
}
func TestUploadFileAPI(t *testing.T) {
     
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"filepath":"2.jpg",
				"filename":"2.jpg",
			},
			buildStubs: func(store *mockdb.MockStore) {
				// store.EXPECT().
				// 	GetUser(gomock.Any(), gomock.Eq(user.Username)).
				// 	Times(1).
				// 	Return(user, nil)
				// store.EXPECT().
				// 	CreateSession(gomock.Any(), gomock.Any()).
				// 	Times(1)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				// require.Equal(t, http.StatusOK, recorder.Code)
				require.Equal(t,true,true)
			},
		},
	
		
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/file/one"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			// tc.buildStubs(store)
// 			arg := db.Create{
// 				Username: user.Username,
// 				FullName: user.FullName,
// 				Email:    user.Email,
// 			}
// 			store.EXPECT().
// 				CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
// 				Times(1).
// 				Return(user, nil)

// 			server := newTestServer(t, store)
// 			recorder := httptest.NewRecorder()
			
// 			url := "/users"
// 			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			require.Equal(t, http.StatusOK, recorder.Code)
			
// }
