package model

import (
	"github.com/spf13/cast"
)

type BaseDto struct {
	CreatedAt JsonTime `json:"createdAt"`
	UpdatedAt JsonTime `json:"updatedAt"`
	CreateBy  string   `json:"createBy"`
	UpdateBy  string   `json:"updateBy"`
}

func FromModel(model BaseModel) BaseDto {
	dto := BaseDto{}
	dto.CreatedAt = JsonTime{
		Time: model.CreatedAt,
	}
	dto.UpdatedAt = JsonTime{
		Time: model.UpdatedAt,
	}
	dto.CreateBy = cast.ToString(model.CreateBy)
	dto.UpdateBy = cast.ToString(model.UpdateBy)
	return dto
}
