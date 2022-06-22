package response

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"data"`
}
