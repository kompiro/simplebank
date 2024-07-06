package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

type getEntryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetEntry(ctx *gin.Context) {
	var req getEntryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	entry, err := server.store.GetEntry(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, entry)
}

type listEntryRequest struct {
	Uri struct {
		AccountID int64 `uri:"account_id" binding:"required,min=1"`
	}
	Query struct {
		PageID   int32 `form:"page_id" binding:"required,min=1"`
		PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
	}
}

func (server *Server) ListEntry(ctx *gin.Context) {
	var req listEntryRequest

	if err := ctx.ShouldBindUri(&req.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindQuery(&req.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.ListEntriesParams{
		AccountID: req.Uri.AccountID,
		Limit:     req.Query.PageSize,
		Offset:    (req.Query.PageID - 1) * req.Query.PageSize,
	}

	accounts, err := server.store.ListEntries(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
