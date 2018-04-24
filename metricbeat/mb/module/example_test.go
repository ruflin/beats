// +build !integration

package module_test

import (
	stdjson "encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/outputs/codec/json"

	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/mb/module"
)

// ExampleWrapper demonstrates how to create a single Wrapper
// from configuration, start the module, and consume events generated by the
// module.
func ExampleWrapper() {
	// Build a configuration object.
	config, err := common.NewConfigFrom(map[string]interface{}{
		"module":     moduleName,
		"metricsets": []string{eventFetcherName},
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a new Wrapper based on the configuration.
	m, err := module.NewWrapper(config, mb.Registry, module.WithMetricSetInfo())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Run the module until done is closed.
	done := make(chan struct{})
	output := m.Start(done)

	// Process events from the output channel until it is closed.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range output {
			event.Fields.Put("metricset.rtt", 111)

			output, err := encodeEvent(event)
			if err == nil {
				fmt.Println(output)
			}
		}
	}()

	// Simulate running for a while.
	time.Sleep(50 * time.Millisecond)

	// When finished with the module, close the done channel. When the Module
	// stops it will automatically close its output channel so that the output
	// for loop stops.
	close(done)
	wg.Wait()

	// Output:
	// {
	//   "@metadata": {
	//     "beat": "noindex",
	//     "type": "doc",
	//     "version": "1.2.3"
	//   },
	//   "@timestamp": "2016-05-10T23:27:58.485Z",
	//   "fake": {
	//     "eventfetcher": {
	//       "metric": 1
	//     }
	//   },
	//   "metricset": {
	//     "module": "fake",
	//     "name": "eventfetcher",
	//     "period": 10000000,
	//     "rtt": 111
	//   }
	// }
}

// ExampleRunner demonstrates how to use Runner to start and stop
// a module.
func ExampleRunner() {
	// A *beat.Beat is injected into a Beater when it runs and contains the
	// Publisher used to publish events. This Beat pointer is created here only
	// for demonstration purposes.
	var b *beat.Beat

	config, err := common.NewConfigFrom(map[string]interface{}{
		"module":     moduleName,
		"metricsets": []string{eventFetcherName},
	})
	if err != nil {
		return
	}

	// Create a new Wrapper based on the configuration.
	m, err := module.NewWrapper(config, mb.Registry, module.WithMetricSetInfo())
	if err != nil {
		return
	}

	connector, err := module.NewConnector(b.Publisher, config, nil)
	if err != nil {
		return
	}

	client, err := connector.Connect()
	if err != nil {
		return
	}

	// Create the Runner facade.
	runner := module.NewRunner(client, m)

	// Start the module and have it publish to a new publisher.Client.
	runner.Start()

	// Stop the module. This blocks until all MetricSets in the Module have
	// stopped and the publisher.Client is closed.
	runner.Stop()
}

func encodeEvent(event beat.Event) (string, error) {
	output, err := json.New(false, "1.2.3").Encode("noindex", &event)
	if err != nil {
		return "", nil
	}

	// FIX: need to parse and re-encode, so fields ordering in final json document
	//      keeps stable.

	var tmp interface{}
	if err := stdjson.Unmarshal(output, &tmp); err != nil {
		panic(err)
	}

	output, err = stdjson.MarshalIndent(tmp, "", "  ")
	return string(output), err
}
