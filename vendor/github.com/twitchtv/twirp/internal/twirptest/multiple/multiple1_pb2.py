# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: multiple1.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='multiple1.proto',
  package='twirp.internal.twirptest.multiple',
  syntax='proto3',
  serialized_pb=_b('\n\x0fmultiple1.proto\x12!twirp.internal.twirptest.multiple\"\x06\n\x04Msg12`\n\x04Svc1\x12X\n\x04Send\x12\'.twirp.internal.twirptest.multiple.Msg1\x1a\'.twirp.internal.twirptest.multiple.Msg1B\nZ\x08multipleb\x06proto3')
)




_MSG1 = _descriptor.Descriptor(
  name='Msg1',
  full_name='twirp.internal.twirptest.multiple.Msg1',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=54,
  serialized_end=60,
)

DESCRIPTOR.message_types_by_name['Msg1'] = _MSG1
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Msg1 = _reflection.GeneratedProtocolMessageType('Msg1', (_message.Message,), dict(
  DESCRIPTOR = _MSG1,
  __module__ = 'multiple1_pb2'
  # @@protoc_insertion_point(class_scope:twirp.internal.twirptest.multiple.Msg1)
  ))
_sym_db.RegisterMessage(Msg1)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\010multiple'))
# @@protoc_insertion_point(module_scope)
