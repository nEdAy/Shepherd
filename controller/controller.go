package controller

import (
	"Shepherd/pkg/config"
	"Shepherd/pkg/redis"
	"Shepherd/pkg/response"
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/resty.v1"
	"net/http"
	"sort"
)

type Dataoke struct {
	cacheKey  string
	cacheTime int
	url       string
	params    map[string]string
}

func getFromDataoke(c *gin.Context, dataoke Dataoke) {
	data, err := getDataByCacheOrSource(dataoke)
	if err != nil {
		response.ErrorWithMsg(c, err.Error())
	} else {
		c.String(http.StatusOK, data)
	}
}

func getDataByCacheOrSource(dataoke Dataoke) (string, error) {
	data, err := redis.Get(dataoke.cacheKey)
	if err != nil {
		data, err := getDataBySource(dataoke.url, dataoke.params)
		if err != nil {
			return "", err
		} else {
			defer redis.Set(dataoke.cacheKey, data, dataoke.cacheTime)
			return data, nil
		}
	} else {
		return string(data), nil
	}
}

func getDataBySource(url string, params map[string]string) (string, error) {
	rankingData, err := resty.R().SetQueryParams(signParams(params)).Get(url)
	if err != nil {
		return "", err
	} else if rankingData.IsSuccess() {
		return rankingData.String(), nil
	} else {
		return "", errors.New("大淘客API异常")
	}
}

func signParams(params map[string]string) map[string]string {
	params["appKey"] = config.Dataoke.AppKey
	var buffer bytes.Buffer
	sortedMap(params, func(key string, value interface{}) {
		buffer.WriteString(key)
		buffer.WriteString("=")
		buffer.WriteString(value.(string))
		buffer.WriteString("&")
	})
	buffer.WriteString("key=")
	buffer.WriteString(config.Dataoke.AppSecret)
	sign := fmt.Sprintf("%x", md5.Sum(buffer.Bytes()))
	params["sign"] = sign
	return params
}

func sortedMap(m map[string]string, f func(k string, v interface{})) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		f(k, m[k])
	}
}
