package partition

import (
	"github.com/Shopify/sarama"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"
)

// init registers the partition MetricSet with the central registry.
func init() {
	if err := mb.Registry.AddMetricSet("kafka", "partition", New); err != nil {
		panic(err)
	}
}

// MetricSet type defines all fields of the partition MetricSet
type MetricSet struct {
	mb.BaseMetricSet
	broker *sarama.Broker
}

// New create a new instance of the partition MetricSet
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {

	config := struct{}{}

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	cfg := sarama.NewConfig()
	cfg.Net.DialTimeout = base.Module().Config().Timeout
	cfg.Net.ReadTimeout = base.Module().Config().Timeout
	cfg.ClientID = "metricbeat"

	broker := sarama.NewBroker(base.Host())
	err := broker.Open(cfg)
	if err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		broker:        broker,
	}, nil
}

// Fetch partition stats list from kafka
func (m *MetricSet) Fetch() ([]common.MapStr, error) {

	response, err := m.broker.GetMetadata(&sarama.MetadataRequest{})
	if err != nil {
		return nil, err
	}

	events := []common.MapStr{}
	for _, topic := range response.Topics {

		for _, partition := range topic.Partitions {

			offsetRequest := &sarama.OffsetRequest{}
			// Get 2 offsets, assumption is newest first in array, oldest second
			offsetRequest.AddBlock(topic.Name, partition.ID, sarama.OffsetNewest, 2)

			offsetResponse, _ := m.broker.GetAvailableOffsets(offsetRequest)
			block := offsetResponse.GetBlock(topic.Name, partition.ID)

			if len(block.Offsets) == 0 {
				continue
			}
			event := common.MapStr{
				"topic": common.MapStr{
					"name":  topic.Name,
					"error": topic.Err,
				},
				"partition": common.MapStr{
					"id":       partition.ID,
					"error":    partition.Err,
					"leader":   partition.Leader,
					"replicas": partition.Replicas,
					"isr":      partition.Isr,
				},
				"offset": common.MapStr{
					"newest": block.Offsets[0],
					"oldest": block.Offsets[1],
					"error":  block.Err,
				},
				"broker": common.MapStr{
					"id":      m.broker.ID(),
					"address": m.broker.Addr(),
				},
			}

			events = append(events, event)
		}
	}

	return events, nil
}
