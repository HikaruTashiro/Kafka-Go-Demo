package kafka_clients

type InvoiceListRequest struct {
	Limit        int    `json:"limit,omitempty"`
	Start        int    `json:"start,omitempty"`
	PaidAtFrom   string `json:"paid_at_from,omitempty"`
	PaidAtTo     string `json:"paid_at_to,omitempty"`
	Query        string `json:"query,omitempty"`
	StatusFilter string `json:"status_filter,omitempty"`
}

type AddressInfo struct {
	Street     string `json:"street,omitempty"`
	Number     string `json:"number,omitempty"`
	District   string `json:"district,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	State      string `json:"state,omitempty"`
	Zip_code   string `json:"zip_code,omitempty"`
	Complement string `json:"complement,omitempty"`
}

type InvoicePayerInfo struct {
	CnpjCpf     string      `json:"cnpj_cpf,omitempty"`
	Name        string      `json:"name,omitempty"`
	PhonePrefix string      `json:"phone_prefix,omitempty"`
	Phone       string      `json:"phone,omitempty"`
	Email       string      `json:"email,omitempty"`
	Address     AddressInfo `json:"address,omitempty"`
}
