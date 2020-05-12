all api ... have end sentense

we could generate our structured data for describing the api(web api)

```go
type apiMethod struct {
	method               string
	host                 string
	path                 string
	apiMethodSummary     string
	apiMethodDescription string
	apiMethodSpec
}
type apiMethodSpec struct {
	apiMethodRequest
	apiMethodResponse
}
type apiMethodRequest struct {
	apiMethodPathParameters  requestField
	apiMethodHeaders         requestField
	apiMethodQueryParameters requestField
}
type requestField struct {
	apiMethodParameter
}
type apiMethodParameter struct {
	name     string
	_type    string
	required bool
	text     string
}

type apiMethodResponse struct {
	apiMethodResponseExample
}
type apiMethodResponseExample struct {
	httpCode                            int
	apiMethodResponseExampleDescription string
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


