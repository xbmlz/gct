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

func TestOffset(t *testing.T) {
	_, err := DateUtils.Offset(time.Now(), "-24")
	if err != nil {
		t.Errorf("offset day failed, %s", err.Error())
		return
	}
}

func TestSubDays(t *testing.T) {
	n := DateUtils.SubDays(time.Now(), time.Now())
	if n != 0 {
		t.Error("value must be 0")
		return
	}
}
