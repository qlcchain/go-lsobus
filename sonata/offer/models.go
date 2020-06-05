package offer

type FindResponse struct {
	/*The number of resources retrieved in the response
	 */
	XResultCount string
	/*The total number of matching resources
	 */
	XTotalCount string

	Payload []*ProductOffering
}

type GetResponse struct {
	Payload *ProductOffering
}

type ProductOffering struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Deleted bool     `json:"deleted,omitempty"`
	Product *Product `json:"product,omitempty"`
}

type Product struct {
	ID             string         `json:"id,omitempty"`
	Name           string         `json:"name,omitempty"`
	Deleted        bool           `json:"deleted,omitempty"`
	Type           string         `json:"type,omitempty"`
	Code           string         `json:"code,omitempty"`
	SKU            string         `json:"sku,omitempty"`
	Provider       string         `json:"provider,omitempty"`
	State          string         `json:"state,omitempty"`
	QuoteSpecs     *Specification `json:"quoteSpecs,omitempty"`
	ProvisionSpecs *Specification `json:"provisionSpecs,omitempty"`
}

type Specification struct {
	Required   []string            `json:"required,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
}

type Property struct {
	Type        string   `json:"type,omitempty"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
}
