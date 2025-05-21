package structure

import (
	"property-service/pkg/infrastructure/log"

	"github.com/go-playground/validator/v10"
)

type ConverterStdImpl[NewStruct, OldStruct any] struct {
	log            log.Logger
	v              *validator.Validate
	converters     map[string]FunctionConverter
	oldManipulator Manipulator[OldStruct]
	newManipulator Manipulator[NewStruct]
}

// NewStdConverter will create a new converter object that will convert two struct types and return a new map.
func NewStdConverter[NewStruct, OldStruct any](
	log log.Logger,
	v *validator.Validate,
) *ConverterStdImpl[NewStruct, OldStruct] {
	newStructManipulator := NewStdManipulator[NewStruct](log)
	oldStructManipulator := NewStdManipulator[OldStruct](log)
	keyNames := oldStructManipulator.keyTypeName
	for k := range newStructManipulator.keyTypeName {
		if _, ok := keyNames[k]; !ok {
			panic("New struct contains values old struct does not.")
		}
	}
	return &ConverterStdImpl[NewStruct, OldStruct]{
		log: log,
		v:   v,
		converters: map[string]FunctionConverter{
			"ObjectID": NewIDConverter(),
		},
		newManipulator: newStructManipulator,
		oldManipulator: oldStructManipulator,
	}
}

func (csi *ConverterStdImpl[NewStruct, OldStruct]) GetOldManipulator() Manipulator[OldStruct] {
	return csi.oldManipulator
}

func (csi *ConverterStdImpl[NewStruct, OldStruct]) Convert(
	old *OldStruct, mappingTag string,
) (map[string]interface{}, error) {
	// Validate the old struct to make sure their is no missing data.
	if validatorErr := csi.v.Struct(*old); validatorErr != nil {
		csi.log.Error("Struct passed is missing required data %+v", validatorErr)
		return nil, validatorErr
	}
	// Flatten the struct to a string key map interface.
	flattened := csi.oldManipulator.Flatten(old)
	// Create new map.
	newMap := make(map[string]interface{}, len(flattened))
	// Get all the keys and tags of the first map.
	newTagsKey := csi.newManipulator.GetAllKeysAndTags()
	types := csi.newManipulator.GetTypes()

	// Iterate thought the keys of the inputted struct
	for key, value := range flattened {
		// keyMap will make the key
		newKey, mapped, err := csi.keyMap(newTagsKey[key], types[key], value, mappingTag)
		if err != nil {
			csi.log.Error("logged a error %+v", err)
			return nil, err
		}
		// Check the results of the keyMap, if nil then this means their was no mapping tag and we should just return a
		if mapped == nil {
			converted, conversionErr := csi.convertMap(types[key], value)
			if conversionErr != nil {
				csi.log.Error("logged a error %+v", err)
				return nil, conversionErr
			}
			newMap[key] = converted
			csi.log.Debug("New Mapped Tag %+v for key %+v with mapping tag %+v", converted, key, mappingTag)
			continue
		}
		newMap[newKey] = mapped
	}
	csi.log.Info("Successfully converted type with new value map returned %+v", newMap)
	return newMap, nil
}

func (csi *ConverterStdImpl[NewStruct, OldStruct]) keyMap(
	tag map[string][]string,
	valueType string,
	value interface{},
	mappingTag string,
) (string, interface{}, error) {
	for k := range tag {
		if k == mappingTag {
			csi.log.Debug("Tag for %+v with value %+v for type %+v", k, value, valueType)
			converted, conversionErr := csi.convertMap(valueType, value)
			return tag[k][0], converted, conversionErr
		}
	}
	return "", nil, nil
}

func (csi *ConverterStdImpl[NewStruct, OldStruct]) convertMap(
	convertType string, key interface{},
) (interface{}, error) {
	switch convertType {
	case "ObjectID":
		return csi.converters[convertType].OldToNew(key)
	default:
		return key, nil
	}
}
