package types

// These are the different supported type groups
const (
	Any               = "@Any"
	Text              = "@Text"
	Math              = "@Math"
	Unmarshal         = "@Unmarshal"
	Marshal           = "@Marshal"
	ReadArray         = "@ReadArray"
	ReadArrayWithType = "@ReadArrayWithType"
	WriteArray        = "@WriteArray"
	ReadIndex         = "@ReadIndex"
	ReadNotIndex      = "@ReadNotIndex"
	ReadMap           = "@ReadMap"
)

// GroupText is an array of the data types that make up the `text` type
var GroupText = []string{
	Generic,
	String,
}

// GroupMath is an array of the data types that make up the `math` type
var GroupMath = []string{
	Number,
	Integer,
	Float,
	Boolean,
}
