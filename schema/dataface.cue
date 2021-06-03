package pollyschema

// A Dataface defines a pattern over a set of signals (queries).
// 
// In conventional programming, an interface definition specifies a collection
// of names and types or function signatures that a type must include in order
// to implement that interface. This allows implementations and consumers to
// "cooperate" via the contract of the interface, without needing to know
// the other exists.
//
// A dataface is a similar kind of contract, mediating between its implementors
// and consumers. Datafaces specify names and signal signatures, and optionally
// properties of the signal query results that must align. Polly packages
// implement a dataface by assigning their declared signals to a dataface
// instance. Consumers then rely only on the dataface, not a particular
// implementation.
// 
// Specifying a dataface is a bit like an abstract SQL join. The dataface
// definition specifies a larger dataset, composed from multiple smaller
// datasets (the signal query results), which relate along some dimension of
// their contents; it is the job of the implementations to pick a relation that
// fulfills the intended semantics.
//
// TODO formalize the logical function of datafaces; draw on relational algebra
// TODO signals have no "signature" until they're well-typed; constraints
// are also tough
// TODO elaborate on exactly how it's expected that datafaces get consumed.
// Necessarily an extension of how signals get consumed.
Dataface: {
    name: string
    // TODO regex-constrain frame name to alphanumeric
    // TODO this is where additional type constraints on the signal would be expressed
    frames: [string]: Signal 
}

// TODO just having dataface specifications - in contrast to implementations -
// be free-hanging like this is antithetical to the key goal of referencing them
// in polly pkg implementations
RED: Dataface & {
    name: "RED"
    frames: {
        requests: Signal
        errors: Signal
        duration: Signal
    }
}

USE: Dataface & {
    name: "USE"
    frames: {
        utilization: Signal
        saturation: Signal
        errors: Signal
    }
}