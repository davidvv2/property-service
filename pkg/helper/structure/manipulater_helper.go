package structure

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"property-service/pkg/errors"
)

func getTagMap(tagNames []string, tag reflect.StructTag) (map[string][]string, error) {
	keyMap := make(map[string][]string, len(tagNames))
	for nameIndex := 0; nameIndex < len(tagNames); nameIndex++ {
		if name, ok := tag.Lookup(tagNames[nameIndex]); ok && len(tagNames[nameIndex]) != 0 {
			keyMap[tagNames[nameIndex]] = strings.Split(name, ",")
		} else {
			return nil, errors.New(
				fmt.Errorf("please remove the white spaces from tag  %+v", tagNames[nameIndex]).Error(),
			)
		}
	}
	return keyMap, nil
}

func getTagNames(tag string) []string {
	re := regexp.MustCompile(`"([^"]+)"`)
	result := strings.TrimSuffix(re.ReplaceAllString(tag, ""), ":")
	keys := strings.Split(result, ": ")
	tagNames := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		tagNames[i] = strings.TrimSpace(keys[i])
	}
	return tagNames
}
