# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: notes.proto
# Protobuf Python Version: 4.25.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0bnotes.proto\x12\x05notes\"\x8c\x01\n\x0cNoteDataType\x12\x0b\n\x03_id\x18\x01 \x01(\t\x12\'\n\x04type\x18\x02 \x01(\x0e\x32\x19.notes.SelectableNoteType\x12\r\n\x05title\x18\x03 \x01(\t\x12\x0f\n\x07message\x18\x04 \x01(\t\x12\x12\n\ncreated_at\x18\x05 \x01(\t\x12\x12\n\nupdated_at\x18\x06 \x01(\t*\x1e\n\x12SelectableNoteType\x12\x08\n\x04RAFT\x10\x00\x42\nZ\x08go/notesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'notes_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\010go/notes'
  _globals['_SELECTABLENOTETYPE']._serialized_start=165
  _globals['_SELECTABLENOTETYPE']._serialized_end=195
  _globals['_NOTEDATATYPE']._serialized_start=23
  _globals['_NOTEDATATYPE']._serialized_end=163
# @@protoc_insertion_point(module_scope)
