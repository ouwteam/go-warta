package helper

import "net/http"

func NewResponse(mReq *http.Request, pRes http.ResponseWriter) *Response {
	return &Response{Req: mReq, Res: pRes}
}

type Response struct {
	Req *http.Request
	Res http.ResponseWriter
}

func (dmsR *Response) writeResponse(statusCode int, json []byte) {
	dmsR.Res.Header().Set("Content-Type", "application/json")
	dmsR.Res.WriteHeader(statusCode)
	dmsR.Res.Write(json)
}

func (dmsR *Response) ResponseOK(json []byte) {
	dmsR.writeResponse(http.StatusOK, json)
}

func (dmsR *Response) ResponseBadRequest(json []byte) {
	dmsR.writeResponse(http.StatusBadRequest, json)
}

func (dmsR *Response) ResponseForbiden(json []byte) {
	dmsR.writeResponse(http.StatusForbidden, json)
}

func (dmsR *Response) ResponseError(json []byte) {
	dmsR.writeResponse(http.StatusInternalServerError, json)
}
