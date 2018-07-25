package main

import (
	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"strconv"
)

func asValue(value string) interface{} {
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}

	if f, err := strconv.Atoi(value); err == nil {
		return f
	}

	if i, err := strconv.Atoi(value); err == nil {
		return i
	}

	return value
}

func populateMetrics(sample *metric.Set, metrics map[string]interface{}, metricsDefinition map[string][]interface{}) {

	for metricName, metricConf := range metricsDefinition {
		rawSource := metricConf[0]
		metricType := metricConf[1].(metric.SourceType)

		var rawMetric interface{}
		var ok bool

		switch source := rawSource.(type) {
		case string:
			rawMetric, ok = metrics[source]

		case func(map[string]interface{}) (float64, bool):
			rawMetric, ok = source(metrics)

		case func(map[string]interface{}) (int, bool):
			rawMetric, ok = source(metrics)

		default:
			//fmt.Println("Invalid Raw Source")
			continue
		}

		if !ok {
			//log.Warn("ERROR")
		}

		err := sample.SetMetric(metricName, rawMetric, metricType)

		if err != nil {
			//fmt.Println("ERROR")
		}
	}
}
