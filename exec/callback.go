package exec

type ContractInvoker interface {
	InvokeInternal(contractAddr string, ownerAddr string, callerAddr string, vmContext *VMContext, code []byte, contractName string, method string, params interface{}, rtnType string) (interface{}, error)
}
