package caesar_test

import (
	"appital-service/business/core/models"
	"appital-service/caesar"
	"testing"
)

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

func Test_Encrypter(t *testing.T) {
	t.Log("Given the to talk to encrypt a text")
	{
		t.Logf("\tWhen handling a key of zero value")
		{
			data := models.NewEncoding{
				Text: "Let's carve him as a dish fit for the gods.",
				Key:  0,
			}

			encodedStr, err := caesar.Encrypter(data)
			if err == nil {
				t.Logf("\t%s\tShould return error for a key zero value", failed)
			}

			if encodedStr != "" {
				t.Logf("\t\tExp: %s", "")
				t.Logf("\t\tgot: %s", encodedStr)
				t.Fatalf("\t%s\tShould return empty string", failed)
			}
			t.Logf("\t%s\tShould not be able to encrpt a text", success)
		}

		t.Logf("\tWhen handling a key of value 1 or greater")
		{
			data := models.NewEncoding{
				Text: "Let's carve him as a dish fit for the gods.",
				Key:  14,
			}

			encodedStr, err := caesar.Encrypter(data)
			if err != nil {
				t.Logf("\t%s\tShould not return error for a key greater than zero", failed)
			}

			expectedStr := "Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg."
			if encodedStr != expectedStr {
				t.Logf("\t\tExp: %s", expectedStr)
				t.Logf("\t\tgot: %s", encodedStr)
				t.Fatalf("\t%s\tShould return a valid encrypted string", failed)
			}

			t.Logf("\t%s\tShould be able to encrpt a text", success)
		}
	}
}

func Test_Decrypter(t *testing.T) {
	t.Log("Given the to talk to decrypt a text")
	{
		t.Logf("\tWhen handling a key of zero value")
		{
			data := models.NewDecoding{
				Text: "Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg.",
				Key:  0,
			}

			encodedStr, err := caesar.Decrypter(data)
			if err == nil {
				t.Logf("\t%s\tShould return error for a key zero value", failed)
			}

			if encodedStr != "" {
				t.Logf("\t\tExp: %s", "")
				t.Logf("\t\tgot: %s", encodedStr)
				t.Fatalf("\t%s\tShould return empty string", failed)
			}
			t.Logf("\t%s\tShould not be able to decrypt a string", success)
		}

		t.Logf("\tWhen handling a key of value 1 or greater")
		{
			data := models.NewDecoding{
				Text: "Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg.",
				Key:  14,
			}

			decodedStr, err := caesar.Decrypter(data)
			if err != nil {
				t.Logf("\t%s\tShould not return error for a key greater than zero", failed)
			}

			expectedStr := "Let's carve him as a dish fit for the gods."
			if decodedStr != expectedStr {
				t.Logf("\t\tExp: %s", expectedStr)
				t.Logf("\t\tgot: %s", decodedStr)
				t.Fatalf("\t%s\tShould return a valid decrypted string", failed)
			}

			t.Logf("\t%s\tShould be able to decrypt a string", success)
		}
	}
}
