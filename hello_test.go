package main

import "testing"

func TestSomething(t *testing.T) {
	data = `{
		"mail": {}
	}`

	order, err := something(data)

	fmt.Println(order)

	if err == nil {
		t.Error("Expected no error, got ", err)
	}
}
