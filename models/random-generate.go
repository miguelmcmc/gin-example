package models

type RandomGenerate struct {
	Seed   string `json:"seed"`
	Lenght int    `json:"lenght" validate:"required"`
}

func (rg *RandomGenerate) Validate() error {
	return ValidateStruct(rg)
}
