// Package argv parse command line string into arguments array using the bash syntax.
package argv

var (
	Unbackquote Expander = func(backquoted string) (string, error) {
		return backquoted, nil
	}
)

// Parse splits cmdline string as array of argument array by the '|' character.
//
// The parsing rules is same as bash. The environment variable will be replaced
// and string surround by '`' will be passed to reverse quote parser.
func Parse(cmdline string, expanders ...Expander) ([][]string, error) {
	var backquoteExpander Expander
	var stringExpander Expander
	if len(expanders) > 0 {
		backquoteExpander = expanders[0]
	}
	if len(expanders) > 1 {
		stringExpander = expanders[1]
	}
	return NewParser(NewScanner(cmdline), backquoteExpander, stringExpander).Parse()
}

// MustParse is like Parse but panics if parsing fails.
func MustParse(cmdline string, expanders ...Expander) [][]string {
	parsed, err := Parse(cmdline, expanders...)
	if err != nil {
		panic(err)
	}
	return parsed
}
