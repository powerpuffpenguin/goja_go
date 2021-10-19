package parse

import (
	"text/template/parse"

	"github.com/powerpuffpenguin/goja"
)

func (f *factory) register() {
	f.Set(`IsEmptyTree`, parse.IsEmptyTree)
	f.Set(`Parse`, parse.Parse)

	f.Set(`NewIdentifier`, parse.NewIdentifier)
	f.Set(`Mode`, Mode)
	f.Accessor(`ParseComments`, f.getParseComments, nil)
	f.Accessor(`SkipFuncCheck`, f.getSkipFuncCheck, nil)
	f.Set(`NodeType`, NodeType)
	f.Accessor(`NodeText`, f.getNodeText, nil)
	f.Accessor(`NodeAction`, f.getNodeAction, nil)
	f.Accessor(`NodeBool`, f.getNodeBool, nil)
	f.Accessor(`NodeChain`, f.getNodeChain, nil)
	f.Accessor(`NodeCommand`, f.getNodeCommand, nil)
	f.Accessor(`NodeDot`, f.getNodeDot, nil)
	f.Accessor(`NodeField`, f.getNodeField, nil)
	f.Accessor(`NodeIdentifier`, f.getNodeIdentifier, nil)
	f.Accessor(`NodeIf`, f.getNodeIf, nil)
	f.Accessor(`NodeList`, f.getNodeList, nil)
	f.Accessor(`NodeNil`, f.getNodeNil, nil)
	f.Accessor(`NodeNumber`, f.getNodeNumber, nil)
	f.Accessor(`NodePipe`, f.getNodePipe, nil)
	f.Accessor(`NodeRange`, f.getNodeRange, nil)
	f.Accessor(`NodeString`, f.getNodeString, nil)
	f.Accessor(`NodeTemplate`, f.getNodeTemplate, nil)
	f.Accessor(`NodeVariable`, f.getNodeVariable, nil)
	f.Accessor(`NodeWith`, f.getNodeWith, nil)
	f.Accessor(`NodeComment`, f.getNodeComment, nil)
	f.Set(`Pos`, Pos)
	f.Set(`New`, parse.New)
	f.Set(`isActionNodePointer`, isActionNodePointer)
	f.Set(`isBoolNodePointer`, isBoolNodePointer)
	f.Set(`isBranchNodePointer`, isBranchNodePointer)
	f.Set(`isChainNodePointer`, isChainNodePointer)
	f.Set(`isCommandNodePointer`, isCommandNodePointer)
	f.Set(`isCommentNodePointer`, isCommentNodePointer)
	f.Set(`isDotNodePointer`, isDotNodePointer)
	f.Set(`isFieldNodePointer`, isFieldNodePointer)
	f.Set(`isIdentifierNodePointer`, isIdentifierNodePointer)
	f.Set(`isIfNodePointer`, isIfNodePointer)
	f.Set(`isListNodePointer`, isListNodePointer)
	f.Set(`isMode`, isMode)
	f.Set(`isNilNodePointer`, isNilNodePointer)
	f.Set(`isNode`, isNode)
	f.Set(`isNodeType`, isNodeType)
	f.Set(`isNumberNodePointer`, isNumberNodePointer)
	f.Set(`isPipeNodePointer`, isPipeNodePointer)
	f.Set(`isPos`, isPos)
	f.Set(`isRangeNodePointer`, isRangeNodePointer)
	f.Set(`isStringNodePointer`, isStringNodePointer)
	f.Set(`isTemplateNodePointer`, isTemplateNodePointer)
	f.Set(`isTextNodePointer`, isTextNodePointer)
	f.Set(`isTreePointer`, isTreePointer)
	f.Set(`isVariableNodePointer`, isVariableNodePointer)
	f.Set(`isWithNodePointer`, isWithNodePointer)
}
func isActionNodePointer(i interface{}) bool {
	_, result := i.(*parse.ActionNode)
	return result
}
func isBoolNodePointer(i interface{}) bool {
	_, result := i.(*parse.BoolNode)
	return result
}
func isBranchNodePointer(i interface{}) bool {
	_, result := i.(*parse.BranchNode)
	return result
}
func isChainNodePointer(i interface{}) bool {
	_, result := i.(*parse.ChainNode)
	return result
}
func isCommandNodePointer(i interface{}) bool {
	_, result := i.(*parse.CommandNode)
	return result
}
func isCommentNodePointer(i interface{}) bool {
	_, result := i.(*parse.CommentNode)
	return result
}
func isDotNodePointer(i interface{}) bool {
	_, result := i.(*parse.DotNode)
	return result
}
func isFieldNodePointer(i interface{}) bool {
	_, result := i.(*parse.FieldNode)
	return result
}
func isIdentifierNodePointer(i interface{}) bool {
	_, result := i.(*parse.IdentifierNode)
	return result
}
func isIfNodePointer(i interface{}) bool {
	_, result := i.(*parse.IfNode)
	return result
}
func isListNodePointer(i interface{}) bool {
	_, result := i.(*parse.ListNode)
	return result
}
func isMode(i interface{}) bool {
	_, result := i.(parse.Mode)
	return result
}
func isNilNodePointer(i interface{}) bool {
	_, result := i.(*parse.NilNode)
	return result
}
func isNode(i interface{}) bool {
	_, result := i.(parse.Node)
	return result
}
func isNodeType(i interface{}) bool {
	_, result := i.(parse.NodeType)
	return result
}
func isNumberNodePointer(i interface{}) bool {
	_, result := i.(*parse.NumberNode)
	return result
}
func isPipeNodePointer(i interface{}) bool {
	_, result := i.(*parse.PipeNode)
	return result
}
func isPos(i interface{}) bool {
	_, result := i.(parse.Pos)
	return result
}
func isRangeNodePointer(i interface{}) bool {
	_, result := i.(*parse.RangeNode)
	return result
}
func isStringNodePointer(i interface{}) bool {
	_, result := i.(*parse.StringNode)
	return result
}
func isTemplateNodePointer(i interface{}) bool {
	_, result := i.(*parse.TemplateNode)
	return result
}
func isTextNodePointer(i interface{}) bool {
	_, result := i.(*parse.TextNode)
	return result
}
func isTreePointer(i interface{}) bool {
	_, result := i.(*parse.Tree)
	return result
}
func isVariableNodePointer(i interface{}) bool {
	_, result := i.(*parse.VariableNode)
	return result
}
func isWithNodePointer(i interface{}) bool {
	_, result := i.(*parse.WithNode)
	return result
}
func Pos(v int) parse.Pos {
	return parse.Pos(v)
}
func (f *factory) getNodeText(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeText)
}
func (f *factory) getNodeAction(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeAction)
}
func (f *factory) getNodeBool(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeBool)
}
func (f *factory) getNodeChain(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeChain)
}
func (f *factory) getNodeCommand(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeCommand)
}
func (f *factory) getNodeDot(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeDot)
}
func (f *factory) getNodeField(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeField)
}
func (f *factory) getNodeIdentifier(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeIdentifier)
}
func (f *factory) getNodeIf(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeIf)
}
func (f *factory) getNodeList(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeList)
}
func (f *factory) getNodeNil(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeNil)
}
func (f *factory) getNodeNumber(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeNumber)
}
func (f *factory) getNodePipe(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodePipe)
}
func (f *factory) getNodeRange(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeRange)
}
func (f *factory) getNodeString(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeString)
}
func (f *factory) getNodeTemplate(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeTemplate)
}
func (f *factory) getNodeVariable(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeVariable)
}
func (f *factory) getNodeWith(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeWith)
}
func (f *factory) getNodeComment(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.NodeComment)
}
func NodeType(v int) parse.NodeType {
	return parse.NodeType(v)
}
func (f *factory) getParseComments(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.ParseComments)
}
func (f *factory) getSkipFuncCheck(call goja.FunctionCall) goja.Value {
	return f.runtime.ToValue(parse.ParseComments << 1)
}
func Mode(v uint) parse.Mode {
	return parse.Mode(v)
}
