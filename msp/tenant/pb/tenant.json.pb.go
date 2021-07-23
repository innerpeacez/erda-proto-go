// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: tenant.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*CreateTenantRequest)(nil)
var _ json.Unmarshaler = (*CreateTenantRequest)(nil)
var _ json.Marshaler = (*CreateTenantResponse)(nil)
var _ json.Unmarshaler = (*CreateTenantResponse)(nil)
var _ json.Marshaler = (*GetTenantRequest)(nil)
var _ json.Unmarshaler = (*GetTenantRequest)(nil)
var _ json.Marshaler = (*GetTenantResponse)(nil)
var _ json.Unmarshaler = (*GetTenantResponse)(nil)
var _ json.Marshaler = (*Tenant)(nil)
var _ json.Unmarshaler = (*Tenant)(nil)

// CreateTenantRequest implement json.Marshaler.
func (m *CreateTenantRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateTenantRequest implement json.Marshaler.
func (m *CreateTenantRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateTenantResponse implement json.Marshaler.
func (m *CreateTenantResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateTenantResponse implement json.Marshaler.
func (m *CreateTenantResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetTenantRequest implement json.Marshaler.
func (m *GetTenantRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetTenantRequest implement json.Marshaler.
func (m *GetTenantRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetTenantResponse implement json.Marshaler.
func (m *GetTenantResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetTenantResponse implement json.Marshaler.
func (m *GetTenantResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Tenant implement json.Marshaler.
func (m *Tenant) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Tenant implement json.Marshaler.
func (m *Tenant) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
