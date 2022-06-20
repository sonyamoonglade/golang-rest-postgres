package middleware

import (
	"context"
	"fmt"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/util"
	"net/http"
	"strconv"
)

type extractMiddleware struct {
	logger *myLogger.Logger
	xtkey  string
}

const (
	UserId = "userId"
)

func NewExtractMiddleware(logger *myLogger.Logger, xtkey string) *extractMiddleware {
	return &extractMiddleware{logger: logger, xtkey: xtkey}
}

func (m *extractMiddleware) Extract(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		q := r.URL.Query()

		if has := q.Has(m.xtkey); has != true {
			util.JsonResponse(w, http.StatusBadRequest, map[string]interface{}{
				"message": fmt.Sprintf("%s has not been provided", m.xtkey),
			})
			m.logger.PrintWithErr(fmt.Sprintf("%s has not been provided", m.xtkey))
			return
		}

		xt, err := strconv.ParseInt(q.Get(m.xtkey), 10, 64)
		if err != nil {
			util.JsonResponse(w, http.StatusBadRequest, map[string]interface{}{
				"message": fmt.Sprintf("incorrect %s format", m.xtkey),
			})
			m.logger.PrintWithInf(fmt.Sprintf("unable to format %s", m.xtkey))
			return
		}
		ctx := context.WithValue(r.Context(), m.xtkey, xt)
		h.ServeHTTP(w, r.WithContext(ctx))

	}
}
