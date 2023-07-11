package factory_test

import (
	"github.com/erfanmomeniii/factory"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Op struct {
	Name  string
	Phone int
}

func Test_model(t *testing.T) {
	f := factory.NewFactory()

	instances := f.Model(Op{}).
		Set("Name", "Erfan").
		Generate(2)

	assert.Equal(t, len(instances), 2)
	assert.IsType(t, instances[0], Op{})
	assert.Equal(t, instances[0].(Op).Name, "Erfan")
	assert.IsType(t, instances[1], Op{})
}
