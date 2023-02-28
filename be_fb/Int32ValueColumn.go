// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Int32ValueColumn struct {
	_tab flatbuffers.Table
}

func GetRootAsInt32ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *Int32ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Int32ValueColumn{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInt32ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *Int32ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Int32ValueColumn{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Int32ValueColumn) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Int32ValueColumn) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Int32ValueColumn) Value(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *Int32ValueColumn) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Int32ValueColumn) MutateValue(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func Int32ValueColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Int32ValueColumnAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func Int32ValueColumnStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func Int32ValueColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}