package protocol

type Login struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type FaceValidate struct {
	UserName string `json:"user_name"`
	Feature  string `json:"feature"`
}
