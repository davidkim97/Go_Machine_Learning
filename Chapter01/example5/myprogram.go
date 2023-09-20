package main

import (
	"fmt"
	"github.com/kniren/gota/dataframe"
	"log"
	"os"
)

func main() {

	// csv 파일 읽기
	irisFile, err := os.Open("/Users/david/dev/Golang/src/golang_ml/data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// csv 파일로부터 데이터프레임 생성.
	// 열의 유형은 추론된다.
	irisDF := dataframe.ReadCSV(irisFile)

	// 데이터 프레임의 필터를 생성.
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// 붓꽃(iris) 품종이 "Iris-versicolor"인 행만 볼 수 있도록
	// 데이터프레임을 필터링함.
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	// 데이터프레임을 다시 필터링함. 하지만 이번에는
	// sepal_width 및 species 열만 선택한다.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})

	// 데이터프레임을 필터링하고 다시 선택함. 하지만 이번에는
	// 처음 세 개의 결과만 보여줌.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)
}
