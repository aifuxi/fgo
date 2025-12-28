package handler

import (
	"context"
	"strings"
	"time"

	"github.com/aifuxi/fgo/config"
	"github.com/aifuxi/fgo/internal/model/dto"
	"github.com/aifuxi/fgo/pkg/logger"
	"github.com/aifuxi/fgo/pkg/response"
	"github.com/aifuxi/fgo/pkg/upload"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (h *UploadHandler) UploadPresign(ctx *gin.Context) {
	var req dto.UploadPresignReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ParamError(ctx, err.Error())
		return
	}

	filename := config.AppConfig.OSS.UploadDir + "/" + req.Name

	result, err := upload.GetClient().Presign(context.TODO(), &oss.PutObjectRequest{
		Bucket:      oss.Ptr(config.AppConfig.OSS.Bucket),
		Key:         oss.Ptr(filename),
		ContentType: oss.Ptr("application/octet-stream"),
	},
		oss.PresignExpires(10*time.Minute),
	)
	if err != nil {
		logger.Sugar.Fatalf("failed to put object presign %v", err)
		response.BusinessError(ctx, err.Error())
		return
	}

	if len(result.SignedHeaders) > 0 {
		//当返回结果包含签名头时，使用签名URL发送Put请求时，需要设置相应的请求头
		logger.Sugar.Infof("signed headers:\n")
		for k, v := range result.SignedHeaders {
			logger.Sugar.Infof("%v: %v\n", k, v)
		}
	}

	response.Success(ctx, dto.UploadPresignResp{
		URL:           strings.SplitN(result.URL, "?", 2)[0],
		Name:          req.Name,
		SignedHeaders: result.SignedHeaders,
		UploadURL:     result.URL,
	})
}
