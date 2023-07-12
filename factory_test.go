package factory_test

import (
	"github.com/erfanmomeniii/factory"
	"github.com/stretchr/testify/assert"
	"testing"
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
}
