package edit

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTrim(t *testing.T) {
	cases := []struct {
		name   string
		before string
		after  string
	}{
		{
			name:   "simple",
			before: `{"foo": "bar", "baz": null}`,
			after:  `{"foo": "bar"}`,
		},
		{
			name:   "nested",
			before: `{"foo": "bar", "baz": { "foo": null}}`,
			after:  `{"foo": "bar", "baz": {}}`,
		},
		{
			name:   "list-msi",
			before: `{"list": [{"foo": null}]}`,
			after:  `{"list": [{}]}`,
		},
		{
			name:   "list-plain",
			before: `{"list": [null]}`,
			after:  `{"list": [null]}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data := unmarshal(c.before)
			trim(data)
			if diff := cmp.Diff(unmarshal(c.after), data); diff != "" {
				fmt.Println(diff)
				t.Fail()
			}
		})
	}
}

func unmarshal(data string) map[string]interface{} {
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		panic(err)
	}
	return m
}
