package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/twitchtv/twirp/exp/gen"
	"github.com/twitchtv/twirp/exp/gen/stringutils"
	"github.com/twitchtv/twirp/exp/gen/typemap"
)

func main() {
	versionFlag := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *versionFlag {
		fmt.Println(gen.Version)
		os.Exit(0)
	}

	g := newGenerator()
	gen.Main(g)
}

func newGenerator() *generator {
	return &generator{output: new(bytes.Buffer)}
}

type generator struct {
	reg    *typemap.Registry
	output *bytes.Buffer
}

func (g *generator) Generate(in *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	genFiles := gen.FilesToGenerate(in)
	g.reg = typemap.New(in.ProtoFile)

	resp := new(plugin.CodeGeneratorResponse)
	for _, f := range genFiles {
		respFile := g.generateFile(f)
		if respFile != nil {
			resp.File = append(resp.File, respFile)
		}
	}

	return resp
}

func (g *generator) generateFile(file *descriptor.FileDescriptorProto) *plugin.CodeGeneratorResponse_File {
	g.P("// Code generated by protoc-gen-twirp_browserjs ", gen.Version, ", DO NOT EDIT.")
	g.P("// source: ", file.GetName())
	g.P("")

	g.P("// _request takes the HTTP method, URL path (this will usually contain the domain name)")
	g.P("// json body that will be sent to the server, a callback on successful requests and a")
	g.P("// callback for requests that error out.")
	g.P("var _request = function(method, path, body, onSuccess, onError) {")
	g.P("  var xhr = new XMLHttpRequest();")
	g.P("  xhr.open(method, path, true);")
	g.P(`  xhr.setRequestHeader("Accept","application/json");`)
	g.P(`  xhr.setRequestHeader("Content-Type","application/json");`)
	g.P("")
	g.P("  xhr.onreadystatechange = function (e) {")
	g.P("    if (xhr.readyState == 4) {")
	g.P("      if (xhr.status == 204 || xhr.status == 205) {")
	g.P("        onSuccess();")
	g.P("      } else if (xhr.status == 200) {")
	g.P("        var value = JSON.parse(xhr.responseText);")
	g.P("        onSuccess(value);")
	g.P("      } else {")
	g.P("        var value = JSON.parse(xhr.responseText);")
	g.P("        onError(value);")
	g.P("      }")
	g.P("    }")
	g.P("  };")
	g.P("")
	g.P("  if (body != null) {")
	g.P("    xhr.send(JSON.stringify(body));")
	g.P("  } else {") // realistically this branch will never happen
	g.P("    xhr.send(null);")
	g.P("  }")
	g.P("};")

	g.P("")

	for _, service := range file.Service {
		g.generateProtobufClient(file, service)
	}

	resp := new(plugin.CodeGeneratorResponse_File)
	resp.Name = proto.String(jsFileName(file))
	resp.Content = proto.String(g.output.String())
	g.output.Reset()

	return resp
}

func (g *generator) generateProtobufClient(file *descriptor.FileDescriptorProto, service *descriptor.ServiceDescriptorProto) {
	g.P("// methods for ", clientName(service))

	comments, err := g.reg.ServiceComments(file, service)
	if err == nil && comments.Leading != "" {
		g.P("/*")
		g.printComments(comments, `    `)
		g.P("*/")
		g.P()
	}

	g.P("")

	for _, method := range service.Method {
		svcName := serviceName(service)
		methName := methodName(method)
		inputName := methodInputName(method)

		comments, err := g.reg.MethodComments(file, service, method)
		if err == nil && comments.Leading != "" {
			g.P("/*")
			g.printComments(comments, `    `)
			g.P("*/")
		}

		// Be careful not to write code that overwrites the input parameter.
		for _, x := range []string{"self", "_sym_db", "full_method", "body",
			"serialize", "deserialize", "resp_str"} {
			if inputName == x {
				inputName = inputName + "_"
			}
		}

		g.P("var ", svcName+"_"+methName, " = function(server_address, ", inputName, ", onSuccess, onError) {")
		g.P(`  var full_method = server_address + "/twirp/" + `, strconv.Quote(fullServiceName(file, service)), ` + "/" + `, strconv.Quote(method.GetName()), ";")
		g.P(`  _request("POST", full_method, `, inputName, ", onSuccess, onError);")
		g.P("};")
	}
}

func (g *generator) P(args ...string) {
	for _, v := range args {
		g.output.WriteString(v)
	}
	g.output.WriteByte('\n')
}

func (g *generator) printComments(comments typemap.DefinitionComments, prefix string) {
	text := strings.TrimSuffix(comments.Leading, "\n")
	for _, line := range strings.Split(text, "\n") {
		g.P(prefix, strings.TrimPrefix(line, " "))
	}
}

func serviceName(service *descriptor.ServiceDescriptorProto) string {
	return stringutils.CamelCase(service.GetName())
}

func clientName(service *descriptor.ServiceDescriptorProto) string {
	return serviceName(service) + "Client"
}

func fullServiceName(file *descriptor.FileDescriptorProto, service *descriptor.ServiceDescriptorProto) string {
	name := serviceName(service)
	if pkg := file.GetPackage(); pkg != "" {
		name = pkg + "." + name
	}
	return name
}

func methodName(method *descriptor.MethodDescriptorProto) string {
	return stringutils.SnakeCase(method.GetName())
}

// methodInputName returns the basename of the input type of a method in snake
// case.
func methodInputName(meth *descriptor.MethodDescriptorProto) string {
	fullName := meth.GetInputType()
	split := strings.Split(fullName, ".")
	return stringutils.SnakeCase(split[len(split)-1])
}

func jsFileName(f *descriptor.FileDescriptorProto) string {
	name := *f.Name
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		name = name[:len(name)-len(ext)]
	}
	name += "_twirp.js"
	return name
}
