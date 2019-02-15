package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type record struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	SafeTitle  string `json:"safe_title"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	News       string `json:"news"`
	Link       string `json:"link"`
}

func (r *record) String() string {
	return r.Title + " - " + r.Img
}

type records struct {
	Access time.Time `json:"access"`
	M      map[int]*record
}

func newRecords() *records {
	return &records{M: make(map[int]*record)}
}

func (rs *records) upsert(r *record) {
	if r.Num < 1 {
		return
	}

	rs.M[r.Num] = r
}

func (rs *records) access() time.Time {
	return rs.Access
}

func (rs *records) setAccess(t time.Time) {
	rs.Access = t
}

func (rs *records) length() int {
	return len(rs.M)
}

func (rs *records) record(num int) (*record, bool) {
	r, ok := rs.M[num]
	return r, ok
}

type access struct {
	root string
	sufx string
}

func newAccess(root, sufx string) *access {
	return &access{root, sufx}
}

func (c *access) count() (int, error) {
	url := c.root + c.sufx
	var res struct{ Num int }

	if err := sendDecode(url, &res); err != nil {
		return 0, fmt.Errorf("access count: %s", err)
	}

	return res.Num, nil
}

func (c *access) record(num int) (*record, error) {
	url := c.root + "/" + strconv.Itoa(num) + c.sufx
	var res record

	if num == 404 {
		res.Num = 404
		res.Title = "Not Found"

		return &res, nil
	}

	if err := sendDecode(url, &res); err != nil {
		return nil, fmt.Errorf("access record: %s", err)
	}

	return &res, nil
}

func updateRecords(expiry, pause time.Duration, a *access, rs *records) error {
	efmt := "update records %s"

	ct, err := a.count()
	if err != nil {
		return fmt.Errorf(efmt, err)
	}

	expired := isExpired(expiry, rs.access(), time.Now())

	bgn := rs.length()
	setAccess := func(_ time.Time) {}
	if bgn == 0 || expired {
		bgn = 1
		setAccess = rs.setAccess
	}

	if rs.length() > 0 && bgn == ct {
		return nil
	}

	for i := bgn; i <= ct; i++ {
		time.Sleep(pause)

		r, err := a.record(i)
		if err != nil {
			return fmt.Errorf(efmt, err)
		}

		rs.upsert(r)
	}

	setAccess(time.Now())

	return nil
}

func runUserInteraction(rs *records) error {
	var entry int
	for {
		fmt.Print("please enter a number (0 to exit): ")

		var txt string
		_, err := fmt.Scanln(&txt)
		if err != nil {
			return fmt.Errorf("user interaction: %s", err)
		}

		entry, err = strconv.Atoi(txt)
		if err != nil {
			fmt.Fprint(os.Stderr, "cannot parse input\n")
			continue
		}

		if entry == 0 {
			break
		}

		r, ok := rs.record(entry)
		if !ok {
			fmt.Println("not found")
			continue
		}

		fmt.Println(r)
	}

	return nil
}
