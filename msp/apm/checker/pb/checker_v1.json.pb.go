// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: checker_v1.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*CreateCheckerV1Request)(nil)
var _ json.Unmarshaler = (*CreateCheckerV1Request)(nil)
var _ json.Marshaler = (*CreateCheckerV1Response)(nil)
var _ json.Unmarshaler = (*CreateCheckerV1Response)(nil)
var _ json.Marshaler = (*UpdateCheckerV1Request)(nil)
var _ json.Unmarshaler = (*UpdateCheckerV1Request)(nil)
var _ json.Marshaler = (*UpdateCheckerV1Response)(nil)
var _ json.Unmarshaler = (*UpdateCheckerV1Response)(nil)
var _ json.Marshaler = (*DeleteCheckerV1Request)(nil)
var _ json.Unmarshaler = (*DeleteCheckerV1Request)(nil)
var _ json.Marshaler = (*DeleteCheckerV1Response)(nil)
var _ json.Unmarshaler = (*DeleteCheckerV1Response)(nil)
var _ json.Marshaler = (*GetCheckerV1Request)(nil)
var _ json.Unmarshaler = (*GetCheckerV1Request)(nil)
var _ json.Marshaler = (*GetCheckerV1Response)(nil)
var _ json.Unmarshaler = (*GetCheckerV1Response)(nil)
var _ json.Marshaler = (*DescribeCheckersV1Request)(nil)
var _ json.Unmarshaler = (*DescribeCheckersV1Request)(nil)
var _ json.Marshaler = (*DescribeCheckersV1Response)(nil)
var _ json.Unmarshaler = (*DescribeCheckersV1Response)(nil)
var _ json.Marshaler = (*DescribeCheckerV1Request)(nil)
var _ json.Unmarshaler = (*DescribeCheckerV1Request)(nil)
var _ json.Marshaler = (*DescribeCheckerV1Response)(nil)
var _ json.Unmarshaler = (*DescribeCheckerV1Response)(nil)
var _ json.Marshaler = (*GetCheckerStatusV1Request)(nil)
var _ json.Unmarshaler = (*GetCheckerStatusV1Request)(nil)
var _ json.Marshaler = (*GetCheckerStatusV1Response)(nil)
var _ json.Unmarshaler = (*GetCheckerStatusV1Response)(nil)
var _ json.Marshaler = (*CheckerStatusV1)(nil)
var _ json.Unmarshaler = (*CheckerStatusV1)(nil)
var _ json.Marshaler = (*GetCheckerIssuesV1Request)(nil)
var _ json.Unmarshaler = (*GetCheckerIssuesV1Request)(nil)
var _ json.Marshaler = (*GetCheckerIssuesV1Response)(nil)
var _ json.Unmarshaler = (*GetCheckerIssuesV1Response)(nil)
var _ json.Marshaler = (*CheckerV1)(nil)
var _ json.Unmarshaler = (*CheckerV1)(nil)
var _ json.Marshaler = (*DescribeResultV1)(nil)
var _ json.Unmarshaler = (*DescribeResultV1)(nil)
var _ json.Marshaler = (*DescribeItemV1)(nil)
var _ json.Unmarshaler = (*DescribeItemV1)(nil)
var _ json.Marshaler = (*CheckerChartV1)(nil)
var _ json.Unmarshaler = (*CheckerChartV1)(nil)

// CreateCheckerV1Request implement json.Marshaler.
func (m *CreateCheckerV1Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateCheckerV1Request implement json.Marshaler.
func (m *CreateCheckerV1Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateCheckerV1Response implement json.Marshaler.
func (m *CreateCheckerV1Response) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateCheckerV1Response implement json.Marshaler.
func (m *CreateCheckerV1Response) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateCheckerV1Request implement json.Marshaler.
func (m *UpdateCheckerV1Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateCheckerV1Request implement json.Marshaler.
func (m *UpdateCheckerV1Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateCheckerV1Response implement json.Marshaler.
func (m *UpdateCheckerV1Response) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateCheckerV1Response implement json.Marshaler.
func (m *UpdateCheckerV1Response) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteCheckerV1Request implement json.Marshaler.
func (m *DeleteCheckerV1Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteCheckerV1Request implement json.Marshaler.
func (m *DeleteCheckerV1Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteCheckerV1Response implement json.Marshaler.
func (m *DeleteCheckerV1Response) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteCheckerV1Response implement json.Marshaler.
func (m *DeleteCheckerV1Response) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCheckerV1Request implement json.Marshaler.
func (m *GetCheckerV1Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCheckerV1Request implement json.Marshaler.
func (m *GetCheckerV1Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCheckerV1Response implement json.Marshaler.
func (m *GetCheckerV1Response) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCheckerV1Response implement json.Marshaler.
func (m *GetCheckerV1Response) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DescribeCheckersV1Request implement json.Marshaler.
func (m *DescribeCheckersV1Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DescribeCheckersV1Request implement json.Marshaler.
func (m *DescribeCheckersV1Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DescribeCheckersV1Response implement json.Marshaler.
func (m *DescribeCheckersV1Response) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DescribeCheckersV1Response implement json.Marshaler.
func (m *DescribeCheckersV1Response) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DescribeCheckerV1Request implement json.Marshaler.
func (m *DescribeCheckerV1Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DescribeCheckerV1Request implement json.Marshaler.
func (m *DescribeCheckerV1Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DescribeCheckerV1Response implement json.Marshaler.
func (m *DescribeCheckerV1Response) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DescribeCheckerV1Response implement json.Marshaler.
func (m *DescribeCheckerV1Response) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCheckerStatusV1Request implement json.Marshaler.
func (m *GetCheckerStatusV1Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCheckerStatusV1Request implement json.Marshaler.
func (m *GetCheckerStatusV1Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCheckerStatusV1Response implement json.Marshaler.
func (m *GetCheckerStatusV1Response) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCheckerStatusV1Response implement json.Marshaler.
func (m *GetCheckerStatusV1Response) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CheckerStatusV1 implement json.Marshaler.
func (m *CheckerStatusV1) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CheckerStatusV1 implement json.Marshaler.
func (m *CheckerStatusV1) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCheckerIssuesV1Request implement json.Marshaler.
func (m *GetCheckerIssuesV1Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCheckerIssuesV1Request implement json.Marshaler.
func (m *GetCheckerIssuesV1Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCheckerIssuesV1Response implement json.Marshaler.
func (m *GetCheckerIssuesV1Response) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCheckerIssuesV1Response implement json.Marshaler.
func (m *GetCheckerIssuesV1Response) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CheckerV1 implement json.Marshaler.
func (m *CheckerV1) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CheckerV1 implement json.Marshaler.
func (m *CheckerV1) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DescribeResultV1 implement json.Marshaler.
func (m *DescribeResultV1) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DescribeResultV1 implement json.Marshaler.
func (m *DescribeResultV1) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DescribeItemV1 implement json.Marshaler.
func (m *DescribeItemV1) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DescribeItemV1 implement json.Marshaler.
func (m *DescribeItemV1) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CheckerChartV1 implement json.Marshaler.
func (m *CheckerChartV1) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CheckerChartV1 implement json.Marshaler.
func (m *CheckerChartV1) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}