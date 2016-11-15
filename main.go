package gotest

import (
	"fmt"
	"runtime"
)

// TestScorer テストの結果を保管する
type TestScorer struct {
	FailedCount int
}

// Scorer テストのスコアを保管するグローバルな変数
var Scorer TestScorer

// Assert アサート関数
var Assert func(bool, string)

func init() {
	Scorer = TestScorer{0}

	Assert = func(testResult bool, msg string) {
		fmt.Printf("test result %v", testResult)
		if !testResult {
			Scorer.FailedCount++
			fmt.Printf("cnt: %v\n", Scorer.FailedCount)
			_, filename, line, _ := runtime.Caller(1)
			fmt.Printf("%v:%v: %v\n", filename, line, msg)
		}
	}
}
