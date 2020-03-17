package elasticsearch

import (
	"log"
	"testing"

	es7 "github.com/elastic/go-elasticsearch/v7"
)

type Test struct {
	Int    int     `json:"int"`
	Float  float64 `json:"float"`
	String string  `json:"string"`
	Slice  []int   `json:"slice"`
}

type TestDoc struct {
	Doc
	Source Test `json:"_source"`
}

type TestSearchHits struct {
	SearchHits
	TestDocs []TestDoc `json:"hits"`
}

type TestSearchResponse struct {
	SearchResponse
	TestSearchHits `json:"hits"`
}

func TestElasticSearch(t *testing.T) {
	config := es7.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	client := NewClient(config)

	test := Test{
		Int:    1,
		Float:  1.2,
		String: "1.3",
		Slice:  []int{2, 3, 4},
	}
	res, err := client.Index("test", "2", test)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", res)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	searchResponse := &TestSearchResponse{}
	err = client.Search("test", query, searchResponse)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("%+v\n", searchResponse)
}
