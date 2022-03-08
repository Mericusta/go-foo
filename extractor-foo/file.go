package extractorfoo

type GoFileInfo struct {
	Name                string                                          // 文件名
	Path                string                                          // 相对项目根目录的路径
	ImportStruct        map[string]map[string]struct{}                  // 该文件引入的外部包
	StructDefinitionMap map[string]map[string]*GoStructMemberDefinition // 该文件定义的结构体
}
