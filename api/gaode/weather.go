package gaode

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"server/global"
	"server/service/common/response"
)

type IPResponse struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Adcode    string `json:"adcode"`
	Rectangle string `json:"rectangle"`
}

type WeatherResponse struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Lives    []struct {
		Province         string `json:"province"`
		City             string `json:"city"`
		Adcode           string `json:"adcode"`
		Weather          string `json:"weather"`
		Temperature      string `json:"temperature"`
		Winddirection    string `json:"winddirection"`
		Windpower        string `json:"windpower"`
		Humidity         string `json:"humidity"`
		Reporttime       string `json:"reporttime"`
		TemperatureFloat string `json:"temperature_float"`
		HumidityFloat    string `json:"humidity_float"`
	} `json:"lives"`
}

type WeatherInfoResponse struct {
	Province         string `json:"province"`
	City             string `json:"city"`
	Adcode           string `json:"adcode"`
	Weather          string `json:"weather"`
	Temperature      string `json:"temperature"`
	Winddirection    string `json:"winddirection"`
	Windpower        string `json:"windpower"`
	Humidity         string `json:"humidity"`
	Reporttime       string `json:"reporttime"`
	TemperatureFloat string `json:"temperature_float"`
	HumidityFloat    string `json:"humidity_float"`
}

// WeatherInfoView 获取实时天气
// @Tags 第三方 API 管理
// @Summary 获取实时天气
// @Description 获取实时天气
// @Param token header string  true  "token"
// @Router /api/gaode/weather [get]
// @Produce json
// @Success 200 {object} response.Response{data=WeatherInfoResponse}
func (GaodeApi) WeatherInfoView(c *gin.Context) {

	var data WeatherInfoResponse
	if !global.Config.Gaode.Enable {
		response.OkWithData(data, c)
		return
	}

	key := global.Config.Gaode.Key

	res, err := http.Get("https://restapi.amap.com/v3/ip?key=" + key)
	if err != nil {
		logrus.Errorf(err.Error())
		response.FailWithMessage("获取定位失败", c)
		return
	}

	byteData, _ := io.ReadAll(res.Body)

	var ipResponse IPResponse
	err = json.Unmarshal(byteData, &ipResponse)
	if err != nil {
		response.FailWithMessage("解析定位失败", c)
		return
	}
	if ipResponse.Status != "1" {
		response.FailWithMessage("请求定位失败", c)
		return
	}

	res, err = http.Get(fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s", key, ipResponse.Adcode))
	if err != nil {
		global.Log.Errorf(err.Error())
		response.FailWithMessage("获取天气信息失败", c)
		return
	}
	byteData, _ = io.ReadAll(res.Body)

	var weatherResponse WeatherResponse
	err = json.Unmarshal(byteData, &weatherResponse)
	if err != nil {
		response.FailWithMessage("解析天气信息失败", c)
		return
	}

	if weatherResponse.Status != "1" {
		global.Log.Errorf(string(byteData))
		response.FailWithMessage("请求天气信息失败", c)
		return
	}

	if len(weatherResponse.Lives) > 0 {
		response.OkWithData(weatherResponse.Lives[0], c)
		return
	}
	response.FailWithMessage("获取天气信息异常", c)
	return
}
