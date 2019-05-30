package rule_engine

import (
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func TestRuleSlice(t *testing.T) {
	Convey("test sort", t, func() {
		l := RuleSlice{}
		l = append(l, &Rule{Name: "1", Priority: 1})
		l = append(l, &Rule{Name: "3", Priority: 3})
		l = append(l, &Rule{Name: "4", Priority: 4})
		l = append(l, &Rule{Name: "2", Priority: 2})

		sort.Sort(l)
		So(l[0].Priority, ShouldEqual, 4)
		So(l[1].Priority, ShouldEqual, 3)
		So(l[2].Priority, ShouldEqual, 2)
		So(l[3].Priority, ShouldEqual, 1)
	})
}
