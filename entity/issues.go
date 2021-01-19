package entity

type Isuues struct {
	ID     int64  `db:"id" json:"id"`
	Issues string `db:"issues" json:"issues"`
}
