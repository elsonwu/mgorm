package mgorm

import (
	"errors"
	// "fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var session *mgo.Session
var db *mgo.Database

func InitDB(connectString, dbName string) error {
	session, err := mgo.Dial(connectString)
	if err != nil {
		return err
	}

	db = session.DB(dbName)
	return nil
}

func DB() *mgo.Database {
	return db
}

func FindAll(model IModel, criteria ICriteria) *Query {
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
	query.query = q
	return query
}

func Find(model IModel, criteria ICriteria) error {
	criteria.SetLimit(1)
	iter := FindAll(model, criteria).Iter()
	defer iter.Close()
	if iter.Next(model) {
		return nil
	}

	return errors.New("Not found")
}

func FindById(model IModel, id string) error {
	criteria := NewCriteria()
	criteria.AddCond("_id", "==", bson.ObjectIdHex(id))
	criteria.SetLimit(1)
	err := FindAll(model, criteria).Query().One(model)

	if nil == err {
		model.Init()
		model.AfterFind()
	}

	return err
}

func Update(model IModel) bool {
	if model.IsNew() {
		model.AddError("the model is a new record")
		return false
	}

	if "" == model.GetId().Hex() {
		model.AddError("the id is empty")
		return false
	}

	err := model.Collection().UpdateId(model.GetId(), model)
	if nil != err {
		model.AddError(err.Error())
		return false
	}

	return true
}

func Insert(model IModel) bool {
	if !model.IsNew() {
		model.AddError("the model is not a new record")
		return false
	}

	err := model.Collection().Insert(model)
	if nil != err {
		model.AddError(err.Error())
		return false
	}

	return true
}

func Save(model IModel) bool {

	if !model.Validate() {
		return false
	}

	err := model.BeforeSave()
	if nil != err {
		model.AddError(err.Error())
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
