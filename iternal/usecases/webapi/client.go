package webapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	urlForBreeds = "https://api.thedogapi.com/v1/breeds?limit=264&page=0" // 264 is not a random number. Its a max id in apis responce (when u get all breeds)
	urlForFacts = "https://dog-api.kinduff.com/api/facts"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		 return nil, errors.New("timeout cant be zero")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c *Client) FindBreed(letter byte) ([]string, error) {
	
	resp, err := c.client.Get(urlForBreeds) 
		if err != nil{
			return nil, err
		}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
		if err != nil{
			return nil, err
	}
	
	var (
		breeds BreedInfo
		dogList []string = make([]string, 0)
	)
	if err := json.Unmarshal(body, &breeds); err != nil{
		return nil, err
	}
	for _, breed := range breeds{
		if breed.Name[0] == letter{
			dogList = append(dogList, breed.Name)
		}
	}
	return dogList, nil
}

func (c *Client) BreedInfo(breed string) (string,string, error) {
	resp, err := c.client.Get(urlForBreeds) 
		if err != nil{
			return "", "", err
		}
	defer resp.Body.Close()	

	body, err := io.ReadAll(resp.Body)
		if err != nil{
			return "", "", err
		}

	var breeds BreedInfo
	if err := json.Unmarshal(body, &breeds); err != nil{
		return "", "", err
	}
	
	for _, value := range breeds{
		if breed == value.Name{
			return fmt.Sprintf("%d %s",value.ID, value.Bredfor),value.Image.Url, nil
		}
	}
	return "Please enter the dog breed corectly", "",errors.New("dont found this breed")
}

func (c *Client) RndFact() (string, error) {
	resp, err := c.client.Get(urlForFacts) 
		if err != nil{
			return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
		if err != nil{
			return "", err
		}

	var facts Fact
	if err := json.Unmarshal(body, &facts); err != nil{
		return "", err
	}
	return facts.Fact[0], nil
}