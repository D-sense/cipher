package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"service/api/handler"
	"service/business/core/models"
	"strings"
	"testing"
)

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

// Test_Encypter is the entry point for testing Encrypter functions.
func Test_Encypter_Decrypter(t *testing.T) {

	t.Parallel()

	t.Run("encrypt200", encrypt200)
	t.Run("decrypt200", decrypt200)
}

// encrypt200 validates deleting a user that does not exist is not a failure.
func encrypt200(t *testing.T) {

	t.Log("Given the need to encrypt a valid request")
	{
		newEncoder := models.NewEncoding{
			Text: "Let's carve him as a dish fit for the gods.",
			Key:  14,
		}
		body, err := json.Marshal(&newEncoder)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/api/v1/encrypt", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()

		h := http.HandlerFunc(handler.Encrypt)
		h.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Fatalf("\t%s\tShould receive a status code of 200 for the response : %v", failed, resp.Body)
		}
		t.Logf("\t%s\tShould receive a status code of 404 for the response.", success)

		got := resp.Body.String()
		expBody := "{\"status\":200,\"data\":\"Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg.\"}"
		if !strings.EqualFold(got, expBody) {
			t.Logf("\t\tGot : %v", got)
			t.Logf("\t\tExp: %v", expBody)
			t.Fatalf("\t%s\tShould get the expected result.", failed)
		}
		t.Logf("\t%s\tShould get the expected result.", success)
	}

	t.Log("Given the need to encrypt an invalid request")
	{
		newEncoder := models.NewEncoding{
			Text: "Let's carve him as a dish fit for the gods.",
			Key:  0,
		}
		body, err := json.Marshal(&newEncoder)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/api/v1/encrypt", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()

		h := http.HandlerFunc(handler.Encrypt)
		h.ServeHTTP(resp, req)

		if resp.Code == http.StatusOK {
			t.Fatalf("\t%s\tShould receive a status code of 400 for the response : %v", failed, resp.Body)
		}
		t.Logf("\t%s\tShould receive a status code of 400 for the response.", success)

		got := resp.Body.String()
		expBody := "{\"status\":400,\"error\":\"validating data: [{\\\"field\\\":\\\"key\\\",\\\"error\\\":\\\"key must be 1 or greater\\\"}]\"}"
		if !strings.EqualFold(got, expBody) {
			t.Logf("\t\tGot : %v", got)
			t.Logf("\t\tExp: %v", expBody)
			t.Fatalf("\t%s\tShould get error message", failed)
		}
		t.Logf("\t%s\tShould get an error message", success)
	}
}

// deleteUserNotFound validates deleting a user that does not exist is not a failure.
func decrypt200(t *testing.T) {

	t.Log("Given the need to decrypt a valid request")
	{
		newEncoder := models.NewDecoding{
			Text: "Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg.",
			Key:  14,
		}
		body, err := json.Marshal(&newEncoder)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/api/v1/decrypt", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()

		h := http.HandlerFunc(handler.Decrypt)
		h.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Fatalf("\t%s\tShould receive a status code of 200 for the response : %v", failed, resp.Body)
		}
		t.Logf("\t%s\tShould receive a status code of 404 for the response.", success)

		got := resp.Body.String()
		expBody := "{\"status\":200,\"data\":\"Let's carve him as a dish fit for the gods.\"}"
		if !strings.EqualFold(got, expBody) {
			t.Logf("\t\tGot : %v", got)
			t.Logf("\t\tExp: %v", expBody)
			t.Fatalf("\t%s\tShould get the expected result.", failed)
		}
		t.Logf("\t%s\tShould get the expected result.", success)
	}

	t.Log("Given the need to decrypt an invalid request")
	{
		newEncoder := models.NewDecoding{
			Text: "Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg.",
			Key:  0,
		}
		body, err := json.Marshal(&newEncoder)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, "/api/v1/decrypt", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()

		h := http.HandlerFunc(handler.Decrypt)
		h.ServeHTTP(resp, req)

		if resp.Code == http.StatusOK {
			t.Fatalf("\t%s\tShould receive a status code of 400 for the response : %v", failed, resp.Body)
		}
		t.Logf("\t%s\tShould receive a status code of 400 for the response.", success)

		got := resp.Body.String()
		expBody := "{\"status\":400,\"error\":\"validating data: [{\\\"field\\\":\\\"key\\\",\\\"error\\\":\\\"key must be 1 or greater\\\"}]\"}"
		if !strings.EqualFold(got, expBody) {
			t.Logf("\t\tGot : %v", got)
			t.Logf("\t\tExp: %v", expBody)
			t.Fatalf("\t%s\tShould get error message", failed)
		}
		t.Logf("\t%s\tShould get an error message", success)
	}
}
