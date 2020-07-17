package data

type Thing struct {
	Name    string `json:"Name" db:"full_name"`
	Address bool   `json:"addr" db:"address"`
	Phone   string `json:"phone" db:"phone_number"`
}
