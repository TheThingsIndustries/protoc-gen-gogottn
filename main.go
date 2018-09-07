// Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package main

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	req := command.Read()
	files := req.GetProtoFile()
	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)

	for _, opt := range []func(*descriptor.FileDescriptorProto){
		vanity.SetBoolFileOption(gogoproto.E_BenchgenAll, false),
		vanity.SetBoolFileOption(gogoproto.E_CompareAll, false),
		vanity.SetBoolFileOption(gogoproto.E_DescriptionAll, false),
		vanity.SetBoolFileOption(gogoproto.E_EnumStringerAll, true),
		vanity.SetBoolFileOption(gogoproto.E_EnumdeclAll, true),
		vanity.SetBoolFileOption(gogoproto.E_EqualAll, true),
		vanity.SetBoolFileOption(gogoproto.E_FaceAll, false),
		vanity.SetBoolFileOption(gogoproto.E_GogoprotoImport, false),
		vanity.SetBoolFileOption(gogoproto.E_GoprotoEnumPrefixAll, true),
		vanity.SetBoolFileOption(gogoproto.E_GoprotoEnumStringerAll, false),
		vanity.SetBoolFileOption(gogoproto.E_GoprotoExtensionsMapAll, true),
		vanity.SetBoolFileOption(gogoproto.E_GoprotoGettersAll, true),
		vanity.SetBoolFileOption(gogoproto.E_GoprotoRegistration, true),
		vanity.SetBoolFileOption(gogoproto.E_GoprotoStringerAll, false),
		vanity.SetBoolFileOption(gogoproto.E_GoprotoUnrecognizedAll, false),
		vanity.SetBoolFileOption(gogoproto.E_GostringAll, false),
		vanity.SetBoolFileOption(gogoproto.E_MarshalerAll, true),
		vanity.SetBoolFileOption(gogoproto.E_OnlyoneAll, false),
		vanity.SetBoolFileOption(gogoproto.E_PopulateAll, true),
		vanity.SetBoolFileOption(gogoproto.E_ProtosizerAll, false),
		vanity.SetBoolFileOption(gogoproto.E_SizerAll, true),
		vanity.SetBoolFileOption(gogoproto.E_StableMarshalerAll, false),
		vanity.SetBoolFileOption(gogoproto.E_StringerAll, true),
		vanity.SetBoolFileOption(gogoproto.E_TestgenAll, false),
		vanity.SetBoolFileOption(gogoproto.E_TypedeclAll, true),
		vanity.SetBoolFileOption(gogoproto.E_UnmarshalerAll, true),
		vanity.SetBoolFileOption(gogoproto.E_UnsafeMarshalerAll, false),
		vanity.SetBoolFileOption(gogoproto.E_UnsafeUnmarshalerAll, false),
		vanity.SetBoolFileOption(gogoproto.E_VerboseEqualAll, false),
	} {
		vanity.ForEachFile(files, opt)
	}
	command.Write(command.Generate(req))
}
