package handlers

import (
	"fmt"
	"net/http"
)

func sendError(w http.ResponseWriter, statusCode int, errorCode string, message string) {
	w.Header().Set("Content-Type", "application/xml")

	w.WriteHeader(statusCode)

	errorXML := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Error>
    <Code>%s</Code>
    <Message>%s</Message>
</Error>`, errorCode, message)

	fmt.Fprint(w, errorXML)
}
