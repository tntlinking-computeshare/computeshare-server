package model

// Dh3tTemplateContent message Dh3tTemplateContent structure
type Dh3tTemplateContent struct {
	Name  string `json:"name"`  //模版占位符
	Value string `json:"value"` //替换值
}

// Template message Template structure
type Template struct {
	Id        string                `json:"id"`        //模版ID
	Variables []Dh3tTemplateContent `json:"variables"` //内容
}

// Dh3t message Dh3t structure
type Dh3t struct {
	Account  string   `json:"account"`  // 账号
	Password string   `json:"password"` // 密码
	Phones   string   `json:"phones"`   // 手机
	Template Template `json:"template"` // 模版
}
