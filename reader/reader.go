package reader

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

func compareFiles(file os.FileInfo, path string) bool {
	fileFromPath, err := os.Stat(path)
	return err == nil && os.SameFile(file, fileFromPath)
}

func getSchemas(filename string, data []interface{}) ([]map[interface{}]interface{}, error) {
	result := make([]map[interface{}]interface{}, len(data))
	for i, item := range data {
		var ok bool
		if result[i], ok = item.(map[interface{}]interface{}); !ok {
			return nil, fmt.Errorf(
				"error in file %s: schema should have type map[interface{}]interface{}",
				filename,
			)
		}
	}
	return result, nil
}

func getSchemasFromFile(filename string) ([]interface{}, error) {
	inputContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to open file %s",
			filename,
		)
	}
	data := map[interface{}]interface{}{}
	err = yaml.Unmarshal(inputContent, &data)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot parse given schema from file %s",
			filename,
		)
	}
	schemas, ok := data["schemas"].([]interface{})
	if !ok {
		return nil, fmt.Errorf(
			"no schemas found in file %s",
			filename,
		)
	}
	return schemas, nil
}

// ReadSingle gets a list of maps describing schemas from a file
func ReadSingle(filename string) ([]map[interface{}]interface{}, error) {
	schemas, err := getSchemasFromFile(filename)
	if err != nil {
		return nil, err
	}
	return getSchemas(filename, schemas)
}

// ReadAll gets a list of maps describing schemas from files contained
// within given file except that restricted name
// args:
//   filename string - path to file containing paths to schemas
//   restricted - path to a file from which schemas should not be read
// return:
//   1. list of maps of schemas
//   2. error during execution
func ReadAll(filename, restricted string) ([]map[interface{}]interface{}, error) {
	restrictedFile, _ := os.Stat(restricted)
	schemas, err := getSchemasFromFile(filename)
	if err != nil {
		return nil, err
	}
	result := []map[interface{}]interface{}{}
	for _, schema := range schemas {
		newFilename, ok := schema.(string)
		if !ok {
			return nil, fmt.Errorf(
				"in config file %s schemas should be filenames",
				filename,
			)
		}
		if !strings.HasPrefix(newFilename, "embed") && !compareFiles(restrictedFile, newFilename) {
			schemas, err := ReadSingle(newFilename)
			if err == nil {
				result = append(result, schemas...)
			}
		}
	}
	return result, nil
}
