package connection

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"dataBase"
	"testing"
)

type handlerTest struct {
	request *http.Request
	wantCode int
}

type postTest struct {
	request *http.Request
	wantCode int
	language string
	content string
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

func TestPostFunctions(t *testing.T){
	db, err := dataBase.ConnectToDB("cryptTest")
	if err != nil {
		t.Fatal("couldn't connect to database")
	}
	defer db.Close()

	POSThandler := POSThandler{dbPointer: db}
	handlePOST := POSThandler.ServeHTTP



}

func TestHandlerFunctions(t *testing.T) {
	tests := handlerTestArray{}
	tests.newTest("GET", "/api/12345678", 404)
	tests.newTest("GET","/api/1234", 404)
	tests.newTest("POST", "/api/12345678", 405)
	tests.newTest("PUT", "/api", 405)
	tests.newTest("DELETE", "/api/1234567", 405)


	db , err:= dataBase.ConnectToDB("cryptTest")
	if err != nil {
		t.Fatal("no database connection")
	}
	defer db.Close()

	getHandler := GEThandler{dbPointer: db }
	handleGET := getHandler.ServeHTTP 

	for _, tt := range tests.testArray {
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleGET)
		handler.ServeHTTP(rr, tt.request)
		fmt.Println(rr.Body) 
		if rr.Code != tt.wantCode {
			t.Errorf("Wanted status %d; got %d",tt.wantCode, rr.Code)
			continue
		}
		fmt.Println(rr.Code , rr.Body)
	}
}

/*
func TestMux(t *testing.T) {
	db, err := dataBase.ConnectToDB("cryptTest")
	if err != nil {
		t.Fatal("no database connection")
	}
	defer db.Close()
	tests := handlerTestArray{}
	testMux := createMux("api", "api/", db)

	for _, tt := range tests.testArray {
		rr := httptest.NewRecorder()
		testMux.ServeHTTP(rr, tt.request)
		if rr.Code != tt.wantCode {
			t.Errorf("Wanted status ‰d; got %d", tt.wantCode, rr.Code)
		}
	}
}

*/
// ######
// some automation tools
// #####
