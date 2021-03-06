package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/DuckLuckBreakout/web/backend/internal/pkg/models"
	"github.com/DuckLuckBreakout/web/backend/internal/pkg/promo_code"
	"github.com/DuckLuckBreakout/web/backend/internal/server/errors"
	"github.com/DuckLuckBreakout/web/backend/internal/server/tools/http_utils"
	"github.com/DuckLuckBreakout/web/backend/internal/server/tools/validator"
	"github.com/DuckLuckBreakout/web/backend/pkg/tools/logger"
)

type PromoCodeHandler struct {
	PromoCodeUCase promo_code.UseCase
}

func NewHandler(UCase promo_code.UseCase) promo_code.Handler {
	return &PromoCodeHandler{
		PromoCodeUCase: UCase,
	}
}

func (h *PromoCodeHandler) ApplyPromoCodeToOrder(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		requireId := http_utils.MustGetRequireId(r.Context())
		if err != nil {
			logger.LogError("promo_code_handler", "ApplyPromoCodeToOrder", requireId, err)
		}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http_utils.SetJSONResponse(w, errors.ErrBadRequest, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	promoCodeGroup := &models.PromoCodeGroup{}
	err = json.Unmarshal(body, promoCodeGroup)
	if err != nil {
		http_utils.SetJSONResponse(w, errors.ErrCanNotUnmarshal, http.StatusBadRequest)
		return
	}

	err = validator.ValidateStruct(promoCodeGroup)
	if err != nil {
		http_utils.SetJSONResponse(w, errors.CreateError(err), http.StatusBadRequest)
		return
	}

	discountedPrice, err := h.PromoCodeUCase.ApplyPromoCodeToOrder(promoCodeGroup)
	if err == errors.ErrProductNotInPromo {
		http_utils.SetJSONResponse(w, errors.CreateError(err), http.StatusConflict)
		return
	} else if err != nil {
		http_utils.SetJSONResponse(w, errors.CreateError(err), http.StatusInternalServerError)
		return
	}

	http_utils.SetJSONResponse(w, discountedPrice, http.StatusOK)
}
