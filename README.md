# gglmm-elasticsearch
## 依赖
+ github.com/elastic/go-elasticsearch/v7
## 使用方法
```golang
type Client struct {
	esClient *es7.Client
}
func NewClient(config es7.Config) *Client
func (client *Client) Index(index string, id string, doc interface{}) (*IndexResponse, error)
func (client *Client) Create(index string, id string, doc interface{}) (*IndexResponse, error)
func (client *Client) Update(index string, id string, doc interface{}) (*IndexResponse, error)
func (client *Client) Get(index string, id string, response interface{}) error
func (client *Client) Delete(index string, id string) (*IndexResponse, error)
func (client *Client) Search(index string, query interface{}, response interface{}) error
```