package rule_engine

import (
	"fmt"
	"testing"
)

type Person struct {
	Name  string
	Age   int
	Role  string
	Tasks []string
}

func TestCase1(t *testing.T) {
	fact1 := &Person{
		"Mike", 7, "", []string{},
	}
	rule1 := NewRule("Age>=6").
		When(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			return p.Age >= 6
		}).
		Then(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			p.Role = "Student"
			return true
		})
	rule2 := NewRule("is student").
		When(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			return p.Role == "Student"
		}).
		Then(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			p.Tasks = append(p.Tasks, "Go to school")
			return true
		})

	rule3 := NewRule("School").
		When(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			for _, task := range p.Tasks {
				if task == "Go to school" {
					return true
				}
			}
			return false
		}).
		Then(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			p.Tasks = append(p.Tasks, "Learning")
			return true
		}).WithOrder(10)
	engine := NewEngine()
	engine.AppendRule(rule3).
		AppendRule(rule1).
		AppendRule(rule2).
		Analyze(fact1)
	fmt.Println(*fact1)
	if fact1.Role != "Student" {
		t.Error("Failed to analyze")
	}
	if len(fact1.Tasks) < 2 {
		t.Error("Failed to analyze")
	}
}

func TestCase2(t *testing.T) {
	fact1 := &Person{
		"Mike", 4, "", []string{},
	}
	rule1 := NewRule("Age>=6").
		When(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			return p.Age >= 6
		}).
		Then(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			p.Role = "Student"
			return true
		})
	rule2 := NewRule("is student").
		When(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			return p.Role == "Student"
		}).
		Then(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			p.Tasks = append(p.Tasks, "Go to school")
			return true
		})

	rule3 := NewRule("School").
		When(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			for _, task := range p.Tasks {
				if task == "Go to school" {
					return true
				}
			}
			return false
		}).
		Then(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			p.Tasks = append(p.Tasks, "Learning")
			return true
		}).WithOrder(10)
	engine := NewEngine()
	engine.AppendRule(rule3).
		AppendRule(rule1).
		AppendRule(rule2).
		Analyze(fact1)
	fmt.Println(*fact1)
	if fact1.Role == "Student" {
		t.Error("Failed to analyze")
	}
	if len(fact1.Tasks) > 0 {
		t.Error("Failed to analyze")
	}
}

func TestPriority(t *testing.T) {
	fact1 := &Person{
		"Mike", 7, "", []string{},
	}
	rule1 := NewRule("task1").
		When(func(f Fact) bool {
			return true
		}).
		Then(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			p.Tasks = append(p.Tasks, "task1")
			return true
		})

	rule2 := NewRule("task2").
		When(func(f Fact) bool {
			return true
		}).
		Then(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			p.Tasks = append(p.Tasks, "task2")
			return true
		}).WithOrder(100)

	rule3 := NewRule("task3").
		When(func(f Fact) bool {
			return true
		}).
		Then(func(f Fact) bool {
			p, ok := f.(*Person)
			if !ok {
				t.Fatal("fact should be a person.")
			}
			p.Tasks = append(p.Tasks, "task3")
			return true
		}).WithOrder(10)
	engine := NewEngine()
	engine.AppendRule(rule3).
		AppendRule(rule1).
		AppendRule(rule2).
		Analyze(fact1)
	fmt.Println(*fact1)
	if fact1.Tasks[1] != "task3" {
		t.Fatal("failed to set the priority")
	}
}
