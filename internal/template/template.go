package template

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Masterminds/sprig"
)

const perm = 0644

// Replace replaces the env vars in file
func Replace(f string) error {
	// Read file
	c, err := read(f)
	if err != nil {
		return err
	}

	// Get env vars
	envMap, err := envToMap()
	if err != nil {
		return err
	}

	// Execute template
	t, err := template.
		New("output").
		Funcs(sprig.GenericFuncMap()).
		Parse(string(c))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, envMap); err != nil {
		return err
	}

	// Write file
	return write(f, buf.Bytes())
}

// DryReplace replaces the env vars for output
func DryReplace(c []byte) ([]byte, error) {
	// Get env vars
	envMap, err := envToMap()
	if err != nil {
		return []byte{}, err
	}

	// Execute template
	t, err := template.
		New("output").
		Funcs(sprig.GenericFuncMap()).
		Parse(string(c))
	if err != nil {
		return []byte{}, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, envMap); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), err
}

// read reads the file content
func read(f string) ([]byte, error) {
	return ioutil.ReadFile(f)
}

// write writes the file content
func write(f string, d []byte) error {
	return ioutil.WriteFile(f, d, perm)
}

func envToMap() (map[string]string, error) {
	envMap := make(map[string]string)
	var err error

	for _, v := range os.Environ() {
		s := strings.Split(v, "=")
		envMap[s[0]] = s[1]
	}

	return envMap, err
}
