package micro

import (
	"context"
	"fmt"
	data "testGin/datasharems/share"
	"testGin/model"
)

// 处理sql请求 sqlStr
func GetResults(req *model.SearchSql) (*model.SearchResponse, error) {
	getCryptoSqlReq := new(data.GetCrySqlRequest)
	getCryptoSqlReq.ConsumerId = req.ConsumerID
	getCryptoSqlReq.ProviderId = req.ProviderID
	getCryptoSqlReq.DumpId = req.DumpID
	getCryptoSqlReq.Sql = req.Sql
	getCryptoSqlReq.Text = req.Text

	cli := GetDataShareClient()
	fmt.Println("api request is ", getCryptoSqlReq)
	response, err := cli.GetCrySql(
		context.TODO(),
		getCryptoSqlReq,
	)
	if err != nil {
		return nil, err
	}

	ret := new(model.SearchResponse)
	ret.Message = response.Message
	ret.FlowID = response.FlowId
	ret.Status = response.Status
	ret.TxID = response.TxId
	return ret, nil
}
