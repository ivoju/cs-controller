package main

import (
    "gopkg.in/resty.v1"
)

// Variable Declaration

var (
    KAFKA_PUB_URL_PRESCREENING_DATA = "http://172.18.133.254:8020/controller/publish/cs-prescreening-data"
    //KAFKA_PUB_URL_BRIGUNA = "http://10.148.0.6:8020/publish/cs-briguna-scored"
)

// Functions Declaration

func InitKafkaConsumer() {
    // kupedes_nongbt_rating
    resty.R().SetBody("{\"topic\":\"cs-kupedes_nongbt_rating-unscored\",\"group\":\"test-group\"}").Post("http://10.148.0.6:8020/subscribe/topic/add")
    resty.R().SetBody("{\"data\":[{\"topic\":\"cs-kupedes_nongbt_rating-unscored\",\"url\":[\"http://0.0.0.0:8000/las/kupedes_nongbt_rating\"]}]}").Post("http://10.148.0.6:8020/subscribe/url/add")
    resty.R().SetBody("{\"topic\":\"cs-kupedes_nongbt_rating-scored\",\"group\":\"test-group\"}").Post("http://10.148.0.6:8020/subscribe/topic/add")
    // briguna
    resty.R().SetBody("{\"topic\":\"cs-briguna-unscored\",\"group\":\"test-group\"}").Post("http://10.148.0.6:8020/subscribe/topic/add")
    resty.R().SetBody("{\"data\":[{\"topic\":\"cs-briguna-unscored\",\"url\":[\"http://0.0.0.0:8000/las/briguna\"]}]}").Post("http://10.148.0.6:8020/subscribe/url/add")
    resty.R().SetBody("{\"topic\":\"cs-briguna-scored\",\"group\":\"test-group\"}").Post("http://10.148.0.6:8020/subscribe/topic/add")
}

func ProduceKafka(data interface{}, url string) {
    resty.R().SetBody(data).Post(url)
}

