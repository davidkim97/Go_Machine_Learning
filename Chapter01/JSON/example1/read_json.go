package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// citiBikeURL은 CitiBike 자전거 공유 정류장의 상황을 알려줌.
const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// stationData 는 citiBikeURL 로 부터 반환된 JSON 문서의 구문을 분석하는 데 사용됨.
type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

// station은 stationData 안의 각 station 문서의 구문을 분석하는 데 사용 됨.
type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_available_keys"`
}

func main() {

	// URL 로부터 JSON 응답 얻기.
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// 응답의 Body를 []byte로 읽는다.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// stationData 유형의 변수 선언.
	var sd stationData
	// stationData 변수로 JSON 데이터를 읽는다.
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
	}

	// 첫 번째 정류장 정보를 출력
	fmt.Printf("%+v\n\n", sd.Data.Stations[0])
}
