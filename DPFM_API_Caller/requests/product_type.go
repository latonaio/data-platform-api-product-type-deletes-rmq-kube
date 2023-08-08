package requests

type ProductType struct {
	ProductType         string `json:"ProductType"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
