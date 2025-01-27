// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: project.proto

package pb

import (
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetProjectRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetProjectResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetProjectsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetProjectsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateProjectRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateProjectResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Project)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TenantRelationship)(nil)

// GetProjectRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetProjectRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			}
		}
	}
	return nil
}

// GetProjectResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetProjectResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Project{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &Project{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Id = val
			case "data.name":
				if m.Data == nil {
					m.Data = &Project{}
				}
				m.Data.Name = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &Project{}
				}
				m.Data.Type = vals[0]
			case "data.createTime":
				if m.Data == nil {
					m.Data = &Project{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.CreateTime = val
			case "data.updateTime":
				if m.Data == nil {
					m.Data = &Project{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.UpdateTime = val
			case "data.isDeleted":
				if m.Data == nil {
					m.Data = &Project{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.IsDeleted = val
			case "data.displayName":
				if m.Data == nil {
					m.Data = &Project{}
				}
				m.Data.DisplayName = vals[0]
			case "data.displayType":
				if m.Data == nil {
					m.Data = &Project{}
				}
				m.Data.DisplayType = vals[0]
			}
		}
	}
	return nil
}

// GetProjectsRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetProjectsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectId":
				list := make([]int64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseInt(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.ProjectId = list
			}
		}
	}
	return nil
}

// GetProjectsResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetProjectsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CreateProjectRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateProjectRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "name":
				m.Name = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// CreateProjectResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateProjectResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Project{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &Project{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Id = val
			case "data.name":
				if m.Data == nil {
					m.Data = &Project{}
				}
				m.Data.Name = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &Project{}
				}
				m.Data.Type = vals[0]
			case "data.createTime":
				if m.Data == nil {
					m.Data = &Project{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.CreateTime = val
			case "data.updateTime":
				if m.Data == nil {
					m.Data = &Project{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.UpdateTime = val
			case "data.isDeleted":
				if m.Data == nil {
					m.Data = &Project{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.IsDeleted = val
			case "data.displayName":
				if m.Data == nil {
					m.Data = &Project{}
				}
				m.Data.DisplayName = vals[0]
			case "data.displayType":
				if m.Data == nil {
					m.Data = &Project{}
				}
				m.Data.DisplayType = vals[0]
			}
		}
	}
	return nil
}

// Project implement urlenc.URLValuesUnmarshaler.
func (m *Project) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "name":
				m.Name = vals[0]
			case "type":
				m.Type = vals[0]
			case "createTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreateTime = val
			case "updateTime":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.UpdateTime = val
			case "isDeleted":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsDeleted = val
			case "displayName":
				m.DisplayName = vals[0]
			case "displayType":
				m.DisplayType = vals[0]
			}
		}
	}
	return nil
}

// TenantRelationship implement urlenc.URLValuesUnmarshaler.
func (m *TenantRelationship) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "workspace":
				m.Workspace = vals[0]
			case "tenantID":
				m.TenantID = vals[0]
			case "displayWorkspace":
				m.DisplayWorkspace = vals[0]
			}
		}
	}
	return nil
}
