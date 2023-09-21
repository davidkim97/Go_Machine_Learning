package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("/Users/david/dev/Golang/src/golang_ml/data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	// rawCSVData는 성공적으로 읽어온 행의 데이터를 저장.
	var rawCSVData [][]string

	// 레코드를 하나씩 읽는다. 예상하지 못한 필드 수를 찾는다.
	for {
		// 열을 읽는다. 파일 종료 지점에 도달했는지 확인.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 값을 읽는 과정에서 오류가 발생하면 오류를 로그에 기록하고 계속 진핸한다.
		if err != nil {
			log.Println(err)
			continue
		}

		// 레코드가 기대한 필드 수를 갖는 경우
		// 데이터 집합에 레코드를 추가
		rawCSVData = append(rawCSVData, record)
	}

	fmt.Println(rawCSVData)
}
