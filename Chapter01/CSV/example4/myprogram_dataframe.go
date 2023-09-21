package main

import (
	"fmt"
	"github.com/kniren/gota/dataframe"
	"log"
	"os"
)

func main() {

	// open csv file
	irisFile, err := os.Open("/Users/david/dev/Golang/src/golang_ml/data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// csv 파일로부터 데이터프레임 생성하기.
	// 열의 유형은 추론됨.
	irisDF := dataframe.ReadCSV(irisFile)
	// 검사를 위해 레코드를 stdout(표준 출력)으로 보여줌.
	// gota 패키지는 적절한 형태로 출력될 수 있도록 데이터프레임의 형식을 지정한다.
	fmt.Println(irisDF)
}
