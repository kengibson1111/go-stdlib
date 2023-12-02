# encoding/json - unmarshal (forward)

This shows how the JSON encoder is very tolerant. In this case, a component was built and deployed with an event type definition. Over time, the event type definition evolves, but the component owner has decided to continue using an older type definition.

The first JSON blob has a field not included in the current event's type definition. No error on the unmarshal. The second JSON blob is the next version of the first JSON blob - 2 versions ahead of the current event's type definition. In this case, a field is no longer used. Again,no error on the unmarshal. If the missing field is required by the component, it can return an error related to the missing data. And the component owner can decide how to handle the three different versions of JSON blobs.
