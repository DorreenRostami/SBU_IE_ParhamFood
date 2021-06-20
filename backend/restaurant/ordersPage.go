package restaurant

import (
	"net/http"

	fh "github.com/DorreenRostami/IE_ParhamFood/filehandler"
	"github.com/labstack/echo/v4"
)

func ChangeOrderStatus(c echo.Context) error {
	var order fh.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profiles := fh.GetProfilesFromFile()
	for i := 0; i < len(profiles.Profiles); i++ {
		if profiles.Profiles[i].ID == order.RestaurantID {
			orders := profiles.Profiles[i].Orders
			for j := 0; j < len(orders); j++ {
				if orders[j].OrderID == order.OrderID {
					orders[j].Status = orders[j].Status + 1
					order = orders[j]
					break
				}
				if j == len(orders)-1 {
					return c.JSON(http.StatusConflict, ResponseMessage{
						StatusCode: http.StatusBadRequest,
						Message:    "Wrong order ID.",
					})
				}
			}
			break
		}
		if i == len(profiles.Profiles)-1 {
			return c.JSON(http.StatusConflict, ResponseMessage{
				StatusCode: http.StatusBadRequest,
				Message:    "Wrong restaurant ID.",
			})
		}
	}
	fh.UpdateProfileFile(profiles)
	return c.JSON(http.StatusOK, order)
}
