Just a demonstration package that constructs a hierarchy of users, based on
their roles, and allows retrieval of subordinate users for a given user ID.

Note: the initially supplied data is malformed for json (roles have a trailing
comma), and malformed for Go structs (users do not have the trailing comma, as
well as the keys all being strings).
I've assumed that the intention was to supply json data, and have manually
cleaned up the roles input.

The tests can be executed with `go test -v ./...`

There is an opportunity to build a binary within cmd/ but I've not provided a
way to pass data without editing the main.go (because it's unclear what the
desire is).

The roles package has unexported methods that take structs, and exported methods
that take what I presume the intended input is. Should the input be changed it's
trivial to write other methods that transform the data into structs.
