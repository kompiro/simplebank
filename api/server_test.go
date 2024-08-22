package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/techschool/simplebank/db/mock"
)

func TestRootAPI(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	body := bytes.NewBuffer([]byte(""))

	store := mockdb.NewMockStore(ctrl)
	server := newTestServer(t, store)
	recorder := httptest.NewRecorder()

	url := "/"
	request, err := http.NewRequest(http.MethodGet, url, body)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	require.Equal(t, "Welcome!", recorder.Body.String())
}
