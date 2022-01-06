package edsl

import "fmt"

type TagStyle int

const (
	TagStyleComma TagStyle = iota
	TagStyleCommaEqSpace
	TagStyleSemiColon
)

type Tagger interface {
	// Comma `name:"key1,key2,,,"`
	Comma(name string) CommaTag
	// CommaEqSpace `name:"key1=value,key2=value1 value2,,,key3,,,"`
	CommaEqSpace(name string) CommaEqSpaceTag
	// SemiColon `name:"key1:value2;key2;key3:value3;;;"`
	SemiColon(name string) SemiColonTag
}

type Tag interface {
	fmt.Stringer

	Name() string
	FirstKey() string
	Value(key string) string
	Values(key string) []string
	Of(tagger Tag) Tag
}

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
	EntryString(key string, values ...string) CommaEqSpaceTag
}

type SemiColonTag interface {
	Tag

	Nil() SemiColonTag
	Key(key string) SemiColonTag
	Entry(key string, value interface{}) SemiColonTag
}
