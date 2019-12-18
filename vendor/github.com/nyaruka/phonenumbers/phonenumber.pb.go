// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: phonenumber.proto

package phonenumbers

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The source from which the country_code is derived. This is not set in the
// general parsing method, but in the method that parses and keeps raw_input.
// New fields could be added upon request.
type PhoneNumber_CountryCodeSource int32

const (
	// Default value returned if this is not set, because the phone number was
	// created using parse, not parseAndKeepRawInput. hasCountryCodeSource will
	// return false if this is the case.
	PhoneNumber_UNSPECIFIED PhoneNumber_CountryCodeSource = 0
	// The country_code is derived based on a phone number with a leading "+",
	// e.g. the French number "+33 1 42 68 53 00".
	PhoneNumber_FROM_NUMBER_WITH_PLUS_SIGN PhoneNumber_CountryCodeSource = 1
	// The country_code is derived based on a phone number with a leading IDD,
	// e.g. the French number "011 33 1 42 68 53 00", as it is dialled from US.
	PhoneNumber_FROM_NUMBER_WITH_IDD PhoneNumber_CountryCodeSource = 5
	// The country_code is derived based on a phone number without a leading
	// "+", e.g. the French number "33 1 42 68 53 00" when defaultCountry is
	// supplied as France.
	PhoneNumber_FROM_NUMBER_WITHOUT_PLUS_SIGN PhoneNumber_CountryCodeSource = 10
	// The country_code is derived NOT based on the phone number itself, but
	// from the defaultCountry parameter provided in the parsing function by the
	// clients. This happens mostly for numbers written in the national format
	// (without country code). For example, this would be set when parsing the
	// French number "01 42 68 53 00", when defaultCountry is supplied as
	// France.
	PhoneNumber_FROM_DEFAULT_COUNTRY PhoneNumber_CountryCodeSource = 20
)

var PhoneNumber_CountryCodeSource_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "FROM_NUMBER_WITH_PLUS_SIGN",
	5:  "FROM_NUMBER_WITH_IDD",
	10: "FROM_NUMBER_WITHOUT_PLUS_SIGN",
	20: "FROM_DEFAULT_COUNTRY",
}
var PhoneNumber_CountryCodeSource_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"FROM_NUMBER_WITH_PLUS_SIGN":    1,
	"FROM_NUMBER_WITH_IDD":          5,
	"FROM_NUMBER_WITHOUT_PLUS_SIGN": 10,
	"FROM_DEFAULT_COUNTRY":          20,
}

func (x PhoneNumber_CountryCodeSource) Enum() *PhoneNumber_CountryCodeSource {
	p := new(PhoneNumber_CountryCodeSource)
	*p = x
	return p
}
func (x PhoneNumber_CountryCodeSource) String() string {
	return proto.EnumName(PhoneNumber_CountryCodeSource_name, int32(x))
}
func (x *PhoneNumber_CountryCodeSource) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhoneNumber_CountryCodeSource_value, data, "PhoneNumber_CountryCodeSource")
	if err != nil {
		return err
	}
	*x = PhoneNumber_CountryCodeSource(value)
	return nil
}
func (PhoneNumber_CountryCodeSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_phonenumber_2c745f30989ef8a6, []int{0, 0}
}

type PhoneNumber struct {
	// The country calling code for this number, as defined by the International
	// Telecommunication Union (ITU). For example, this would be 1 for NANPA
	// countries, and 33 for France.
	CountryCode *int32 `protobuf:"varint,1,req,name=country_code,json=countryCode" json:"country_code,omitempty"`
	// The National (significant) Number, as defined in International
	// Telecommunication Union (ITU) Recommendation E.164, without any leading
	// zero. The leading-zero is stored separately if required, since this is an
	// uint64 and hence cannot store such information. Do not use this field
	// directly: if you want the national significant number, call the
	// getNationalSignificantNumber method of PhoneNumberUtil.
	//
	// For countries which have the concept of an "area code" or "national
	// destination code", this is included in the National (significant) Number.
	// Although the ITU says the maximum length should be 15, we have found longer
	// numbers in some countries e.g. Germany.
	// Note that the National (significant) Number does not contain the National
	// (trunk) prefix. Obviously, as a uint64, it will never contain any
	// formatting (hyphens, spaces, parentheses), nor any alphanumeric spellings.
	NationalNumber *uint64 `protobuf:"varint,2,req,name=national_number,json=nationalNumber" json:"national_number,omitempty"`
	// Extension is not standardized in ITU recommendations, except for being
	// defined as a series of numbers with a maximum length of 40 digits. It is
	// defined as a string here to accommodate for the possible use of a leading
	// zero in the extension (organizations have complete freedom to do so, as
	// there is no standard defined). Other than digits, some other dialling
	// characters such as "," (indicating a wait) may be stored here.
	Extension *string `protobuf:"bytes,3,opt,name=extension" json:"extension,omitempty"`
	// In some countries, the national (significant) number starts with one or
	// more "0"s without this being a national prefix or trunk code of some kind.
	// For example, the leading zero in the national (significant) number of an
	// Italian phone number indicates the number is a fixed-line number.  There
	// have been plans to migrate fixed-line numbers to start with the digit two
	// since December 2000, but it has not happened yet. See
	// http://en.wikipedia.org/wiki/%2B39 for more details.
	//
	// These fields can be safely ignored (there is no need to set them) for most
	// countries. Some limited number of countries behave like Italy - for these
	// cases, if the leading zero(s) of a number would be retained even when
	// dialling internationally, set this flag to true, and also set the number of
	// leading zeros.
	//
	// Clients who use the parsing functionality of the i18n phone
	// number libraries will have these fields set if necessary automatically.
	ItalianLeadingZero   *bool  `protobuf:"varint,4,opt,name=italian_leading_zero,json=italianLeadingZero" json:"italian_leading_zero,omitempty"`
	NumberOfLeadingZeros *int32 `protobuf:"varint,8,opt,name=number_of_leading_zeros,json=numberOfLeadingZeros,def=1" json:"number_of_leading_zeros,omitempty"`
	// This field is used to store the raw input string containing phone numbers
	// before it was canonicalized by the library. For example, it could be used
	// to store alphanumerical numbers such as "1-800-GOOG-411".
	RawInput *string `protobuf:"bytes,5,opt,name=raw_input,json=rawInput" json:"raw_input,omitempty"`
	// The source from which the country_code is derived.
	CountryCodeSource *PhoneNumber_CountryCodeSource `protobuf:"varint,6,opt,name=country_code_source,json=countryCodeSource,enum=phonenumbers.PhoneNumber_CountryCodeSource" json:"country_code_source,omitempty"`
	// The carrier selection code that is preferred when calling this phone number
	// domestically. This also includes codes that need to be dialed in some
	// countries when calling from landlines to mobiles or vice versa. For
	// example, in Columbia, a "3" needs to be dialed before the phone number
	// itself when calling from a mobile phone to a domestic landline phone and
	// vice versa.
	//
	// Note this is the "preferred" code, which means other codes may work as
	// well.
	PreferredDomesticCarrierCode *string  `protobuf:"bytes,7,opt,name=preferred_domestic_carrier_code,json=preferredDomesticCarrierCode" json:"preferred_domestic_carrier_code,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *PhoneNumber) Reset()         { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()    {}
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_phonenumber_2c745f30989ef8a6, []int{0}
}
func (m *PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneNumber.Unmarshal(m, b)
}
func (m *PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneNumber.Marshal(b, m, deterministic)
}
func (dst *PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneNumber.Merge(dst, src)
}
func (m *PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_PhoneNumber.Size(m)
}
func (m *PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneNumber proto.InternalMessageInfo

const Default_PhoneNumber_NumberOfLeadingZeros int32 = 1

func (m *PhoneNumber) GetCountryCode() int32 {
	if m != nil && m.CountryCode != nil {
		return *m.CountryCode
	}
	return 0
}

func (m *PhoneNumber) GetNationalNumber() uint64 {
	if m != nil && m.NationalNumber != nil {
		return *m.NationalNumber
	}
	return 0
}

func (m *PhoneNumber) GetExtension() string {
	if m != nil && m.Extension != nil {
		return *m.Extension
	}
	return ""
}

func (m *PhoneNumber) GetItalianLeadingZero() bool {
	if m != nil && m.ItalianLeadingZero != nil {
		return *m.ItalianLeadingZero
	}
	return false
}

func (m *PhoneNumber) GetNumberOfLeadingZeros() int32 {
	if m != nil && m.NumberOfLeadingZeros != nil {
		return *m.NumberOfLeadingZeros
	}
	return Default_PhoneNumber_NumberOfLeadingZeros
}

func (m *PhoneNumber) GetRawInput() string {
	if m != nil && m.RawInput != nil {
		return *m.RawInput
	}
	return ""
}

func (m *PhoneNumber) GetCountryCodeSource() PhoneNumber_CountryCodeSource {
	if m != nil && m.CountryCodeSource != nil {
		return *m.CountryCodeSource
	}
	return PhoneNumber_UNSPECIFIED
}

func (m *PhoneNumber) GetPreferredDomesticCarrierCode() string {
	if m != nil && m.PreferredDomesticCarrierCode != nil {
		return *m.PreferredDomesticCarrierCode
	}
	return ""
}

func init() {
	proto.RegisterType((*PhoneNumber)(nil), "phonenumbers.PhoneNumber")
	proto.RegisterEnum("phonenumbers.PhoneNumber_CountryCodeSource", PhoneNumber_CountryCodeSource_name, PhoneNumber_CountryCodeSource_value)
}

func init() { proto.RegisterFile("phonenumber.proto", fileDescriptor_phonenumber_2c745f30989ef8a6) }

var fileDescriptor_phonenumber_2c745f30989ef8a6 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0x71, 0x9a, 0x40, 0x32, 0xa9, 0xda, 0x64, 0x89, 0x84, 0x05, 0x01, 0xdc, 0x5e, 0xb0,
	0x84, 0x64, 0x51, 0x4e, 0x15, 0x37, 0x6a, 0x3b, 0xd4, 0x52, 0xea, 0x44, 0x9b, 0x58, 0x08, 0x38,
	0xac, 0x16, 0x7b, 0x12, 0x56, 0x4a, 0x76, 0xa2, 0xb5, 0xa3, 0x02, 0x7f, 0x85, 0x7f, 0xc3, 0x2f,
	0x43, 0x89, 0x4d, 0xeb, 0xb6, 0xc7, 0xfd, 0xde, 0x7b, 0x33, 0xb3, 0x0f, 0xfa, 0x9b, 0x1f, 0xa4,
	0x51, 0x6f, 0xd7, 0xdf, 0xd1, 0x78, 0x1b, 0x43, 0x05, 0xb1, 0xc3, 0x1a, 0xca, 0x4f, 0xff, 0x36,
	0xa1, 0x3b, 0xdd, 0x81, 0x78, 0x0f, 0xd8, 0x09, 0x1c, 0xa6, 0xb4, 0xd5, 0x85, 0xf9, 0x25, 0x52,
	0xca, 0xd0, 0xb6, 0x9c, 0x86, 0xdb, 0xe2, 0xdd, 0x8a, 0xf9, 0x94, 0x21, 0x7b, 0x03, 0xc7, 0x5a,
	0x16, 0x8a, 0xb4, 0x5c, 0x89, 0x72, 0x8c, 0xdd, 0x70, 0x1a, 0x6e, 0x93, 0x1f, 0xfd, 0xc7, 0xd5,
	0xac, 0x21, 0x74, 0xf0, 0x67, 0x81, 0x3a, 0x57, 0xa4, 0xed, 0x03, 0xc7, 0x72, 0x3b, 0xfc, 0x16,
	0xb0, 0x77, 0x30, 0x50, 0x85, 0x5c, 0x29, 0xa9, 0xc5, 0x0a, 0x65, 0xa6, 0xf4, 0x52, 0xfc, 0x46,
	0x43, 0x76, 0xd3, 0xb1, 0xdc, 0x36, 0x67, 0x95, 0x36, 0x2e, 0xa5, 0xaf, 0x68, 0x88, 0x9d, 0xc3,
	0xb3, 0x72, 0x9f, 0xa0, 0xc5, 0x9d, 0x4c, 0x6e, 0xb7, 0x1d, 0xcb, 0x6d, 0x7d, 0xb0, 0xce, 0xf8,
	0xa0, 0x74, 0x4c, 0x16, 0xb5, 0x60, 0xce, 0x5e, 0x40, 0xc7, 0xc8, 0x6b, 0xa1, 0xf4, 0x66, 0x5b,
	0xd8, 0xad, 0xfd, 0x25, 0x6d, 0x23, 0xaf, 0xa3, 0xdd, 0x9b, 0x7d, 0x83, 0xa7, 0xf5, 0x2f, 0x8b,
	0x9c, 0xb6, 0x26, 0x45, 0xfb, 0xb1, 0x63, 0xb9, 0x47, 0xef, 0xdf, 0x7a, 0xf5, 0xba, 0xbc, 0x5a,
	0x55, 0x9e, 0x7f, 0xdb, 0xc9, 0x6c, 0x1f, 0xe1, 0xfd, 0xf4, 0x3e, 0x62, 0x21, 0xbc, 0xde, 0x18,
	0x5c, 0xa0, 0x31, 0x98, 0x89, 0x8c, 0xd6, 0x98, 0x17, 0x2a, 0x15, 0xa9, 0x34, 0x46, 0xa1, 0x29,
	0x2b, 0x7e, 0xb2, 0xbf, 0x67, 0x78, 0x63, 0x0b, 0x2a, 0x97, 0x5f, 0x9a, 0x76, 0xc3, 0x4e, 0xff,
	0x58, 0xd0, 0x7f, 0xb0, 0x8f, 0x1d, 0x43, 0x37, 0x89, 0x67, 0xd3, 0xd0, 0x8f, 0x46, 0x51, 0x18,
	0xf4, 0x1e, 0xb1, 0x57, 0xf0, 0x7c, 0xc4, 0x27, 0x57, 0x22, 0x4e, 0xae, 0x2e, 0x42, 0x2e, 0x3e,
	0x47, 0xf3, 0x4b, 0x31, 0x1d, 0x27, 0x33, 0x31, 0x8b, 0x3e, 0xc5, 0x3d, 0x8b, 0xd9, 0x30, 0x78,
	0xa0, 0x47, 0x41, 0xd0, 0x6b, 0xb1, 0x13, 0x78, 0x79, 0x5f, 0x99, 0x24, 0xf3, 0x5a, 0x18, 0x6e,
	0xc2, 0x41, 0x38, 0xfa, 0x98, 0x8c, 0xe7, 0xc2, 0x9f, 0x24, 0xf1, 0x9c, 0x7f, 0xe9, 0x0d, 0x2e,
	0x1c, 0x18, 0xa6, 0xb4, 0xf6, 0x96, 0x44, 0xcb, 0x15, 0x7a, 0xea, 0xec, 0x5c, 0xdf, 0x69, 0xed,
	0xf2, 0xe0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xdc, 0xb0, 0x94, 0x88, 0x02, 0x00, 0x00,
}