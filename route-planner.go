package main

type Point struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

func isNodeHasDependencies(node string, inverted map[string]string) bool {
	_, k := inverted[node]
	return k
}

func GetRouteStartAndEnd(route [][]string) Point {
	sourceMap := make(map[string]string, 0)
	invertedSourceMap := make(map[string]string, 0)

	for _, p := range route {
		source, destination := p[0], p[1]
		sourceMap[source] = destination
		invertedSourceMap[destination] = source
	}
	head := getIndependentNode(sourceMap, invertedSourceMap)
	tail := getIndependentNode(invertedSourceMap, sourceMap)
	return Point{Source: head, Destination: tail}
}

func getIndependentNode(sourceMap map[string]string, invertedSourceMap map[string]string) string {
	for key, _ := range sourceMap {
		if _, k := invertedSourceMap[key]; !k {
			return key
		}
	}
	return ""
}
