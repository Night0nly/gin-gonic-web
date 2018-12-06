package service

import "testing"

func TestChiasenhacCollector_Search(t *testing.T) {
	collector := NewChiasenhacCollector()
	songs, err := collector.Search("Dem ngay xa em")
	if err != nil {
		t.Fatal("Fail test")
	}
	if len(*songs) < 1 {
		t.Fatal("Fail getting any song from the internet")
	}
}
