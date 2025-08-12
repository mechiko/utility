package utility

import (
	"fmt"
	"reflect"
	"strings"
)

// массив строк имен структуры в input
// exclude строка со списком через запятую имен полей исключаемых
// prefix строка префикс для генерируемых имен например "rule."
func StructFieldNames(input any, exclude string, prefix string) (out []string) {
	v := reflect.ValueOf(input)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	t := v.Type()
	// Build exclusion set from comma-separated list
	ex := map[string]struct{}{}
	for _, e := range strings.Split(exclude, ",") {
		e = strings.TrimSpace(e)
		if e != "" {
			ex[e] = struct{}{}
		}
	}
	out = make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// Consider only exported fields
		if f.PkgPath != "" {
			continue
		}
		name := f.Name
		if _, skip := ex[name]; skip {
			continue
		}
		out = append(out, fmt.Sprintf("%s%s", prefix, name))
	}
	return out
}
