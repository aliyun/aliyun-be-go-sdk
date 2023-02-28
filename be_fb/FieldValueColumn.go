// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import "strconv"

type FieldValueColumn byte

const (
	FieldValueColumnNONE                   FieldValueColumn = 0
	FieldValueColumnInt8ValueColumn        FieldValueColumn = 1
	FieldValueColumnUInt8ValueColumn       FieldValueColumn = 2
	FieldValueColumnInt16ValueColumn       FieldValueColumn = 3
	FieldValueColumnUInt16ValueColumn      FieldValueColumn = 4
	FieldValueColumnInt32ValueColumn       FieldValueColumn = 5
	FieldValueColumnUInt32ValueColumn      FieldValueColumn = 6
	FieldValueColumnInt64ValueColumn       FieldValueColumn = 7
	FieldValueColumnUInt64ValueColumn      FieldValueColumn = 8
	FieldValueColumnFloatValueColumn       FieldValueColumn = 9
	FieldValueColumnDoubleValueColumn      FieldValueColumn = 10
	FieldValueColumnStringValueColumn      FieldValueColumn = 11
	FieldValueColumnMultiInt8ValueColumn   FieldValueColumn = 12
	FieldValueColumnMultiUInt8ValueColumn  FieldValueColumn = 13
	FieldValueColumnMultiInt16ValueColumn  FieldValueColumn = 14
	FieldValueColumnMultiUInt16ValueColumn FieldValueColumn = 15
	FieldValueColumnMultiInt32ValueColumn  FieldValueColumn = 16
	FieldValueColumnMultiUInt32ValueColumn FieldValueColumn = 17
	FieldValueColumnMultiInt64ValueColumn  FieldValueColumn = 18
	FieldValueColumnMultiUInt64ValueColumn FieldValueColumn = 19
	FieldValueColumnMultiFloatValueColumn  FieldValueColumn = 20
	FieldValueColumnMultiDoubleValueColumn FieldValueColumn = 21
	FieldValueColumnMultiStringValueColumn FieldValueColumn = 22
	FieldValueColumnUnknownValueColumn     FieldValueColumn = 23
)

var EnumNamesFieldValueColumn = map[FieldValueColumn]string{
	FieldValueColumnNONE:                   "NONE",
	FieldValueColumnInt8ValueColumn:        "Int8ValueColumn",
	FieldValueColumnUInt8ValueColumn:       "UInt8ValueColumn",
	FieldValueColumnInt16ValueColumn:       "Int16ValueColumn",
	FieldValueColumnUInt16ValueColumn:      "UInt16ValueColumn",
	FieldValueColumnInt32ValueColumn:       "Int32ValueColumn",
	FieldValueColumnUInt32ValueColumn:      "UInt32ValueColumn",
	FieldValueColumnInt64ValueColumn:       "Int64ValueColumn",
	FieldValueColumnUInt64ValueColumn:      "UInt64ValueColumn",
	FieldValueColumnFloatValueColumn:       "FloatValueColumn",
	FieldValueColumnDoubleValueColumn:      "DoubleValueColumn",
	FieldValueColumnStringValueColumn:      "StringValueColumn",
	FieldValueColumnMultiInt8ValueColumn:   "MultiInt8ValueColumn",
	FieldValueColumnMultiUInt8ValueColumn:  "MultiUInt8ValueColumn",
	FieldValueColumnMultiInt16ValueColumn:  "MultiInt16ValueColumn",
	FieldValueColumnMultiUInt16ValueColumn: "MultiUInt16ValueColumn",
	FieldValueColumnMultiInt32ValueColumn:  "MultiInt32ValueColumn",
	FieldValueColumnMultiUInt32ValueColumn: "MultiUInt32ValueColumn",
	FieldValueColumnMultiInt64ValueColumn:  "MultiInt64ValueColumn",
	FieldValueColumnMultiUInt64ValueColumn: "MultiUInt64ValueColumn",
	FieldValueColumnMultiFloatValueColumn:  "MultiFloatValueColumn",
	FieldValueColumnMultiDoubleValueColumn: "MultiDoubleValueColumn",
	FieldValueColumnMultiStringValueColumn: "MultiStringValueColumn",
	FieldValueColumnUnknownValueColumn:     "UnknownValueColumn",
}

var EnumValuesFieldValueColumn = map[string]FieldValueColumn{
	"NONE":                   FieldValueColumnNONE,
	"Int8ValueColumn":        FieldValueColumnInt8ValueColumn,
	"UInt8ValueColumn":       FieldValueColumnUInt8ValueColumn,
	"Int16ValueColumn":       FieldValueColumnInt16ValueColumn,
	"UInt16ValueColumn":      FieldValueColumnUInt16ValueColumn,
	"Int32ValueColumn":       FieldValueColumnInt32ValueColumn,
	"UInt32ValueColumn":      FieldValueColumnUInt32ValueColumn,
	"Int64ValueColumn":       FieldValueColumnInt64ValueColumn,
	"UInt64ValueColumn":      FieldValueColumnUInt64ValueColumn,
	"FloatValueColumn":       FieldValueColumnFloatValueColumn,
	"DoubleValueColumn":      FieldValueColumnDoubleValueColumn,
	"StringValueColumn":      FieldValueColumnStringValueColumn,
	"MultiInt8ValueColumn":   FieldValueColumnMultiInt8ValueColumn,
	"MultiUInt8ValueColumn":  FieldValueColumnMultiUInt8ValueColumn,
	"MultiInt16ValueColumn":  FieldValueColumnMultiInt16ValueColumn,
	"MultiUInt16ValueColumn": FieldValueColumnMultiUInt16ValueColumn,
	"MultiInt32ValueColumn":  FieldValueColumnMultiInt32ValueColumn,
	"MultiUInt32ValueColumn": FieldValueColumnMultiUInt32ValueColumn,
	"MultiInt64ValueColumn":  FieldValueColumnMultiInt64ValueColumn,
	"MultiUInt64ValueColumn": FieldValueColumnMultiUInt64ValueColumn,
	"MultiFloatValueColumn":  FieldValueColumnMultiFloatValueColumn,
	"MultiDoubleValueColumn": FieldValueColumnMultiDoubleValueColumn,
	"MultiStringValueColumn": FieldValueColumnMultiStringValueColumn,
	"UnknownValueColumn":     FieldValueColumnUnknownValueColumn,
}

func (v FieldValueColumn) String() string {
	if s, ok := EnumNamesFieldValueColumn[v]; ok {
		return s
	}
	return "FieldValueColumn(" + strconv.FormatInt(int64(v), 10) + ")"
}