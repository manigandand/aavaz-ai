package respond

import (
	"fmt"
	"reflect"
	"time"

	"compress/gzip"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

var (
	DEFAULT_PAGE_LIMIT = 10
	decoder            *schema.Decoder
)

// Response holds the handlerfunc response
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta Meta        `json:"meta"`
}

// Meta holds the status of the request informations
type Meta struct {
	Status  int    `json:"status_code"`
	Message string `json:"error_message,omitempty"`

	// Pagination
	Total    int    `json:"total,omitempty"`
	Count    int    `json:"count,omitempty"`
	Previous string `json:"previous,omitempty"`
	Current  string `json:"current,omitempty"`
	Next     string `json:"next,omitempty"`
}

// Page holds the paginate informations
type Page struct {
	Offset int      `schema:"offset" url:"offset"`
	Limit  int      `schema:"limit" url:"limit"`
	Topics []string `schema:"topics" url:"topics"`
	Sort   string   `schema:"sort" url:"sort"`
}

func init() {
	decoder = schema.NewDecoder()
	decoder.ZeroEmpty(true)
	decoder.IgnoreUnknownKeys(true)
	decoder.RegisterConverter(time.Time{}, parseFilterTime)
}

func parseFilterTime(date string) reflect.Value {
	if s, err := time.Parse(time.RFC3339, date); err == nil {
		return reflect.ValueOf(s)
	}

	return reflect.Value{}
}

// NewPage decodes the pagiante information from the request
func NewPage(r *http.Request) *Page {
	page := &Page{Limit: DEFAULT_PAGE_LIMIT}
	if err := decoder.Decode(page, r.Form); err != nil {
		log.Println(err)
	}
	if page.Limit == 0 || page.Limit > DEFAULT_PAGE_LIMIT {
		page.Limit = DEFAULT_PAGE_LIMIT
	}
	return page
}

func (r *Response) Send(w http.ResponseWriter) {
	gz := gzip.NewWriter(w)
	defer gz.Close()
	buf, err := json.Marshal(r)
	if err != nil {
		// Fail(w, )
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(r.Meta.Status)
	if r.Meta.Status != http.StatusNoContent {
		if _, err := gz.Write(buf); err != nil {
			log.Error("respond.With.error: ", err)
		}
	}
}

// Paginate returns the response container with paginated inforamtions
func Paginate(w http.ResponseWriter, r *http.Request, data interface{}, p *Page, total, count int) {
	var (
		previous,
		current,
		next string
	)

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	current = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)
	u, err := url.Parse(current)
	if err != nil {
		// return nil
	}

	prevQuery := u.Query()
	nextQuery := u.Query()
	// previous
	if p.Offset != 0 {
		prevQuery.Set("limit", fmt.Sprintf("%d", p.Limit))
		prevQuery.Set("offset", fmt.Sprintf("%d", p.Offset-p.Limit))
		u.RawQuery = prevQuery.Encode()
		previous = u.String()
	}
	// next link
	if p.Offset < (total-1) && p.Offset+count < total {
		nextQuery.Set("limit", fmt.Sprintf("%d", p.Limit))
		nextQuery.Set("offset", fmt.Sprintf("%d", p.Offset+p.Limit))
		u.RawQuery = nextQuery.Encode()
		next = u.String()
	}

	res := &Response{
		Data: data,
		Meta: Meta{
			Status:   http.StatusOK,
			Total:    total,
			Count:    count,
			Previous: previous,
			Current:  current,
			Next:     next,
		},
	}
	res.Send(w)
}
