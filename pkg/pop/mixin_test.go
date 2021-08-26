package pop

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMixin(t *testing.T) {
	p, err := Load("./testdata/test.cue")
	if err != nil {
		t.Fatal(err)
	}

	got, err := p.Mixin()
	if err != nil {
		t.Fatal(err)
	}

	var want Mixin
	data, err := os.ReadFile("./testdata/test.json")
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(data, &want); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, *got); diff != "" {
		fmt.Println(diff)
		t.Fail()
	}
}
