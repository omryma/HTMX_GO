package main

var data []Configuration

type Configuration struct {
	ID         string
	title      string
	appId      string
	expression string
}

func init() {
	data = []Configuration{
		{
			ID:         "1234abc",
			title:      "WeekendConf",
			appId:      "PX2000",
			expression: "UA: PhantomJS",
		},
	}
}
