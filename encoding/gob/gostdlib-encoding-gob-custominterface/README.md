# encoding/gob - custominterface

The custom type would be registered in the encoder and decoder gob processes. The type implements a method on a custom interface. The interface is available in the encoder and decoder processes. The encoder encodes through the interface. The decoder decodes through the interface and calls the interface's method.
