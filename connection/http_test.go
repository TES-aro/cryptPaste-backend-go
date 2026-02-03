package connection

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type handlerTest struct {
	request *http.Request
	wantCode int
}

type handlerTestArray struct {
	testArray []handlerTest
}

func (h *handlerTestArray) newTest(method, url string, want int){
	request , err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	test := handlerTest{request, want }
	h.testArray = append(h.testArray, test)
}
/*
func TestHandlerFunctions(t *testing.T) {
	tests := handlerTestArray{}
	tests.newTest("GET", "/api/12345678", 200)
	tests.newTest("GET","/api/1234", 404)
	tests.newTest("POST", "/api/12345678", 405)

	for _, tt := range tests.testArray {
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleGET)
		handler.ServeHTTP(rr, tt.request)
		if rr.Code != tt.wantCode {
			t.Errorf("Wanted status %d; got %d",tt.wantCode, rr.Code)
			continue
		}
		fmt.Println(rr.Code , rr.Body)
	}
}
*/
func TestMux(t *testing.T) {
	tests := handlerTestArray{}
	testMux := createMux()

	for _, tt := range tests.testArray {
		rr := httptest.NewRecorder()
		testMux.ServeHTTP(rr, tt.request)
		if rr.Code != tt.wantCode {
			t.Errorf("Wanted status ‰d; got %d", tt.wantCode, rr.Code)
		}
}

// ######
// some automation tools
// #####
