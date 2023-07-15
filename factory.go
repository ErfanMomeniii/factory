package factory

import (
	"math/rand"
	"reflect"
)

type Factory struct {
	structure interface{}
	desired   map[string]any
}

func NewFactory() *Factory {
	return &Factory{
		structure: make([]reflect.StructField, 0),
		desired:   make(map[string]any),
	}
}

func (factory *Factory) Model(t any) *Factory {
	myType := reflect.TypeOf(t)
	if myType.Kind() != reflect.Struct {
		return nil
	}

	factory.structure = t

	return factory
}

func (factory *Factory) Set(field string, value any) *Factory {
	if f, b := reflect.TypeOf(factory.structure).FieldByName(field); !b || reflect.TypeOf(value) != f.Type {
		return factory
	}

	factory.desired[field] = value

	return factory
}

func (factory *Factory) Generate(count int) []interface{} {
	var answer []interface{}
	for i := 0; i < count; i++ {
		instance := generate(factory)
		answer = append(answer, instance)
	}

	return answer
}

func generate(factory *Factory) interface{} {
	structField := reflect.New(reflect.TypeOf(factory.structure)).Elem()

	for i := 0; i < structField.NumField(); i++ {
		if v, ok := factory.desired[reflect.TypeOf(factory.structure).Field(i).Name]; ok {
			structField.Field(i).Set(reflect.ValueOf(v))
		} else {
			setRandomValue(structField.Field(i))
		}
	}

	return structField.Interface()
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		result[i] = charset[randomIndex]
	}

	return string(result)
}

func setRandomValue(t reflect.Value) {
	sign := 1
	lowBit := 0
	if rand.Intn(1e1) == 1 {
		lowBit = 1
		sign = -1
	}

	switch t.Type().String() {
	case "bool":
		t.SetBool(rand.Intn(2) == 1)
	case "string":
		t.SetString(generateRandomString(rand.Int() % 100))
	case "int":
		t.SetInt(int64(rand.Int() * sign))
	case "int8":
		t.SetInt(int64(rand.Intn(1e7) * sign))
	case "int16":
		t.SetInt(int64(rand.Intn(1e15) * sign))
	case "int32":
		t.SetInt(int64(rand.Int31() * int32(sign)))
	case "int64":
		t.SetInt(rand.Int63() * int64(sign))
	case "uint":
		t.SetUint(uint64(rand.Int()*2 + lowBit))
	case "uint8":
		t.SetUint(uint64(rand.Intn(1e8)))
	case "uint16":
		t.SetUint(uint64(rand.Intn(1e16)))
	case "uint32":
		t.SetUint(uint64(rand.Uint32()))
	case "uint64":
		t.SetUint(rand.Uint64())
	case "float32":
		t.SetFloat(float64(rand.Float32()))
	case "float64":
		t.SetFloat(rand.Float64())
	case "complex64":
		t.SetComplex(complex(rand.Float64(), rand.Float64()))
	case "complex128":
		t.SetComplex(complex(rand.Float64()*rand.Float64(), rand.Float64()*rand.Float64()))
	default:
		if t.Kind() == reflect.Struct {
			for i := 0; i < t.NumField(); i++ {
				setRandomValue(t.Field(i))
			}
		}
	}

	return
}
