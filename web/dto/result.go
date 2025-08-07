package dto

type Result[T any] struct {
	Status   string `json:"status" bson:"status" description:"状态码" example:"0000"`
	Describe string `json:"describe,omitempty" bson:"describe,omitempty" description:"描述" example:"成功"`
	Payload  T      `json:"payload,omitempty" bson:"payload,omitempty" description:"数据" example:"{}"`
}
