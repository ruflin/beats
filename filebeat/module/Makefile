

run:
	$(MAKE) -C ..
	# Load template
	curl -XPUT "localhost:9200/_ingest/pipeline/nginx" -d @nginx/error/pipeline.json
	# Load pipeline
	../filebeat -c ../filebeat.yml -c nginx/acccess/config.yml -e
