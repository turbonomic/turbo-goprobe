// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ActionExecution.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	ActionExecution.proto
	CommonDTO.proto
	Discovery.proto
	IdentityMetadata.proto
	MediationMessage.proto
	NonMarketEntityDTO.proto
	ProfileDTO.proto
	SupplyChain.proto

It has these top-level messages:
	ActionPolicyDTO
	ActionExecutionDTO
	ActionItemDTO
	EntityDTO
	CommodityDTO
	VStoragePartitionData
	GroupDTO
	ContextData
	NotificationDTO
	CustomTypeDef
	ServerEntityPropDef
	PropertyHandler
	EntityIdentityData
	FlowDTO
	ValidationResponse
	DiscoveryResponse
	CustomAccountDefEntry
	AccountDefEntry
	ErrorDTO
	CustomMetadataDTO
	AccountValue
	DerivedTargetSpecificationDTO
	EntityIdentityMetadata
	MediationClientMessage
	MediationServerMessage
	ActionRequest
	ActionResult
	ActionProgress
	ActionResponse
	ContainerInfo
	KeepAlive
	Ack
	InitializationContent
	ValidationRequest
	DiscoveryRequest
	ProbeInfo
	SetProperties
	NonMarketEntityDTO
	EntityProfileDTO
	CommodityProfileDTO
	DeploymentProfileDTO
	TemplateDTO
	TemplateCommodity
	Provider
	ExternalEntityLink
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ActionResponseState int32

const (
	ActionResponseState_PENDING_ACCEPT ActionResponseState = 0
	ActionResponseState_ACCEPTED       ActionResponseState = 1
	ActionResponseState_REJECTED       ActionResponseState = 2
	ActionResponseState_IN_PROGRESS    ActionResponseState = 3
	ActionResponseState_SUCCEEDED      ActionResponseState = 4
	ActionResponseState_FAILED         ActionResponseState = 5
	ActionResponseState_RECOMMENDED    ActionResponseState = 6
	ActionResponseState_DISABLED       ActionResponseState = 7
	ActionResponseState_QUEUED         ActionResponseState = 8
)

var ActionResponseState_name = map[int32]string{
	0: "PENDING_ACCEPT",
	1: "ACCEPTED",
	2: "REJECTED",
	3: "IN_PROGRESS",
	4: "SUCCEEDED",
	5: "FAILED",
	6: "RECOMMENDED",
	7: "DISABLED",
	8: "QUEUED",
}
var ActionResponseState_value = map[string]int32{
	"PENDING_ACCEPT": 0,
	"ACCEPTED":       1,
	"REJECTED":       2,
	"IN_PROGRESS":    3,
	"SUCCEEDED":      4,
	"FAILED":         5,
	"RECOMMENDED":    6,
	"DISABLED":       7,
	"QUEUED":         8,
}

func (x ActionResponseState) Enum() *ActionResponseState {
	p := new(ActionResponseState)
	*p = x
	return p
}
func (x ActionResponseState) String() string {
	return proto1.EnumName(ActionResponseState_name, int32(x))
}
func (x *ActionResponseState) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(ActionResponseState_value, data, "ActionResponseState")
	if err != nil {
		return err
	}
	*x = ActionResponseState(value)
	return nil
}
func (ActionResponseState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Defines how action can be executed by the probe
type ActionPolicyDTO_ActionCapability int32

const (
	// The action does not make sense for the probe. Actions of this type should not be shown
	// to the user.
	ActionPolicyDTO_NOT_SUPPORTED ActionPolicyDTO_ActionCapability = 1
	// The probe does not support executing the action, but it may still make sense to display
	// actions of this type to the user. For example, VSTORAGE resize is not supported in VCenter,
	// but many users still wish to see actions of this type.
	ActionPolicyDTO_NOT_EXECUTABLE ActionPolicyDTO_ActionCapability = 2
	// The probe supports execution of the action. Actions of this type are displayed to the user.
	ActionPolicyDTO_SUPPORTED ActionPolicyDTO_ActionCapability = 3
)

var ActionPolicyDTO_ActionCapability_name = map[int32]string{
	1: "NOT_SUPPORTED",
	2: "NOT_EXECUTABLE",
	3: "SUPPORTED",
}
var ActionPolicyDTO_ActionCapability_value = map[string]int32{
	"NOT_SUPPORTED":  1,
	"NOT_EXECUTABLE": 2,
	"SUPPORTED":      3,
}

func (x ActionPolicyDTO_ActionCapability) Enum() *ActionPolicyDTO_ActionCapability {
	p := new(ActionPolicyDTO_ActionCapability)
	*p = x
	return p
}
func (x ActionPolicyDTO_ActionCapability) String() string {
	return proto1.EnumName(ActionPolicyDTO_ActionCapability_name, int32(x))
}
func (x *ActionPolicyDTO_ActionCapability) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(ActionPolicyDTO_ActionCapability_value, data, "ActionPolicyDTO_ActionCapability")
	if err != nil {
		return err
	}
	*x = ActionPolicyDTO_ActionCapability(value)
	return nil
}
func (ActionPolicyDTO_ActionCapability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type ActionItemDTO_ActionType int32

const (
	ActionItemDTO_NONE                   ActionItemDTO_ActionType = 0
	ActionItemDTO_START                  ActionItemDTO_ActionType = 1
	ActionItemDTO_MOVE                   ActionItemDTO_ActionType = 2
	ActionItemDTO_SUSPEND                ActionItemDTO_ActionType = 3
	ActionItemDTO_TERMINATE              ActionItemDTO_ActionType = 4
	ActionItemDTO_SPAWN                  ActionItemDTO_ActionType = 5
	ActionItemDTO_ADD_PROVIDER           ActionItemDTO_ActionType = 6
	ActionItemDTO_CHANGE                 ActionItemDTO_ActionType = 7
	ActionItemDTO_REMOVE_PROVIDER        ActionItemDTO_ActionType = 8
	ActionItemDTO_PROVISION              ActionItemDTO_ActionType = 9
	ActionItemDTO_RECONFIGURE            ActionItemDTO_ActionType = 10
	ActionItemDTO_RESIZE                 ActionItemDTO_ActionType = 11
	ActionItemDTO_RESIZE_CAPACITY        ActionItemDTO_ActionType = 12
	ActionItemDTO_WARN                   ActionItemDTO_ActionType = 13
	ActionItemDTO_RECONFIGURE_THRESHOLD  ActionItemDTO_ActionType = 14
	ActionItemDTO_DELETE                 ActionItemDTO_ActionType = 15
	ActionItemDTO_RIGHT_SIZE             ActionItemDTO_ActionType = 16
	ActionItemDTO_RESERVE_ON_PM          ActionItemDTO_ActionType = 17
	ActionItemDTO_RESERVE_ON_DS          ActionItemDTO_ActionType = 18
	ActionItemDTO_RESIZE_FOR_EFFICIENCY  ActionItemDTO_ActionType = 19
	ActionItemDTO_RESIZE_FOR_PERFORMANCE ActionItemDTO_ActionType = 20
	ActionItemDTO_CROSS_TARGET_MOVE      ActionItemDTO_ActionType = 21
	ActionItemDTO_MOVE_TOGETHER          ActionItemDTO_ActionType = 22
	ActionItemDTO_RESERVE_ON_DA          ActionItemDTO_ActionType = 23
)

var ActionItemDTO_ActionType_name = map[int32]string{
	0:  "NONE",
	1:  "START",
	2:  "MOVE",
	3:  "SUSPEND",
	4:  "TERMINATE",
	5:  "SPAWN",
	6:  "ADD_PROVIDER",
	7:  "CHANGE",
	8:  "REMOVE_PROVIDER",
	9:  "PROVISION",
	10: "RECONFIGURE",
	11: "RESIZE",
	12: "RESIZE_CAPACITY",
	13: "WARN",
	14: "RECONFIGURE_THRESHOLD",
	15: "DELETE",
	16: "RIGHT_SIZE",
	17: "RESERVE_ON_PM",
	18: "RESERVE_ON_DS",
	19: "RESIZE_FOR_EFFICIENCY",
	20: "RESIZE_FOR_PERFORMANCE",
	21: "CROSS_TARGET_MOVE",
	22: "MOVE_TOGETHER",
	23: "RESERVE_ON_DA",
}
var ActionItemDTO_ActionType_value = map[string]int32{
	"NONE":                   0,
	"START":                  1,
	"MOVE":                   2,
	"SUSPEND":                3,
	"TERMINATE":              4,
	"SPAWN":                  5,
	"ADD_PROVIDER":           6,
	"CHANGE":                 7,
	"REMOVE_PROVIDER":        8,
	"PROVISION":              9,
	"RECONFIGURE":            10,
	"RESIZE":                 11,
	"RESIZE_CAPACITY":        12,
	"WARN":                   13,
	"RECONFIGURE_THRESHOLD":  14,
	"DELETE":                 15,
	"RIGHT_SIZE":             16,
	"RESERVE_ON_PM":          17,
	"RESERVE_ON_DS":          18,
	"RESIZE_FOR_EFFICIENCY":  19,
	"RESIZE_FOR_PERFORMANCE": 20,
	"CROSS_TARGET_MOVE":      21,
	"MOVE_TOGETHER":          22,
	"RESERVE_ON_DA":          23,
}

func (x ActionItemDTO_ActionType) Enum() *ActionItemDTO_ActionType {
	p := new(ActionItemDTO_ActionType)
	*p = x
	return p
}
func (x ActionItemDTO_ActionType) String() string {
	return proto1.EnumName(ActionItemDTO_ActionType_name, int32(x))
}
func (x *ActionItemDTO_ActionType) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(ActionItemDTO_ActionType_value, data, "ActionItemDTO_ActionType")
	if err != nil {
		return err
	}
	*x = ActionItemDTO_ActionType(value)
	return nil
}
func (ActionItemDTO_ActionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type ActionItemDTO_CommodityAttribute int32

const (
	ActionItemDTO_Capacity    ActionItemDTO_CommodityAttribute = 0
	ActionItemDTO_Limit       ActionItemDTO_CommodityAttribute = 1
	ActionItemDTO_Reservation ActionItemDTO_CommodityAttribute = 2
)

var ActionItemDTO_CommodityAttribute_name = map[int32]string{
	0: "Capacity",
	1: "Limit",
	2: "Reservation",
}
var ActionItemDTO_CommodityAttribute_value = map[string]int32{
	"Capacity":    0,
	"Limit":       1,
	"Reservation": 2,
}

func (x ActionItemDTO_CommodityAttribute) Enum() *ActionItemDTO_CommodityAttribute {
	p := new(ActionItemDTO_CommodityAttribute)
	*p = x
	return p
}
func (x ActionItemDTO_CommodityAttribute) String() string {
	return proto1.EnumName(ActionItemDTO_CommodityAttribute_name, int32(x))
}
func (x *ActionItemDTO_CommodityAttribute) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(ActionItemDTO_CommodityAttribute_value, data, "ActionItemDTO_CommodityAttribute")
	if err != nil {
		return err
	}
	*x = ActionItemDTO_CommodityAttribute(value)
	return nil
}
func (ActionItemDTO_CommodityAttribute) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 1}
}

// ActionPolicyDTO defines action execution policy for an EntityType.
// Every EntityType can be associated with a list of actions and execution modes.
// E.g., 'Resize' action can be defined as 'Recommend' for Virtual Machines.
type ActionPolicyDTO struct {
	// Entity type which the policy is applied to
	EntityType *EntityDTO_EntityType `protobuf:"varint,1,req,name=entityType,enum=proto.EntityDTO_EntityType" json:"entityType,omitempty"`
	// Action policy elements (for every supported action type)
	PolicyElement    []*ActionPolicyDTO_ActionPolicyElement `protobuf:"bytes,2,rep,name=policyElement" json:"policyElement,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *ActionPolicyDTO) Reset()                    { *m = ActionPolicyDTO{} }
func (m *ActionPolicyDTO) String() string            { return proto1.CompactTextString(m) }
func (*ActionPolicyDTO) ProtoMessage()               {}
func (*ActionPolicyDTO) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActionPolicyDTO) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

func (m *ActionPolicyDTO) GetPolicyElement() []*ActionPolicyDTO_ActionPolicyElement {
	if m != nil {
		return m.PolicyElement
	}
	return nil
}

// Policy applied to a particular action type
type ActionPolicyDTO_ActionPolicyElement struct {
	// Action type
	ActionType *ActionItemDTO_ActionType `protobuf:"varint,1,req,name=actionType,enum=proto.ActionItemDTO_ActionType" json:"actionType,omitempty"`
	// Action execution capability
	ActionCapability *ActionPolicyDTO_ActionCapability `protobuf:"varint,2,req,name=actionCapability,enum=proto.ActionPolicyDTO_ActionCapability" json:"actionCapability,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *ActionPolicyDTO_ActionPolicyElement) Reset()         { *m = ActionPolicyDTO_ActionPolicyElement{} }
func (m *ActionPolicyDTO_ActionPolicyElement) String() string { return proto1.CompactTextString(m) }
func (*ActionPolicyDTO_ActionPolicyElement) ProtoMessage()    {}
func (*ActionPolicyDTO_ActionPolicyElement) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *ActionPolicyDTO_ActionPolicyElement) GetActionType() ActionItemDTO_ActionType {
	if m != nil && m.ActionType != nil {
		return *m.ActionType
	}
	return ActionItemDTO_NONE
}

func (m *ActionPolicyDTO_ActionPolicyElement) GetActionCapability() ActionPolicyDTO_ActionCapability {
	if m != nil && m.ActionCapability != nil {
		return *m.ActionCapability
	}
	return ActionPolicyDTO_NOT_SUPPORTED
}

// ActionExecutionDTO may contain one or more action items related to the overall action to be
// executed.
type ActionExecutionDTO struct {
	// Overall action type.  In most cases this action type will be the same as the action type
	// specified in the action items contained within this.
	ActionType *ActionItemDTO_ActionType `protobuf:"varint,1,req,name=actionType,enum=proto.ActionItemDTO_ActionType" json:"actionType,omitempty"`
	// One or more action items related to the overall action to be executed
	ActionItem []*ActionItemDTO `protobuf:"bytes,2,rep,name=actionItem" json:"actionItem,omitempty"`
	// Action Execution Progress
	Progress *int64 `protobuf:"varint,3,opt,name=progress" json:"progress,omitempty"`
	// There may be an action orchestration workflow
	// associated with the execution of an action
	Workflow         *NonMarketEntityDTO `protobuf:"bytes,4,opt,name=workflow" json:"workflow,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ActionExecutionDTO) Reset()                    { *m = ActionExecutionDTO{} }
func (m *ActionExecutionDTO) String() string            { return proto1.CompactTextString(m) }
func (*ActionExecutionDTO) ProtoMessage()               {}
func (*ActionExecutionDTO) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ActionExecutionDTO) GetActionType() ActionItemDTO_ActionType {
	if m != nil && m.ActionType != nil {
		return *m.ActionType
	}
	return ActionItemDTO_NONE
}

func (m *ActionExecutionDTO) GetActionItem() []*ActionItemDTO {
	if m != nil {
		return m.ActionItem
	}
	return nil
}

func (m *ActionExecutionDTO) GetProgress() int64 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

func (m *ActionExecutionDTO) GetWorkflow() *NonMarketEntityDTO {
	if m != nil {
		return m.Workflow
	}
	return nil
}

// This message holds values necessary for executing actions on Service Entity discovered with SDK probe
type ActionItemDTO struct {
	// ActionType of the action
	ActionType *ActionItemDTO_ActionType `protobuf:"varint,1,req,name=actionType,enum=proto.ActionItemDTO_ActionType" json:"actionType,omitempty"`
	// ActionItem's uuid
	Uuid *string `protobuf:"bytes,2,req,name=uuid" json:"uuid,omitempty"`
	// EntityDTO for the Target ServiceEntity on which the action should be applied
	TargetSE *EntityDTO `protobuf:"bytes,3,req,name=targetSE" json:"targetSE,omitempty"`
	// EntityDTO for the service entity which hosts the targetSE
	HostedBySE *EntityDTO `protobuf:"bytes,4,opt,name=hostedBySE" json:"hostedBySE,omitempty"`
	// EntityDTO for the ServiceEntity by which the Target ServiceEntity is hosted now (for Move action only)
	CurrentSE *EntityDTO `protobuf:"bytes,5,opt,name=currentSE" json:"currentSE,omitempty"`
	// EntityDTO for the ServiceEntity by which the Target ServiceEntity will be hosted
	// after applied the action (for Move action only)
	NewSE *EntityDTO `protobuf:"bytes,6,opt,name=newSE" json:"newSE,omitempty"`
	// CommodityDTO for the commodity on which the action
	// should be applied (current values)
	CurrentComm *CommodityDTO `protobuf:"bytes,7,opt,name=currentComm" json:"currentComm,omitempty"`
	// CommodityDTO for the commodity on which the action
	// should be applied (after the action is applied)
	NewComm *CommodityDTO `protobuf:"bytes,8,opt,name=newComm" json:"newComm,omitempty"`
	// CommodityAttribute enum notifying type of changed attribute
	CommodityAttribute *ActionItemDTO_CommodityAttribute `protobuf:"varint,10,opt,name=commodityAttribute,enum=proto.ActionItemDTO_CommodityAttribute" json:"commodityAttribute,omitempty"`
	// Information for providers of the targetSE.
	Providers []*ActionItemDTO_ProviderInfo `protobuf:"bytes,11,rep,name=providers" json:"providers,omitempty"`
	// Profile related information used in executing a deploy action
	EntityProfileDTO *EntityProfileDTO `protobuf:"bytes,12,opt,name=entityProfileDTO" json:"entityProfileDTO,omitempty"`
	// Context data used in executing actions
	ContextData      []*ContextData `protobuf:"bytes,13,rep,name=contextData" json:"contextData,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ActionItemDTO) Reset()                    { *m = ActionItemDTO{} }
func (m *ActionItemDTO) String() string            { return proto1.CompactTextString(m) }
func (*ActionItemDTO) ProtoMessage()               {}
func (*ActionItemDTO) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ActionItemDTO) GetActionType() ActionItemDTO_ActionType {
	if m != nil && m.ActionType != nil {
		return *m.ActionType
	}
	return ActionItemDTO_NONE
}

func (m *ActionItemDTO) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *ActionItemDTO) GetTargetSE() *EntityDTO {
	if m != nil {
		return m.TargetSE
	}
	return nil
}

func (m *ActionItemDTO) GetHostedBySE() *EntityDTO {
	if m != nil {
		return m.HostedBySE
	}
	return nil
}

func (m *ActionItemDTO) GetCurrentSE() *EntityDTO {
	if m != nil {
		return m.CurrentSE
	}
	return nil
}

func (m *ActionItemDTO) GetNewSE() *EntityDTO {
	if m != nil {
		return m.NewSE
	}
	return nil
}

func (m *ActionItemDTO) GetCurrentComm() *CommodityDTO {
	if m != nil {
		return m.CurrentComm
	}
	return nil
}

func (m *ActionItemDTO) GetNewComm() *CommodityDTO {
	if m != nil {
		return m.NewComm
	}
	return nil
}

func (m *ActionItemDTO) GetCommodityAttribute() ActionItemDTO_CommodityAttribute {
	if m != nil && m.CommodityAttribute != nil {
		return *m.CommodityAttribute
	}
	return ActionItemDTO_Capacity
}

func (m *ActionItemDTO) GetProviders() []*ActionItemDTO_ProviderInfo {
	if m != nil {
		return m.Providers
	}
	return nil
}

func (m *ActionItemDTO) GetEntityProfileDTO() *EntityProfileDTO {
	if m != nil {
		return m.EntityProfileDTO
	}
	return nil
}

func (m *ActionItemDTO) GetContextData() []*ContextData {
	if m != nil {
		return m.ContextData
	}
	return nil
}

// This message holds provider information for the target SE of this ActionItem.
type ActionItemDTO_ProviderInfo struct {
	EntityType       *EntityDTO_EntityType `protobuf:"varint,1,req,name=entityType,enum=proto.EntityDTO_EntityType" json:"entityType,omitempty"`
	Ids              []string              `protobuf:"bytes,2,rep,name=ids" json:"ids,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *ActionItemDTO_ProviderInfo) Reset()                    { *m = ActionItemDTO_ProviderInfo{} }
func (m *ActionItemDTO_ProviderInfo) String() string            { return proto1.CompactTextString(m) }
func (*ActionItemDTO_ProviderInfo) ProtoMessage()               {}
func (*ActionItemDTO_ProviderInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *ActionItemDTO_ProviderInfo) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

func (m *ActionItemDTO_ProviderInfo) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto1.RegisterType((*ActionPolicyDTO)(nil), "proto.ActionPolicyDTO")
	proto1.RegisterType((*ActionPolicyDTO_ActionPolicyElement)(nil), "proto.ActionPolicyDTO.ActionPolicyElement")
	proto1.RegisterType((*ActionExecutionDTO)(nil), "proto.ActionExecutionDTO")
	proto1.RegisterType((*ActionItemDTO)(nil), "proto.ActionItemDTO")
	proto1.RegisterType((*ActionItemDTO_ProviderInfo)(nil), "proto.ActionItemDTO.ProviderInfo")
	proto1.RegisterEnum("proto.ActionResponseState", ActionResponseState_name, ActionResponseState_value)
	proto1.RegisterEnum("proto.ActionPolicyDTO_ActionCapability", ActionPolicyDTO_ActionCapability_name, ActionPolicyDTO_ActionCapability_value)
	proto1.RegisterEnum("proto.ActionItemDTO_ActionType", ActionItemDTO_ActionType_name, ActionItemDTO_ActionType_value)
	proto1.RegisterEnum("proto.ActionItemDTO_CommodityAttribute", ActionItemDTO_CommodityAttribute_name, ActionItemDTO_CommodityAttribute_value)
}

func init() { proto1.RegisterFile("ActionExecution.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 998 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x41, 0x6f, 0xe2, 0x46,
	0x14, 0xc7, 0x17, 0x08, 0x09, 0x3c, 0x20, 0x4c, 0x26, 0x9b, 0x5d, 0x2f, 0x3d, 0x2c, 0xe5, 0xd0,
	0x46, 0xab, 0x16, 0x55, 0xd1, 0xee, 0xa9, 0x52, 0x57, 0x8e, 0xfd, 0x00, 0x57, 0xc1, 0x76, 0xc7,
	0x26, 0xe9, 0x56, 0xaa, 0x2c, 0x02, 0x93, 0xad, 0xb5, 0x60, 0x23, 0x33, 0x24, 0x9b, 0x6f, 0xd2,
	0x43, 0x2f, 0x3d, 0xf7, 0x1b, 0xf5, 0xd0, 0xcf, 0x52, 0xcd, 0x98, 0x60, 0x02, 0xc9, 0x5e, 0x72,
	0xf2, 0xcc, 0xbc, 0xdf, 0xfb, 0xbf, 0xf1, 0x7b, 0x33, 0x6f, 0xe0, 0x48, 0x1f, 0x89, 0x30, 0x8e,
	0xf0, 0x33, 0x1f, 0x2d, 0xe4, 0xa0, 0x3d, 0x4b, 0x62, 0x11, 0xd3, 0xa2, 0xfa, 0x34, 0xea, 0x46,
	0x3c, 0x9d, 0xc6, 0x91, 0xe9, 0x3b, 0xe9, 0x7a, 0x83, 0xb8, 0x49, 0x7c, 0x15, 0x4e, 0x78, 0xb6,
	0xa2, 0xd9, 0x71, 0xd4, 0x1f, 0x26, 0x9f, 0xb8, 0xc0, 0x48, 0x84, 0xe2, 0x76, 0x65, 0x69, 0xfd,
	0x59, 0x80, 0x7a, 0xaa, 0xee, 0xc6, 0x93, 0x70, 0x24, 0x2d, 0xf4, 0x47, 0x00, 0xae, 0x30, 0xff,
	0x76, 0xc6, 0xb5, 0x5c, 0x33, 0x7f, 0xbc, 0x7f, 0xf2, 0x55, 0xca, 0xb7, 0x33, 0x7f, 0x5c, 0x21,
	0x6c, 0x0d, 0xa7, 0x2e, 0xd4, 0x66, 0x4a, 0x09, 0x27, 0x7c, 0xca, 0x23, 0xa1, 0xe5, 0x9b, 0x85,
	0xe3, 0xca, 0xc9, 0x9b, 0xa5, 0xff, 0x46, 0xac, 0x7b, 0xf3, 0xa5, 0x07, 0xbb, 0x2f, 0xd0, 0xf8,
	0x27, 0x07, 0x87, 0x0f, 0x60, 0xf4, 0x3d, 0xc0, 0x50, 0x2d, 0xaf, 0x6d, 0xf3, 0xf5, 0xbd, 0x30,
	0x96, 0xe0, 0xd3, 0x2c, 0x48, 0xba, 0xd5, 0xcc, 0x85, 0x7a, 0x40, 0xd2, 0x99, 0x31, 0x9c, 0x0d,
	0x2f, 0xc3, 0x49, 0x28, 0x6e, 0xb5, 0xbc, 0x92, 0xf9, 0xf6, 0x8b, 0xbb, 0xcd, 0x70, 0xb6, 0x25,
	0xd0, 0xea, 0x01, 0xd9, 0xa4, 0xe8, 0x01, 0xd4, 0x6c, 0xc7, 0x0f, 0xbc, 0x81, 0xeb, 0x3a, 0xcc,
	0x47, 0x93, 0xe4, 0x28, 0x85, 0x7d, 0xb9, 0x84, 0xbf, 0xa2, 0x31, 0xf0, 0xf5, 0xd3, 0x33, 0x24,
	0x79, 0x5a, 0x83, 0x72, 0x86, 0x14, 0x5a, 0xff, 0xe5, 0x80, 0x6e, 0x14, 0x5e, 0x56, 0xe7, 0xc9,
	0xbf, 0xfd, 0xf6, 0x4e, 0x40, 0x72, 0xcb, 0xf2, 0x3c, 0x7f, 0x48, 0x80, 0xad, 0x71, 0xb4, 0x01,
	0xa5, 0x59, 0x12, 0x7f, 0x4c, 0xf8, 0x7c, 0xae, 0x15, 0x9a, 0xb9, 0xe3, 0x02, 0x5b, 0xcd, 0xe9,
	0x3b, 0x28, 0xdd, 0xc4, 0xc9, 0xa7, 0xab, 0x49, 0x7c, 0xa3, 0xed, 0x34, 0x73, 0xc7, 0x95, 0x93,
	0x57, 0x4b, 0xbd, 0xed, 0x73, 0xc7, 0x56, 0x68, 0xeb, 0xdf, 0x32, 0xd4, 0xee, 0x05, 0x7c, 0xfa,
	0xbf, 0x51, 0xd8, 0x59, 0x2c, 0xc2, 0xb1, 0x2a, 0x63, 0x99, 0xa9, 0x31, 0xfd, 0x0e, 0x4a, 0x62,
	0x98, 0x7c, 0xe4, 0xc2, 0x43, 0xad, 0xd0, 0xcc, 0x1f, 0x57, 0x4e, 0xc8, 0xe6, 0x61, 0x66, 0x2b,
	0x82, 0xfe, 0x00, 0xf0, 0x47, 0x3c, 0x17, 0x7c, 0x7c, 0x7a, 0xeb, 0xe1, 0xf2, 0x6f, 0xb6, 0xf9,
	0x35, 0x86, 0xb6, 0xa1, 0x3c, 0x5a, 0x24, 0x09, 0x8f, 0x64, 0x80, 0xe2, 0x23, 0x0e, 0x19, 0x42,
	0xbf, 0x81, 0x62, 0xc4, 0x6f, 0x3c, 0xd4, 0x76, 0x1f, 0x61, 0x53, 0x33, 0x7d, 0x07, 0x95, 0xa5,
	0x93, 0xbc, 0xe0, 0xda, 0x9e, 0xa2, 0x0f, 0x97, 0xb4, 0xba, 0xf3, 0xe3, 0xa5, 0xc3, 0x3a, 0x47,
	0xbf, 0x87, 0xbd, 0x88, 0xdf, 0x28, 0x97, 0xd2, 0xe3, 0x2e, 0x77, 0x0c, 0xbd, 0x00, 0x3a, 0xba,
	0x33, 0xe8, 0x42, 0x24, 0xe1, 0xe5, 0x42, 0x70, 0x0d, 0x9a, 0xb9, 0xad, 0x6b, 0x70, 0x97, 0x7a,
	0x63, 0x0b, 0x67, 0x0f, 0x48, 0xd0, 0xf7, 0x50, 0x9e, 0x25, 0xf1, 0x75, 0x38, 0xe6, 0xc9, 0x5c,
	0xab, 0xa8, 0x53, 0xf6, 0xf5, 0x83, 0x7a, 0xee, 0x92, 0xb2, 0xa2, 0xab, 0x98, 0x65, 0x3e, 0xd4,
	0x00, 0x92, 0xf6, 0x95, 0xac, 0x9d, 0x69, 0x55, 0xf5, 0x47, 0x2f, 0xef, 0xa5, 0x2c, 0x33, 0xb3,
	0x2d, 0x07, 0xfa, 0x16, 0x2a, 0xa3, 0x38, 0x12, 0xfc, 0xb3, 0x30, 0x87, 0x62, 0xa8, 0xd5, 0xd4,
	0x3e, 0xe8, 0x2a, 0x23, 0x2b, 0x0b, 0x5b, 0xc7, 0x1a, 0xbf, 0x43, 0x75, 0x7d, 0x57, 0x4f, 0xeb,
	0x88, 0x04, 0x0a, 0xe1, 0x78, 0xae, 0x2e, 0x5a, 0x99, 0xc9, 0x61, 0xeb, 0xef, 0x02, 0x40, 0x76,
	0x80, 0x69, 0x09, 0x76, 0x6c, 0xc7, 0x46, 0xf2, 0x8c, 0x96, 0xa1, 0xe8, 0xf9, 0x3a, 0xf3, 0x49,
	0x4e, 0x2e, 0xf6, 0x9d, 0x73, 0xd9, 0x16, 0x2a, 0xb0, 0xe7, 0x0d, 0x3c, 0x17, 0x6d, 0x93, 0x14,
	0x64, 0x8f, 0xf0, 0x91, 0xf5, 0x2d, 0x5b, 0xf7, 0x91, 0xec, 0x28, 0x07, 0x57, 0xbf, 0xb0, 0x49,
	0x91, 0x12, 0xa8, 0xea, 0xa6, 0x19, 0xb8, 0xcc, 0x39, 0xb7, 0x4c, 0x64, 0x64, 0x97, 0x02, 0xec,
	0x1a, 0x3d, 0xdd, 0xee, 0x22, 0xd9, 0xa3, 0x87, 0x50, 0x67, 0x28, 0x05, 0x33, 0xa0, 0x24, 0xc5,
	0xd4, 0xcc, 0xb3, 0x1c, 0x9b, 0x94, 0x69, 0x1d, 0x2a, 0x0c, 0x0d, 0xc7, 0xee, 0x58, 0xdd, 0x01,
	0x43, 0x02, 0x52, 0x80, 0xa1, 0x67, 0xfd, 0x86, 0xa4, 0x92, 0x0a, 0xc8, 0x71, 0x60, 0xe8, 0xae,
	0x6e, 0x58, 0xfe, 0x07, 0x52, 0x95, 0x9b, 0xbc, 0xd0, 0x99, 0x4d, 0x6a, 0xf4, 0x15, 0x1c, 0xad,
	0xf9, 0x06, 0x7e, 0x8f, 0xa1, 0xd7, 0x73, 0xce, 0x4c, 0xb2, 0x2f, 0x55, 0x4c, 0x3c, 0x43, 0x1f,
	0x49, 0x9d, 0xee, 0x03, 0x30, 0xab, 0xdb, 0xf3, 0x03, 0xa5, 0x4a, 0x64, 0x67, 0x64, 0xe8, 0x21,
	0x3b, 0xc7, 0xc0, 0xb1, 0x03, 0xb7, 0x4f, 0x0e, 0x36, 0x96, 0x4c, 0x8f, 0xd0, 0x54, 0x5c, 0xc5,
	0xee, 0x38, 0x2c, 0xc0, 0x4e, 0xc7, 0x32, 0x2c, 0xb4, 0x8d, 0x0f, 0xe4, 0x90, 0x36, 0xe0, 0xc5,
	0x9a, 0xc9, 0x45, 0xd6, 0x71, 0x58, 0x5f, 0xb7, 0x0d, 0x24, 0xcf, 0xe9, 0x11, 0x1c, 0x18, 0xcc,
	0xf1, 0xbc, 0xc0, 0xd7, 0x59, 0x17, 0xfd, 0x40, 0xe5, 0xf3, 0x48, 0x06, 0x50, 0x89, 0xf0, 0x9d,
	0x2e, 0xfa, 0x3d, 0x64, 0xe4, 0xc5, 0x66, 0x4c, 0x9d, 0xbc, 0x6c, 0xfd, 0x04, 0x74, 0xfb, 0xa0,
	0xd3, 0x2a, 0x94, 0x64, 0x5f, 0x1f, 0x85, 0xe2, 0x36, 0x2d, 0xd7, 0x59, 0x38, 0x0d, 0x05, 0xc9,
	0xa9, 0xdc, 0xf1, 0x39, 0x4f, 0xae, 0x87, 0xb2, 0xac, 0x24, 0xff, 0xe6, 0xaf, 0xd5, 0xab, 0xc5,
	0xf8, 0x7c, 0x16, 0x47, 0x73, 0xee, 0x89, 0xa1, 0x90, 0x1d, 0x6a, 0x5f, 0x96, 0xd2, 0xb2, 0xbb,
	0x81, 0x6e, 0x18, 0xe8, 0xfa, 0xe4, 0x99, 0x54, 0x4d, 0xc7, 0xea, 0x69, 0xa8, 0x42, 0x89, 0xe1,
	0xcf, 0x68, 0xc8, 0x59, 0x5e, 0x0a, 0x5b, 0xb6, 0x2c, 0x5a, 0x97, 0xa1, 0xe7, 0xa5, 0x27, 0xc0,
	0x1b, 0x18, 0x06, 0xa2, 0x89, 0x26, 0xd9, 0x91, 0xd9, 0xed, 0xe8, 0xd6, 0x19, 0x9a, 0xa4, 0x78,
	0x57, 0xc0, 0x7e, 0x1f, 0x6d, 0x69, 0xdc, 0x95, 0x52, 0xa6, 0xe5, 0xc9, 0xe7, 0xc5, 0x24, 0x7b,
	0x12, 0xfd, 0x65, 0x80, 0x03, 0x34, 0x49, 0xe9, 0xf4, 0x04, 0x5e, 0x8f, 0xe2, 0x69, 0xfb, 0x7a,
	0x2a, 0x16, 0xc9, 0x65, 0xdc, 0x9e, 0x4d, 0x86, 0xe2, 0x2a, 0x4e, 0xa6, 0x6d, 0x75, 0x91, 0xa3,
	0xf6, 0x58, 0xc4, 0xa7, 0xf5, 0x8d, 0xc7, 0xe7, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xf2,
	0xa0, 0xec, 0x87, 0x08, 0x00, 0x00,
}
