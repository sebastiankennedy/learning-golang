package models

import (
	"goblog/pkg/types"
)

type BaseModel struct {
	ID uint64
}

func (b BaseModel) GetStringId() string {
	return types.Uint64ToString(b.ID)
}
