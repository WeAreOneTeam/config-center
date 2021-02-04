package model

type UpdateProfileBody struct{
	Value string `json:"value"`

	Description string `json:"description"`

	Status string `json:"status"`
}
