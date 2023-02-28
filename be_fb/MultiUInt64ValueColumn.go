// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MultiUInt64ValueColumn struct {
	_tab flatbuffers.Table
}

func GetRootAsMultiUInt64ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *MultiUInt64ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MultiUInt64ValueColumn{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMultiUInt64ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *MultiUInt64ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MultiUInt64ValueColumn{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MultiUInt64ValueColumn) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MultiUInt64ValueColumn) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MultiUInt64ValueColumn) Value(obj *MultiUInt64Value, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *MultiUInt64ValueColumn) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MultiUInt64ValueColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MultiUInt64ValueColumnAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func MultiUInt64ValueColumnStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MultiUInt64ValueColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}