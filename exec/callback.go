package exec

type ContractInvoker interface {
	InvokeInternal(vmContext *VMContext, code []byte, contractName string, method string, params interface{}, rtnType string) (interface{}, error)
}
