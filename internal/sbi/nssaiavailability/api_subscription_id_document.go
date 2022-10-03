/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nssaiavailability

import (
	"net/http"

	"github.com/gin-gonic/gin"

	nssf_context "github.com/free5gc/nssf/internal/context"
	"github.com/free5gc/nssf/internal/logger"
	"github.com/free5gc/nssf/internal/sbi/producer"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

func HTTPNSSAIAvailabilityUnsubscribe(c *gin.Context) {
	scopes := // Due to conflict of route matching, 'subscriptions' in the route is replaced with the existing wildcard ':nfId'
		[]string{"nnssf-nssaiavailability"}
	_, oauth_err := openapi.CheckOAuth(c.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && nssf_context.NSSF_Self().OAuth == true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}

	nfID := c.Param("nfId")
	if nfID != "subscriptions" {
		c.JSON(http.StatusNotFound, gin.H{})
		logger.Nssaiavailability.Infof("404 Not Found")
		return
	}

	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	rsp := producer.HandleNSSAIAvailabilityUnsubscribe(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.HandlerLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
