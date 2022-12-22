package monkey_math

import (
  "fmt"
  "math"
  "regexp"

  "AdityaHegde/adventofcode/utils"
)

type expression struct {
  variable  string
  operation rune
  left      *expression
  right     *expression
  value     float64
  hasValue  bool
  parent    *expression
}

func (e *expression) solve() {
  if e.hasValue || e.left == nil || !e.left.hasValue || e.right == nil || !e.right.hasValue {
    return
  }
  switch e.operation {
  case '+':
    e.value = e.left.value + e.right.value
  case '-':
    e.value = e.left.value - e.right.value
  case '*':
    e.value = e.left.value * e.right.value
  case '/':
    e.value = e.left.value / e.right.value
  }
  e.hasValue = true
  return
}

func (e *expression) fillIn() *expression {
  if e.left == nil || e.right == nil {
    return nil
  }
  known := e.left
  unknown := e.right
  if !e.left.hasValue {
    unknown = e.left
    known = e.right
  }
  switch e.operation {
  case '+':
    unknown.value = e.value - known.value
  case '-':
    unknown.value = e.value + known.value
  case '*':
    unknown.value = e.value / known.value
  case '/':
    unknown.value = e.value * known.value
  }
  return unknown
}

func (e *expression) print() string {
  if e.left == nil && e.right == nil {
    return "x"
  }
  leftStr := ""
  if e.left.hasValue {
    leftStr = fmt.Sprintf("%f", e.left.value)
  } else {
    leftStr = e.left.print()
  }
  oprnStr := ""
  switch e.operation {
  case '+':
    oprnStr = "+"
  case '-':
    oprnStr = "-"
  case '*':
    oprnStr = "*"
  case '/':
    oprnStr = "/"
  }
  rightStr := ""
  if e.right.hasValue {
    rightStr = fmt.Sprintf("%f", e.right.value)
  } else {
    rightStr = e.right.print()
  }
  return fmt.Sprintf("(%s)%s(%s)", leftStr, oprnStr, rightStr)
}

type problem struct {
  expressions map[string]*expression
  root        *expression
  ignore      bool
}

func newProblem(ignore bool) *problem {
  return &problem{
    expressions: map[string]*expression{},
    ignore:      ignore,
  }
}

var expressionParseRegex = regexp.MustCompile(`(\w*): (?:(\d*)|(\w*) (.) (\w*))$`)

func (p *problem) getOrCreate(variable string, parent *expression) *expression {
  expr, ok := p.expressions[variable]
  if ok {
    if parent != nil {
      expr.parent = parent
    }
    return expr
  }
  e := &expression{
    variable: variable,
    parent:   parent,
  }
  p.expressions[variable] = e
  return e
}

func (p *problem) parseLine(line string) {
  parts := expressionParseRegex.FindStringSubmatch(line)
  lhsExpr := p.getOrCreate(parts[1], nil)
  if parts[1] == "root" {
    p.root = lhsExpr
  } else if parts[1] == "humn" && p.ignore {
    return
  }

  if len(parts[4]) == 0 {
    lhsExpr.value = float64(utils.Int(parts[2]))
    lhsExpr.hasValue = true
  } else {
    lhsExpr.left = p.getOrCreate(parts[3], lhsExpr)
    lhsExpr.operation = rune(parts[4][0])
    lhsExpr.right = p.getOrCreate(parts[5], lhsExpr)
    lhsExpr.solve()
  }

  for lhsExpr.hasValue && lhsExpr.parent != nil {
    lhsExpr = lhsExpr.parent
    lhsExpr.solve()
  }
}

func (p *problem) parse(lines []string) {
  for _, line := range lines {
    p.parseLine(line)
  }
}

func partOne(lines []string) int {
  p := newProblem(false)
  p.parse(lines)
  return int(math.Round(p.root.value))
}

func partTwo(lines []string) int {
  p := newProblem(true)
  p.parse(lines)

  // select the starting sub tree
  start := p.root.left
  startVal := p.root.right.value
  if !p.root.right.hasValue {
    start = p.root.right
    startVal = p.root.left.value
  }
  expr := start
  val := startVal
  expr.value = val

  for expr != nil {
    expr = expr.fillIn()
    if expr != nil {
      val = expr.value
    }
  }

  fmt.Printf("%f=%s\n", startVal, start.print())

  return int(math.Round(val))
}
