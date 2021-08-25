package prometheus

import "list"

// #FlattenAlerts takes a map of Prometheus alerts as they're wrapped in the
// PollyPackage spec and converts them to the form Prometheus expects: separate lists
// of alerts, keyed by their group.
//
// An input like this:
//   {
//     Alert1: { group: "foo", alert: { alert: Alert1, ...}}
//     Alert2: { group: "bar", alert: { alert: Alert2, ...}}
//     Alert3: { group: "foo", alert: { alert: Alert3, ...}}
//   }
//
// Is converted into this:
//   {
//     foo: [{ alert: Alert1, ... }, { alert: Alert3, ... }]
//     bar: [{ alert: Alert2, ... }]
//   }
#FlattenAlerts: {
    arg: [string]: { group: string, alert: {...} }
    out: [string]: [...]
    _inter: [string]: [string]: {...}
    for n, v in arg {
        _inter: "\(v.group)": "\(n)": v.alert
    }
    for g, a in _inter {
        out: "\(g)": [ for v in a {v}, ...]
    }
}

#LabelMatcher: {
	Label: string
	Value: string
	Op: "=" | "!=" | "=~" | "!=~"
    str: "\(Label)\(Op)\"\(Value)\""
}

#LabelMatcherExact: {
	Label: string
	Value: string
	Op: *"=" | "!="
    str Label + Op + Value
}

#LabelMatcherDisjunct: {
	Label: string
	Value: [string, string, ...string]
	Op: *"=~" | "!=~"
    str: Label + Op + strings.Join(Value, "|")
}

#LabelMatcherRegex: {
	Label: string
	Value: string
	Op: *"=~" | "!=~"
    str: Label + Op + Value
}

#LabelMatcherGroup: [L=string]: #LabelMatcher & {
	Label: L
}