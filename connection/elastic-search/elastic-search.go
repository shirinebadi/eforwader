package elasticsearch

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testAWS/internal/config"

	"testAWS/internal/model"
)

type ElasticSearchClient struct {
	Config     config.ElasticConfig
	DomainName string
}

func (es *ElasticSearchClient) QueryRule(index string) (*model.Rule, error) {
	client := &http.Client{}

	rule := &model.Rule{}
	req, err := http.NewRequest("GET", string(es.Config.Hosts+index+"/_search?pretty=true"), nil)
	if err != nil {
		return rule, err
	}
	req.SetBasicAuth(es.Config.Username, es.Config.Password)

	resp, err := client.Do(req)
	if err != nil {
		return rule, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var response model.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		log.Println(string(body))
		return rule, err
	}

	return &response.Hits.Hits[0].Source, nil
}
