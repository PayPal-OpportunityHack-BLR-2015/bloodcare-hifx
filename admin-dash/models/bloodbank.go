package models

import (
	"fmt"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
)

//BloodBank model
type BloodBank struct {
	Id   int
	Name string
	// Add Fields here
}

func (a *BloodBank) String() string {
	return fmt.Sprintf("id:%s", a.Id)
}

func BloodBanks(page int, cs *services.Cassandra) (*BloodBank, *app.Msg, error) {
	const (
		BBANK_FETCH_SQL = "SELECT *  FROM bloodbanks"
	)

	return &BloodBank{}, nil, nil
}

func InsertBloodBank(name string, cs *services.Cassandra) (string, *app.Msg) {
	const (
		BBANK_INSERT_SQL = "INSERT INTO bloodbanks (userId, id, filepath, stage) VALUES (?, ?, ?, ?)"
	)
	if err := cs.Query(BBANK_INSERT_SQL, name).Exec(); err != nil {
		return "", app.NewErrMsg(err.Error())
	}
	return id, nil
}

func UpdateBloodBank(id int, name string, cs *services.Cassandra) (string, *app.Msg) {
	const (
		BBANK_UPDATE_SQL = ""
	)
	if err := cs.Query(BBANK_UPDATE_SQL, id, name).Exec(); err != nil {
		return "", app.NewErrMsg(err.Error())
	}
	return id, nil
}

func FetchBloodBankDetails(jobId, userId string, cs *services.Cassandra) (string, *app.Msg) {
	const (
		JOB_FETCH_SQL = "SELECT (stage) FROM bloodcare_jobs WHERE id=? AND userId=?"
	)
	var stage string
	iter := cs.Query(JOB_FETCH_SQL, jobId, userId).Iter()

	ok := iter.Scan(&stage)
	if !ok {
		return "", app.NewErrMsg("Invalid Job Id")
	}

	return stage, nil
}
