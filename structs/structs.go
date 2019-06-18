package structs

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Credential struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}

type OrderDash struct {
	OrderDate     time.Time
	OrderApproved int
	OrderRejected int
	OrderPending  int
}

type OrderDashAll struct {
	OrderStatus int
	OrderDash   struct {
		OrderDate   time.Time
		OrderStatus int
		OrderTotal  int
	}
}

type TbOrder struct {
	Order_Id                string
	Product_Id              int
	Order_Customer          string
	Order_Customer_Phone    string
	Order_Start_Date        time.Time
	Order_Finish_Date       time.Time
	Order_Price             int
	Order_Down_Payment      int
	Order_Remaining_Payment int
	Order_Pick_Up           string
	Order_Destination       string
	Order_Status            int
	Order_Notes             string
	Created_At              time.Time
	Updated_At              time.Time
	Created_By              string
	Updated_By              string
}
