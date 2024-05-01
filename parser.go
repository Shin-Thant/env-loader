package envloader

import (
	"fmt"
	"reflect"
)

type stringParser struct {
	parserMap parserMap
}

func newStringParser() *stringParser {
	return &stringParser{
		parserMap: createParserMap(),
	}
}

func (c *stringParser) Parse(target string, into reflect.Kind) (interface{}, error) {
	parser, found := c.parserMap[into]
	if !found {
		return nil, fmt.Errorf("unsupported type %v", into)
	}
	return parser(target)
}

type parserMethod func(val string) (interface{}, error)
type parserMap map[reflect.Kind]parserMethod

func createParserMap() parserMap {
	parsers := parsers{}

	parserMap := make(map[reflect.Kind]parserMethod)
	parserMap[reflect.Interface] = parsers.ParseInterface
	parserMap[reflect.String] = parsers.ParseString
	parserMap[reflect.Int] = parsers.ParseInt
	parserMap[reflect.Int8] = parsers.ParseInt8
	parserMap[reflect.Int16] = parsers.ParseInt16
	parserMap[reflect.Int32] = parsers.ParseInt32
	parserMap[reflect.Int64] = parsers.ParseInt64
	parserMap[reflect.Float32] = parsers.ParseFloat32
	parserMap[reflect.Float64] = parsers.ParseFloat64
	parserMap[reflect.Bool] = parsers.ParseBool

	return parserMap
}
