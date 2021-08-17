package internal

import (
	"fmt"
	"strings"

	"github.com/anqur/procstruct/edsl"
)

type Tagger struct{}

func (Tagger) Comma(name string) edsl.CommaTag {
	return CommaTag{Name: name}
}

func (Tagger) CommaEqSpace(name string) edsl.CommaEqSpaceTag {
	return CommaEqSpaceTag{Name: name}
}

func (t Tagger) SemiComma(name string) edsl.SemiCommaTag {
	return SemiCommaTag{Name: name}
}

type CommaTag struct {
	Name string

	items []string
}

func (c CommaTag) String() string {
	return fmt.Sprintf("%s:%q", c.Name, strings.Join(c.items, ","))
}

func (c CommaTag) Nil() edsl.CommaTag {
	c.items = append(c.items, "")
	return c
}

func (c CommaTag) Key(key string) edsl.CommaTag {
	c.items = append(c.items, key)
	return c
}

type cesEntry struct {
	Key  string
	Vals []string
}

type CommaEqSpaceTag struct {
	Name string

	entries []*cesEntry
}

func (c CommaEqSpaceTag) String() string {
	var entries []string
	for _, entry := range c.entries {
		if entry == nil {
			entries = append(entries, "")
			continue
		}
		if len(entry.Vals) == 0 {
			entries = append(entries, entry.Key)
			continue
		}
		entries = append(
			entries,
			entry.Key+"="+strings.Join(entry.Vals, " "),
		)
	}
	return fmt.Sprintf("%s:%q", c.Name, strings.Join(entries, ","))
}

func (c CommaEqSpaceTag) Nil() edsl.CommaEqSpaceTag {
	c.entries = append(c.entries, nil)
	return c
}

func (c CommaEqSpaceTag) Key(key string) edsl.CommaEqSpaceTag {
	c.entries = append(c.entries, &cesEntry{Key: key})
	return c
}

func (c CommaEqSpaceTag) Entry(
	key string,
	values ...interface{},
) edsl.CommaEqSpaceTag {
	var vals []string
	for _, val := range values {
		vals = append(vals, fmt.Sprintf("%v", val))
	}
	c.entries = append(c.entries, &cesEntry{Key: key, Vals: vals})
	return c
}

type scEntry struct {
	Key string
	Val string
}

type SemiCommaTag struct {
	Name string

	entries []*scEntry
}

func (s SemiCommaTag) String() string {
	var entries []string
	for _, entry := range s.entries {
		if entry == nil {
			entries = append(entries, "")
			continue
		}
		if entry.Val == "" {
			entries = append(entries, entry.Key)
			continue
		}
		entries = append(entries, entry.Key+":"+entry.Val)
	}
	return fmt.Sprintf("%s:%q", s.Name, strings.Join(entries, ";"))
}

func (s SemiCommaTag) Nil() edsl.SemiCommaTag {
	s.entries = append(s.entries, nil)
	return s
}

func (s SemiCommaTag) Key(key string) edsl.SemiCommaTag {
	s.entries = append(s.entries, &scEntry{Key: key})
	return s
}

func (s SemiCommaTag) Entry(key string, value interface{}) edsl.SemiCommaTag {
	s.entries = append(
		s.entries,
		&scEntry{Key: key, Val: fmt.Sprintf("%v", value)},
	)
	return s
}
