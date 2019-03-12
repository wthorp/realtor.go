package realtor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Session represents a set of connections
type Session struct {
	Jar http.CookieJar
}

//NewSession returns a session
func NewSession() (*Session, error) {
	s := &Session{}
	return s, s.http("GET", "https://www.realtor.com/", nil, nil) //populate cookieJar
}

// Search returns place information by name
func (s *Session) Search(req ListingRequest) (*ListingResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	rdr := bytes.NewReader(b)
	res := &ListingResponse{}
	u := "https://www.realtor.com/search_result.json"
	err = s.http("POST", u, rdr, res)
	return res, err
}

// FindPlace returns place information by name
func (s *Session) FindPlace(place string) (*Place, error) {
	u := "https://www.realtor.com/api/v1/geo-landing/parser/suggest/?input=%s&area_types=neighborhood&area_types=city&area_types=postal_code&limit=1&includeState=false"
	//u := "https://www.realtor.com/api/v1/geo-landing/parser/suggest/?input=%s&area_types=address&area_types=neighborhood&area_types=city&area_types=county&area_types=postal_code&area_types=street&area_types=building&area_types=school&area_types=building&area_types=school_district&limit=1&latitude=35.175201416015625&longitude=-79.3833999633789&includeState=false"
	r := &PlaceResponse{}
	err := s.http("GET", fmt.Sprintf(u, url.PathEscape(place)), nil, r)
	if err != nil {
		return nil, err
	}
	if len(r.Result) == 0 {
		return nil, fmt.Errorf("Place not found")
	}
	return &r.Result[0], err
}

//http makes web request and provides JSON serialization
func (s *Session) http(httpType, u string, post io.Reader, output interface{}) error {
	req, err := http.NewRequest(httpType, u, post)
	if err != nil {
		return err
	}
	//try at least a little to look like normal traffic
	req.Header.Set("Host", "www.realtor.com")
	req.Header.Set("Origin", "https://www.realtor.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	if output != nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
	client := &http.Client{Jar: s.Jar} //use cookiejar for any session / XSS checks
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	if output == nil {
		return nil
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(output)
	if err != nil {
		b, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("JSON err: %s\n", string(b))
	}
	return err
}
