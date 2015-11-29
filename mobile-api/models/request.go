package models

import (
	"strconv"
	"time"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/services"
)

type BloodRequest struct {
	ReqId       string `json:"req_id,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	Date        string `json:"date,omitempty"`
	Blood       string `json:"blood,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Description string `json:"description,omitempty"`
	Lat         string `json:"lat,omitempty"`
	Lng         string `json:"lng,omitempty"`
	PlaceId     string `json:"place_id,omitempty"`
}
type BloodRequests []*BloodRequest

func (r *BloodRequest) String() string {
	return r.UserId
}

func (r *BloodRequest) ValidateBloodReq() (bool, string) {
	if len(r.UserId) == 0 {
		return false, "UserId is empty"
	}
	if len(r.Date) == 0 {
		return false, "Date is empty"
	}
	if IsValidBloodType(r.Blood) == false {
		return false, "Invalid blood type"
	}
	if len(r.Phone) == 0 {
		return false, "Phone is empty"
	}
	if len(r.Lat) == 0 {
		return false, "Lat is empty"
	}
	if len(r.Lng) == 0 {
		return false, "Lng is empty"
	}

	return true, ""
}

func (r *BloodRequest) ValidateDeleteReq() (bool, string) {
	if len(r.ReqId) == 0 {
		return false, "ReqId is empty"
	}
	if len(r.UserId) == 0 {
		return false, "UserId is empty"
	}
	return true, ""
}

func (r *BloodRequest) ValidateGetReq() (bool, string) {
	if len(r.ReqId) == 0 {
		return false, "ReqId is empty"
	}
	return true, ""
}

func CreateBloodRequest(r *BloodRequest, ms *services.MySQL) (bool, *app.Msg, error) {

	status, error := r.ValidateBloodReq()
	if status == false {
		return false, app.NewErrMsg(error), nil
	}
	query := "INSERT INTO requests(user_id, date_of_requirement, location, place_id, blood, comments, mobile) VALUES(?,?,POINT(" + r.Lat + ", " + r.Lng + "),?,?,?,?)"
	const createdFormat = "2006-01-02 15:04:05" //"Jan 2, 2006 at 3:04pm (MST)"

	utcTime, _ := strconv.ParseInt(r.Date, 10, 64)
	utcTime1 := time.Unix(utcTime, 0).Format(createdFormat)
	_, dbError := ms.Exec(query, r.UserId, utcTime1, r.PlaceId, r.Blood, r.Description, r.Phone)
	if dbError != nil {
		return false, nil, dbError
	}
	return false, app.NewErrMsg(error), nil
}

func DeleteBloodRequest(r *BloodRequest, ms *services.MySQL) (bool, *app.Msg, error) {

	status, error := r.ValidateDeleteReq()
	if status == false {
		return false, app.NewErrMsg(error), nil
	}
	query := "DELETE FROM requests WHERE id = ,?)"
	_, dbError := ms.Exec(query, r.ReqId)
	if dbError != nil {
		return false, nil, dbError
	}
	return false, app.NewErrMsg(error), nil
}

func GetBloodRequest(r *BloodRequest, ms *services.MySQL) (bool, *BloodRequest, error) {

	// status, _ := r.ValidateGetReq()
	// if status == false {
	// 	return nil, nil //TODO return validation err
	// }
	query := "SELECT * FROM requests WHERE id = , ?)"
	dbRows, dbError := ms.Query(query, r.ReqId)
	if dbError == nil {
		return false, r, dbError
	}
	defer dbRows.Close()

	if dbRows.Next() {
		var (
			result      BloodRequest
			reqId       string
			userId      string
			date        string
			blood       string
			description string
			placeId     string
		)
		dbRows.Scan(&reqId, &userId, &date, &blood, &description, &placeId)
		result = BloodRequest{
			ReqId:       reqId,
			UserId:      userId,
			Date:        date,
			Blood:       blood,
			Description: description,
			PlaceId:     placeId,
		}
		// ReqId UserId Date Blood Phone Description Lat Lng PlaceId
		// dbRows.Scan(r.ReqId, r.UserId, r.Date, r.Blood, r.Description, r.PlaceId)
		return true, &result, nil
	}
	return false, r, nil
}
