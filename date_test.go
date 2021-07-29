package gct

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	_, err := DateUtils.Format(time.Now(), "yyyy-MM-dd")
	if err != nil {
		t.Errorf("format failed, %s", err.Error())
		return
	}
}

func TestParse(t *testing.T) {
	_, err := DateUtils.Parse("2006-01-02", "yyyy-MM-dd")
	if err != nil {
		t.Errorf("parse failed, %s", err.Error())
		return
	}
}
