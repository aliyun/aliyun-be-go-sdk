// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Int16ValueColumn struct {
	_tab flatbuffers.Table
}

func GetRootAsInt16ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *Int16ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Int16ValueColumn{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInt16ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *Int16ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Int16ValueColumn{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Int16ValueColumn) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Int16ValueColumn) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Int16ValueColumn) Value(j int) int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt16(a + flatbuffers.UOffsetT(j*2))
	}
	return 0
}

func (rcv *Int16ValueColumn) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Int16ValueColumn) MutateValue(j int, n int16) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt16(a+flatbuffers.UOffsetT(j*2), n)
	}
	return false
}

func Int16ValueColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Int16ValueColumnAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func Int16ValueColumnStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(2, numElems, 2)
}
func Int16ValueColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
