package main

const METORO_API_URL_ENV_VAR = "METORO_API_URL"
const METORO_AUTH_TOKEN_ENV_VAR = "METORO_AUTH_TOKEN"

type GetLogsRequest struct {
	// Required: Start time of when to get the logs in seconds since epoch
	StartTime int64 `json:"startTime"`
	// Required: End time of when to get the logs in seconds since epoch
	EndTime int64 `json:"endTime"`
	// The filters to apply to the logs, so for example, if you want to get logs for a specific service
	// you can pass in a filter like {"service_name": ["microservice_a"]}
	Filters map[string][]string `json:"filters"`
	// ExcludeFilters are filters that should be excluded from the logs
	// For example, if you want to get logs for all services except microservice_a you can pass in
	// {"service_name": ["microservice_a"]}
	ExcludeFilters map[string][]string `json:"excludeFilters"`
	// Previous page endTime in nanoseconds, used to get the next page of logs if there are more logs than the page size
	// If omitted, the first page of logs will be returned
	PrevEndTime *int64 `json:"prevEndTime"`
	// Regexes are used to filter logs based on a regex inclusively
	Regexes []string `json:"regexes"`
	// ExcludeRegexes are used to filter logs based on a regex exclusively
	ExcludeRegexes []string `json:"excludeRegexes"`
	Ascending      bool     `json:"ascending"`
	// The cluster/environments to get the logs for. If empty, all clusters will be included
	Environments []string `json:"environments"`
}

type Log struct {
	// The time that the log line was emitted in milliseconds since the epoch
	Time int64 `json:"time"`
	// The severity of the log line
	Severity string `json:"severity"`
	// The log message
	Message string `json:"message"`
	// The attributes of the log line
	LogAttributes map[string]string `json:"logAttributes"`
	// The attributes of the resource that emitted the log line
	ResourceAttributes map[string]string `json:"resourceAttributes"`
	// Service name
	ServiceName string `json:"serviceName"`
	// Environment
	Environment string `json:"environment"`
}

type GetLogsResponse struct {
	// The logs that match the filters
	Logs []Log `json:"logs"`
}

type GetTracesRequest struct {
	ServiceNames   []string            `json:"serviceNames"`
	StartTime      int64               `json:"startTime"`
	EndTime        int64               `json:"endTime"`
	Filters        map[string][]string `json:"filters"`
	ExcludeFilters map[string][]string `json:"excludeFilters"`
	PrevEndTime    *int64              `json:"prevEndTime"`
	Regexes        []string            `json:"regexes"`
	ExcludeRegexes []string            `json:"excludeRegexes"`
	Ascending      bool                `json:"ascending"`
	Environments   []string            `json:"environments"`
}

type GetMetricRequest struct {
	// MetricName is the name of the metric to get
	MetricName string `json:"metricName"`
	// Required: Start time of when to get the logs in seconds since epoch
	StartTime int64 `json:"startTime"`
	// Required: End time of when to get the logs in seconds since epoch
	EndTime int64 `json:"endTime"`
	// The filters to apply to the logs, so for example, if you want to get logs for a specific service
	// you can pass in a filter like {"service_name": ["microservice_a"]}
	Filters map[string][]string `json:"filters"`
	// The filters to exclude from the logs, so for example, if you want to exclude logs for a specific service
	// you can pass in a filter like {"service_name": ["microservice_a"]}
	ExcludeFilters map[string][]string `json:"excludeFilters"`
	// Splits is a list of attributes to split the metrics by, for example, if you want to split the metrics by service
	// you can pass in a list like ["service_name"]
	Splits []string `json:"splits"`
	// Aggregation is the operation to apply to the metrics, for example, if you want to sum the metrics you can pass in "sum"
	Aggregation Aggregation `json:"aggregation"`
	// IsRate is a flag to indicate if the metric is a rate metric
	IsRate bool `json:"isRate"`
	// Functions is the list of functions to apply to the metric, in the same order that they appear in this array!!
	Functions []MetricFunction `json:"functions"`
	// LimitResults is a flag to indicate if the results should be limited.
	LimitResults bool `json:"limitResults"`
	// BucketSize is the size of each datapoint bucket in seconds
	BucketSize int64 `json:"bucketSize"`
}

type GetProfileRequest struct {
	// Required: ServiceName to get profiling for
	ServiceName string `json:"serviceName"`

	// Optional: ContainerNames to get profiling for
	ContainerNames []string `json:"containerNames"`

	// Required: Timestamp to get profiling after this time
	// Seconds since epoch
	StartTime int64 `json:"startTime"`

	// Required: Timestamp to get profiling this time
	// Seconds since epoch
	EndTime int64 `json:"endTime"`
}
