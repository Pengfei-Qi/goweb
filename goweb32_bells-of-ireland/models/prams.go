package models

//定义请求时的结构体
//binding 的参数可参考: https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme

const (
	OrderByTime  = "time"
	OrderByScore = "score"
)

// PramsSignUp 注册时参数
type PramsSignUp struct {
	Email      string `json:"email" binding:"required,email"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// PramsLogin 登陆时参数
type PramsLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// PramsVoteData 投票时参数
type PramsVoteData struct {
	//用户ID 从请求头中获取
	PostId    string `json:"post_id" binding:"required"`       //post id 帖子ID
	Direction int8   `json:"direction" binding:"oneof=1 0 -1"` //投票类型, 赞成票(1) 反对票(-1) 取消投票(0)
}

// ParamPostData 请求数据
type ParamPostData struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}

// ParamCommunityPostData 请求数据
type ParamCommunityPostData struct {
	*ParamPostData
	CommunityID int64 `json:"community_id" form:"community_id"`
}
