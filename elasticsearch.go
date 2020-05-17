package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	es7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// Client --
type Client struct {
	esClient *es7.Client
}

// NewClient --
func NewClient(config es7.Config) *Client {
	esClient, err := es7.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{esClient: esClient}
}

// Index --
func (client *Client) Index(index string, id string, doc interface{}) (*IndexResponse, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	if err := json.NewEncoder(buf).Encode(doc); err != nil {
		return nil, err
	}

	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       buf,
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), client.esClient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	indexResponse := &IndexResponse{}
	if err := json.NewDecoder(res.Body).Decode(indexResponse); err != nil {
		return nil, err
	}
	indexResponse.HTTPStatusCode = res.StatusCode
	return indexResponse, nil
}

// Create --
func (client *Client) Create(index string, id string, doc interface{}) (*IndexResponse, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	if err := json.NewEncoder(buf).Encode(doc); err != nil {
		return nil, err
	}

	req := esapi.CreateRequest{
		Index:      index,
		DocumentID: id,
		Body:       buf,
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), client.esClient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	indexResponse := &IndexResponse{}
	if err := json.NewDecoder(res.Body).Decode(indexResponse); err != nil {
		return nil, err
	}
	indexResponse.HTTPStatusCode = res.StatusCode
	return indexResponse, nil
}

// Update --
func (client *Client) Update(index string, id string, doc interface{}) (*IndexResponse, error) {
	query := make(map[string]interface{})
	query["doc"] = doc
	buf := bytes.NewBuffer(make([]byte, 0))
	if err := json.NewEncoder(buf).Encode(query); err != nil {
		return nil, err
	}

	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: id,
		Body:       buf,
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), client.esClient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	indexResponse := &IndexResponse{}
	if err := json.NewDecoder(res.Body).Decode(indexResponse); err != nil {
		return nil, err
	}
	indexResponse.HTTPStatusCode = res.StatusCode
	return indexResponse, nil
}

// Get --
func (client *Client) Get(index string, id string, response interface{}) error {
	req := esapi.GetRequest{
		Index:      index,
		DocumentID: id,
	}
	res, err := req.Do(context.Background(), client.esClient)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return err
	}
	return nil
}

// Delete --
func (client *Client) Delete(index string, id string) (*IndexResponse, error) {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), client.esClient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	indexResponse := &IndexResponse{}
	if err := json.NewDecoder(res.Body).Decode(indexResponse); err != nil {
		return nil, err
	}
	indexResponse.HTTPStatusCode = res.StatusCode
	return indexResponse, nil
}

// Search --
func (client *Client) Search(index string, query interface{}, response interface{}) error {
	buf := bytes.NewBuffer(make([]byte, 0))
	if err := json.NewEncoder(buf).Encode(query); err != nil {
		return err
	}

	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  buf,
	}
	res, err := req.Do(context.Background(), client.esClient)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return err
	}
	return nil
}
