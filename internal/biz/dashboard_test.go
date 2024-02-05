package biz

import (
	"fmt"
	pb "github.com/mohaijiang/computeshare-server/api/dashboard/v1"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestDashboard_Prometheus(t *testing.T) {
	basicUrl := "http://61.172.179.73:9090"
	path := "/api/v1/query"
	StorageSpaceTotal := "sum(SeaweedFS_volumeServer_total_disk_size) by (exported_instance)"
	params := url.Values{}
	parseURL, err := url.Parse(basicUrl + path)
	if err != nil {
		fmt.Errorf("url.Parse err is %s", err)
		panic(t)
	}
	params.Set("query", StorageSpaceTotal)
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	resp, err := http.Get(urlPathWithParams)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	result := gjson.GetBytes(bytes, "data.result.0.value.1")
	i := result.Int()
	fmt.Println(ByteCountIEC(i))
}

func TestDashboard_Prometheus_List(t *testing.T) {
	basicUrl := "http://61.172.179.73:9090"
	path := "/api/v1/query"
	ProvidersCorrespondingVolumesNum = "sum by (instance)(SeaweedFS_volumeServer_volumes{type=\"volume\"})"
	params := url.Values{}
	parseURL, err := url.Parse(basicUrl + path)
	if err != nil {
		fmt.Errorf("url.Parse err is %s", err)
		panic(t)
	}
	params.Set("query", ProvidersCorrespondingVolumesNum)
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	resp, err := http.Get(urlPathWithParams)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	getManyBytes := gjson.GetBytes(bytes, "data.result").Array()
	for _, manyByte := range getManyBytes {
		var storagesProviders pb.StoragesProvidersListReply_StoragesProviders
		storagesProviders.Instance = manyByte.Get("metric.instance").String()
		storagesProviders.VolumeNum = int32(manyByte.Get("value.1").Int())
		fmt.Println(storagesProviders)
	}
}
