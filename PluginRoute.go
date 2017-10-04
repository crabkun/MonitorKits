package MonitorKits


type routeNode struct {
	Method string
	Path string
	FuncName string
}

type PluginRoute []routeNode

func (t *PluginRoute)Add(method string,path string,funcName string){
	if *t==nil{
		*t=make([]routeNode,0)
	}
	*t=append(*t,routeNode{
		Method:method,
		Path:path,
		FuncName:funcName,
	})
}