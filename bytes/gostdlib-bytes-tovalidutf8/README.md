# bytes - tovalidutf8

ToValidUTF8() treats the input as UTF-8 encoded bytes even though some of the bytes may be invalid. For each sequence of consecutive invalid bytes, ToValidUTF8() will replace the invalid bytes with the bytes in the replacement slice.

Functionally, I don't know how much value this brings in real-world use. The invalid bytes may not conform to UTF-8 encoding, but I don't think they would be included without a reason. I don't know what the business cases would be.
