package json

import (
	"net"

	"encoding/json"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	if err := mb.Registry.AddMetricSet("udp", "json", New); err != nil {
		panic(err)
	}
}

// MetricSet type defines all fields of the MetricSet
// As a minimum it must inherit the mb.BaseMetricSet fields, but can be extended with
// additional entries. These variables can be used to persist data or configuration between
// multiple fetch calls.
type MetricSet struct {
	mb.BaseMetricSet
}

// New create a new instance of the MetricSet
// Part of new is also setting up the configuration by processing additional
// configuration entries if needed.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {

	config := struct{}{}

	logp.Warn("EXPERIMENTAL: The udp json metricset is experimental")

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
	}, nil
}

func (m *MetricSet) Run(reporter mb.PushReporter) {
	listener, err := net.ListenPacket("udp", ":8080")

	if err != nil {
		logp.Err("error happened: %s", err)
		return
	}
	defer listener.Close()

	go func() {
		select {
		case <-reporter.Done():
			// TODO: Good way to do it?
			listener.Close()
		}
	}()
	buffer := make([]byte, 1000)

	for {
		length, _, err := listener.ReadFrom(buffer)
		if err != nil {
			logp.Err("Error reading from buffer: %v", err.Error())
			// TODO: what happens on error?
			return
		}
		data := buffer[:length]
		var entry interface{}
		err = json.Unmarshal(data, &entry)
		if err != nil {
			logp.Err("Error decoding json: %v", err)
		}
		event := common.MapStr{
			"data": entry,
		}
		// TODO: introduce namespace
		reporter.Event(event)
	}
}
