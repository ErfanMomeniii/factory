package factory_test

import (
	"testing"

	"github.com/erfanmomeniii/factory"
	"github.com/stretchr/testify/assert"
)

type Info struct {
	Name  string
	Phone int
}

func Test_model(t *testing.T) {
	f := factory.NewFactory()

	instances := f.Model(Info{}).
		Set("Name", "Erfan").
		Generate(2)

	assert.Equal(t, len(instances), 2)
	assert.IsType(t, instances[0], Info{})
	assert.Equal(t, instances[0].(Info).Name, "Erfan")
	assert.IsType(t, instances[1], Info{})

	instances = f.Model(Info{}).
		Set("Family", "Erfan").
		Generate(1)
	assert.Equal(t, len(instances), 1)
}
