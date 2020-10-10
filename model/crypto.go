package model

import (
	"errors"
	"reflect"

	"github.com/sirupsen/logrus"
)

// 查询sql
type SearchSql struct {
	ConsumerID string `json:"consumerId" example:"202009021838" format:"string"`
	ProviderID string `json:"providerId" example:"202009021840" format:"string"`
	Sql        string `json:"sql" example:"select * from t_aos;" format:"string"`
	DumpID     string `json:"dumpId" example:""`
	Text       string `json:"text" example:"美国国家科学院" format:"string"`
}

var (
	ErrInvalid = errors.New("request params is empty")
)

// 校验查询sql传参
func (s SearchSql) Validation() error {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			logrus.Error(ErrInvalid)
			return ErrInvalid
		}
	}
	return nil
}

type SearchResponse struct {
	FlowID  string `json:"flowId" format:"string"`
	TxID    string `json:"txId" format:"string"`
	Status  int32  `json:"code" example:"200" format:"int32"`
	Message string `json:"message" example:"Search Success" format:"string"`
}
