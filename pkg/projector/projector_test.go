package projector_test

import (
	"testing"

	"github.com/mrohner94/compare/pkg/projector"
)

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"baz": "foo1",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(
		&projector.Config{
			Args:      []string{},
			Operation: projector.Print,
			Pwd:       pwd,
			Config:    "Hello, config",
		},
		data,
	)
}

func test(t *testing.T, proj *projector.Projector, key, value string) {
	v, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("expected to find value \"%v\"", value)
	}

	if value != v {
		t.Errorf("expected to find %v but received %v", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)
	test(t, proj, "foo", "bar3")
	test(t, proj, "baz", "foo1")
}

func TestSetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)
	test(t, proj, "foo", "bar3")
	proj.SetValue("foo", "bar4")
	test(t, proj, "foo", "bar4")
	proj.SetValue("baz", "bugs")
	test(t, proj, "baz", "bugs")

	proj = getProjector("/", data)
	test(t, proj, "baz", "foo1")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)
	test(t, proj, "foo", "bar3")
	proj.RemoveValue("foo")
	test(t, proj, "foo", "bar2")

	proj.RemoveValue("baz")
	test(t, proj, "baz", "foo1")
}
