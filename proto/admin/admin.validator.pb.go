// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/admin/admin.proto

package admin

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	time "time"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

func (this *CommonRsp) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Empty) Validate() error {
	return nil
}
func (this *CreateBrandReq) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if !(len(this.Name) < 255) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '255'`, this.Name))
	}
	if this.CoverUrl == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CoverUrl", fmt.Errorf(`value '%v' must not be an empty string`, this.CoverUrl))
	}
	if !(len(this.CoverUrl) < 255) {
		return github_com_mwitkow_go_proto_validators.FieldError("CoverUrl", fmt.Errorf(`value '%v' must have a length smaller than '255'`, this.CoverUrl))
	}
	if !(this.SortOrder > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("SortOrder", fmt.Errorf(`value '%v' must be greater than '0'`, this.SortOrder))
	}
	if !(this.CreatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.CreatedBy))
	}
	if !(this.UpdatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.UpdatedBy))
	}
	if this.Remark == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Remark", fmt.Errorf(`value '%v' must not be an empty string`, this.Remark))
	}
	if !(len(this.Remark) < 255) {
		return github_com_mwitkow_go_proto_validators.FieldError("Remark", fmt.Errorf(`value '%v' must have a length smaller than '255'`, this.Remark))
	}
	return nil
}
func (this *CreateBrandRsp) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateBrandData) Validate() error {
	return nil
}
func (this *UpdateBrandReq) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if !(len(this.Name) < 255) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '255'`, this.Name))
	}
	if this.CoverUrl == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CoverUrl", fmt.Errorf(`value '%v' must not be an empty string`, this.CoverUrl))
	}
	if !(len(this.CoverUrl) < 255) {
		return github_com_mwitkow_go_proto_validators.FieldError("CoverUrl", fmt.Errorf(`value '%v' must have a length smaller than '255'`, this.CoverUrl))
	}
	if !(this.SortOrder > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("SortOrder", fmt.Errorf(`value '%v' must be greater than '0'`, this.SortOrder))
	}
	if !(this.CreatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.CreatedBy))
	}
	if !(this.UpdatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.UpdatedBy))
	}
	if this.Remark == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Remark", fmt.Errorf(`value '%v' must not be an empty string`, this.Remark))
	}
	if !(len(this.Remark) < 255) {
		return github_com_mwitkow_go_proto_validators.FieldError("Remark", fmt.Errorf(`value '%v' must have a length smaller than '255'`, this.Remark))
	}
	return nil
}
func (this *GetBrandByIDReq) Validate() error {
	if !(this.BrandId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("BrandId", fmt.Errorf(`value '%v' must be greater than '0'`, this.BrandId))
	}
	return nil
}
func (this *GetBrandByIDRsp) Validate() error {
	if this.CreatedTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedTime", err)
		}
	}
	if this.UpdatedTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedTime", err)
		}
	}
	return nil
}
func (this *MultiGetBrandReq) Validate() error {
	if !(this.Page > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Page", fmt.Errorf(`value '%v' must be greater than '0'`, this.Page))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if !(len(this.Name) < 255) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '255'`, this.Name))
	}
	if this.StartTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartTime", err)
		}
	}
	if this.EndTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EndTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EndTime", err)
		}
	}
	return nil
}
func (this *MultiGetBrandRsp) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *MultiGetBrandData) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *MultiGetBrandContent) Validate() error {
	if this.CreatedTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedTime", err)
		}
	}
	if this.UpdatedTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedTime", err)
		}
	}
	return nil
}
func (this *RemoveBrandReq) Validate() error {
	if !(this.BrandId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("BrandId", fmt.Errorf(`value '%v' must be greater than '0'`, this.BrandId))
	}
	return nil
}
