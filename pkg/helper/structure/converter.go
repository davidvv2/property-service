package structure

var _ Converter[any] = (*ConverterStdImpl[any, any])(nil)

type Converter[OldStruct any] interface {
	Convert(old *OldStruct, mappingTag string) (map[string]interface{}, error)
	GetOldManipulator() Manipulator[OldStruct]
}
