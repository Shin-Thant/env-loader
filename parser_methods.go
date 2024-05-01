package envloader

import "strconv"

type parsers struct{}

func (parsers) ParseString(val string) (interface{}, error) {
	return val, nil
}
func (parsers) ParseInterface(val string) (interface{}, error) {
	return val, nil
}
func (parsers) ParseBool(val string) (interface{}, error) {
	return strconv.ParseBool(val)
}
func (parsers) ParseInt(val string) (interface{}, error) {
	return strconv.Atoi(val)
}
func (parsers) ParseInt8(val string) (interface{}, error) {
	return strconv.ParseInt(val, 10, 8)
}
func (parsers) ParseInt16(val string) (interface{}, error) {
	return strconv.ParseInt(val, 10, 16)
}
func (parsers) ParseInt32(val string) (interface{}, error) {
	return strconv.ParseInt(val, 10, 32)
}
func (parsers) ParseInt64(val string) (interface{}, error) {
	return strconv.ParseInt(val, 10, 64)
}
func (parsers) ParseFloat32(val string) (interface{}, error) {
	return strconv.ParseFloat(val, 32)
}
func (parsers) ParseFloat64(val string) (interface{}, error) {
	return strconv.ParseFloat(val, 64)
}
