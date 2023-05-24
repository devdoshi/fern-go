package ir

type ExampleEndpointCall struct {
	Docs                   *string                  `json:"docs"`
	Name                   *Name                    `json:"name"`
	Url                    string                   `json:"url"`
	RootPathParameters     []*ExamplePathParameter  `json:"rootPathParameters"`
	ServicePathParameters  []*ExamplePathParameter  `json:"servicePathParameters"`
	EndpointPathParameters []*ExamplePathParameter  `json:"endpointPathParameters"`
	ServiceHeaders         []*ExampleHeader         `json:"serviceHeaders"`
	EndpointHeaders        []*ExampleHeader         `json:"endpointHeaders"`
	QueryParameters        []*ExampleQueryParameter `json:"queryParameters"`
	Request                *ExampleRequestBody      `json:"request"`
	Response               *ExampleResponse         `json:"response"`
}