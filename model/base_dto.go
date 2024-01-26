package model

type BaseDto struct {
	CreatedAt JsonTime `json:"created_at"`
	UpdatedAt JsonTime `json:"updated_at"`
	CreateBy  int64    `json:"create_by"`
	UpdateBy  int64    `json:"update_by"`
}

func FromModel(model BaseModel) BaseDto {
	dto := BaseDto{}
	dto.CreatedAt = JsonTime{
		Time: model.CreatedAt,
	}
	dto.UpdatedAt = JsonTime{
		Time: model.UpdatedAt,
	}
	dto.CreateBy = model.CreateBy
	dto.UpdateBy = model.UpdateBy
	return dto
}
