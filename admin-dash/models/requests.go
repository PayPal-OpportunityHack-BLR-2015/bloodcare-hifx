package models

import (
	"fmt"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/services"
	"github.com/paulmach/go.geo"
)

//Request model
type Request struct {
	ID       string
	UserID 	 string
	UserName string
	Date string
	Location *geo.Point
	Place string
	Blood string
	Comments string
	Mobile string
	Created string
}
type Requests []*Request

func (r *Request) String() string {
	//TODO: Fix :(
	return fmt.Sprintf("id:%s uid:% uname:% date:% loc:% place:% blood:% comments:% mobile:% created:% ",  r.ID, r.UserID, r.UserName, r.Date, r.Location, r.Place, r.Blood, r.Comments, r.Mobile, r.Created)
}

func ListRequests(db *services.MySQL) (*Requests, error) {
	const (
		REQUEST_LIST_SQL = "SELECT requests.id, users.id, users.name, date_of_requirement, requests.location, requests.blood, requests.comments, requests.mobile, requests.created FROM `requests` JOIN users ON requests.user_id = users.id"
	)

	var (
		results  Requests
		id       string
		userId 	string
		userName string
		name     string
		date	 string
		location *geo.Point
		blood    string
		comments   string
		mobile   string
		created  string
	)
	rows, err := db.Query(REQUEST_LIST_SQL)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&id, &userId,  &name, &date, &location, &blood, &comments, &mobile, &created)
		results = append(results,
			&Request{
				ID:id,
				UserID: userId,
				UserName: name,
				Date: date,
				Location: location,
				Blood:blood,
				Comments:comments,
				Mobile:mobile,
				Created:created,
			})

		fmt.Println(Request{
			ID:id,
			UserID: userId,
			UserName: userName,
			Date: date,
			Location: location,
			Blood:blood,
			Comments:comments,
			Mobile:mobile,
			Created:created,
		})
	}
	return &results, nil

}

func FetchRequestDetails(id int, db *services.MySQL) (string, *app.Msg) {
	const (
		JOB_FETCH_SQL = "SELECT (stage) FROM bloodcare_jobs WHERE id=? AND userId=?"
	)
	return "", nil
}
