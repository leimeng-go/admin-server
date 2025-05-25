package {{.pkg}}

import (
    "github.com/zeromicro/go-zero/core/stores/sqlx"
    "errors"
)

var ErrNotFound = sqlx.ErrNotFound
var ErrNoRowsUpdate = errors.New("update db no rows change")