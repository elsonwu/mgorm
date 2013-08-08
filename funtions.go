package mgorm

import (
	// "errors"
	// "fmt"
	// "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func FindAll(model IModel, criteria *Criteria) *Query {
	if !model.HasInited() {
		model.Init()
	}

	q := model.Collection().Find(criteria.GetConditions())
	criteriaSelects := criteria.GetSelect()
	if 0 < len(criteriaSelects) {
		selects := map[string]bool{}
		for _, field := range criteriaSelects {
			selects[field] = true
		}

		q.Select(selects)
	}

	if 0 < criteria.GetLimit() {
		q.Limit(criteria.GetLimit())
	}

	if 0 < criteria.GetOffset() {
		q.Skip(criteria.GetOffset())
	}

	if nil != criteria.GetSort() {
		sort := criteria.GetSort()
		sortStr := []string{}
		for key, value := range sort {
			if 0 < value {
				sortStr = append(sortStr, key)
			} else {
				sortStr = append(sortStr, "-"+key)
			}
		}

		q.Sort(sortStr...)
	}

	query := new(Query)
	query.SetQuery(q)
	return query
}

func FindById(model IModel, id string) error {
	criteria := NewCriteria()
	criteria.AddCond("_id", "==", bson.ObjectIdHex(id))
	criteria.SetLimit(1)
	err := FindAll(model, criteria).GetQuery().One(model)

	if nil == err {
		model.Init()
		model.AfterFind()
	}

	return err
}

func Update(model IModel) bool {
	if model.IsNew() {
		model.ErrorHandler().AddError("the model is a new record")
		return false
	}

	if "" == model.GetId().Hex() {
		model.ErrorHandler().AddError("the id is empty")
		return false
	}

	err := model.Collection().UpdateId(model.GetId(), model)
	if nil != err {
		model.ErrorHandler().AddError(err.Error())
		return false
	}

	return true
}

func Insert(model IModel) bool {
	if !model.IsNew() {
		model.ErrorHandler().AddError("the model is not a new record")
		return false
	}

	err := model.Collection().Insert(model)
	if nil != err {
		model.ErrorHandler().AddError(err.Error())
		return false
	}

	return true
}

func Save(model IModel) bool {

	if !model.BeforeSave() {
		return false
	}

	res := false
	if model.IsNew() {
		res = Insert(model)
	} else {
		res = Update(model)
	}

	if res {
		model.AfterSave()
	}

	return res
}
