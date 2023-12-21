package db

import (
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/qustavo/dotsql"

	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
)

func trackerTypeBinder(r *model.TrackerType) []any {
	return []any{&r.Id, &r.Tier, &r.ContextEligible, &r.Key}
}

func companyBinder(r *model.Company) []any {
	return []any{&r.Id, &r.Name}
}

func titleBinder(r *model.Title) []any {
	return []any{&r.Id, &r.CompanyId, &r.Title, &r.Level}
}

func FindTrackerTypes(handle DbHandle) ([]model.TrackerType, error) {
	dot, err := dotsql.LoadFromFile("sql/search.sql")
	if err != nil {
		log.Fatal("Could not load search.sql", err)
		return nil, errors.New("Could not load search.sql for finding all tracker types")
	}
	err, rtn := NamedQuery[model.TrackerType](handle, dot, trackerTypeBinder, "trackertype:list")

	if err != nil {
		log.Fatal("Could not query database", err)
		return nil, err
	}

	return rtn, nil
}

func FindContextEligibleTrackerTypes(handle DbHandle) ([]model.TrackerType, error) {
	dot, err := dotsql.LoadFromFile("sql/search.sql")
	if err != nil {
		log.Fatal("Could not load search.sql", err)
		return nil, errors.New("Could not load search.sql for finding all tracker types")
	}
	err, rtn := NamedQuery[model.TrackerType](handle, dot, trackerTypeBinder, "trackertype:contexteligible", true)

	if err != nil {
		log.Fatal("Could not query database", err)
		return nil, err
	}

	return rtn, nil
}

func FindCompanyById(handle DbHandle, companyId int) (*model.Company, error) {
	dot, err := dotsql.LoadFromFile("sql/search.sql")
	if err != nil {
		return nil, err
	}
	err, rtn := NamedQuery[model.Company](handle, dot, companyBinder, "company:byid", companyId)

	if len(rtn) != 1 {
		return nil, err
	}
	if err != nil {
		log.Fatal("Could not query database", err)
		return nil, err
	}
	return &rtn[0], nil
}

func FindCompanyByName(handle DbHandle, companyName string) ([]model.Company, error) {
	dot, err := dotsql.LoadFromFile("sql/search.sql")
	if err != nil {
		return nil, err
	}
	err, rtn := NamedQuery[model.Company](handle, dot, companyBinder, "company:byname", companyName)

	if err != nil {
		log.Fatal("Could not query database", err)
		return nil, err
	}

	return rtn, nil
}

func FindTitleById(handle DbHandle, titleId int) (*model.Title, error) {
	dot, err := dotsql.LoadFromFile("sql/search.sql")
	if err != nil {
		return nil, err
	}
	err, rtn := NamedQuery[model.Title](handle, dot, titleBinder, "title:byid", titleId)

	if len(rtn) != 1 {
		return nil, err
	}
	if err != nil {
		log.Fatal("Could not query database", err)
		return nil, err
	}

	return &rtn[0], nil
}

func FindTitleByName(handle DbHandle, title string, companyId int) ([]model.Title, error) {
	dot, err := dotsql.LoadFromFile("sql/search.sql")
	if err != nil {
		return nil, err
	}
	err, rtn := NamedQuery[model.Title](handle, dot, titleBinder, "title:byname", title, companyId)

	if err != nil {
		log.Fatal("Could not query database", err)
		return nil, err
	}

	return rtn, nil
}

func NamedQuery[T any](
	handle DbHandle,
	dot *dotsql.DotSql,
	binder func(*T) []any,
	query string,
	args ...any) (error, []T) {
	rows, err := dot.Query(handle.Db(), query, args...)
	if err != nil {
		return err, nil
	}
	defer func() {
		_ = rows.Close()
	}()
	var results []T
	for rows.Next() {
		var result T
		cols := binder(&result)
		err = rows.Scan(cols...)
		if err != nil {
			return err, nil
		}
		results = append(results, result)
	}
	return nil, results
}
