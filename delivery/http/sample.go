package delivery

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (d *delivery) SampleGinHandler(ginCtx *gin.Context) {
	data, err := d.service.SampleProcess(ginCtx)
	if err != nil {
		// Do something here
		d.logger.Error(ginCtx.Request.Context(), "SampleGinHandler error", err)
		return
	}

	d.logger.Info(context.Background(), "SampleGinHandler Info", zap.Any("message", "Success"))
	responseData(ginCtx, http.StatusOK, httpResponse{
		Message: "Success",
		Data:    data,
	})
}

func (d *delivery) SampleError(ginCtx *gin.Context) {
	err := errors.New("just a simple error for testing purposes")
	d.logger.Error(ginCtx.Request.Context(), "SampleGinError", err)
	responseError(ginCtx, err)
}
