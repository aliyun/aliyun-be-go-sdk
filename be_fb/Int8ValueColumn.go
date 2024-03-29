// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Int8ValueColumn struct {
	_tab flatbuffers.Table
}

func GetRootAsInt8ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *Int8ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Int8ValueColumn{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInt8ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *Int8ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Int8ValueColumn{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Int8ValueColumn) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Int8ValueColumn) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Int8ValueColumn) Value(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Int8ValueColumn) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Int8ValueColumn) MutateValue(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func Int8ValueColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Int8ValueColumnAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func Int8ValueColumnStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func Int8ValueColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
