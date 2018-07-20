package main

import (
	// "fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	data := `{
		"mail": {
			"headers": [
			  {
					"name": "To",
					"value": "something"
				}
			],
			"common_headers": {
				"from": [
         "John Doe <john@example.com>"
				],
				"to": [
         "John Doe <john@example.com>"
				]
			}
		},
		"content": "When Monday, Jul 9th BETWEEN 11:45am-12:15pm Your Order"
	}`

	_, err := something(data)

	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}
