package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tj/assert"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func (r *Result) ToComputeInstanceRds() []ComputeInstanceRds {

	var result []ComputeInstanceRds
	for _, v := range r.Values {
		if len(v) == 2 {
			timestamp := v[0].(float64)
			statsTime := time.UnixMilli(int64(timestamp * 1000))
			value := v[1].(string)
			cpuUsage, err := strconv.ParseFloat(value, 64)
			if err != nil {
				cpuUsage = 0
			}
			result = append(result, ComputeInstanceRds{
				ID:        uuid.New().String(),
				CpuUsage:  float32(cpuUsage),
				StatsTime: statsTime,
			})
		}
	}

	return result
}

func TestPrometheus(t *testing.T) {
	// Prometheus 查询 API 地址
	prometheusURL := "http://61.172.179.73:9090/api/v1/query_range"

	// PromQL 查询语句
	// cpu (sum by(instance) (irate(node_cpu_seconds_total{instance="76fe0a88-1960-4966-9beb-41b1c1251595",job="node", mode!="idle"}[10m15s])) / on(instance) group_left sum by (instance)((irate(node_cpu_seconds_total{instance="76fe0a88-1960-4966-9beb-41b1c1251595",job="node"}[10m15s])))) * 100
	// memory  100 - ((avg_over_time(node_memory_MemAvailable_bytes{instance="76fe0a88-1960-4966-9beb-41b1c1251595",job="node"}[10m15s]) * 100) / avg_over_time(node_memory_MemTotal_bytes{instance="76fe0a88-1960-4966-9beb-41b1c1251595",job="node"}[10m15s]))
	query := "(sum by(instance) (irate(node_cpu_seconds_total{instance=\"76fe0a88-1960-4966-9beb-41b1c1251595\",job=\"node\", mode!=\"idle\"}[10m15s])) / on(instance) group_left sum by (instance)((irate(node_cpu_seconds_total{instance=\"76fe0a88-1960-4966-9beb-41b1c1251595\",job=\"node\"}[10m15s])))) * 100"

	// 构建查询参数
	params := fmt.Sprintf("query=%s", url.QueryEscape(query))
	queryURL := fmt.Sprintf("%s?%s&start=1704437472.534&end=1704441072.534&step=60", prometheusURL, params)

	queryURL = fmt.Sprintf("%s?%s", prometheusURL, "query=%28sum%28increase%28node_cpu_seconds_total%7Bmode%3D%27system%27%2Cinstance%3D%2276fe0a88-1960-4966-9beb-41b1c1251595%22%7D%5B10m%5D%29%29by%28instance%29%29+%2F+%28sum%28increase%28node_cpu_seconds_total%5B10m%5D%29%29by%28instance%29%29++*100%0A&start=1704437735.053&end=1704441335.053&step=14")
	// 发起 GET 请求
	response, err := http.Get(queryURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var queryResult PrometheusQueryResult
	if err := json.Unmarshal(body, &queryResult); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// 打印查询结果
	fmt.Println("Prometheus Query Result:")
	fmt.Println(queryResult.Data.Result[0].ToComputeInstanceRds())
}

func TestTime(t *testing.T) {
	s := 1704439933.053
	d := time.Now()
	s1 := time.UnixMilli(int64(s * 1000))
	fmt.Println(s1)
	fmt.Println(d)
	fmt.Println("1704439933.053")
	fmt.Println(time.Now().Unix())
}

func TestComputeInstanceUsercase_GetLast24HInstanceStats(t *testing.T) {
	client := NewComputeInstanceUsercase(
		nil, nil, nil, nil, nil, nil, nil, nil, nil)
	data, err := client.GetLast24HInstanceStats(context.Background(), "76fe0a88-1960-4966-9beb-41b1c1251595")
	assert.NoError(t, err)
	fmt.Println(data)
}
