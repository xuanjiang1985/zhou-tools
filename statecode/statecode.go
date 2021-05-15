package statecode

// 成功状态
const (
	// 请求成功
	Success = 0

	// 通用状态：无效|禁用
	StateDistable = 0
	// 通用状态：有效|启用
	StateEnable = 1
)

// 错误状态
const (
	// 参数为空
	ErrorParamEmpty = 2000
	// 参数不允许
	ErrorParamNotAllow = 2001
	// 数据库查询错误
	ErrorDBSelect = 2002
	// 数据库更新错误
	ErrorDBUpdate = 2003
	// 数据库插入错误
	ErrorDBInsert = 2004
	// 网络请求错误
	ErrorInterface = 2005
)
