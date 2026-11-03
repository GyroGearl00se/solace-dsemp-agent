package controllers

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

// DeepEqualPointerValues checks if two values are equal, properly dereferencing
// pointers (including nested pointers like **bool or *string) before comparison.
func DeepEqualPointerValues(a, b reflect.Value) bool {
	if !a.IsValid() && !b.IsValid() {
		return true
	}
	if !a.IsValid() || !b.IsValid() {
		return false
	}

	// Unwrap pointers until we reach a non-pointer or a nil pointer
	for a.Kind() == reflect.Ptr && !a.IsNil() && b.Kind() == reflect.Ptr && !b.IsNil() {
		a = a.Elem()
		b = b.Elem()
	}

	// If one is still a pointer and the other is not (or one is nil and the other is not)
	if a.Kind() == reflect.Ptr || b.Kind() == reflect.Ptr {
		return a.Kind() == b.Kind() && a.IsNil() && b.IsNil()
	}

	// If types differ after pointer unwrapping
	if a.Type() != b.Type() {
		return false
	}

	switch a.Kind() {
	case reflect.Slice:
		if a.IsNil() && b.IsNil() {
			return true
		}
		if a.Len() != b.Len() {
			return false
		}
		for i := 0; i < a.Len(); i++ {
			if !DeepEqualPointerValues(a.Index(i), b.Index(i)) {
				return false
			}
		}
		return true

	case reflect.Struct:
		for i := 0; i < a.NumField(); i++ {
			if !DeepEqualPointerValues(a.Field(i), b.Field(i)) {
				return false
			}
		}
		return true

	default:
		return reflect.DeepEqual(a.Interface(), b.Interface())
	}
}

// NeedsUpdate compares non-zero fields in the desired state against the actual broker state.
// Pointers are properly dereferenced so two different pointers holding identical values
// are treated as equal.
func NeedsUpdate(broker, state interface{}) bool {
	if state == nil || broker == nil {
		return false
	}

	stateVal := reflect.ValueOf(state)
	brokerVal := reflect.ValueOf(broker)

	if stateVal.Kind() == reflect.Ptr {
		stateVal = stateVal.Elem()
	}
	if brokerVal.Kind() == reflect.Ptr {
		brokerVal = brokerVal.Elem()
	}

	if stateVal.Kind() != reflect.Struct || brokerVal.Kind() != reflect.Struct {
		return !DeepEqualPointerValues(stateVal, brokerVal)
	}

	typeOf := stateVal.Type()
	for i := 0; i < stateVal.NumField(); i++ {
		stateField := stateVal.Field(i)
		fieldName := typeOf.Field(i).Name

		// If field is not specified in desired state, don't trigger update
		if stateField.IsZero() {
			continue
		}

		brokerField := brokerVal.FieldByName(fieldName)
		if !brokerField.IsValid() {
			continue
		}

		if !DeepEqualPointerValues(stateField, brokerField) {
			logrus.WithFields(logrus.Fields{
				"field": fieldName,
				"state": stateField.Interface(),
			}).Debug("Field difference detected between state and broker")
			return true
		}
	}
	return false
}

