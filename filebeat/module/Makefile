

run:
	#$(MAKE) -C ..
	# Load template
	curl -XPUT "localhost:9200/_ingest/pipeline/nginx.access" -d @nginx/access/_meta/pipeline/default.json
	# Load pipeline
	../filebeat -c ../filebeat.yml -c nginx/access/config.yml -e


generate:
	python template.py --fields=./nginx/access/_meta/fields.yml --target=./nginx/access/template.json
