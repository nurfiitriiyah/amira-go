package controller

import (
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (idb *InDB) GetDashboard(c *gin.Context) {

	var (
		DashboardLabel           []time.Time
		DashboardArrayCountTotal []int

		DashboardArrayApproved []int
		DashboardArrayRejected []int
		DashboardArrayPending  []int

		//DashboardArrayCount int
		Dashboard structs.OrderDash
		result    gin.H
	)

	type DashDetailGraph struct {
		Approve []int
		Reject  []int
		Pending []int
	}

	type DashAll struct {
		GraphCountTotal []int
		GraphLabel      []time.Time
		DashDetailGraph DashDetailGraph
	}
	orderDash, err := idb.DB.Table("tb_orders").Select("order_start_date,sum(if(order_status = 1, 1, 0)) as 'Approve',sum(if(order_status = 2, 1, 0)) as 'Rejected',sum(if(order_status = 3, 1, 0)) as 'Pending'").Group("date(order_start_date)").Order("order_start_date asc").Rows()
	orderDashLabel, err := idb.DB.Table("tb_orders").Select("order_start_date").Group("date(order_start_date)").Order("order_start_date asc").Rows()
	orderCountTotalData, err := idb.DB.Table("tb_orders").Select("sum(if(order_status = 1, 1, 0)) as 'Approve',sum(if(order_status = 2, 1, 0)) as 'Rejected',sum(if(order_status = 3, 1, 0)) as 'Pending'").Rows()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
	} else {
		for orderCountTotalData.Next() {
			err := orderCountTotalData.Scan(&Dashboard.OrderApproved, &Dashboard.OrderRejected, &Dashboard.OrderPending)

			if err != nil {
				fmt.Println(err)
			}
			DashboardArrayCountTotal = append(DashboardArrayCountTotal, Dashboard.OrderApproved)
			DashboardArrayCountTotal = append(DashboardArrayCountTotal, Dashboard.OrderRejected)
			DashboardArrayCountTotal = append(DashboardArrayCountTotal, Dashboard.OrderPending)

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
			err := orderDash.Scan(&Dashboard.OrderDate, &Dashboard.OrderApproved, &Dashboard.OrderRejected, &Dashboard.OrderPending)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, err)
				c.Abort()
			} else {
				DashboardArrayApproved = append(DashboardArrayApproved, Dashboard.OrderApproved)
				DashboardArrayRejected = append(DashboardArrayRejected, Dashboard.OrderRejected)
				DashboardArrayPending = append(DashboardArrayPending, Dashboard.OrderPending)
			}

		}

		result = gin.H{
			"status": "ok",
			"result": DashAll{
				GraphCountTotal: DashboardArrayCountTotal,
				GraphLabel:      DashboardLabel,
				DashDetailGraph: DashDetailGraph{
					Approve: DashboardArrayApproved,
					Pending: DashboardArrayPending,
					Reject:  DashboardArrayRejected,
				},
			},
		}

		c.JSON(http.StatusOK, result)
	}

}
