declare module "stdgo/text/template/parse" {
    import {
        Float32, Float64,
        Int64, Int32, Int16, Int8, Int,
        Uint64, Uint32, Uint16, Uint8, Uint,
        Number, NumberLike,
        Byte, Bytes, Rune, Runes,
        Float32Slice, Float64Slice,
        Int64Slice, Int32Slice, Int16Slice, Int8Slice, IntSlice,
        Uint64Slice, Uint32Slice, Uint16Slice, Uint8Slice, UintSlice,
        Error,
        ReadChannel, WriteChannel, Channel,
        Slice, Map,
        Uintptr, Native,
        Complex128,
    } from "stdgo/builtin";

    function IsEmptyTree(n: Node): boolean
    function Parse(name: string, text: string, leftDelim: string, rightDelim: string, ...funcs: Array<Map<string, any>>): Map<string, TreePointer>

    interface ActionNodePointer extends NodeType, Pos {
        readonly __ActionNodePointer: ActionNodePointer
        NodeType: NodeType
        Pos: Pos

        Line: Int       // The line number in the input. Deprecated: Kept for compatibility.
        Pipe: PipeNodePointer // The pipeline in the action.
        // contains filtered or unexported fields

        Copy(): Node
        String(): string
    }
    interface BoolNodePointer extends NodeType, Pos {
        readonly __BoolNodePointer: BoolNodePointer
        NodeType: NodeType
        Pos: Pos

        True: boolean // The value of the boolean constant.
        // contains filtered or unexported fields

        Copy(): Node
        String(): string
    }
    interface BranchNodePointer extends NodeType, Pos {
        readonly __BranchNodePointer: BranchNodePointer
        NodeType: NodeType
        Pos: Pos

        Line: Int       // The line number in the input. Deprecated: Kept for compatibility.
        Pipe: PipeNodePointer // The pipeline to be evaluated.
        List: ListNodePointer // What to execute if the value is non-empty.
        ElseList: ListNodePointer // What to execute if the value is empty (nil if absent).
        // contains filtered or unexported fields

        Copy(): Node
        String(): string
    }
    interface ChainNodePointer extends NodeType, Pos {
        readonly __ChainNodePointer: ChainNodePointer
        NodeType: NodeType
        Pos: Pos

        Node: Node
        Field: Array<string> // The identifiers in lexical order.
        // contains filtered or unexported fields

        Add(field: string): void
        Copy(): Node
        String(): string
    }
    interface CommandNodePointer extends NodeType, Pos {
        readonly __CommandNodePointer: CommandNodePointer
        NodeType: NodeType
        Pos: Pos

        Args: Array<Node> // Arguments in lexical order: Identifier, field, or constant.
        // contains filtered or unexported fields

        Copy(): Node
        String(): string
    }
    interface CommentNodePointer extends NodeType, Pos {
        readonly __CommentNodePointer: CommentNodePointer
        NodeType: NodeType
        Pos: Pos

        Text: string // Comment text.
        // contains filtered or unexported fields
        Copy(): Node
        String(): string
    }
    interface DotNodePointer extends NodeType, Pos {
        readonly __DotNodePointer: DotNodePointer
        NodeType: NodeType
        Pos: Pos
        // contains filtered or unexported fields
        Copy(): Node
        String(): string
        Type(): NodeType
    }
    interface FieldNodePointer extends NodeType, Pos {
        readonly __FieldNodePointer: FieldNodePointer
        NodeType: NodeType
        Pos: Pos

        Ident: Array<string> // The identifiers in lexical order.
        // contains filtered or unexported fields
        Copy(): Node
        String(): string
    }
    function NewIdentifier(ident: string): IdentifierNodePointer
    interface IdentifierNodePointer extends NodeType, Pos {
        readonly __IdentifierNodePointer: IdentifierNodePointer
        NodeType: NodeType
        Pos: Pos

        Ident: string // The identifier's name.
        // contains filtered or unexported fields
        Copy(): Node
        SetPos(pos: Pos): IdentifierNodePointer
        SetTree(t: TreePointer): IdentifierNodePointer
        String(): string
    }
    interface IfNodePointer extends BranchNodePointer {
        readonly __IfNodePointer: IfNodePointer
    }
    interface ListNodePointer extends NodeType, Pos {
        readonly __FieldNodePointer: FieldNodePointer
        NodeType: NodeType
        Pos: Pos

        Nodes: Array<Node> // The element nodes in lexical order.
        // contains filtered or unexported fields
        Copy(): Node
        CopyList(): ListNodePointer
        String(): string
    }
    function Mode(v: uint | NumberLike): Mode
    interface Mode extends Number {
        readonly __Mode: Mode
    }
    const ParseComments = Mode(1)
    const SkipFuncCheck = Mode(1 << 1)

    interface NilNodePointer extends NodeType, Pos {
        readonly __NilNodePointer: NilNodePointer
        NodeType: NodeType
        Pos: Pos
        // contains filtered or unexported fields
        Copy(): Node
        String(): string
        Type(): NodeType
    }
    interface Node extends Native {
        Type(): NodeType
        String(): string
        // Copy does a deep copy of the Node and all its components.
        // To avoid type assertions, some XxxNodes also have specialized
        // CopyXxx methods that return *XxxNode.
        Copy(): Node
        Position(): Pos // byte position of start of node in full original input string
        // contains filtered or unexported methods
    }
    function NodeType(v: Int | NumberLike): NodeType
    interface NodeType extends Number {
        readonly __NodeType: NodeType
        Type(): NodeType
    }
    const NodeText = NodeType(0)      // Plain text.
    const NodeAction = NodeType(1)                  // A non-control action such as a field evaluation.
    const NodeBool = NodeType(2)                    // A boolean constant.
    const NodeChain = NodeType(3)                   // A sequence of field accesses.
    const NodeCommand = NodeType(4)                 // An element of a pipeline.
    const NodeDot = NodeType(5)                     // The cursor, dot.
    const NodeField = NodeType(6)      // A field or method name.
    const NodeIdentifier = NodeType(7) // An identifier; always a function name.
    const NodeIf = NodeType(8)         // An if action.
    const NodeList = NodeType(9)       // A list of Nodes.
    const NodeNil = NodeType(10)        // An untyped nil constant.
    const NodeNumber = NodeType(11)     // A numerical constant.
    const NodePipe = NodeType(12)       // A pipeline of commands.
    const NodeRange = NodeType(13)      // A range action.
    const NodeString = NodeType(14)     // A string constant.
    const NodeTemplate = NodeType(15)   // A template invocation action.
    const NodeVariable = NodeType(16)   // A $ variable.
    const NodeWith = NodeType(17)       // A with action.
    const NodeComment = NodeType(18)    // A comment.
    interface NumberNodePointer extends NodeType, Pos {
        readonly __NumberNodePointer: NumberNodePointer
        NodeType: NodeType
        Pos: Pos

        IsInt: boolean       // Number has an integral value.
        IsUint: boolean       // Number has an unsigned integral value.
        IsFloat: boolean       // Number has a floating-point value.
        IsComplex: boolean       // Number is complex.
        Int64: Int64      // The signed integer value.
        Uint64: Uint64     // The unsigned integer value.
        Float64: Float64    // The floating-point value.
        Complex128: Complex128 // The complex value.
        Text: string     // The original textual representation from the input.
        // contains filtered or unexported fields

        Copy(): Node
        String(): string
    }
    interface PipeNodePointer extends NodeType, Pos {
        readonly __PipeNodePointer: PipeNodePointer
        NodeType: NodeType
        Pos: Pos

        Line: Int             // The line number in the input. Deprecated: Kept for compatibility.
        IsAssign: boolean            // The variables are being assigned, not declared.
        Decl: Array<VariableNodePointer> // Variables in lexical order.
        Cmds: Array<CommandNodePointer>  // The commands in lexical order.
        // contains filtered or unexported fields
        Copy(): Node
        CopyPipe(): PipeNodePointer
        String(): string
    }

    function Pos(v: Int | NumberLike): Pos
    interface Pos extends Number {
        readonly __Pos: Pos
        Position(): Pos
    }
    interface RangeNodePointer extends BranchNodePointer {
        readonly __RangeNodePointer: RangeNodePointer
    }
    interface StringNodePointer extends NodeType, Pos {
        readonly __StringNodePointer: StringNodePointer
        NodeType: NodeType
        Pos: Pos

        Quoted: string // The original text of the string, with quotes.
        Text: string // The string, after quote processing.
        // contains filtered or unexported fields
        Copy(): Node
        String(): string
    }
    interface TemplateNodePointer extends NodeType, Pos {
        readonly __TemplateNodePointer: TemplateNodePointer
        NodeType: NodeType
        Pos: Pos

        Line: Int       // The line number in the input. Deprecated: Kept for compatibility.
        Name: string    // The name of the template (unquoted).
        Pipe: PipeNodePointer // The command to evaluate as dot for the template.
        // contains filtered or unexported fields
        Copy(): Node
        String(): string
    }
    interface TextNodePointer extends NodeType, Pos {
        readonly __TextNodePointer: TextNodePointer
        NodeType: NodeType
        Pos: Pos

        Text: Bytes // The text; may span newlines.
        // contains filtered or unexported fields
        Copy(): Node
        String(): string
    }
    function New(name: string, ...funcs: Array<Map<string, any>>): TreePointer
    interface TreePointer extends Native {
        readonly __TreePointer: TreePointer
        Name: string    // name of the template represented by the tree.
        ParseName: string    // name of the top-level template during parsing, for error messages.
        Root: ListNodePointer // top-level root of the tree.
        Mode: Mode      // parsing mode.
        // contains filtered or unexported fields
        Copy(): TreePointer
        /** return (location, context string) */
        ErrorContext(n: Node): [string, string]
        Parse(text: string, leftDelim: string, rightDelim: string, treeSet: Map<string, TreePointer>, ...funcs: Array<Map<string, any>>): TreePointer
    }
    interface VariableNodePointer extends NodeType, Pos {
        readonly __VariableNodePointer: VariableNodePointer
        NodeType: NodeType
        Pos: Pos

        Ident: Array<string> // Variable name and fields in lexical order.
        // contains filtered or unexported fields
        Copy(): Node
        String(): string
    }
    interface WithNodePointer extends BranchNodePointer {
        readonly __WithNode: WithNode
    }
    function isActionNodePointer(v: any): v is ActionNodePointer
    function isBoolNodePointer(v: any): v is BoolNodePointer
    function isBranchNodePointer(v: any): v is BranchNodePointer
    function isChainNodePointer(v: any): v is ChainNodePointer
    function isCommandNodePointer(v: any): v is CommandNodePointer
    function isCommentNodePointer(v: any): v is CommentNodePointer
    function isDotNodePointer(v: any): v is DotNodePointer
    function isFieldNodePointer(v: any): v is FieldNodePointer
    function isIdentifierNodePointer(v: any): v is IdentifierNodePointer
    function isIfNodePointer(v: any): v is IfNodePointer
    function isListNodePointer(v: any): v is ListNodePointer
    function isMode(v: any): v is Mode
    function isNilNodePointer(v: any): v is NilNodePointer
    function isNode(v: any): v is Node
    function isNodeType(v: any): v is NodeType
    function isNumberNodePointer(v: any): v is NumberNodePointer
    function isPipeNodePointer(v: any): v is PipeNodePointer
    function isPos(v: any): v is Pos
    function isRangeNodePointer(v: any): v is RangeNodePointer
    function isStringNodePointer(v: any): v is StringNodePointer
    function isTemplateNodePointer(v: any): v is TemplateNodePointer
    function isTextNodePointer(v: any): v is TextNodePointer
    function isTreePointer(v: any): v is TreePointer
    function isVariableNodePointer(v: any): v is VariableNodePointer
    function isWithNodePointer(v: any): v is WithNodePointer
}