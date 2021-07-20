package controllers

import (
	"behealth-api/api/responses"
	"behealth-api/api/services"
	"behealth-api/infrastructure"
	"behealth-api/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UtilityController -> utitley controller struct
type UtilityController struct {
	logger infrastructure.Logger
	env    infrastructure.Env
	bucket services.StorageBucketService
}

// NewUtilityController -> construct
func NewUtilityController(logger infrastructure.Logger,
	env infrastructure.Env,
	bucket services.StorageBucketService) UtilityController {
	return UtilityController{
		logger: logger,
		env:    env,
		bucket: bucket,
	}
}

// Input -> add binding model for input
type Input struct {
	Model string `form:"model,omitempty" binding:"required"`
}

// Response -> response for the util scope
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    string      `json:"data"`
	Path    string      `json:"path"`
	Value   interface{} `json:"attributes"`
}

const storageURL string = "https://storage.googleapis.com/"

// FileUploadHandler -> handles file upload
func (uc UtilityController) FileUploadHandler(ctx *gin.Context) {
	file, uploadFile, err := ctx.Request.FormFile("file")
	if err != nil {
		uc.logger.Zap.Error("Error Get File from request :: ", err.Error())
		responses.ErrorJSON(ctx, http.StatusBadRequest, "Failed to get file form request")
		return
	}

	var input Input
	err = ctx.ShouldBind(&input)
	if err != nil {
		uc.logger.Zap.Error("Error getting the model name  :: ", err.Error())
		responses.ErrorJSON(ctx, http.StatusBadRequest, "Failed to get model name ")
		return
	}

	group := input.Model
	fileExtension := filepath.Ext(uploadFile.Filename)
	fileName := utils.GenerateRandomFileName() + fileExtension
	originalFileName := group + "/images/original/" + fileName

	// File type
	file1, _, _ := ctx.Request.FormFile("file")
	fileHeader := make([]byte, 512)
	if _, err := file1.Read(fileHeader); err != nil {
		uc.logger.Zap.Error("Error File Read upload File::", err.Error())
		responses.ErrorJSON(ctx, http.StatusBadRequest, "Failed to read upload  File")
		return
	}
	fileType := http.DetectContentType(fileHeader)
	if fileType == "image/png" || fileType == "image/jpg" || fileType == "image/jpeg" || fileType == "image/gif" {
		uploadedOriginalURL, err := uc.bucket.UploadFile(ctx.Request.Context(), file, originalFileName, fileType)
		if err != nil {
			uc.logger.Zap.Error("Error Failed to upload File::", err.Error())
			responses.ErrorJSON(ctx, http.StatusBadRequest, "Failed to upload File")
			return
		}

		response := &Response{
			Success: true,
			Message: "Uploaded Successfully",
			Data:    storageURL + uc.env.StorageBucketName + "/" + uploadedOriginalURL,
			Path:    uploadedOriginalURL,
			Value: map[string]string{
				"original_image_url":  storageURL + uc.env.StorageBucketName + "/" + uploadedOriginalURL,
				"original_image_path": uploadedOriginalURL,
			}}
		ctx.JSON(http.StatusOK, response)
		return
	}

	originalFileName = group + "/files/" + fileName
	uploadedFileURL, err := uc.bucket.UploadFile(ctx.Request.Context(), file, originalFileName, fileType)
	if err != nil {
		uc.logger.Zap.Error("Error Failed to upload File::", err.Error())
		responses.ErrorJSON(ctx, http.StatusBadRequest, "Failed to upload file ")
		return
	}
	response := &Response{
		Success: true,
		Message: "Uploaded Successfully",
		Data:    storageURL + uc.env.StorageBucketName + "/" + uploadedFileURL,
		Path:    uploadedFileURL,
	}
	ctx.JSON(http.StatusOK, response)
}

// DeleteFile -> struct
type DeleteFile struct {
	FilePath string `json:"file_path" binding:"required"`
}

// DeleteuploadedFileHandler -> handler to delete the uploaded file
func (uc UtilityController) DeleteuploadedFileHandler(ctx *gin.Context) {
	var b DeleteFile
	err := ctx.BindJSON(&b)
	if err != nil {
		uc.logger.Zap.Error("Error binding the filepath  :: ", err.Error())
		responses.ErrorJSON(ctx, http.StatusBadRequest, " Error binding the filepath")
		return
	}
	err = uc.bucket.RemoveObject(b.FilePath)
	if err != nil {
		uc.logger.Zap.Error("Error deleting file :: ", err.Error())
		responses.ErrorJSON(ctx, http.StatusBadRequest, " Error deleting file ")
		return
	}
	response := &Response{
		Success: true,
		Message: "Deleted Successfully",
	}
	ctx.JSON(http.StatusOK, response)
}
