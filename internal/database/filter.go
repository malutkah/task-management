package database

const (
	ConditionEqual   = "="
	ConditionUnequal = "!="
	ConditionGreater = ">"
	ConditionLess    = "<"
)

func (c Condition) Eval(conditionValue any) bool {
	if c.Equals == ConditionEqual {
		return c.Value == conditionValue
	} else if c.Equals == ConditionUnequal {
		return c.Value != conditionValue
	} else if c.Equals == ConditionGreater {
		switch colValue := c.Value.(type) {
		case int:
			if conditionValueInt, ok := conditionValue.(int); ok {
				return conditionValueInt > colValue
			}
		case float64:
			if conditionValueFloat, ok := conditionValue.(float64); ok {
				return conditionValueFloat > colValue
			}
		}
	} else if c.Equals == ConditionLess {
		switch v := c.Value.(type) {
		case int:
			if valInt, ok := conditionValue.(int); ok {
				return valInt < v
			}
		case float64:
			if valFloat, ok := conditionValue.(float64); ok {
				return valFloat < v
			}
		}
	}
	return false
}

func (t *Table) filterColumns(cols []Expression, i int) []any {
	row := t.rows[i]
	
	if cols[0].Column == "*" {
		return append([]any(nil), row...)
	}
	res := make([]any, len(cols))
	
	for i, col := range cols {
		res[i] = row[t.getColumnIndex(col.Column)]
	}
	
	return res
}

func (t *Table) rowMatch(condition []Condition, i int) bool {
	if t.rows[i] == nil {
		return false
	}
	
	if condition == nil {
		return true
	}
	
	for _, c := range condition {
		j := t.getColumnIndex(c.Column)
		if !c.Eval(t.rows[i][j]) {
			return false
		}
	}
	return true
}
