package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/tidwall/sjson"
)

type doaRepository struct{}

func NewDoaRepository() (rep *doaRepository) {
	return &doaRepository{}
}

type AddNodeInput struct {
	Title    string `json:"title"`
	NodeID   string `json:"nodeId"`
	ParentID string `json:"parentId"`
	Label    string `json:"label"`
	Depth    int    `json:"depth,string"`
	Path     string `json:"path"`
}

type DistrictInsertInput struct {
	Code         string `json:"code"`
	Title        string `json:"title"`
	NodeID       string `json:"nodeId"`
	ParentCode   string `json:"parentCode"`
	Label        string `json:"label"`
	Depth        int    `json:"depth,string"`
	Path         string `json:"path"`
	FirstLetter  string `json:"firstLetter"`
	IsDeprecated string `json:"isDeprecated"`
}

var DoaHost = "http://doa.programmerfamily.com"

func (r *doaRepository) AddNode(data []byte) (err error) {
	input := AddNodeInput{}
	err = json.Unmarshal(data, &input)
	if err != nil {
		return err
	}

	entity := DistrictInsertInput{
		Code:         input.NodeID,
		Depth:        input.Depth,
		Label:        input.Label,
		ParentCode:   input.ParentID,
		Path:         input.Path,
		Title:        input.Title,
		IsDeprecated: "0",
	}

	url := fmt.Sprintf("%s/api/rent/v1/district/insert", DoaHost)
	err = Curl(url, entity, nil)
	if err != nil {
		return err
	}

	return nil
}
func (r *doaRepository) GetNode(nodeId string, output interface{}) (err error) {
	data := fmt.Sprintf(`{"code":"%s"}`, nodeId)
	url := fmt.Sprintf("%s/api/rent/v1/district/get_by_code", DoaHost)
	err = Curl(url, data, output)
	if err != nil {
		return err
	}
	return err
}
func (r *doaRepository) GetAllByPathPrefixWithDepth(parentPath string, depth int, output interface{}) (err error) {
	return err
}
func (r *doaRepository) CountByPathPrefix(path string, output interface{}) (err error) {
	return err
}
func (r *doaRepository) GetAllNodeByNodeIds(nodeIds []string, output interface{}) (err error) {
	data, err := sjson.Set("", "codeList", nodeIds)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/api/rent/v1/district/get_all_by_code", DoaHost)
	err = Curl(url, data, output)
	if err != nil {
		return err
	}
	return err
}

func Curl(url string, body interface{}, out interface{}) (err error) {
	client := resty.New()
	resp, err := client.R().EnableTrace().SetBody(body).Post(url)
	if err != nil {
		return err
	}
	httpCode := resp.StatusCode()
	if httpCode != http.StatusOK {
		err = errors.Errorf("%s httpCode:%d", url, httpCode)
		return err
	}
	if out != nil {
		err = json.Unmarshal(resp.Body(), out)
		if err != nil {
			return err
		}
	}
	return nil
}
