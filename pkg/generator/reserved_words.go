package generator

// Reserved words for various programming languages

// CppReservedWords contains C++ keywords and reserved identifiers
var CppReservedWords = []string{
	"alignas", "alignof", "and", "and_eq", "asm", "auto", "bitand", "bitor",
	"bool", "break", "case", "catch", "char", "char8_t", "char16_t", "char32_t",
	"class", "compl", "concept", "const", "consteval", "constexpr", "constinit",
	"const_cast", "continue", "co_await", "co_return", "co_yield", "decltype",
	"default", "delete", "do", "double", "dynamic_cast", "else", "enum", "explicit",
	"export", "extern", "false", "float", "for", "friend", "goto", "if", "inline",
	"int", "long", "mutable", "namespace", "new", "noexcept", "not", "not_eq",
	"nullptr", "operator", "or", "or_eq", "private", "protected", "public",
	"register", "reinterpret_cast", "requires", "return", "short", "signed",
	"sizeof", "static", "static_assert", "static_cast", "struct", "switch",
	"template", "this", "thread_local", "throw", "true", "try", "typedef",
	"typeid", "typename", "union", "unsigned", "using", "virtual", "void",
	"volatile", "wchar_t", "while", "xor", "xor_eq",
}

// GoReservedWords contains Go keywords and built-in identifiers
var GoReservedWords = []string{
	"break", "case", "chan", "const", "continue", "default", "defer", "else",
	"fallthrough", "for", "func", "go", "goto", "if", "import", "interface",
	"map", "package", "range", "return", "select", "struct", "switch", "type",
	"var", "append", "bool", "byte", "cap", "close", "complex", "complex64",
	"complex128", "copy", "delete", "error", "false", "float32", "float64",
	"imag", "int", "int8", "int16", "int32", "int64", "iota", "len", "make",
	"new", "nil", "panic", "print", "println", "real", "recover", "rune",
	"string", "true", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
}

// CSharpReservedWords contains C# keywords
var CSharpReservedWords = []string{
	"abstract", "as", "base", "bool", "break", "byte", "case", "catch", "char",
	"checked", "class", "const", "continue", "decimal", "default", "delegate",
	"do", "double", "else", "enum", "event", "explicit", "extern", "false",
	"finally", "fixed", "float", "for", "foreach", "goto", "if", "implicit",
	"in", "int", "interface", "internal", "is", "lock", "long", "namespace",
	"new", "null", "object", "operator", "out", "override", "params", "private",
	"protected", "public", "readonly", "ref", "return", "sbyte", "sealed",
	"short", "sizeof", "stackalloc", "static", "string", "struct", "switch",
	"this", "throw", "true", "try", "typeof", "uint", "ulong", "unchecked",
	"unsafe", "ushort", "using", "virtual", "void", "volatile", "while",
}

// V8ReservedWords contains JavaScript/TypeScript keywords
var V8ReservedWords = []string{
	"break", "case", "catch", "class", "const", "continue", "debugger",
	"default", "delete", "do", "else", "export", "extends", "false", "finally",
	"for", "function", "if", "import", "in", "instanceof", "new", "null",
	"return", "super", "switch", "this", "throw", "true", "try", "typeof",
	"var", "void", "while", "with", "yield", "let", "static", "enum", "await",
	"implements", "interface", "package", "private", "protected", "public",
}

// PythonReservedWords contains Python keywords
var PythonReservedWords = []string{
	"False", "None", "True", "and", "as", "assert", "async", "await", "break",
	"class", "continue", "def", "del", "elif", "else", "except", "finally",
	"for", "from", "global", "if", "import", "in", "is", "lambda", "nonlocal",
	"not", "or", "pass", "raise", "return", "try", "while", "with", "yield",
}

// LuaReservedWords contains Lua keywords
var LuaReservedWords = []string{
	"and", "break", "do", "else", "elseif", "end", "false", "for", "function",
	"goto", "if", "in", "local", "nil", "not", "or", "repeat", "return",
	"then", "true", "until", "while",
}

// DReservedWords contains D language keywords
var DReservedWords = []string{
	"abstract", "alias", "align", "asm", "assert", "auto", "body", "bool",
	"break", "byte", "case", "cast", "catch", "cdouble", "cent", "cfloat",
	"char", "class", "const", "continue", "creal", "dchar", "debug", "default",
	"delegate", "delete", "deprecated", "do", "double", "else", "enum",
	"export", "extern", "false", "final", "finally", "float", "for", "foreach",
	"foreach_reverse", "function", "goto", "idouble", "if", "ifloat",
	"immutable", "import", "in", "inout", "int", "interface", "invariant",
	"ireal", "is", "lazy", "long", "macro", "mixin", "module", "new", "nothrow",
	"null", "out", "override", "package", "pragma", "private", "protected",
	"public", "pure", "real", "ref", "return", "scope", "shared", "short",
	"static", "struct", "super", "switch", "synchronized", "template", "this",
	"throw", "true", "try", "typedef", "typeid", "typeof", "ubyte", "ucent",
	"uint", "ulong", "union", "unittest", "ushort", "version", "void",
	"volatile", "wchar", "while", "with", "__FILE__", "__MODULE__",
	"__LINE__", "__FUNCTION__", "__PRETTY_FUNCTION__", "__gshared",
	"__traits", "__vector", "__parameters",
}

// RustReservedWords contains Rust keywords and reserved identifiers
var RustReservedWords = []string{
	"as", "async", "await", "break", "const", "continue", "crate", "dyn",
	"else", "enum", "extern", "false", "fn", "for", "if", "impl", "in",
	"let", "loop", "match", "mod", "move", "mut", "pub", "ref", "return",
	"self", "Self", "static", "struct", "super", "trait", "true", "type",
	"unsafe", "use", "where", "while", "abstract", "become", "box", "do",
	"final", "macro", "override", "priv", "typeof", "unsized", "virtual",
	"yield", "try", "union",
}
