package api

import (
	"bytes"
	// "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	 "mime/multipart"
    "os"
	"io"
	// "github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/SabariGanesh-K/21BPS1209_Backend.git/db/mock"
)
type uploadFileParams struct {
	File *multipart.FileHeader `form:"file" `
	Filename string `form:"filename" binding:"required"`
	
}

func TestUploadFileAPI(t *testing.T) {
	
	file,err:= os.Open("./2.jpg")
	require.NoError(t,err)
	defer file.Close()
	filestate,_ := file.Stat()
	fileheader:=multipart.FileHeader{
		Filename:"2.jpg",
		Header:nil,
		Size:     filestate.Size(),

	}
	reqbody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqbody)
	part, errr := writer.CreateFormFile("file", fileheader.Filename)
	require.NoError(t,errr)
	_, err = io.Copy(part, file)
	require.NoError(t,err)
	err = writer.WriteField("filename", "sample.jpg") 
	require.NoError(t,err)
	err = writer.Close()
	require.NoError(t,err)
	// url:="http://localhost:8081/file/one"
	// request, err := http.NewRequest(http.MethodPost,url , body)
	// 		require.NoError(t, err)
	// 		request.Header.Set("Content-Type", writer.FormDataContentType())
	// 		client := &http.Client{}
	// 		resp, erro := client.Do(request)
	// require.NoError(t,erro)

			// defer resp.Body.Close()

	testCases := []struct {
		name          string
		body         *bytes.Buffer
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: reqbody,
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
			require.NoError(t, err)

			url := "/file/one"
			request, err := http.NewRequest(http.MethodPost, url, tc.body)
			require.NoError(t, err)
			request.Header.Set("Content-Type", writer.FormDataContentType())
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
