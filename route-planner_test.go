package main

import (
	"reflect"
	"testing"
)

func TestGetRouteStartAndEnd(t *testing.T) {
	route := [][]string{{"A", "B"}, {"B", "C"}, {"D", "E"}}
	expected := Point{Source: "A", Destination: "E"}
	result := GetRouteStartAndEnd(route)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
func TestGetRouteStartAndEndExample(t *testing.T) {
	route := [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}
	expected := Point{Source: "SFO", Destination: "EWR"}
	result := GetRouteStartAndEnd(route)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
func TestGetRouteStartAndEndExample2(t *testing.T) {
	route := [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}
	expected := Point{Source: "SFO", Destination: "EWR"}
	result := GetRouteStartAndEnd(route)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
