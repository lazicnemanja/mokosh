package argsparser

import "testing"

func TestParse(t *testing.T) {

	args := []string{"./file.mwb", "model-name", "1.0.3"}
	keys := []string{"path", "name", "version"}

	res := Parse(args, keys)

	for i, v := range keys {

		if val, ok := res[v]; ok {
			if args[i] != val {
				t.Errorf("Parse(args, keys) = %q, want %q", val, args[i])
			}
		} else {
			t.Errorf("Missing key in map %q", v)
		}

	}
}
