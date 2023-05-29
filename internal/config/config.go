package config

// ElasticConfig represents values that is used to configure a ElasticSearch connection.
type ElasticConfig struct {
	Hosts    string `env-required:"false" env:"ElASTIC_SEARCH_HOSTS" json:"elsatic_search_hosts"`
	Region   string `env_required:"false" env:"ELASTIC_SEARCH_REGION" json:"elastic_search_region"`
	Username string `env-required:"false" env:"ELASTIC_SEARCH_USERNAME" json:"elastic_search_username"`
	Password string `env-required:"false" env:"ELASTIC_SEARCH_PASSWORD" json:"elastic_search_password"`
	Timeout  int    `env-required:"false" env:"ELASTIC_SEARCH_TIMEOUT" json:"elastic_search_timeout"`
}
