package gee

//前缀树节点信息
type node struct {
	pattern 	string  //待匹配路由，例如 /p/:lang
	part 		string  //路由中的一部分，例如 :lang
	children	[]*node //子节点，例如[doc,intro]
	isWild		bool	//是否精确匹配，part含有 : 或 * 时为true
}
