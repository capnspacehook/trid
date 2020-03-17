# trid

## Totally Random Identifiers

trid is a GUID library built for speed, size, and anonymity. 

Cryptographically secure random GUIDs can be generated at extremely high speeds and very low memory usage:

```
BenchmarkNew-4                   6769496               174 ns/op              16 B/op          1 allocs/op
BenchmarkToString-4             33261237                36.9 ns/op            32 B/op          1 allocs/op
BenchmarkFromBytes-4            478937239                2.48 ns/op            0 B/op          0 allocs/op
BenchmarkFromString-4           44093491                26.2 ns/op             0 B/op          0 allocs/op
```

GUIDs are 12 bytes in length raw, and 20 characters in length encoded as strings. Base32 encoding is used
so that the non alphanumeric characters of base64 don't cause a problem.

trid was heavily inspired by `github.com/rs/xid`, much of the code is taken from there. This repository was
necessary for me as I do not want PIDs, timestamps, or machine identifiers in my GUIDs, and I don't care if
my GUIDs are sortable.
