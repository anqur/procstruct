package internal

import (
	"fmt"
	"strings"

	"github.com/anqur/procstruct/edsl"
	"github.com/fatih/structtag"
)

func ParseTag(text string) (ret []edsl.Tag) {
	tags, err := structtag.Parse(text)
	if err != nil {
		panic(err)
	}

	for _, tag := range tags.Tags() {
		name := tag.Key
		raw := tag.Value()
		style := TagStyles.Get(name)

		switch style {
		case edsl.TagStyleComma:
			t := Tagger{}.Comma(name)
			for _, key := range strings.Split(raw, ",") {
				t = t.Key(key)
			}
			ret = append(ret, t)
		case edsl.TagStyleCommaEqSpace:
			t := Tagger{}.CommaEqSpace(name)
			for _, rawEntry := range strings.Split(raw, ",") {
				entry := strings.SplitN(rawEntry, "=", 2)
				if len(entry) == 1 {
					t = t.Key(entry[0])
					continue
				}
				var vals []interface{}
				for _, v := range strings.Split(entry[1], " ") {
					vals = append(vals, v)
				}
				t = t.Entry(entry[0], vals...)
			}
			ret = append(ret, t)
		case edsl.TagStyleSemiColon:
			t := Tagger{}.SemiColon(name)
			for _, rawEntry := range strings.Split(raw, ";") {
				entry := strings.SplitN(rawEntry, ":", 2)
				if len(entry) == 1 {
					t = t.Key(entry[0])
					continue
				}
				t = t.Entry(entry[0], entry[1])
			}
			ret = append(ret, t)
		default:
			panic(fmt.Errorf("unknown tag style: %v", style))
		}
	}

	return
}

type Tagger struct{}

func (Tagger) Comma(name string) edsl.CommaTag {
	return CommaTag{name: name}
}

func (Tagger) CommaEqSpace(name string) edsl.CommaEqSpaceTag {
	return CommaEqSpaceTag{name: name}
}

func (t Tagger) SemiColon(name string) edsl.SemiColonTag {
	return SemiCommaTag{name: name}
}

type CommaTag struct {
	name  string
	items []string
}

func (c CommaTag) Name() string { return c.name }

func (c CommaTag) FirstKey() string {
	if len(c.items) == 0 {
		return ""
	}
	return c.items[0]
}

func (c CommaTag) Value(string) string { return c.FirstKey() }

func (c CommaTag) String() string {
	return fmt.Sprintf("%s:%q", c.name, strings.Join(c.items, ","))
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
	name string

	entries []*cesEntry
}

func (c CommaEqSpaceTag) Name() string { return c.name }

func (c CommaEqSpaceTag) FirstKey() string {
	if len(c.entries) == 0 {
		return ""
	}
	if entry := c.entries[0]; entry != nil {
		return entry.Key
	}
	return ""
}

func (c CommaEqSpaceTag) Value(key string) string {
	for _, entry := range c.entries {
		if entry == nil {
			continue
		}
		if entry.Key != key {
			continue
		}
		if len(entry.Vals) == 0 {
			return ""
		}
		return entry.Vals[0]
	}
	return ""
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
	return fmt.Sprintf("%s:%q", c.name, strings.Join(entries, ","))
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
	name string

	entries []*scEntry
}

func (s SemiCommaTag) Name() string { return s.name }

func (s SemiCommaTag) FirstKey() string {
	if len(s.entries) == 0 {
		return ""
	}
	if entry := s.entries[0]; entry != nil {
		return entry.Key
	}
	return ""
}

func (s SemiCommaTag) Value(key string) string {
	for _, entry := range s.entries {
		if entry == nil {
			continue
		}
		if entry.Key != key {
			continue
		}
		return entry.Val
	}
	return ""
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
	return fmt.Sprintf("%s:%q", s.name, strings.Join(entries, ";"))
}

func (s SemiCommaTag) Nil() edsl.SemiColonTag {
	s.entries = append(s.entries, nil)
	return s
}

func (s SemiCommaTag) Key(key string) edsl.SemiColonTag {
	s.entries = append(s.entries, &scEntry{Key: key})
	return s
}

func (s SemiCommaTag) Entry(key string, value interface{}) edsl.SemiColonTag {
	s.entries = append(
		s.entries,
		&scEntry{Key: key, Val: fmt.Sprintf("%v", value)},
	)
	return s
}
