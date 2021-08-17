package edsl

import "fmt"

type Tagger interface {
	// AsIs is just as-is.
	AsIs(text string) AsIsTag
	// Comma `name:"key1,key2,,,"`
	Comma(name string) CommaTag
	// CommaEqSpace `name:"key1=value,key2=value1 value2,,,key3,,,"`
	CommaEqSpace(name string) CommaEqSpaceTag
	// SemiComma `name:"key1:value2;key2;key3:value3;;;"`
	SemiComma(name string) SemiCommaTag
}

type Tag interface {
	fmt.Stringer

	Name() string
	FirstKey() string
}

type AsIsTag interface{ Tag }

type CommaTag interface {
	Tag

	Nil() CommaTag
	Key(key string) CommaTag
}

type CommaEqSpaceTag interface {
	Tag

	Nil() CommaEqSpaceTag
	Key(key string) CommaEqSpaceTag
	Entry(key string, values ...interface{}) CommaEqSpaceTag
}

type SemiCommaTag interface {
	Tag

	Nil() SemiCommaTag
	Key(key string) SemiCommaTag
	Entry(key string, value interface{}) SemiCommaTag
}
