package envloader

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func LoadEnv(accepter interface{}) {
	ptrVal := reflect.ValueOf(accepter)
	if ptrVal.Type().Kind() != reflect.Pointer {
		log.Fatalln("[EnvLoaderError] : argument must be pointer to a struct")
	}

	val := ptrVal.Elem()
	if val.Type().Kind() != reflect.Struct {
		log.Fatalln("[EnvLoaderError] : argument must be pointer to a struct")
	}

	fieldCount := val.NumField()
	if fieldCount < 0 {
		return
	}

	valType := val.Type()
	stringParser := newStringParser()

	for i := 0; i < fieldCount; i++ {
		fieldName := valType.Field(i).Name
		envVal, found := os.LookupEnv(fieldName)
		if !found {
			continue
		}
		field := val.Field(i)

		parsedResult, err := stringParser.Parse(envVal, field.Kind())
		if err != nil {
			fmt.Println("[EnvLoaderError] :", err)
			continue
		}
		field.Set(reflect.ValueOf(parsedResult))
	}
}
