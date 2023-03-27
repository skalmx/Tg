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

// func test(){
// 	resp, err := http.Get("https://api.thedogapi.com/v1/breeds?limit=100&page=0")
// 		if err != nil{
// 			log.Print(err)
// 		}
	
// 	body, err := io.ReadAll(resp.Body)
// 		if err != nil{
// 			log.Print(err)
// 		}
// 	fmt.Println(string(body))	
	
// 	var responce BreedInfo
// 	if err := json.Unmarshal(body, &responce); err != nil{
// 		log.Print(err)
// 	}
// 	for _, r := range responce{
// 		fmt.Print(r.Lifetime)
// 	}
// }