package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MjHistorysModel = (*customMjHistorysModel)(nil)

type (
	// MjHistorysModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMjHistorysModel.
	MjHistorysModel interface {
		mjHistorysModel
		withSession(session sqlx.Session) MjHistorysModel
	}

	customMjHistorysModel struct {
		*defaultMjHistorysModel
	}
)

// NewMjHistorysModel returns a model for the database table.
func NewMjHistorysModel(conn sqlx.SqlConn) MjHistorysModel {
	return &customMjHistorysModel{
		defaultMjHistorysModel: newMjHistorysModel(conn),
	}
}

func (m *customMjHistorysModel) withSession(session sqlx.Session) MjHistorysModel {
	return NewMjHistorysModel(sqlx.NewSqlConnFromSession(session))
}
