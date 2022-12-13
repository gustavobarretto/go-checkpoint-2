package domain

type Patient struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	RG string `json:"rg"`
	RegistryDate string `json:"registry_date"`
}