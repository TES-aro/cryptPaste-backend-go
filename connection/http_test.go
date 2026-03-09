package connection

import (
	"bytes"
	"dataBase"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// en tiedä miten tekisin nätimmin :/
var workingID string

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

func (h *handlerTestArray) newPost(language, text string, want int){
	body := JSONpaste{language, text}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("error creating newPost: %v \n", err.Error())
		return
	}
	bodyBuffer := bytes.NewBuffer(jsonBody)
	req, err := http.NewRequest("POST", "/api",
		bodyBuffer)
	req.Header.Set("Content-Type", "application/json")
	newTest := handlerTest{req, want}
	h.testArray = append(h.testArray, newTest)
}

func postBody(language, text, url string) *http.Request{
	body := JSONpaste{language, text }
	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("error creating JSON for POST request: %v \n", err.Error())
		return nil
	}
	bodyBuffer := bytes.NewBuffer(jsonBody)
	req, err := http.NewRequest(http.MethodPost, url, bodyBuffer)
	if err != nil {
		fmt.Printf("error creating POST req: %v \n", err.Error())
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	return req
}

func TestInitDB(t *testing.T){
	db, err := dataBase.ConnectToDB("cryptTest")
	if err != nil {
		fmt.Println("  BIG ERROR INITING DB!") 
		t.FailNow()
	}
	dataBase.InitTable(db)
}
func TestPostFunctions(t *testing.T){
	db, err := dataBase.ConnectToDB("cryptTest")
	if err != nil {
		t.Fatal("couldn't connect to database")
	}
	defer db.Close()

	POSThandler := POSThandler{dbPointer: db}
	handlePOST := POSThandler.ServeHTTP

	ta := handlerTestArray{}
	ta.newPost("java", "some random text here", 200)
	ta.newPost("aivan liian pitkä nimi ohjelmointikielelle menee tänne", "mitä väliä?", 400)
	ta.newPost("", "", 200)

	for _, tt := range ta.testArray {
		rr := httptest.NewRecorder()
		//kieron kierteinen tämä, mutta handlerit hämmentää.
		handler := http.HandlerFunc(handlePOST)
		handler.ServeHTTP(rr, tt.request)
	}

	//erityistapaus, joka otetaan tateen

	req := postBody("go","fmt.Println(\"hello world\")", "/api")

	record := httptest.NewRecorder()
	handler := http.HandlerFunc(handlePOST)
	handler.ServeHTTP(record, req)

	workingID = decode(record.Body)

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
		if rr.Code != tt.wantCode {
			t.Errorf("Wanted status %d; got %d",tt.wantCode, rr.Code)
			continue
		}
	}
	//special working case
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleGET)

	workingUrl := "/api/"+workingID
	fmt.Printf("working GET url: %s\n", workingUrl)
	req, err := http.NewRequest(http.MethodGet,
		workingUrl, nil)
	if err != nil {
		t.Fatal("failed GET request") 
	}
	handler.ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Error("improper code for fetching existing entry")
	}
}


func TestMux(t *testing.T) {
	db, err := dataBase.ConnectToDB("cryptTest")
	if err != nil {
		t.Fatal("no database connection")
	}
	defer db.Close()
	tests := handlerTestArray{}
	testMux := CreateMux("/api", "/api/", db)

	tests.newPost("lorem ipsum", "and still working!", 200)
	tests.newTest("POST", "/api/"+workingID, 405)
	tests.newTest("GET", "/api/"+workingID, 200)
	tests.newTest("GET","/api", 405) 

	for _, tt := range tests.testArray {
		rr := httptest.NewRecorder()
		testMux.ServeHTTP(rr, tt.request)
		if rr.Code != tt.wantCode {
			t.Error("Wanted status " + strconv.Itoa(tt.wantCode) +",  got " + strconv.Itoa(rr.Code))
		}
	}
}


func TestCleanDB(t *testing.T) {
	db, err := dataBase.ConnectToDB("cryptTest")
	if err != nil {
		t.Fatalf("failed to connect to db\n")
	}
	defer db.Close()
	db.Exec( "DROP TABLES pasteCrypt")
}
// ######
// some automation tools
// #####
func decode(resp *bytes.Buffer) string {
	bytes := resp.Bytes()
	var jsonID JSONid
	err := json.Unmarshal(bytes, &jsonID)
	if err != nil {
		fmt.Println("error decoding buffer: " + err.Error())
	}
	fmt.Printf("decoded buffer to: %#v \n", jsonID)
	return jsonID.ID
}
