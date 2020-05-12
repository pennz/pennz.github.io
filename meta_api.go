package main

import "fmt"

type APIMethod struct {
	method               string
	host                 string
	path                 string
	APIMethodSummary     string
	APIMethodDescription string
	APIMethodSpec
}
type APIMethodSpec struct {
	APIMethodRequest
	APIMethodResponse
}
type APIMethodRequest struct {
	APIMethodPathParameters  requestField
	APIMethodHeaders         requestField
	APIMethodQueryParameters requestField
}
type requestField struct {
	APIMethodParameter
}
type APIMethodParameter struct {
	name     string
	_type    string
	required bool
	text     string
}

type APIMethodResponse struct {
	APIMethodResponseExample
}
type APIMethodResponseExample struct {
	httpCode                            int
	APIMethodResponseExampleDescription string
}

func main() {
	fmt.Println("vim-go")
}
