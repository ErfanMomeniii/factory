package factory

import (
	"math/rand"
	"reflect"
)

type Factory struct {
	structure interface{}
	defines   map[string]any
}

func NewFactory() *Factory {
	return &Factory{
		structure: make([]reflect.StructField, 0),
		defines:   make(map[string]any),
	}
}

func (f *Factory) Model(t any) *Factory {
	myType := reflect.TypeOf(t)

	if myType.Kind() != reflect.Struct {
		return nil
	}

	f.structure = t

	return f
}

func (f *Factory) Set(field string, value any) *Factory {
	//TODO error handling for type of value
	f.defines[field] = value
	return f
}

func (f *Factory) Generate(count int) []interface{} {
	var answer []interface{}
	for i := 0; i < count; i++ {
		instance := generate(f)
		answer = append(answer, instance)
	}
	return answer
}

func generate(factory *Factory) interface{} {
	structField := reflect.New(reflect.TypeOf(factory.structure)).Elem()

	for i := 0; i < structField.NumField(); i++ {
		if v, ok := factory.defines[reflect.TypeOf(factory.structure).Field(i).Name]; ok {
			structField.Field(i).Set(reflect.ValueOf(v))
		} else {
			structField.Field(i).Set(reflect.ValueOf(generateRandomValue(structField.Field(i).Type())))
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

func generateRandomValue(t reflect.Type) any {
	sign := 1
	lowBit := 0
	if rand.Intn(1e1) == 1 {
		lowBit = 1
		sign = -1
	}

	switch t.String() {
	case "bool":
		return rand.Intn(2) == 1
	case "string":
		return generateRandomString(rand.Int() % 100)
	case "int":
		return rand.Int() * sign
	case "int8":
		return rand.Intn(1e7) * sign
	case "int16":
		return rand.Intn(1e15) * sign
	case "int32":
		return rand.Int31() * int32(sign)
	case "int64":
		return rand.Int63() * int64(sign)
	case "uint":
		return uint64(rand.Int()*2 + lowBit)
	case "uint8":
		return uint64(rand.Intn(1e8))
	case "uint16":
		return uint64(rand.Intn(1e16))
	case "uint32":
		return uint64(rand.Uint32())
	case "uint64":
		return rand.Uint64()
	case "float32":
		return float64(rand.Float32())
	case "float64":
		return rand.Float64()
	case "complex64":
		return complex(rand.Float64(), rand.Float64())
	case "complex128":
		return complex(rand.Float64()*rand.Float64(), rand.Float64()*rand.Float64())
	}
	// TODO extract struct before return
	return nil
}
