package controller

import (
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (idb *InDB) GetDashboard(c *gin.Context) {

	type DetailAtrray struct {
		OrderTotal int
	}

	type allArray struct {
		DashboardLabel int
		DetailAtrray   []DetailAtrray
	}
	var (
		DashboardLabel       []time.Time
		DashboardArray       []structs.OrderDash
		DashboardArraySingle []structs.OrderDash
		//DashboardArrayAll []allArray

		Dashboard structs.OrderDash
		result    gin.H
	)

	orderDash, err := idb.DB.Table("tb_orders").Select("order_start_date,order_status,count(*) as total").Group("order_status,date(order_start_date)").Order("order_status,order_start_date asc").Rows()
	orderDashLabel, err := idb.DB.Table("tb_orders").Select("order_start_date").Group("date(order_start_date)").Order("order_start_date asc").Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
	}

	for orderDashLabel.Next() {
		err := orderDashLabel.Scan(&Dashboard.OrderDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		} else {
			DashboardLabel = append(DashboardLabel, Dashboard.OrderDate)
		}
	}

	for orderDash.Next() {
		err := orderDash.Scan(&Dashboard.OrderDate, &Dashboard.OrderStatus, &Dashboard.OrderTotal)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
		} else {
			DashboardArray = append(DashboardArray, Dashboard)
		}

	}

	for i, num := range DashboardArray {
		if i >= 1 {
			LastStatus := DashboardArray[i-1].OrderStatus
			CurrentStatus := num.OrderStatus
			if LastStatus == CurrentStatus {
				DashboardArraySingle = append(DashboardArraySingle, num)
			} else {

			}
		} else {
			state := allArray{
				DashboardLabel: DashboardArray[i-1].OrderStatus,
				DetailAtrray: []DetailAtrray{
					{1},
				},
			}
			fmt.Println(state)
			DashboardArraySingle = nil
			DashboardArraySingle = append(DashboardArraySingle, num)
			//DashboardArrayAll = append(DashboardArrayAll, state)

		}
	}

	result = gin.H{
		"label":        DashboardLabel,
		"result":       DashboardArray,
		"resultSingle": DashboardArraySingle,
		"count":        len(DashboardArray),
	}
	c.JSON(http.StatusOK, result)
}
