package main

import (
	"encoding/json"
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

	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sd stationData
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
		return
	}

	// 데이터를 마샬링한다(다시 JSON으로 만든다).
	outputData, err := json.Marshal(sd)
	if err != nil {
		log.Fatal(err)
	}

	// JSON 형식으로 생성된 데이터를 파일에 저장한다.
	if err := ioutil.WriteFile("citibike.json", outputData, 0644); err != nil {
		log.Fatal(err)
	}
}
