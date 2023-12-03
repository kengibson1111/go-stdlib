# encoding/json - unmarshal (backward)

This also shows JSON encoder tolerance. In this case, components are still publishing version 1 events. But there are subscribers with version 2 and version 3 event type definitions. No error on the unmarshal, and the subscriber component's owner can decide if/when to return an application error for previous versions of events.
