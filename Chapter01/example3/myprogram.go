/*
	예상하지 못한 타입 처리하기

1. Go 에서 CSV 데이터가 [][]string으로 읽혀진다는 것을 확인.
2. Go는 정적으로 타입(유형)을 지정하기 때문에 CSV 필드에 대해 엄격하게 검사를 수행 할 수 있음.
3. 처리를 위해 각 필드의 구문을 분석해 값을 읽을 때 이 작업을 수행할 수 있음.
4. 한 열의 다른 값들과 일치하지 않는 임의의 필드를 갖는 지저분한 데이터를 생각해보자.
*/

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// CSV 레코드에서 필드의 타입을 확인하기 위해 성공적으로 읽어온 값을 저장하는 구조체 변수 생성.
// CSVRecord는 CSV 파일에서 성공적으로 읽어온 행을 저장한다.
type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

func main() {

	f, err := os.Open("/Users/david/dev/Golang/src/golang_ml/data/iris_mixed_types.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	// CSV에서 성공적으로 읽어온 레코드를 저장하는 값을 생성.
	var csvData []CSVRecord

	line := 1

	// 레코드를 읽고 예상하지 못한 타입을 찾는다.
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 열을 저장하기 위한 CSVRecord 생성
		var csvRecord CSVRecord

		// 기대하는 타입을 기반으로 레코드의 각 값을 읽는다.
		for idx, value := range record {

			// 문자열 행에 대해 문아열로 레코드의 값을 읽는다.
			if idx == 4 {

				// 값이 빈 문자열이 아닌지 확인. 해당 값이 빈 문자열인 경우
				// 구문 분석을 처리하는 루프를 중단한다.
				if value == "" {
					log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}

				// CSVRecord 에 문자열값을 추가
				csvRecord.Species = value
				continue
			}

			// 문자열 행이 아닌 경우 레코드의 값을 float64로 읽는다.
			var floatValue float64

			// 레코드의 값이 float로 읽혀지지 않으면 로그에 기록.
			// 구문 분석 처리 루프를 중단.
			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Unexpected type in column %d\n", idx)
				csvRecord.ParseError = fmt.Errorf("Could not parse float")
				break
			}

			// CSVRecord의 해당 필드에 float 값을 추가.
			switch idx {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.PetalLength = floatValue
			case 3:
				csvRecord.PetalWidth = floatValue
			}
		}

		// 앞에서 생성해둔 csvData에 성공적으로 읽어온 레코드 추가.
		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}

		line++
	}

	fmt.Printf("Successfully parsed %d lines\n", len(csvData))
}
