package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"gitea.programmerfamily.com/go/pathtree"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	templatemaputil "github.com/suifengpiao14/templatemap/util"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type doaRepository struct {
	pathtree.EmptyTreeRpository
}

func NewDoaRepository() (rep *doaRepository) {
	return &doaRepository{}
}

type DistrictInsertInput struct {
	Code         string `json:"code"`
	Title        string `json:"title"`
	ParentCode   string `json:"parentCode"`
	Label        string `json:"label"`
	Depth        int    `json:"depth,string"`
	Path         string `json:"path"`
	FirstLetter  string `json:"firstLetter"`
	IsDeprecated string `json:"isDeprecated"`
}

var DoaHost = "http://doa.programmerfamily.com"

func (r *doaRepository) AddNode(node pathtree.TreeNodeI) (err error) {
	input := DistrictInsertInput{
		IsDeprecated: "0",
	}
	b, err := json.Marshal(node)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &input)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/api/rent/v1/district/insert", DoaHost)
	_, err = Curl(url, input)
	if err != nil {
		return err
	}
	return nil
}
func (r *doaRepository) GetNode(nodeId string, output interface{}) (err error) {
	args := fmt.Sprintf(`{"code":"%s"}`, nodeId)
	url := fmt.Sprintf("%s/api/rent/v1/district/get_by_code", DoaHost)
	resStr, err := Curl(url, args)
	if err != nil {
		return err
	}
	data, err := DataFill(resStr)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), output)
	if err != nil {
		return err
	}
	return nil
}
func (r *doaRepository) GetAllByPathPrefix(parentPath string, depth int, output interface{}) (err error) {
	args := fmt.Sprintf(`{"pathPrefix":"%s%%","depth":"%d"}`, parentPath, depth)
	url := fmt.Sprintf("%s/api/rent/v1/district/get_by_path_prefix_limit_depth", DoaHost)
	resStr, err := Curl(url, args)
	if err != nil {
		return err
	}
	data, err := DataFill(resStr)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), output)
	if err != nil {
		return err
	}
	return nil
}

func (r *doaRepository) BatchUpdatePath(input []byte) (err error) {
	url := fmt.Sprintf("%s/api/rent/v1/district/batchUpdatePath", DoaHost)
	_, err = Curl(url, input)
	if err != nil {
		return err
	}
	return nil
}

func (r *doaRepository) CountByPathPrefix(path string, depth int, output interface{}) (err error) {
	data := fmt.Sprintf(`{"pathPrefix":"%s%%","depth":"%d"}`, path, depth)
	url := fmt.Sprintf("%s/api/rent/v1/district/count_by_path_prefix", DoaHost)
	resStr, err := Curl(url, data)
	if err != nil {
		return err
	}
	count, err := strconv.ParseInt(resStr, 10, 64)
	if err != nil {
		return err
	}
	rv := reflect.Indirect(reflect.ValueOf(output))
	if !rv.CanSet() {
		err := errors.Errorf("output reflect.value CanSet is false")
		return err
	}
	rv.SetInt(count)
	return nil
}
func (r *doaRepository) GetAllNodeByNodeIds(nodeIds []string, output interface{}) (err error) {
	args, err := sjson.Set("", "codeList", nodeIds)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/api/rent/v1/district/get_all_by_code", DoaHost)
	resStr, err := Curl(url, args)
	data, err := DataFill(resStr)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), output)
	if err != nil {
		return err
	}
	return nil
}

func Curl(url string, body interface{}) (out string, err error) {
	client := resty.New()
	resp, err := client.R().EnableTrace().SetBody(body).Post(url)
	if err != nil {
		return "", err
	}
	httpCode := resp.StatusCode()
	if httpCode != http.StatusOK {
		err = errors.Errorf("%s httpCode:%d", url, httpCode)
		return "", err
	}
	out = string(resp.Body())
	return out, nil
}

var JsonKeyMap = map[string]string{
	"code":       "nodeId",
	"parentCode": "parentId",
	"label":      "label",
	"depth":      "depth",
	"path":       "path",
}

// DataFill 数据填充，当json数据中的key 和目标结构体json key 不一致时，需要转换key，此处采用复制方式达到目的，方便兼容多结构体json key
func DataFill(jsonStr string) (out string, err error) {
	out = strings.TrimSpace(jsonStr)
	isArr := false
	if out == "" {
		return "", nil
	}
	if out[0] == '[' {
		isArr = true
		out = templatemaputil.Row2Column(out)
	}
	for srcKey, dstKey := range JsonKeyMap {
		if srcKey == dstKey {
			continue
		}
		result := gjson.Get(out, srcKey)
		if !result.Exists() {
			err = errors.Errorf("db key:%s not exists", srcKey)
			return "", err
		}
		out, err = sjson.SetRaw(out, dstKey, result.Raw)
		if err != nil {
			return "", err
		}
	}
	if isArr {
		out = templatemaputil.Column2Row(out)
	}
	return out, nil
}
