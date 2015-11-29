package models

import (
	"fmt"

	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/app"
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/mobile-api/services"
)

type BloodRequest struct {
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

func CreateBloodRequest(r *BloodRequest, ms *services.MySQL) (bool, *app.Msg, error) {

	status, error := r.ValidateBloodReq()
	fmt.Println(error)
	if status == false {
		return false, app.NewErrMsg(error), nil
	}
	query := "INSERT INTO requests(user_id, date_of_requirement, location, place_id, blood, comments, mobile) VALUES(?,?,POINT(" + r.Lat + ", " + r.Lng + "),?,?,?,?)"
	fmt.Println(query)
	_, dbError := ms.Exec(query, r.UserId, r.Date, r.PlaceId, r.Blood, r.Description, r.Phone)
	if dbError != nil {
		return false, nil, dbError
	}
	return false, app.NewErrMsg(error), nil

}
