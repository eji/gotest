package gotest

import "testing"

func TestHoge(t *testing.T) {
	Assert(4 == 5, "4 equal 5")
	if Scorer.FailedCount != 1 {
		t.Errorf("test cnt %v", Scorer.FailedCount)
	}

	Assert(true, "true is true")
	if Scorer.FailedCount != 1 {
		t.Errorf("ここに来るのはおかしい")
	}

	Assert(false, "false is false")
	if Scorer.FailedCount != 2 {
		t.Errorf("カウントアップされるべき")
	}
}
