package webapi

type BreedInfo  []struct {
	W Weight `json:"weight"`
	H Height `json:"height"`
	ID int `json:"id"`
	Name string `json:"name"`
	Bredfor string `json:"bred_for"`
	Breedgr string `json:"breed_group"`
	Lifetime string `json:"life_span"`
	Temperament string `json:"temperament"`
	Origin string `json:"origin"`
	ReferenceId string `json:"reference_image_id"`
	Image Image `json:"image"`
}
type Weight struct {
	Imperial  string `json:"imperial"`
	Metric string `json:"metric"`
}
type Height struct {
	Imperial  string `json:"imperial"`
	Metric string `json:"metric"`
}

type Image struct{
	ID string `json:"id"`
	Width int `json:"width"`
	Height int `json:"height"`
	Url  string `json:"url"`
}

type Fact struct{
	Fact []string `json:"facts"`
	OK bool `json:"success"`
}
