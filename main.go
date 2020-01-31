package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

func main() {
	err := loadConfig()
	// fmt.Println(config)

	stationID := findStationIDFrom("07-479") // H스퀘어
	resp, err := http.Get(urlBusArrivalServiceStation + fmt.Sprintf("?serviceKey=%s&stationId=%s", getServiceKey(), stationID))
	if err != nil {
		panic(err)
	}
	// defer resp.Body.Close()
	var sr BusArrivalStationResponse
	xmlDec := xml.NewDecoder(resp.Body)
	xmlDec.Decode(&sr)
	for _, item := range sr.MsgBody.BusArrivalList {
		fmt.Println("routeId", item.RouteID) // TODO: routeId to bus No. from routeYYYYMMDD.txt
		fmt.Println("PredictTime1", item.PredictTime1)
		fmt.Println("PredictTime2", item.PredictTime2)
		fmt.Println("----")
	}

	resp.Body.Close()

}
