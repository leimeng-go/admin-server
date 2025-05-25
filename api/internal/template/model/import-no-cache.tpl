import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}
    "admin-server/api/internal/pkg/globalkeys"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
    {{if .containsPQ}}"github.com/lib/pq"{{end}}
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"

	{{.third}}
)
