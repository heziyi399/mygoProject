package dao

import "fmt"

type Condition struct {
	where     []string
	args      []interface{}
	limit     string
	limitArgs []interface{}
}

func NewCondition() *Condition {
	return &Condition{
		where:     make([]string, 0, 0),
		args:      make([]interface{}, 0, 0),
		limitArgs: make([]interface{}, 0, 0),
	}
}
func (this *Condition) Eq(column string, val interface{}) *Condition {
	this.where = append(this.where, fmt.Sprintf("%s=?", column))
	this.args = append(this.args, val)
	return this
}
func (this *Condition) Like(column string, val interface{}) *Condition {
	this.where = append(this.where, fmt.Sprintf("%s like?", column))
	this.args = append(this.args, val)
	return this
}
