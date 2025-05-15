package structure

type FunctionConverter interface {
	OldToNew(interface{}) (interface{}, error)
	NewToOld(interface{}) (interface{}, error)
}
