package config

const (
	//Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// ElasticSearch
	ElasticIndex = "hongniang_profile"

	// RpcEndpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	//rate limiting
	Qps = 20
)
