package models

type Correspondance struct {
	AddressType  string `json:"addressType"`
	Number       string `json:"number"`
	Way          string `json:"way"`
	Complement   string `json:"complement"`
	City         string `json:"city"`
	PostalCode   string `json:"postalCode"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber  string `json:"phoneNumber"`
	FixNumber    string `json:"fixNumber"`
	FaxNumber    string `json:"faxNumber"`
}
