// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: controller/services/application.proto

package services // import "github.com/argoproj/argo-cd/controller/services"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ResourcesQuery struct {
	ApplicationName      string   `protobuf:"bytes,1,opt,name=applicationName,proto3" json:"applicationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourcesQuery) Reset()         { *m = ResourcesQuery{} }
func (m *ResourcesQuery) String() string { return proto.CompactTextString(m) }
func (*ResourcesQuery) ProtoMessage()    {}
func (*ResourcesQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_112230d8cacba6aa, []int{0}
}
func (m *ResourcesQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourcesQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourcesQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResourcesQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcesQuery.Merge(dst, src)
}
func (m *ResourcesQuery) XXX_Size() int {
	return m.Size()
}
func (m *ResourcesQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcesQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcesQuery proto.InternalMessageInfo

func (m *ResourcesQuery) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

type ResourcesTreeResponse struct {
	Items                []*v1alpha1.ResourceNode `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ResourcesTreeResponse) Reset()         { *m = ResourcesTreeResponse{} }
func (m *ResourcesTreeResponse) String() string { return proto.CompactTextString(m) }
func (*ResourcesTreeResponse) ProtoMessage()    {}
func (*ResourcesTreeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_112230d8cacba6aa, []int{1}
}
func (m *ResourcesTreeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourcesTreeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourcesTreeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ResourcesTreeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcesTreeResponse.Merge(dst, src)
}
func (m *ResourcesTreeResponse) XXX_Size() int {
	return m.Size()
}
func (m *ResourcesTreeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcesTreeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcesTreeResponse proto.InternalMessageInfo

func (m *ResourcesTreeResponse) GetItems() []*v1alpha1.ResourceNode {
	if m != nil {
		return m.Items
	}
	return nil
}

type ControlledResourcesResponse struct {
	Items                []*v1alpha1.ResourceState `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ControlledResourcesResponse) Reset()         { *m = ControlledResourcesResponse{} }
func (m *ControlledResourcesResponse) String() string { return proto.CompactTextString(m) }
func (*ControlledResourcesResponse) ProtoMessage()    {}
func (*ControlledResourcesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_112230d8cacba6aa, []int{2}
}
func (m *ControlledResourcesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ControlledResourcesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ControlledResourcesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ControlledResourcesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlledResourcesResponse.Merge(dst, src)
}
func (m *ControlledResourcesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ControlledResourcesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlledResourcesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ControlledResourcesResponse proto.InternalMessageInfo

func (m *ControlledResourcesResponse) GetItems() []*v1alpha1.ResourceState {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourcesQuery)(nil), "github.com.argoproj.argo_cd.controller.services.ResourcesQuery")
	proto.RegisterType((*ResourcesTreeResponse)(nil), "github.com.argoproj.argo_cd.controller.services.ResourcesTreeResponse")
	proto.RegisterType((*ControlledResourcesResponse)(nil), "github.com.argoproj.argo_cd.controller.services.ControlledResourcesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApplicationService service

type ApplicationServiceClient interface {
	ResourcesTree(ctx context.Context, in *ResourcesQuery, opts ...grpc.CallOption) (*ResourcesTreeResponse, error)
	ControlledResources(ctx context.Context, in *ResourcesQuery, opts ...grpc.CallOption) (*ControlledResourcesResponse, error)
}

type applicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServiceClient(cc *grpc.ClientConn) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) ResourcesTree(ctx context.Context, in *ResourcesQuery, opts ...grpc.CallOption) (*ResourcesTreeResponse, error) {
	out := new(ResourcesTreeResponse)
	err := c.cc.Invoke(ctx, "/github.com.argoproj.argo_cd.controller.services.ApplicationService/ResourcesTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) ControlledResources(ctx context.Context, in *ResourcesQuery, opts ...grpc.CallOption) (*ControlledResourcesResponse, error) {
	out := new(ControlledResourcesResponse)
	err := c.cc.Invoke(ctx, "/github.com.argoproj.argo_cd.controller.services.ApplicationService/ControlledResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApplicationService service

type ApplicationServiceServer interface {
	ResourcesTree(context.Context, *ResourcesQuery) (*ResourcesTreeResponse, error)
	ControlledResources(context.Context, *ResourcesQuery) (*ControlledResourcesResponse, error)
}

func RegisterApplicationServiceServer(s *grpc.Server, srv ApplicationServiceServer) {
	s.RegisterService(&_ApplicationService_serviceDesc, srv)
}

func _ApplicationService_ResourcesTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourcesQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).ResourcesTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.argoproj.argo_cd.controller.services.ApplicationService/ResourcesTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).ResourcesTree(ctx, req.(*ResourcesQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_ControlledResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourcesQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).ControlledResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.argoproj.argo_cd.controller.services.ApplicationService/ControlledResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).ControlledResources(ctx, req.(*ResourcesQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.argoproj.argo_cd.controller.services.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResourcesTree",
			Handler:    _ApplicationService_ResourcesTree_Handler,
		},
		{
			MethodName: "ControlledResources",
			Handler:    _ApplicationService_ControlledResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/services/application.proto",
}

func (m *ResourcesQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourcesQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ApplicationName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApplication(dAtA, i, uint64(len(m.ApplicationName)))
		i += copy(dAtA[i:], m.ApplicationName)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ResourcesTreeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourcesTreeResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintApplication(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ControlledResourcesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ControlledResourcesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintApplication(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintApplication(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourcesQuery) Size() (n int) {
	var l int
	_ = l
	l = len(m.ApplicationName)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResourcesTreeResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovApplication(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ControlledResourcesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovApplication(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApplication(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApplication(x uint64) (n int) {
	return sovApplication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourcesQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourcesQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourcesQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResourcesTreeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourcesTreeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourcesTreeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &v1alpha1.ResourceNode{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ControlledResourcesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ControlledResourcesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ControlledResourcesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &v1alpha1.ResourceState{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApplication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplication
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApplication
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApplication
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthApplication
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApplication
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApplication(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApplication = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplication   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("controller/services/application.proto", fileDescriptor_application_112230d8cacba6aa)
}

var fileDescriptor_application_112230d8cacba6aa = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x6d, 0xbe, 0x0f, 0x05, 0x23, 0x2a, 0x44, 0x84, 0x52, 0xa1, 0x94, 0x01, 0xa1, 0x1b, 0x13,
	0x5a, 0x77, 0x82, 0x88, 0x0a, 0xfe, 0x80, 0x14, 0x9c, 0xba, 0x12, 0x54, 0xd2, 0xcc, 0x65, 0x1a,
	0x3b, 0x9d, 0x84, 0x24, 0x2d, 0xb8, 0xf0, 0x45, 0x7c, 0x00, 0x9f, 0xc5, 0xa5, 0x6b, 0x57, 0xd2,
	0x27, 0x11, 0x5b, 0xd2, 0xce, 0x48, 0x19, 0xa8, 0xba, 0xbb, 0x84, 0x9c, 0x9f, 0x9c, 0x73, 0x83,
	0x77, 0x84, 0x4a, 0x9d, 0x51, 0x49, 0x02, 0x86, 0x59, 0x30, 0x43, 0x29, 0xc0, 0x32, 0xae, 0x75,
	0x22, 0x05, 0x77, 0x52, 0xa5, 0x54, 0x1b, 0xe5, 0x14, 0x61, 0xb1, 0x74, 0xdd, 0x41, 0x87, 0x0a,
	0xd5, 0xa7, 0xdc, 0xc4, 0x4a, 0x1b, 0xf5, 0x30, 0x1e, 0xee, 0x45, 0x44, 0x67, 0x14, 0xd4, 0x53,
	0x54, 0x2e, 0x66, 0x00, 0xe6, 0x01, 0xe3, 0x61, 0x57, 0x44, 0x4c, 0xf7, 0x62, 0xc6, 0xb5, 0xcc,
	0x09, 0xb1, 0x61, 0x83, 0x27, 0xba, 0xcb, 0x1b, 0x2c, 0x86, 0x14, 0x0c, 0x77, 0x10, 0x4d, 0xb4,
	0x83, 0x7d, 0xbc, 0x1e, 0x82, 0x55, 0x03, 0x23, 0xc0, 0x5e, 0x0d, 0xc0, 0x3c, 0x92, 0x3a, 0xde,
	0xc8, 0x20, 0x5b, 0xbc, 0x0f, 0x65, 0x54, 0x43, 0xf5, 0x95, 0xf0, 0xfb, 0x71, 0x30, 0xc4, 0x5b,
	0x53, 0xec, 0xb5, 0x01, 0x08, 0xc1, 0x6a, 0x95, 0x5a, 0x20, 0xb7, 0x78, 0x49, 0x3a, 0xe8, 0xdb,
	0x32, 0xaa, 0xfd, 0xaf, 0xaf, 0x36, 0xcf, 0x68, 0xd1, 0x03, 0x75, 0x2f, 0xa6, 0x5f, 0x7e, 0x69,
	0x36, 0x18, 0xef, 0x97, 0x7a, 0x81, 0x96, 0x8a, 0x20, 0x9c, 0xb0, 0x06, 0x4f, 0x78, 0xfb, 0xc4,
	0xa7, 0x12, 0x4d, 0x1d, 0x4c, 0xd5, 0xef, 0xf2, 0xea, 0xe7, 0x7f, 0xa0, 0xde, 0x76, 0xdc, 0x79,
	0xf9, 0xe6, 0xfb, 0x3f, 0x4c, 0x8e, 0x66, 0xb7, 0xdb, 0x93, 0x56, 0xc8, 0x33, 0xc2, 0x6b, 0xb9,
	0x38, 0xc8, 0x21, 0x5d, 0xb0, 0x58, 0x9a, 0xaf, 0xa2, 0x72, 0xfa, 0x73, 0x82, 0x6c, 0x1f, 0x41,
	0x89, 0xbc, 0x20, 0xbc, 0x39, 0x27, 0xb3, 0xdf, 0x5b, 0xbc, 0x5c, 0x98, 0xa0, 0xa0, 0xba, 0xa0,
	0x74, 0x7c, 0xf0, 0x3a, 0xaa, 0xa2, 0xb7, 0x51, 0x15, 0x7d, 0x8c, 0xaa, 0xe8, 0x86, 0x15, 0x2d,
	0xfa, 0x9c, 0xcf, 0xd5, 0x59, 0x1e, 0x6f, 0xf5, 0xde, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x70,
	0xc5, 0x84, 0xa3, 0x7a, 0x03, 0x00, 0x00,
}