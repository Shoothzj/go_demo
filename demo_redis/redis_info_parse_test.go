package demo_redis

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestRedisInfoLines(t *testing.T) {
	content, err := ioutil.ReadFile("redis_info.txt")
	if err != nil {
		//Do something
	}
	lines := strings.Split(string(content), "\n")
	redisInfoMetricsLines(lines)
}

func redisInfoMetricsLines(lines []string) {
	fieldClass := ""

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			fmt.Println("line is empty")
			continue
		}
		if len(line) > 0 && strings.HasPrefix(line, "# ") {
			fieldClass = line[2:]
			continue
		}

		if (len(line) < 2) || (!strings.Contains(line, ":")) {
			continue
		}
		split := strings.SplitN(line, ":", 2)
		fieldKey := split[0]
		fieldValue := split[1]

		switch fieldClass {
		case "Clients":
			metricClientInfo(fieldKey, fieldValue)
		case "Cluster":
			metricClusterInfo(fieldKey, fieldValue)
		case "Replication":
			metricClientInfo(fieldKey, fieldValue)
		case "Stats":
			metricClientInfo(fieldKey, fieldValue)
		case "Memory":
			metricClientInfo(fieldKey, fieldValue)
		}
	}
}

func metricClientInfo(key, value string) {
	fmt.Printf("key is %s value is %f \n", key, ConvStr2Float(value))
}

func metricClusterInfo(key, value string) {
	switch key {
	case "cluster_enabled":
		fmt.Printf("cluster key is %s value is %f \n", key, ConvStr2Float(value))
	}
}

func ConvStr2Float(str string) float64 {
	if str == "" {
		return 0
	}
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return result
}
