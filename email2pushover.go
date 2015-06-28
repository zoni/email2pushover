package main

import (
	"bytes"
	"fmt"
	"github.com/famz/RFC2047"
	"github.com/gregdel/pushover"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"log"
	"net/mail"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

const VERSION = "1.0.1"

var (
	app       = kingpin.New("email2pushover", "email2pushover sends pushover notifications from mail read on stdin.")
	headers   = app.Flag("headers", "Comma-separated list of headers to display in notification").Short('H').Default("subject,from").String()
	title     = app.Flag("title", "The notification title").Short('T').Default("Email").String()
	token     = app.Flag("token", "Your application token").Short('t').Required().String()
	recipient = app.Flag("recipient", "Recipient's key (may be a user or delivery group)").Short('r').Required().String()
)

// Capitalize converts the first character of the given string to uppercase.
func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

// ParseHeaderList parses a comma-separated list of headers into a slice
// of prettified headers.
func ParseHeaderList(s string) []string {
	headers := make([]string, 0)
	for _, header := range strings.Split(s, ",") {
		header = Capitalize(strings.TrimSpace(header))
		headers = append(headers, header)
	}
	return headers
}

// extractHeadersFromMail reads an email message and returns a map of header->value
// items from the given list of headers. Any headers that are missing from the email
// will be returned with an empty string value.
func extractHeadersFromMail(r io.Reader, headers []string) (extracted map[string]string, err error) {
	extracted = make(map[string]string)

	message, err := mail.ReadMessage(r)
	if err != nil {
		return
	}

	for _, header := range headers {
		extracted[header] = RFC2047.Decode(message.Header.Get(header))
	}
	return
}

// constructMessageBody creates a formatted message body from the supplied headers
// in the given order.
func constructMessageBody(headers map[string]string, order []string) string {
	var buffer bytes.Buffer
	for _, header := range order {
		value, present := headers[header]
		if !present {
			panic(fmt.Sprintf("Header '%s' missing from headers", header))
		}
		buffer.WriteString(fmt.Sprintf("%s: %s\n", header, value))
	}
	return strings.TrimSpace(buffer.String())
}

func main() {
	app.Version(VERSION)
	kingpin.MustParse(app.Parse(os.Args[1:]))
	headerlist := ParseHeaderList(*headers)

	headers, err := extractHeadersFromMail(os.Stdin, headerlist)
	if err != nil {
		log.Fatal(err)
	}

	pushover_app := pushover.New(*token)
	pushover_recipient := pushover.NewRecipient(*recipient)
	pushover_message := &pushover.Message{
		Message: constructMessageBody(headers, headerlist),
		Title:   *title,
	}

	_, err = pushover_app.SendMessage(pushover_message, pushover_recipient)
	if err != nil {
		log.Fatal(err)
	}
}
