package core

import (
	"encoding/json"
	"fmt"
	cms20190101 "github.com/alibabacloud-go/cms-20190101/v8/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"strings"
)

func TrafficCount(client *cms20190101.Client, describeMetricListRequest *cms20190101.DescribeMetricListRequest) (err error) {

	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		res, err := client.DescribeMetricListWithOptions(describeMetricListRequest, runtime)
		fmt.Println(res)
		if err != nil {
			return err
		}

		return nil
	}()

	if tryErr != nil {
		sdkErr := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			sdkErr = _t
		} else {
			sdkErr.Message = tea.String(tryErr.Error())
		}

		// Diagnostic address
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(sdkErr.Data)))
		_ = d.Decode(&data)

		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, err = util.AssertAsString(sdkErr.Message)
		if err != nil {
			return err
		}
	}
	return
}
