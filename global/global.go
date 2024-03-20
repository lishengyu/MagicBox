package global

type ClickHandler func(string) string

type F_handlers struct {
	Handlers map[int]ClickHandler
	FuncList map[int]string
	Count    int
}

var GFuncList F_handlers

func init() {
	GFuncList.Handlers = make(map[int]ClickHandler)
	GFuncList.FuncList = make(map[int]string)
}

func Register(p ClickHandler, str string) {
	GFuncList.Handlers[GFuncList.Count] = p
	GFuncList.FuncList[GFuncList.Count] = str
	GFuncList.Count++
}
