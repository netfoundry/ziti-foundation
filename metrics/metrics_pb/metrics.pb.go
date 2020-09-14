// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics.proto

package metrics_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ContentType int32

const (
	ContentType_Zero        ContentType = 0
	ContentType_MetricsType ContentType = 1007
)

var ContentType_name = map[int32]string{
	0:    "Zero",
	1007: "MetricsType",
}

var ContentType_value = map[string]int32{
	"Zero":        0,
	"MetricsType": 1007,
}

func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}

func (ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}

type MetricsMessage struct {
	SourceId             string                                     `protobuf:"bytes,2,opt,name=sourceId,proto3" json:"sourceId,omitempty"`
	Timestamp            *timestamp.Timestamp                       `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Tags                 map[string]string                          `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IntValues            map[string]int64                           `protobuf:"bytes,5,rep,name=intValues,proto3" json:"intValues,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	FloatValues          map[string]float64                         `protobuf:"bytes,6,rep,name=floatValues,proto3" json:"floatValues,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Meters               map[string]*MetricsMessage_Meter           `protobuf:"bytes,7,rep,name=meters,proto3" json:"meters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Histograms           map[string]*MetricsMessage_Histogram       `protobuf:"bytes,8,rep,name=histograms,proto3" json:"histograms,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IntervalCounters     map[string]*MetricsMessage_IntervalCounter `protobuf:"bytes,9,rep,name=intervalCounters,proto3" json:"intervalCounters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Timers               map[string]*MetricsMessage_Timer           `protobuf:"bytes,10,rep,name=timers,proto3" json:"timers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *MetricsMessage) Reset()         { *m = MetricsMessage{} }
func (m *MetricsMessage) String() string { return proto.CompactTextString(m) }
func (*MetricsMessage) ProtoMessage()    {}
func (*MetricsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}

func (m *MetricsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsMessage.Unmarshal(m, b)
}
func (m *MetricsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsMessage.Marshal(b, m, deterministic)
}
func (m *MetricsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsMessage.Merge(m, src)
}
func (m *MetricsMessage) XXX_Size() int {
	return xxx_messageInfo_MetricsMessage.Size(m)
}
func (m *MetricsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsMessage proto.InternalMessageInfo

func (m *MetricsMessage) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *MetricsMessage) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *MetricsMessage) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MetricsMessage) GetIntValues() map[string]int64 {
	if m != nil {
		return m.IntValues
	}
	return nil
}

func (m *MetricsMessage) GetFloatValues() map[string]float64 {
	if m != nil {
		return m.FloatValues
	}
	return nil
}

func (m *MetricsMessage) GetMeters() map[string]*MetricsMessage_Meter {
	if m != nil {
		return m.Meters
	}
	return nil
}

func (m *MetricsMessage) GetHistograms() map[string]*MetricsMessage_Histogram {
	if m != nil {
		return m.Histograms
	}
	return nil
}

func (m *MetricsMessage) GetIntervalCounters() map[string]*MetricsMessage_IntervalCounter {
	if m != nil {
		return m.IntervalCounters
	}
	return nil
}

func (m *MetricsMessage) GetTimers() map[string]*MetricsMessage_Timer {
	if m != nil {
		return m.Timers
	}
	return nil
}

type MetricsMessage_Meter struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	M1Rate               float64  `protobuf:"fixed64,2,opt,name=m1Rate,proto3" json:"m1Rate,omitempty"`
	M5Rate               float64  `protobuf:"fixed64,3,opt,name=m5Rate,proto3" json:"m5Rate,omitempty"`
	M15Rate              float64  `protobuf:"fixed64,4,opt,name=m15Rate,proto3" json:"m15Rate,omitempty"`
	MeanRate             float64  `protobuf:"fixed64,5,opt,name=meanRate,proto3" json:"meanRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsMessage_Meter) Reset()         { *m = MetricsMessage_Meter{} }
func (m *MetricsMessage_Meter) String() string { return proto.CompactTextString(m) }
func (*MetricsMessage_Meter) ProtoMessage()    {}
func (*MetricsMessage_Meter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0, 7}
}

func (m *MetricsMessage_Meter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsMessage_Meter.Unmarshal(m, b)
}
func (m *MetricsMessage_Meter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsMessage_Meter.Marshal(b, m, deterministic)
}
func (m *MetricsMessage_Meter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsMessage_Meter.Merge(m, src)
}
func (m *MetricsMessage_Meter) XXX_Size() int {
	return xxx_messageInfo_MetricsMessage_Meter.Size(m)
}
func (m *MetricsMessage_Meter) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsMessage_Meter.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsMessage_Meter proto.InternalMessageInfo

func (m *MetricsMessage_Meter) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MetricsMessage_Meter) GetM1Rate() float64 {
	if m != nil {
		return m.M1Rate
	}
	return 0
}

func (m *MetricsMessage_Meter) GetM5Rate() float64 {
	if m != nil {
		return m.M5Rate
	}
	return 0
}

func (m *MetricsMessage_Meter) GetM15Rate() float64 {
	if m != nil {
		return m.M15Rate
	}
	return 0
}

func (m *MetricsMessage_Meter) GetMeanRate() float64 {
	if m != nil {
		return m.MeanRate
	}
	return 0
}

type MetricsMessage_Histogram struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Max                  int64    `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	Mean                 float64  `protobuf:"fixed64,3,opt,name=mean,proto3" json:"mean,omitempty"`
	Min                  int64    `protobuf:"varint,4,opt,name=min,proto3" json:"min,omitempty"`
	StdDev               float64  `protobuf:"fixed64,5,opt,name=stdDev,proto3" json:"stdDev,omitempty"`
	Variance             float64  `protobuf:"fixed64,6,opt,name=variance,proto3" json:"variance,omitempty"`
	P50                  float64  `protobuf:"fixed64,7,opt,name=p50,proto3" json:"p50,omitempty"`
	P75                  float64  `protobuf:"fixed64,8,opt,name=p75,proto3" json:"p75,omitempty"`
	P95                  float64  `protobuf:"fixed64,9,opt,name=p95,proto3" json:"p95,omitempty"`
	P99                  float64  `protobuf:"fixed64,10,opt,name=p99,proto3" json:"p99,omitempty"`
	P999                 float64  `protobuf:"fixed64,11,opt,name=p999,proto3" json:"p999,omitempty"`
	P9999                float64  `protobuf:"fixed64,12,opt,name=p9999,proto3" json:"p9999,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsMessage_Histogram) Reset()         { *m = MetricsMessage_Histogram{} }
func (m *MetricsMessage_Histogram) String() string { return proto.CompactTextString(m) }
func (*MetricsMessage_Histogram) ProtoMessage()    {}
func (*MetricsMessage_Histogram) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0, 8}
}

func (m *MetricsMessage_Histogram) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsMessage_Histogram.Unmarshal(m, b)
}
func (m *MetricsMessage_Histogram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsMessage_Histogram.Marshal(b, m, deterministic)
}
func (m *MetricsMessage_Histogram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsMessage_Histogram.Merge(m, src)
}
func (m *MetricsMessage_Histogram) XXX_Size() int {
	return xxx_messageInfo_MetricsMessage_Histogram.Size(m)
}
func (m *MetricsMessage_Histogram) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsMessage_Histogram.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsMessage_Histogram proto.InternalMessageInfo

func (m *MetricsMessage_Histogram) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetMean() float64 {
	if m != nil {
		return m.Mean
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetMin() int64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetStdDev() float64 {
	if m != nil {
		return m.StdDev
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetVariance() float64 {
	if m != nil {
		return m.Variance
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetP50() float64 {
	if m != nil {
		return m.P50
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetP75() float64 {
	if m != nil {
		return m.P75
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetP95() float64 {
	if m != nil {
		return m.P95
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetP99() float64 {
	if m != nil {
		return m.P99
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetP999() float64 {
	if m != nil {
		return m.P999
	}
	return 0
}

func (m *MetricsMessage_Histogram) GetP9999() float64 {
	if m != nil {
		return m.P9999
	}
	return 0
}

type MetricsMessage_Timer struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Max                  int64    `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	Mean                 float64  `protobuf:"fixed64,3,opt,name=mean,proto3" json:"mean,omitempty"`
	Min                  int64    `protobuf:"varint,4,opt,name=min,proto3" json:"min,omitempty"`
	StdDev               float64  `protobuf:"fixed64,5,opt,name=stdDev,proto3" json:"stdDev,omitempty"`
	Variance             float64  `protobuf:"fixed64,6,opt,name=variance,proto3" json:"variance,omitempty"`
	P50                  float64  `protobuf:"fixed64,7,opt,name=p50,proto3" json:"p50,omitempty"`
	P75                  float64  `protobuf:"fixed64,8,opt,name=p75,proto3" json:"p75,omitempty"`
	P95                  float64  `protobuf:"fixed64,9,opt,name=p95,proto3" json:"p95,omitempty"`
	P99                  float64  `protobuf:"fixed64,10,opt,name=p99,proto3" json:"p99,omitempty"`
	P999                 float64  `protobuf:"fixed64,11,opt,name=p999,proto3" json:"p999,omitempty"`
	P9999                float64  `protobuf:"fixed64,12,opt,name=p9999,proto3" json:"p9999,omitempty"`
	M1Rate               float64  `protobuf:"fixed64,13,opt,name=m1Rate,proto3" json:"m1Rate,omitempty"`
	M5Rate               float64  `protobuf:"fixed64,14,opt,name=m5Rate,proto3" json:"m5Rate,omitempty"`
	M15Rate              float64  `protobuf:"fixed64,15,opt,name=m15Rate,proto3" json:"m15Rate,omitempty"`
	MeanRate             float64  `protobuf:"fixed64,16,opt,name=meanRate,proto3" json:"meanRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsMessage_Timer) Reset()         { *m = MetricsMessage_Timer{} }
func (m *MetricsMessage_Timer) String() string { return proto.CompactTextString(m) }
func (*MetricsMessage_Timer) ProtoMessage()    {}
func (*MetricsMessage_Timer) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0, 9}
}

func (m *MetricsMessage_Timer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsMessage_Timer.Unmarshal(m, b)
}
func (m *MetricsMessage_Timer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsMessage_Timer.Marshal(b, m, deterministic)
}
func (m *MetricsMessage_Timer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsMessage_Timer.Merge(m, src)
}
func (m *MetricsMessage_Timer) XXX_Size() int {
	return xxx_messageInfo_MetricsMessage_Timer.Size(m)
}
func (m *MetricsMessage_Timer) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsMessage_Timer.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsMessage_Timer proto.InternalMessageInfo

func (m *MetricsMessage_Timer) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MetricsMessage_Timer) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *MetricsMessage_Timer) GetMean() float64 {
	if m != nil {
		return m.Mean
	}
	return 0
}

func (m *MetricsMessage_Timer) GetMin() int64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *MetricsMessage_Timer) GetStdDev() float64 {
	if m != nil {
		return m.StdDev
	}
	return 0
}

func (m *MetricsMessage_Timer) GetVariance() float64 {
	if m != nil {
		return m.Variance
	}
	return 0
}

func (m *MetricsMessage_Timer) GetP50() float64 {
	if m != nil {
		return m.P50
	}
	return 0
}

func (m *MetricsMessage_Timer) GetP75() float64 {
	if m != nil {
		return m.P75
	}
	return 0
}

func (m *MetricsMessage_Timer) GetP95() float64 {
	if m != nil {
		return m.P95
	}
	return 0
}

func (m *MetricsMessage_Timer) GetP99() float64 {
	if m != nil {
		return m.P99
	}
	return 0
}

func (m *MetricsMessage_Timer) GetP999() float64 {
	if m != nil {
		return m.P999
	}
	return 0
}

func (m *MetricsMessage_Timer) GetP9999() float64 {
	if m != nil {
		return m.P9999
	}
	return 0
}

func (m *MetricsMessage_Timer) GetM1Rate() float64 {
	if m != nil {
		return m.M1Rate
	}
	return 0
}

func (m *MetricsMessage_Timer) GetM5Rate() float64 {
	if m != nil {
		return m.M5Rate
	}
	return 0
}

func (m *MetricsMessage_Timer) GetM15Rate() float64 {
	if m != nil {
		return m.M15Rate
	}
	return 0
}

func (m *MetricsMessage_Timer) GetMeanRate() float64 {
	if m != nil {
		return m.MeanRate
	}
	return 0
}

type MetricsMessage_IntervalCounter struct {
	IntervalLength       uint64                           `protobuf:"varint,1,opt,name=intervalLength,proto3" json:"intervalLength,omitempty"`
	Buckets              []*MetricsMessage_IntervalBucket `protobuf:"bytes,2,rep,name=buckets,proto3" json:"buckets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *MetricsMessage_IntervalCounter) Reset()         { *m = MetricsMessage_IntervalCounter{} }
func (m *MetricsMessage_IntervalCounter) String() string { return proto.CompactTextString(m) }
func (*MetricsMessage_IntervalCounter) ProtoMessage()    {}
func (*MetricsMessage_IntervalCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0, 10}
}

func (m *MetricsMessage_IntervalCounter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsMessage_IntervalCounter.Unmarshal(m, b)
}
func (m *MetricsMessage_IntervalCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsMessage_IntervalCounter.Marshal(b, m, deterministic)
}
func (m *MetricsMessage_IntervalCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsMessage_IntervalCounter.Merge(m, src)
}
func (m *MetricsMessage_IntervalCounter) XXX_Size() int {
	return xxx_messageInfo_MetricsMessage_IntervalCounter.Size(m)
}
func (m *MetricsMessage_IntervalCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsMessage_IntervalCounter.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsMessage_IntervalCounter proto.InternalMessageInfo

func (m *MetricsMessage_IntervalCounter) GetIntervalLength() uint64 {
	if m != nil {
		return m.IntervalLength
	}
	return 0
}

func (m *MetricsMessage_IntervalCounter) GetBuckets() []*MetricsMessage_IntervalBucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

type MetricsMessage_IntervalBucket struct {
	IntervalStartUTC     int64             `protobuf:"varint,1,opt,name=intervalStartUTC,proto3" json:"intervalStartUTC,omitempty"`
	Values               map[string]uint64 `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetricsMessage_IntervalBucket) Reset()         { *m = MetricsMessage_IntervalBucket{} }
func (m *MetricsMessage_IntervalBucket) String() string { return proto.CompactTextString(m) }
func (*MetricsMessage_IntervalBucket) ProtoMessage()    {}
func (*MetricsMessage_IntervalBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0, 11}
}

func (m *MetricsMessage_IntervalBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsMessage_IntervalBucket.Unmarshal(m, b)
}
func (m *MetricsMessage_IntervalBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsMessage_IntervalBucket.Marshal(b, m, deterministic)
}
func (m *MetricsMessage_IntervalBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsMessage_IntervalBucket.Merge(m, src)
}
func (m *MetricsMessage_IntervalBucket) XXX_Size() int {
	return xxx_messageInfo_MetricsMessage_IntervalBucket.Size(m)
}
func (m *MetricsMessage_IntervalBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsMessage_IntervalBucket.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsMessage_IntervalBucket proto.InternalMessageInfo

func (m *MetricsMessage_IntervalBucket) GetIntervalStartUTC() int64 {
	if m != nil {
		return m.IntervalStartUTC
	}
	return 0
}

func (m *MetricsMessage_IntervalBucket) GetValues() map[string]uint64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterEnum("metrics.pb.ContentType", ContentType_name, ContentType_value)
	proto.RegisterType((*MetricsMessage)(nil), "metrics.pb.MetricsMessage")
	proto.RegisterMapType((map[string]float64)(nil), "metrics.pb.MetricsMessage.FloatValuesEntry")
	proto.RegisterMapType((map[string]*MetricsMessage_Histogram)(nil), "metrics.pb.MetricsMessage.HistogramsEntry")
	proto.RegisterMapType((map[string]int64)(nil), "metrics.pb.MetricsMessage.IntValuesEntry")
	proto.RegisterMapType((map[string]*MetricsMessage_IntervalCounter)(nil), "metrics.pb.MetricsMessage.IntervalCountersEntry")
	proto.RegisterMapType((map[string]*MetricsMessage_Meter)(nil), "metrics.pb.MetricsMessage.MetersEntry")
	proto.RegisterMapType((map[string]string)(nil), "metrics.pb.MetricsMessage.TagsEntry")
	proto.RegisterMapType((map[string]*MetricsMessage_Timer)(nil), "metrics.pb.MetricsMessage.TimersEntry")
	proto.RegisterType((*MetricsMessage_Meter)(nil), "metrics.pb.MetricsMessage.Meter")
	proto.RegisterType((*MetricsMessage_Histogram)(nil), "metrics.pb.MetricsMessage.Histogram")
	proto.RegisterType((*MetricsMessage_Timer)(nil), "metrics.pb.MetricsMessage.Timer")
	proto.RegisterType((*MetricsMessage_IntervalCounter)(nil), "metrics.pb.MetricsMessage.IntervalCounter")
	proto.RegisterType((*MetricsMessage_IntervalBucket)(nil), "metrics.pb.MetricsMessage.IntervalBucket")
	proto.RegisterMapType((map[string]uint64)(nil), "metrics.pb.MetricsMessage.IntervalBucket.ValuesEntry")
}

func init() {
	proto.RegisterFile("metrics.proto", fileDescriptor_6039342a2ba47b72)
}

var fileDescriptor_6039342a2ba47b72 = []byte{
	// 754 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe4, 0x56, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x7e, 0xd3, 0x38, 0x49, 0x33, 0x7e, 0x9b, 0xe6, 0x5d, 0xbd, 0x20, 0xcb, 0x17, 0x2a, 0x54,
	0x55, 0x34, 0x48, 0xee, 0x07, 0x0a, 0xad, 0x2b, 0x54, 0x21, 0xc2, 0x57, 0x11, 0xb9, 0x84, 0xc0,
	0x01, 0xb8, 0x38, 0xe9, 0x36, 0xb5, 0x1a, 0xdb, 0x91, 0xbd, 0x89, 0xe8, 0x85, 0x0b, 0xff, 0x8d,
	0x9f, 0x81, 0xc4, 0x2f, 0xe0, 0x27, 0xc0, 0xee, 0xec, 0xda, 0xd9, 0xa4, 0xf9, 0x42, 0x1c, 0x39,
	0x65, 0xe7, 0xe3, 0x79, 0x66, 0x76, 0x66, 0x3c, 0x1b, 0xd8, 0x08, 0x28, 0x8b, 0xfd, 0x6e, 0xe2,
	0x0c, 0xe2, 0x88, 0x45, 0x04, 0x32, 0xb1, 0x63, 0xdf, 0xe9, 0x45, 0x51, 0xaf, 0x4f, 0xf7, 0xd0,
	0xd2, 0x19, 0x5e, 0xec, 0x31, 0x3f, 0xa0, 0x09, 0xf3, 0x82, 0x81, 0x74, 0xbe, 0xfb, 0xed, 0x3f,
	0xa8, 0x34, 0xa5, 0x7f, 0x93, 0x26, 0x89, 0xd7, 0xa3, 0xc4, 0x86, 0xf5, 0x24, 0x1a, 0xc6, 0x5d,
	0x7a, 0x76, 0x6e, 0xad, 0x6d, 0xe5, 0xee, 0x95, 0x5b, 0x99, 0x4c, 0x8e, 0xa1, 0x9c, 0x31, 0x58,
	0x79, 0x6e, 0x34, 0x0f, 0x6d, 0x47, 0xc6, 0x70, 0xd2, 0x18, 0x4e, 0x3b, 0xf5, 0x68, 0x8d, 0x9d,
	0x39, 0xd2, 0x60, 0x5e, 0x2f, 0xb1, 0x8c, 0xad, 0x3c, 0x07, 0x6d, 0x3b, 0xe3, 0x24, 0x9d, 0xc9,
	0xf8, 0x4e, 0x9b, 0xbb, 0x3d, 0x0b, 0x59, 0x7c, 0xdd, 0x42, 0x04, 0x79, 0x01, 0x65, 0x3f, 0x64,
	0xef, 0xbc, 0xfe, 0x90, 0x26, 0x56, 0x01, 0xe1, 0xbb, 0x0b, 0xe0, 0x67, 0xa9, 0xaf, 0xe4, 0x18,
	0x63, 0x49, 0x13, 0xcc, 0x8b, 0x7e, 0xe4, 0xa5, 0x54, 0x45, 0xa4, 0xba, 0xbf, 0x80, 0xea, 0xf9,
	0xd8, 0x5b, 0x92, 0xe9, 0x78, 0x72, 0x0a, 0x45, 0x0e, 0xa5, 0x71, 0x62, 0x95, 0x90, 0x69, 0x67,
	0x01, 0x53, 0x13, 0x1d, 0x25, 0x89, 0x42, 0x91, 0x57, 0x00, 0x97, 0x7e, 0xc2, 0xa2, 0x5e, 0xec,
	0x05, 0x89, 0xb5, 0x8e, 0x1c, 0xb5, 0x05, 0x1c, 0x2f, 0x33, 0x67, 0xc9, 0xa3, 0xa1, 0xc9, 0x47,
	0xa8, 0xf2, 0x7b, 0xd2, 0x78, 0xe4, 0xf5, 0x1b, 0xd1, 0x30, 0xc4, 0xac, 0xca, 0xc8, 0xb8, 0xbf,
	0xb8, 0x54, 0x13, 0x10, 0xc9, 0x7b, 0x83, 0x49, 0xdc, 0x54, 0x34, 0x92, 0x73, 0xc2, 0xd2, 0x9b,
	0xb6, 0xd1, 0x51, 0xdd, 0x54, 0xa2, 0xec, 0x23, 0x28, 0x67, 0x4d, 0x25, 0x55, 0xc8, 0x5f, 0xd1,
	0x6b, 0x2b, 0x87, 0x93, 0x25, 0x8e, 0xe4, 0x7f, 0x28, 0x8c, 0x44, 0x49, 0xd5, 0xb4, 0x49, 0xe1,
	0x64, 0xed, 0x38, 0x67, 0x3f, 0x82, 0xca, 0x64, 0x3b, 0x97, 0xa1, 0xf3, 0x3a, 0xfa, 0x14, 0xaa,
	0xd3, 0x1d, 0x5c, 0x86, 0xcf, 0xe9, 0xf8, 0x0f, 0x60, 0x6a, 0x7d, 0x9b, 0x01, 0x7d, 0xa8, 0x43,
	0xcd, 0xc3, 0xad, 0x65, 0x03, 0xa0, 0x93, 0x77, 0x61, 0x73, 0xaa, 0xa1, 0x33, 0x02, 0x9c, 0x4c,
	0x06, 0xd8, 0x5e, 0x65, 0x3a, 0xf4, 0x20, 0x11, 0xdc, 0x9a, 0xd9, 0xe3, 0x19, 0xa1, 0x1e, 0x4f,
	0x86, 0xaa, 0xad, 0x3e, 0x36, 0x53, 0x25, 0xd3, 0x06, 0xe0, 0xcf, 0x4a, 0x86, 0x44, 0x3a, 0xf9,
	0x97, 0x1c, 0x14, 0xb0, 0x8e, 0xa2, 0x67, 0x5d, 0x11, 0x1c, 0x99, 0x79, 0xcf, 0x51, 0x20, 0xb7,
	0xf9, 0x07, 0x79, 0xd0, 0xf2, 0x58, 0xda, 0x4a, 0x25, 0xa1, 0xbe, 0x8e, 0xfa, 0xbc, 0xd2, 0xa3,
	0x44, 0x2c, 0x28, 0x05, 0x07, 0xd2, 0x60, 0xa0, 0x21, 0x15, 0xc5, 0x0a, 0x0c, 0xa8, 0x17, 0xa2,
	0xa9, 0x80, 0xa6, 0x4c, 0xb6, 0x7f, 0xe6, 0xa0, 0x9c, 0x15, 0x7b, 0x4e, 0x26, 0xfc, 0xde, 0x81,
	0xf7, 0x49, 0x4d, 0xa4, 0x38, 0x12, 0x02, 0x86, 0x60, 0x50, 0x19, 0xe0, 0x19, 0xbd, 0xfc, 0x10,
	0x63, 0x0b, 0x2f, 0x3f, 0x14, 0x99, 0x26, 0xec, 0xfc, 0x29, 0x1d, 0xa9, 0xa8, 0x4a, 0x12, 0xf9,
	0x8c, 0xbc, 0xd8, 0xf7, 0xc2, 0x2e, 0xe5, 0x6b, 0x0b, 0xf3, 0x49, 0x65, 0xc1, 0x32, 0xa8, 0xef,
	0xf3, 0x1d, 0x24, 0xd4, 0xe2, 0x88, 0x9a, 0xa3, 0x3a, 0xdf, 0x28, 0x52, 0x73, 0x54, 0x47, 0x8d,
	0x5b, 0xe7, 0x1b, 0x41, 0x6a, 0x5c, 0xa5, 0x71, 0xf9, 0xf7, 0xac, 0x34, 0xae, 0xc8, 0x90, 0xff,
	0xb8, 0x96, 0x29, 0x33, 0x14, 0x67, 0x71, 0x3b, 0xf1, 0xeb, 0x5a, 0xff, 0xca, 0x6f, 0x03, 0x05,
	0xfb, 0xfb, 0x1a, 0x14, 0xb0, 0x39, 0x7f, 0xe3, 0xed, 0xb5, 0x29, 0xdb, 0x98, 0x33, 0x65, 0x95,
	0x79, 0x53, 0xb6, 0x39, 0x7f, 0xca, 0xaa, 0x53, 0x53, 0xf6, 0x19, 0x36, 0xa7, 0x3e, 0x33, 0xb2,
	0x03, 0x95, 0x74, 0x33, 0xbf, 0xa6, 0x61, 0x8f, 0x5d, 0x62, 0xd5, 0x8d, 0xd6, 0x94, 0x96, 0x34,
	0xa0, 0xd4, 0x19, 0x76, 0xaf, 0x28, 0x4b, 0x78, 0x0b, 0x56, 0x78, 0x2d, 0x11, 0xfb, 0x04, 0x11,
	0xad, 0x14, 0x69, 0x7f, 0xcd, 0xe1, 0xea, 0xd5, 0x6c, 0xa4, 0x36, 0x7e, 0x63, 0xde, 0x30, 0x2f,
	0x66, 0x6f, 0xdb, 0x0d, 0xd5, 0xf7, 0x1b, 0x7a, 0xfe, 0xd4, 0x16, 0x47, 0xf2, 0x95, 0xcd, 0x63,
	0x0a, 0xf5, 0x95, 0x53, 0x70, 0xf4, 0xf7, 0x56, 0x91, 0xd8, 0x2e, 0x98, 0xbf, 0xb1, 0xc4, 0x0d,
	0x6d, 0x69, 0xd4, 0x76, 0xc1, 0x6c, 0x44, 0x3c, 0x42, 0xc8, 0xda, 0xd7, 0x03, 0x4a, 0xd6, 0xc1,
	0x78, 0x4f, 0xe3, 0xa8, 0xfa, 0x0f, 0x27, 0x31, 0x55, 0x22, 0xc2, 0x50, 0xfd, 0x51, 0xea, 0x14,
	0xf1, 0x1f, 0xcc, 0x83, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xca, 0x9d, 0xc1, 0x92, 0x50, 0x09,
	0x00, 0x00,
}
