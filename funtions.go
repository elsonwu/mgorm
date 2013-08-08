package model

import (
	"errors"
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

func Update(model IModel) error {
	if model.IsNew() {
		return errors.New("the model is a new record")
	}

	if "" == model.GetId().Hex() {
		return errors.New("the id is empty")
	}

	return model.Collection().UpdateId(model.GetId(), model)
}

func Insert(model IModel) error {
	if !model.IsNew() {
		return errors.New("the model is not a new record")
	}

	return model.Collection().Insert(model)
}

func Save(model IModel) error {
	if model.IsNew() {
		return Insert(model)
	} else {
		return Update(model)
	}
}
