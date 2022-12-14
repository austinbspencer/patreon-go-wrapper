package patreon

// AddressFields is all fields in the Address Attributes struct
var AddressFields = getObjectFields(AddressAttributes{})

// Address represents a Patreon's shipping address.
type Address struct {
	Type          string            `json:"type"`
	ID            string            `json:"id"`
	Attributes    AddressAttributes `json:"attributes"`
	Relationships struct {
		Campaigns *CampaignsRelationship `json:"campaigns,omitempty"`
		User      *UserRelationship      `json:"user,omitempty"`
	} `json:"relationships"`
}

// AddressAttributes is the attributes struct for the Address
type AddressAttributes struct {
	Addressee   string   `json:"addressee"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	CreatedAt   NullTime `json:"created_at"`
	Line1       string   `json:"line_1"`
	Line2       string   `json:"line_2"`
	PhoneNumber string   `json:"phone_number"`
	PostalCode  string   `json:"postal_code"`
	State       string   `json:"state"`
}
