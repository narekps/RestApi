package models

type Data struct {
	A      int `json:"A"`
	B      int `json:"B"`
	C      int `json:"C"`
	Nroots int `json:"Nroots"`
}

// Calculate data result
func (data *Data) Calculate() {
	d := (data.B * data.B) - (4 * data.A * data.C)
	if d > 0 {
		data.Nroots = 2
	} else if d == 0 {
		data.Nroots = 1
	} else {
		data.Nroots = 0
	}
}
