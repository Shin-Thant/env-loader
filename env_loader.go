package envloader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var envCache = make(map[string]map[string]string)

const _KEY_INDEX = 0
const _VALUE_INDEX = 1

type LoadEnvOptions struct {
	EnvPath string
}

func LoadEnv(accepter interface{}, options *LoadEnvOptions) error {
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
		return nil
	}

	valType := val.Type()
	stringParser := newStringParser()

	targetFindPath := ""
	if options != nil {
		targetFindPath = options.EnvPath
	}
	envFilePath, foundEnvFile, err := getEnvFilePath(targetFindPath)
	if err != nil {
		return err
	}

	_, cacheFound := envCache[envFilePath]
	if foundEnvFile && !cacheFound {
		err = loadEnvFileIntoCache(envFilePath)
		if err != nil {
			return err
		}
	}

	for i := 0; i < fieldCount; i++ {
		fieldName := valType.Field(i).Name
		envVal, found := getEnv(fieldName, envFilePath)
		if !found {
			continue
		}
		field := val.Field(i)

		parsedResult, err := stringParser.Parse(envVal, field.Kind())
		if err != nil {
			fmt.Println("[EnvLoaderError] :", err)
			continue
		}
		if field.Type().Kind() != reflect.Interface && reflect.TypeOf(parsedResult).Kind() != field.Type().Kind() {
			continue
		}
		field.Set(reflect.ValueOf(parsedResult))
	}

	return nil
}

func getEnv(key string, cachedFileKey string) (string, bool) {
	if cachedFileKey == "" {
		return os.LookupEnv(key)
	}
	foundVal, found := envCache[cachedFileKey][key]
	return foundVal, found
}

func loadEnvFileIntoCache(path string) error {
	envCache[path] = make(map[string]string)

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("couldn't open .env file at `%s`: %v", path, err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		splittedLine := strings.SplitN(line, "=", 2)
		if len(splittedLine) != 2 {
			continue
		}
		envCache[path][strings.Trim(splittedLine[_KEY_INDEX], " ")] = strings.Trim(splittedLine[_VALUE_INDEX], " ")
	}
	return nil
}

func getEnvFilePath(path string) (string, bool, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", false, fmt.Errorf("couldn't get root dir: %v", err)
	}

	if path == "" {
		defaultEnvPath := rootDir + "/.env"
		_, err = os.Stat(defaultEnvPath)
		if err != nil {
			return "", false, nil
		}
		return defaultEnvPath, true, nil
	}

	customPath, err := filepath.Abs(filepath.Join(rootDir, path))
	if err != nil {
		return "", false, fmt.Errorf("couldn't get absolute path using %s: %v", path, err)
	}
	_, err = os.Stat(customPath)
	if err != nil {
		return "", false, fmt.Errorf("couldn't find %s file at %s", path, customPath)
	}
	return customPath, true, nil
}
