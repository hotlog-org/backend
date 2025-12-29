package models

type FieldType string

const (
	FieldTypeString   FieldType = "string"
	FieldTypeNumber   FieldType = "number"
	FieldTypeBoolean  FieldType = "boolean"
	FieldTypeDatetime FieldType = "datetime"
	FieldTypeEnum     FieldType = "enum"
	FieldTypeArray    FieldType = "array"
	FieldTypeJSON     FieldType = "json"
	FieldTypeObject   FieldType = "object"
)

type ComponentType string
