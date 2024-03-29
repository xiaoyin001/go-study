package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
)

// 客户端的Http请求(如有违规或者侵权，请联系我及时删除)

func TestRequestHttp(t *testing.T) {
	// 接口请求URL
	apiUrl := "http://apis.juhe.cn/simpleWeather/query"

	// 初始化参数
	param := url.Values{}

	// 接口请求参数
	param.Set("city", "上海")     // 要查询的城市名称/id，城市名称如：温州、上海、北京
	param.Set("key", "聚合接口Key") // 接口请求Key

	// 发送请求
	data, err := Get(apiUrl, param)
	if err != nil {
		// 请求异常，根据自身业务逻辑进行调整修改
		fmt.Errorf("请求异常:\r\n%v", err)
	} else {
		var netReturn map[string]interface{}
		jsonerr := json.Unmarshal(data, &netReturn)
		if jsonerr != nil {
			// 解析JSON异常，根据自身业务逻辑进行调整修改
			fmt.Errorf("请求异常:%v", jsonerr)
		} else {
			errorCode := netReturn["error_code"]
			reason := netReturn["reason"]
			data := netReturn["result"]
			// 当前天气信息
			realtime := data.(map[string]interface{})["realtime"]

			if errorCode.(float64) == 0 {
				// 请求成功，根据自身业务逻辑进行调整修改
				fmt.Printf("温度：%v\n湿度：%v\n天气：%v\n风向：%v\n风力：%v\n空气质量：%v",
					realtime.(map[string]interface{})["temperature"],
					realtime.(map[string]interface{})["humidity"],
					realtime.(map[string]interface{})["info"],
					realtime.(map[string]interface{})["direct"],
					realtime.(map[string]interface{})["power"],
					realtime.(map[string]interface{})["aqi"],
				)
			} else {
				// 查询失败，根据自身业务逻辑进行调整修改
				fmt.Printf("请求失败:%v_%v", errorCode.(float64), reason)
			}
		}
	}
}

// get 方式发起网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
