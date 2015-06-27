package main

import (
	"gopkg.in/stretchr/testify.v1/assert"
	"os"
	"testing"
)

func TestCapitalize(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Hello", Capitalize("hello"))
	assert.Equal("Hello", Capitalize("Hello"))
	assert.Equal("Hello world", Capitalize("hello world"))
	assert.Equal("Èllo", Capitalize("èllo"))
}

func TestParseHeaderList(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		[]string{"From", "To", "Subject"},
		ParseHeaderList("	from,to , subject"))
}

func TestExtractHeadersFromMail(t *testing.T) {
	assert := assert.New(t)

	email, err := os.Open("test.eml")
	if err != nil {
		t.Fatal(err)
	}

	expected := make(map[string]string)
	expected["From"] = "Fictious sender <sender@example.com>"
	expected["To"] = "\"Fictious recipient\" <recipient@example.com>"
	expected["Subject"] = "Example subject"
	expected["No_such_header"] = ""

	headerlist := ParseHeaderList("from,to,subject,no_such_header")
	headers, err := extractHeadersFromMail(email, headerlist)
	assert.Nil(err)
	assert.Equal(expected, headers)
}

func TestConstructMessageBody(t *testing.T) {
	assert := assert.New(t)

	headers := make(map[string]string)
	headers["From"] = "Fictious sender <sender@example.com>"
	headers["To"] = "\"Fictious recipient\" <recipient@example.com>"
	headers["Subject"] = "Example subject"

	expected_body := "From: Fictious sender <sender@example.com>\n" +
		"To: \"Fictious recipient\" <recipient@example.com>\n" +
		"Subject: Example subject"

	assert.Equal(expected_body, constructMessageBody(headers))
}
