import "encoding/json"

// TODO: include full schema?
// TODO: figure out why it stops working when uncommented (cue eval works tho)
// prometheusAlerts: v0: {...}
// prometheusRules: v0: {...}
// grafanaDashboards: v0: {...}

// TODO: handle different versions, deal with migrations, etc, etc
mixin: {
    // TODO: change to ruleConvert() once syntax sugar is available
    _alerts: #ruleConvert&{_, #rules: prometheusAlerts.v0}
    _rules: #ruleConvert&{_, #rules: prometheusRules.v0}
    _dashboards: #dashboardConvert&{#dashboards: grafanaDashboards.v0}

    {
        prometheusAlerts: _alerts
        prometheusRules: _rules
        grafanaDashboards: _dashboards
    }
}

#ruleConvert: {
    #rules: [string]: {group: string, alert: {...}, ...}

    [for n, r in (#flattenAlerts & {arg: #rules}).out {
        name: n, rules: r
    }]
}

#dashboardConvert: {
    #dashboards: [string]: _

    for k, v in #dashboards {
        "\(k).json": json.Indent(json.Marshal(v), "", "  ")
    }
}

#flattenAlerts: {
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