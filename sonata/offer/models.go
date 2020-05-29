package offer

type FindResponse struct {
	Data []*ProductOffering `json:"data,omitempty"`
	Mata *Meta              `json:"mata,omitempty"`
}

type GetResponse struct {
	Data *ProductOffering `json:"data,omitempty"`
}

type Meta struct {
	Total   int `json:"total,omitempty"`
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
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
