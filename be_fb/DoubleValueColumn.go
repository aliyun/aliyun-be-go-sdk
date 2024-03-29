// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package be_fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DoubleValueColumn struct {
	_tab flatbuffers.Table
}

func GetRootAsDoubleValueColumn(buf []byte, offset flatbuffers.UOffsetT) *DoubleValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DoubleValueColumn{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDoubleValueColumn(buf []byte, offset flatbuffers.UOffsetT) *DoubleValueColumn {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DoubleValueColumn{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DoubleValueColumn) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DoubleValueColumn) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DoubleValueColumn) Value(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *DoubleValueColumn) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DoubleValueColumn) MutateValue(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func DoubleValueColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DoubleValueColumnAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func DoubleValueColumnStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func DoubleValueColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
