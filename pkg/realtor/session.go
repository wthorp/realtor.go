package realtor

import (
	"encoding/json"
	"fmt"
	"io"
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
	return s, s.http("GET", "https://www.realtor.com/", nil, false) //populate cookieJar
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
	req.Header.Set("Host", "www.realtor.com")
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Content-Length", "1095")
	//req.Header.Set("X-NewRelic-ID", "VwEPVF5XGwsIUlhRAAkH")
	req.Header.Set("Origin", "https://www.realtor.com")
	//req.Header.Set("X-CSRF-Token", "23m/fNRBtBCrPKcQ0f1uJ6RaDIzzESJkcm4BEWZzmd39F9sQ59MRsvFlIM60ufPUZguAQaYtsKy27ny4vDLCjQ==")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	if output != nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
	//req.Header.Set("Referer", "https://www.realtor.com/realestateandhomes-search/Pittsburgh_PA?pos=40.191244,-80.196962,40.70489,-79.715967,11")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	//req.Header.Set("Cookie", "split_tcv=26; __vst=2641a911-b178-46bb-80c7-409f5a330994; __ssn=de5aef00-0683-4371-ad67-0ce12f2cf987; __ssnstarttime=1548108554; AMCVS_8853394255142B6A0A4C98A4%40AdobeOrg=1; ajs_user_id=null; ajs_group_id=null; ajs_anonymous_id=%22305a7d0a-a9f6-4d46-a642-9705ff924e88%22; _ga=GA1.2.228950573.1548108554; _gcl_au=1.1.1863942703.1548108555; threshold_value=43; automation=false; clstr=v; clstr_tcv=69; bcc=false; bcvariation=SRPBCRR%3Av1%3Adesktop; userStatus=return_user; ab_srp_viewtype=ab-mixed-view; srp.viewType=map; lithium_display_name=false; REMEMBER_ME=em=wL5sdU5X0TjgHyzagT9e9FBamv9mZOzshyv5o77be1k=&regID=5c466912244220236306a3a7&nm=bill3000&pud=&pID=5c466912244220236306a3a3&stat=1&act=True&yob=&opt=20&exp=7/30/2019 12:51:30 AM&up=1/22/2019 12:51:30 AM&; _sdsat_userID=5c466912244220236306a3a7; saveAware=5c466912244220236306a3a7; url_search_data=search_city%3DMount+Lebanon%26search_state_id%3DPA; _fbp=fb.1.1551024053334.473121530; AMCV_8853394255142B6A0A4C98A4%40AdobeOrg=-179204249%7CMCIDTS%7C17954%7CMCMID%7C00307026630246207453158312857637115281%7CMCAAMLH-1551628845%7C7%7CMCAAMB-1551744046%7CRKhpRz8krg2tLO6pguXWp5olkAcUniQYPHaMWWgdJ3xzPWQmdj0y%7CMCCIDH%7C-1444499312%7CMCOPTOUT-1551146446s%7CNONE%7CMCAID%7CNONE; _gid=GA1.2.579674677.1551139249; _rv_mpr_ids=4583105483%2C4909694390%2C3912819056%2C5118819911%2C5099031696%2C4134369854; previousUrl=%2Frealestateandhomes-search%2FPittsburgh_PA; srchID=a7c6f6515766459b9b1b789b212787e1; split=n; header_slugs=gs%3DPittsburgh_PA%26lo%3DPittsburgh%26st%3Dcity%2Cgs%3DAllegheny-County_PA%26lo%3DAllegheny%26st%3Dcity; criteria=loc%3DPittsburgh%2C+PA%26locSlug%3DPittsburgh_PA%26lat%3D40.451021%26long%3D-79.958261%26status%3D1%26pg%3D1%26pgsz%3D44%26sprefix%3D%2Frealestateandhomes-search%26city%3DPittsburgh%26state_id%3DPA%26county_fips%3D42003%26county_fips_multi%3D42003-42125-42129; _4c_=XVNNd9owEPwreTpjkGRZtriRpK9JWwIhpLQnnizJ2MVgP0lAPuC%2Fd0UgSesLmtnRatndeUW70qxRnyQJIWnCKI6J6KCleXao%2F4pspcPPFvVRbDLMSSFEBjqpuNGFEkrHmuQF44KjDno65kmyjCUJ41wcOkivz%2Fe1KeSm9p9kglHK4UmQVWeV%2FD%2BeZinE7e4keA9kSUr%2FlQYGpKo9SV%2FRxtaQsvS%2Bdf1eb7fbda2RtW9sVzWrXjgb56U3cq3LZmVc5Iy0quyNK%2B9dvrGLcj4eQEGq0QYSEdFNuxSwfwEUxRjDubWN3ig%2F989t0OxMfuH0EgJfm2ZRm9trIClnRApCopykWcR4nkcZVmnEsCgSGcdYCAY3flauguKOVyKSCooZZWI%2FvLq9nj7sgUgYgOHt9R7jGKeYch5jyjjFKUtikmQxCU3gcUpIQjMC4sFg%2BOMmCsPlNIxln76Rl0cuZQwzvp98L9vJS7a0C%2Bp%2FjHi72PyatUlTLwfqcV3d%2Fx7fyOFsttDf4qeX8ex%2Bpf%2FgZ0gDZd3sI8LgEwKeBmo0no4ep8fchHHGuNvfje6%2BhDeh6nAM3SutcWVT6%2FlW1pvQNBYD7eZb2AL00Yf5sRGAL6UzSq5grugOYK620lbSVw3sFnqYjC%2BvJpP%2BlvS1cUvftGFgtfMWgtsgt83OmYCuSgtTvkgJsA2sN5pVaw1BgNYUxtqjKpRS%2BVDWp2U5kWCLDz4M%2F2iMsAZ1o2CbAIGZOmhqq8XC2KHxZQMGAix1FQqWNXrzxCc7aA9Q1dK5Sp3%2FwuG82oRykpEsjjNonId1zjjD4QPF9sM1KslTTQqMtYyJNBj8J5VOdNgiXqh3U4HHUwrpqDilI9lbtsPhLw%3D%3D; AWSALB=U1MhIegrBCwqQtU1yQmzryK/mBtIYUrp3cm5zQSCW0o/67CXNlvXa2+tiIDI3VP/9Eik7DKG8Urz4EJRujM2ZhlX66VG7Fp+47k+WHEreBzTJg20uCzxStemaVzX; _rdc-next_session=Z2dML3VwdEdLM3hwbTdpUGpPUWJCbGJuNjBEOFh3bXhJMlRld0ZSVElaSWhCYU12NUF2WFZhVTRyNWdhVnFCU2tLcU9jWWlueXdnRkcxSi9SOXo0Q3VETWdvSjlkYUxxdmNEQmwzeGkrT0lvYXhQUW5tTE1lY1NvTWZvRlVBQUVPUTY5MWU0dnZuTEp6bnJGY3NCVDlCQmRBa1JlRnlKUklxc1llMGU5anIvQlMvYmM5N0FpRmlpRjBZRXIzVnd4MjBwVXNRS2xpYzgwWTJGeWhjUVF2LzhGRW5tK1JwU1FYY0tUYUZlVkJzRnhYbGY4aGFoaUhMU0srWlBYTzEzTTFKNDc1QUdpckdmN09NOWcycUJUcDZDd3k4YnpsYTMyeklHT0RhbU05L1k2OWtjUFlLOVcyMjEyRmF4WVp1Yks1V3lBbWF2ai9mT0VCVllMTXlEaUZONDc1ZjZzV1NmVDBaUnU0cnF0dEk0PS0tOW1MdHBIRU9meExiVWwwL1czVGFQdz09--d597b0a3850f139ca3ec9e0d0ff9c488756d03ae; GED_PLAYLIST_ACTIVITY=W3sidSI6ImF1b1giLCJ0c2wiOjE1NTExNzU2OTYsIm52IjowLCJ1cHQiOjE1NTExNzU0MTYsImx0IjoxNTUxMTc1NTQ3fV0.")

	client := &http.Client{Jar: s.Jar}
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	if output == nil {
		return nil
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(output)
}
