// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type SignDemoReq struct {
	Msg string `json:"msg"`
}

type SignDemoResp struct {
	Msg string `json:"msg"`
}