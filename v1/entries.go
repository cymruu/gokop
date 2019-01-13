package v1

import (
	"net/url"
	"strconv"

	"github.com/cymruu/gokop/v1/models"
)

func (w *WykopAPIV1) Index(entryID int64) (*models.Entry, error) {
	resp := new(models.Entry)
	req := w.request("entries/index", OpMethodParams(strconv.FormatInt(entryID, 10)))
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
