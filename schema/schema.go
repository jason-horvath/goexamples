package schema

type ExampleHtmlData struct {
	Heading     string
	Description string
	ListItems   []string
}

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ShoeSize  int    `json:"shoe_size"`
	HasOrder  bool   `json:"has_order"`
}

type OutputJason struct {
	Message      string `json:"message"`
	SlicePrinted string `json:"slice_printed"`
	CustomerData []Customer
}
