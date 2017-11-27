package main

import (
	"os"

	"time"

	filebeat "github.com/elastic/beats/filebeat/cmd"
	heartbeat "github.com/elastic/beats/heartbeat/cmd"
	metricbeat "github.com/elastic/beats/metricbeat/cmd"
)

func main() {
	go func() {
		if err := metricbeat.RootCmd.Execute(); err != nil {
			os.Exit(1)
		}
	}()
	// Timeout otherwise multiple beats try to write to the meta.json file at the same time
	time.Sleep(1 * time.Second)
	go func() {
		if err := filebeat.RootCmd.Execute(); err != nil {
			os.Exit(1)
		}
	}()
	time.Sleep(1 * time.Second)
	if err := heartbeat.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	// Not using auditbeat for now as there is probably an issue as both use same registry for modules?
	//time.Sleep(1 * time.Second)
	//if err := auditbeat.RootCmd.Execute(); err != nil {
	//	os.Exit(1)
	//}
} //
