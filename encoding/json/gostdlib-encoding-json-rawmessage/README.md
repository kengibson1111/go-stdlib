# encoding/json - rawmessage (unmarshal)

This shows how to delay unmarshaling until a lower-level type is identified. Then unmarshaling continues. The example first unmarshals the entire JSON payload, but the raw message in the Color struct instructs the unmarshaler that more work is required for each instance of the Color struct.
