// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UInt16ValueColumn struct {
	_tab flatbuffers.Table
}

func GetRootAsUInt16ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *UInt16ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UInt16ValueColumn{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUInt16ValueColumn(buf []byte, offset flatbuffers.UOffsetT) *UInt16ValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UInt16ValueColumn{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UInt16ValueColumn) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UInt16ValueColumn) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UInt16ValueColumn) Value(j int) uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint16(a + flatbuffers.UOffsetT(j*2))
	}
	return 0
}

func (rcv *UInt16ValueColumn) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *UInt16ValueColumn) MutateValue(j int, n uint16) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint16(a+flatbuffers.UOffsetT(j*2), n)
	}
	return false
}

func UInt16ValueColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UInt16ValueColumnAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func UInt16ValueColumnStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(2, numElems, 2)
}
func UInt16ValueColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}