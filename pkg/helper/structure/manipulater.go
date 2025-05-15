package structure

import "reflect"

var _ Manipulator[any] = (*ManipulatorStdImpl[any])(nil)

type Manipulator[Struct any] interface {
	GetAllKeysAndTags() map[string]map[string][]string
	GetTagsByKey(key string) (map[string][]string, error)
	GetTagValues(key, tag string) ([]string, error)
	GetTagValueWithValidator(key, tag string,
		validator func(tag []string, variableType reflect.Kind) ([]string, error),
	) ([]string, error)
	GetTypes() map[string]string
	GetAllTagWithValidator(tag string,
		validator func(tag []string, variableType reflect.Kind) ([]string, error),
	) (map[string][]string, error)
	Flatten(str *Struct) map[string]interface{}
}
