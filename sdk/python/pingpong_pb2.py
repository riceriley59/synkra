# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: pingpong.proto
# Protobuf Python Version: 6.31.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    6,
    31,
    0,
    '',
    'pingpong.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0epingpong.proto\x12\x08pingpong\"\x1e\n\x0bPingRequest\x12\x0f\n\x07message\x18\x01 \x01(\t\"\x1c\n\tPongReply\x12\x0f\n\x07message\x18\x01 \x01(\t2@\n\x08PingPong\x12\x34\n\x04Ping\x12\x15.pingpong.PingRequest\x1a\x13.pingpong.PongReply\"\x00\x42\x30Z.github.com/riceriley59/synkra/proto;pingpongpbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'pingpong_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z.github.com/riceriley59/synkra/proto;pingpongpb'
  _globals['_PINGREQUEST']._serialized_start=28
  _globals['_PINGREQUEST']._serialized_end=58
  _globals['_PONGREPLY']._serialized_start=60
  _globals['_PONGREPLY']._serialized_end=88
  _globals['_PINGPONG']._serialized_start=90
  _globals['_PINGPONG']._serialized_end=154
# @@protoc_insertion_point(module_scope)
