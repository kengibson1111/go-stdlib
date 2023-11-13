# encoding/gob - basic

This shows how to use basic functionality. In this sample, you have to imagine one type being available in one process for encoding and another version of the same type being available in another process. Bytes are streamed and even though the encoded type is different from the decoded type, the Decoder gets as much info as is available without an error. Like protobufs, but it is built into golang.
