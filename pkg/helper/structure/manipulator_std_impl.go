package structure

import (
	"errors"
	"reflect"

	"property-service/pkg/infrastructure/log"
)

type ManipulatorStdImpl[Struct any] struct {
	log         log.Logger
	keyTypes    map[string]reflect.Type
	keyTypeName map[string]string
	tags        map[string]map[string][]string
	keys        []string
	numFields   int
}

// NewStdManipulator Will create a new manipulator object from the standard library using reflection.
func NewStdManipulator[Struct any](log log.Logger) *ManipulatorStdImpl[Struct] {
	var model Struct
	structReflect := reflect.TypeOf(model)
	numFields := structReflect.NumField()
	keyTypes := make(map[string]reflect.Type, numFields)
	keyTypeName := make(map[string]string, numFields)
	tag := make(map[string]map[string][]string, numFields)
	keys := make([]string, numFields)
	for i := 0; i < numFields; i++ {
		field := structReflect.Field(i)
		keyTypes[field.Name] = field.Type
		keyTypeName[field.Name] = field.Type.Name()
		tagNames := getTagNames(string(field.Tag))
		var tagMapErr error
		tag[field.Name], tagMapErr = getTagMap(tagNames, field.Tag)
		if tagMapErr != nil {
			log.Panic("Error has occurred while mapping tags %+v on %+v ", tagMapErr, field.Name)
		}
		keys[i] = field.Name
	}
	return &ManipulatorStdImpl[Struct]{
		log:         log,
		numFields:   numFields,
		keyTypes:    keyTypes,
		keyTypeName: keyTypeName,
		keys:        keys,
		tags:        tag,
	}
}

// GetAllTags implements Manipulator.
func (msi *ManipulatorStdImpl[Struct]) GetAllKeysAndTags() map[string]map[string][]string {
	return msi.tags
}

// GetTagByKey implements Manipulator.
func (msi *ManipulatorStdImpl[Struct]) GetTagsByKey(key string) (map[string][]string, error) {
	value, exist := msi.tags[key]
	if !exist {
		return nil, errors.New("the key does not exist")
	}
	return value, nil
}

func (msi *ManipulatorStdImpl[Struct]) GetTagValues(key, tag string) ([]string, error) {
	value, exist := msi.tags[key][tag]
	if !exist {
		return nil, errors.New("the key or tags does not exist")
	}
	return value, nil
}

func (msi *ManipulatorStdImpl[Struct]) GetTagValueWithValidator(
	key, tag string, validator func(tag []string, variableType reflect.Kind) ([]string, error),
) ([]string, error) {
	if len(msi.tags[key][tag]) == 0 {
		return []string{}, nil
	}
	return validator(msi.tags[key][tag], msi.keyTypes[key].Kind())
}

// GetTypes implements Manipulator.
func (msi *ManipulatorStdImpl[Struct]) GetTypes() map[string]string {
	return msi.keyTypeName
}

func (msi *ManipulatorStdImpl[Struct]) GetAllTagWithValidator(
	tag string, validator func(tag []string, variableType reflect.Kind) ([]string, error),
) (map[string][]string, error) {
	var conversionErr error
	var finishedTags = make(map[string][]string)
	for i := 0; i < len(msi.keys); i++ {
		var tags []string
		tags, conversionErr = msi.GetTagValueWithValidator(msi.keys[i], tag, validator)
		if conversionErr != nil {
			return nil, conversionErr
		}
		if len(tags) != 0 {
			finishedTags[msi.keys[i]] = tags
		}
		msi.log.Debug("tags %+v and errors %+v", tags, conversionErr)
	}
	msi.log.Debug("finished tag %+v", finishedTags)
	return finishedTags, conversionErr
}

// Flatten implements Manipulator.
func (msi *ManipulatorStdImpl[Struct]) Flatten(str *Struct) map[string]interface{} {
	flattened := make(map[string]interface{}, len(msi.keyTypes))
	structReflect := reflect.ValueOf(*str)
	for i := 0; i < len(msi.keyTypes); i++ {
		field := structReflect.Field(i)
		keyName := msi.keys[i]
		msi.log.Debug("reflected field:  %+v", field)
		typeData := msi.keyTypes[keyName]
		if field.CanConvert(typeData) {
			flattened[keyName] = field.Interface()
		}
	}
	return flattened
}
