// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MatchRecords struct {
	_tab flatbuffers.Table
}

func GetRootAsMatchRecords(buf []byte, offset flatbuffers.UOffsetT) *MatchRecords {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MatchRecords{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMatchRecords(buf []byte, offset flatbuffers.UOffsetT) *MatchRecords {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MatchRecords{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MatchRecords) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MatchRecords) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MatchRecords) FieldName(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *MatchRecords) FieldNameLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MatchRecords) RecordColumns(obj *FieldValueColumnTable, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *MatchRecords) RecordColumnsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MatchRecords) DocCount() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MatchRecords) MutateDocCount(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func (rcv *MatchRecords) Tracers(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *MatchRecords) TracersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MatchRecords) TableName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MatchRecordsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func MatchRecordsAddFieldName(builder *flatbuffers.Builder, fieldName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(fieldName), 0)
}
func MatchRecordsStartFieldNameVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MatchRecordsAddRecordColumns(builder *flatbuffers.Builder, recordColumns flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(recordColumns), 0)
}
func MatchRecordsStartRecordColumnsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MatchRecordsAddDocCount(builder *flatbuffers.Builder, docCount uint64) {
	builder.PrependUint64Slot(2, docCount, 0)
}
func MatchRecordsAddTracers(builder *flatbuffers.Builder, tracers flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(tracers), 0)
}
func MatchRecordsStartTracersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MatchRecordsAddTableName(builder *flatbuffers.Builder, tableName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(tableName), 0)
}
func MatchRecordsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
