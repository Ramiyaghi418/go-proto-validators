// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/enum.proto

package validator_examples

import (
	fmt "fmt"
	_ "github.com/Ramiyaghi418/go-proto-validators"
	github_com_Ramiyaghi418_go_proto_validators "github.com/Ramiyaghi418/go-proto-validators"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SomeMsg) Validate() error {
	if _, ok := Action_name[int32(this.Do)]; !ok {
		return github_com_Ramiyaghi418_go_proto_validators.FieldError("Do", fmt.Errorf(`value '%v' must be a valid Action field`, this.Do))
	}
	return nil
}
