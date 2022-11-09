package handler

import (
	"appital-service/business/core/models"
	"appital-service/caesar"
	"appital-service/foundation/web"
	"net/http"
)

// Encrypt encrypts a string using the Caesar Cipher.
func Encrypt(w http.ResponseWriter, r *http.Request) {
	// unmarshalling data into struct
	var ne models.NewEncoding
	if err := web.Decode(r, &ne); err != nil {
		resp := web.Error{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		web.Respond(w, resp, http.StatusBadRequest)
		return
	}

	encoded, err := caesar.Encrypter(ne)
	if err != nil {
		resp := web.Error{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		web.Respond(w, resp, http.StatusBadRequest)
		return
	}

	statusCode := http.StatusOK
	resp := web.Body{
		Status: statusCode,
		Data:   encoded,
	}
	web.Respond(w, resp, statusCode)
}

// Decrypt decrypts a string using the Caesar Cipher.
func Decrypt(w http.ResponseWriter, r *http.Request) {
	// unmarshalling data into struct
	var ne models.NewDecoding
	if err := web.Decode(r, &ne); err != nil {
		resp := web.Error{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		web.Respond(w, resp, http.StatusBadRequest)
		return
	}

	decoded, err := caesar.Decrypter(ne)
	if err != nil {
		resp := web.Error{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		web.Respond(w, resp, http.StatusBadRequest)
		return
	}

	statusCode := http.StatusOK
	resp := web.Body{
		Status: statusCode,
		Data:   decoded,
	}
	web.Respond(w, resp, http.StatusOK)
}
