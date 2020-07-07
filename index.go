package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

const (
	POST_INDEX = "post"
	ES_URL     = "http://10.128.0.2:9200"
)
