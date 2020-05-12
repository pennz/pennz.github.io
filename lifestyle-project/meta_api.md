all api ... have end sentense

we could generate our structured data for describing the api(web api)

```go
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
```

so the structure doc-{[api-method:{method:,host,path,api-method-summary,api-method-description,api-method-spec:{api-method-quest:{}, api-method-response:{[api-method-response-example:{response-example}]}}}]}

- api-method method host path
    - api-method-summary
        - SUMMARY
    - api-method-description
        - DESCRIPTION
    - api-method-spec
        - api-method-request /end
            - api-method-path-parameters
                - api-method-parameter name type required 
                    - DESCRIPTION
            - api-method-headers
                - api-method-parameter name type required 
                    - DESCRIPTION

            - api-method-query-parameters
                - api-method-parameter 
                    - DESCRIPTION
                - api-method-parameter 
                    - DESCRIPTION

        - api-method-response /end
            - api-method-response-example
                - api-method-response-example-description
                - EXAMPLE_STRING
            - api-method-response-example
                - api-method-response-example-description
                - EXAMPLE_STRING


