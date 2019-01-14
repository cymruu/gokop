package v1

import (
	"net/url"
	"strconv"

	"github.com/cymruu/gokop/v1/models"
)

func (w *WykopAPIV1) Index(entryID int64, params ...OptionalParamV1) (*models.Entry, error) {
	params = append(params, OpMethodParams(strconv.FormatInt(entryID, 10)))
	resp := new(models.Entry)
	req := w.request("entries/index", params...)
	err := w.MakeRequest(req, &resp)
	return resp, err
}

func (w *WykopAPIV1) AddComment(entry *models.Entry, body string) (*models.OK_ID, error) {
	resp := new(models.OK_ID)
	req := w.request("entries/addcomment", OpMethodParams(strconv.FormatInt(entry.ID, 10)))
	post := url.Values{}
	post.Add("body", body)
	req.SetPostParams(post)
	err := w.MakeRequest(req, &resp)
	return resp, err
}
