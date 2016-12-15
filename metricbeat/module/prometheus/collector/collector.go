package collector

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"strconv"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/mb/parse"
)

const (
	defaultScheme = "http"
	defaultPath   = "/metrics"
)

var (
	debugf = logp.MakeDebug("prometheus-collector")

	hostParser = parse.URLHostParserBuilder{
		DefaultScheme: defaultScheme,
		DefaultPath:   defaultPath,
	}.Build()
)

func init() {
	if err := mb.Registry.AddMetricSet("prometheus", "collector", New, hostParser); err != nil {
		panic(err)
	}
}

type MetricSet struct {
	mb.BaseMetricSet
	client *http.Client
}

func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	logp.Warn("EXPERIMENTAL: The prometheus collector metricset is experimental")

	return &MetricSet{
		BaseMetricSet: base,
		client:        &http.Client{Timeout: base.Module().Config().Timeout},
	}, nil
}

func (m *MetricSet) Fetch() (common.MapStr, error) {

	req, err := http.NewRequest("GET", m.HostData().SanitizedURI, nil)
	if m.HostData().User != "" || m.HostData().Password != "" {
		req.SetBasicAuth(m.HostData().User, m.HostData().Password)
	}
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)

	event := common.MapStr{}

	// Iterate through all events to gather data
	for scanner.Scan() {
		line := scanner.Text()
		// Skip comment lines
		if line[0] == '#' {
			continue
		}

		e := decodeLine(line)
		event.Update(e)
	}

	return event, err
}

func decodeLine(line string) common.MapStr {
	// Separate key and value
	split := strings.Split(line, " ")
	key := split[0]

	value := split[1]

	// skip entries without a value
	if value == "NaN" {
		return nil
	}
	// Check if value is int or float

	// Split key
	startLabels := strings.Index(line, "{")
	endLabels := strings.Index(line, "}")

	keyValuePairs := map[string]string{}
	fullKey := strings.Replace(key, `"`, "", -1)

	// Handle labels
	if startLabels != -1 {
		// Overwrite key, as key contained labels until now too
		key = line[0:startLabels]

		// Extract labels
		keyValuePairs = extractLabels(line[startLabels+1 : endLabels])
	}

	values := common.MapStr{
		"value": convertValue(value),
		"key":   key,
	}

	for k, v := range keyValuePairs {
		values[k] = convertValue(v)
	}

	event := common.MapStr{
		fullKey: values,
	}
	return event
}

func extractLabels(labelsString string) map[string]string {

	keyValuePairs := map[string]string{}

	// Extract labels
	labels := strings.Split(labelsString, ",")
	for _, label := range labels {
		// TODO: convert integerts / float
		keyValue := strings.Split(label, "=")
		// Remove " from value
		keyValue[1] = keyValue[1][1 : len(keyValue[1])-1]
		keyValuePairs[keyValue[0]] = keyValue[1]
	}

	return keyValuePairs
}

func convertValue(value string) interface{} {

	if i, err := strconv.ParseInt(value, 10, 64); err == nil {
		return i
	}

	if f, err := strconv.ParseFloat(value, 64); err == nil {
		return f
	}

	return value
}
