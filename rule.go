package rule_engine

const (
	DefaultOrderValue = 5
)

// Fact object
type Fact interface{}

// ConditionFunc is the condition section of the rule,
// which is a function of checking if the fact matches the rule.
type ConditionFunc func(fact Fact) bool

// ActionFunc is the part of a rule.
// The function would be executed when the fact match the associated rule.
// The return value represents if the fact has been updated in the action.
type ActionFunc func(fact Fact) bool

// Rule is to define the rules of rule engine.
type Rule struct {
	Name           string
	ConditionFun   ConditionFunc
	PositiveAction ActionFunc
	Order          int // The higher priority rule would be checked first.
	processed      bool
}

// NewRule is to create a new rule (can be used as "method chain" style)
// See also, the tests.
func NewRule(name string) *Rule {
	return &Rule{
		Name:  name,
		Order: DefaultOrderValue,
	}
}

func (r *Rule) reset() {
	r.processed = false
}

// When is to define the condition part of a rule.
func (r *Rule) When(cond ConditionFunc) *Rule {
	r.ConditionFun = cond
	return r
}

// Then is to define the action part of a rule
func (r *Rule) Then(action ActionFunc) *Rule {
	r.PositiveAction = action
	return r
}

// WithOrder is to set execution order of the rule
// The smaller value would be executed first.
func (r *Rule) WithOrder(order int) *Rule {
	r.Order = order
	return r
}

// RuleSlice is used to sort the rules with order by implemented sortable interface.
type RuleSlice []*Rule

func (rs RuleSlice) Len() int {
	return len(rs)
}

func (rs RuleSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs RuleSlice) Less(i, j int) bool {
	return rs[i].Order < rs[j].Order
}
