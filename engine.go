// Package rule_engine is the Rule Engine Implementation.
package rule_engine

import "sort"

// Engine is to analyze the fact with rules
type Engine struct {
	fact  Fact
	rules []*Rule
}

// NewEngine is to create a new Engine instance
func NewEngine() *Engine {
	return &Engine{}
}

// AppendRule is to append a rule. It can be used as "method chain" pattern
// engine.AppendRule(rule1).AppendRule(rule2)
func (eng *Engine) AppendRule(rule *Rule) *Engine {
	eng.rules = append(eng.rules, rule)
	return eng
}

func (eng *Engine) reset() {
	for _, rule := range eng.rules {
		rule.reset()
	}
}

// Analyze the fact with the rules.
func (eng *Engine) Analyze(fact Fact) {
	sort.Sort(RuleSlice(eng.rules))
	eng.reset()
	needAnalyzing := true
	for {
		if !needAnalyzing {
			return
		}
		needAnalyzing = false
		for _, rule := range eng.rules {
			if rule.processed {
				continue
			}
			if rule.ConditionFun(fact) {
				rule.processed = true
				updated := rule.PositiveAction(fact)
				if updated {
					needAnalyzing = true

				}
			}
		}
	}
}
