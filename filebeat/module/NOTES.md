* User agent plugin should be shipped by default? https://www.elastic.co/guide/en/elasticsearch/plugins/master/ingest-user-agent.html
* Naming suggestion: Module = nginx, dataset = access
    * Note on the data set: Independent which format is used (json or not), the end structure should always be the same (potentially some missing fields)
* Should multiple processors be loaded for each dataset
* Geoip processor shipped by default?
* Message line should only be removed when successful -> error pipeline needed -> error should not block ingest
* Should we "rename" metricset in metricbeat to dataset to consistency?
* Each data set has an additional type to specify if it was json, default, ...
* Should we use pipeline versioning?
* Should JSON processing happen on filebeat or elasticsearch side? -> filebeat side has filtering advantage
* Where should we add the dataset fields? Config file, ingest pipeline, code?
