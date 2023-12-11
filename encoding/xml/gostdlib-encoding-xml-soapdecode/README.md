# encoding/xml - SOAP decode

This shows how to decode a simple SOAP envelope with a header and body. The XSI and XSD namespaces were included as attrs in the envelope, but never used. Just trying to show that the golang parser isn't going to choke on an unused namespace. At some point, you need to be able to get to individual tokens in a header or body and then apply validation rules against those tokens. This sample gets to the individual tokens, but does no validation. That would be an extension to the sample.
