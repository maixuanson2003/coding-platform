package dto

import "reflect"

type Mapper[T any, R any] struct {
	Fields []string
}

func (m *Mapper[T, R]) EntityToResponse(entity T) R {
	var Response R
	entityValue := reflect.ValueOf(entity)
	responseValue := reflect.ValueOf(&Response).Elem()

	for _, item := range m.Fields {

		entityField := entityValue.FieldByName(item)
		respField := responseValue.FieldByName(item)

		if entityField.IsValid() && respField.IsValid() && respField.CanSet() {
			respField.Set(entityField)
		}
	}
	return Response

}
