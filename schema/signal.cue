package pollyschema

// A Signal is a named, parameterized query.
//
// Each signal is expressed in exactly query language - e.g. PromQL, ANSI SQL,
// LogQL, CQL, GraphQL, etc.
//
// It is intended that return typing will be added in the future.
Signal: {
    // The name of the signal. TODO constrain with regex
	name: string
    // lang indicates the language in which the query is written.
    lang: string
    // params is the set of parameters taken by the signal.
	params: {...}
    // query is the actual query string, including parameter variables that
    // need to be interpolated.
	query: string
}