package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"encoding/json"

	"github.com/athom/goset"
)

type FuncTestCase func(input string, url string, check func(*testing.T, []byte)) FuncTestCase

type runner struct {
	f FuncTestCase
}

func (this *runner) run(t *testing.T) FuncTestCase {
	r := route()
	this.f = func(input string, url string, check func(*testing.T, []byte)) FuncTestCase {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", url, strings.NewReader(input))
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fail()
		}

		rsp := w.Body.Bytes()
		check(t, rsp)
		return this.f
	}
	return this.f
}

// 1. As a user, I need an API to create a friend connection between two email addresses.
//	The API should receive the following JSON request:
//
//	{
//		friends:
//		[
//			'andy@example.com',
//			'john@example.com'
//		]
//	}
//	The API should return the following JSON response on success:
//
//	{
//		"success": true
//	}
func TestMakeFriends(t *testing.T) {
	input := `
{
  "friends":
    [
      "andy@example.com",
      "john@example.com"
    ]
}
`
	var rr runner
	rr.run(t)(input, "/v1/friends/connect", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})
}

// 2. As a user, I need an API to retrieve the friends list for an email address.
//	The API should receive the following JSON request:
//	{
//		email: 'andy@example.com'
//	}
//	The API should return the following JSON response on success:
//	{
//		"success": true,
//		"friends" :
//		[
//			'john@example.com'
//		],
//		"count" : 1
//	}
func TestListFriends(t *testing.T) {
	input1 := `
{
  "friends":
    [
      "andy@example.com",
      "john@example.com"
    ]
}
`

	input2 := `
{
  "email": "andy@example.com"
}
`

	var rr runner
	rr.run(t)(input1, "/v1/friends/connect", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})(input2, "/v1/friends/find", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool     `json:"success"`
			Friends []string `json:"friends"`
			Count   int      `json:"count"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
		if !goset.IsIncluded(output.Friends, "john@example.com") {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
		if output.Count != 1 {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})
}

// 3. As a user, I need an API to retrieve the common friends list between two email addresses.
//	The API should receive the following JSON request:
//
//	{
//		friends:
//		[
//			'andy@example.com',
//			'john@example.com'
//		]
//	}
//	The API should return the following JSON response on success:
//
//	{
//		"success": true,
//		"friends" :
//		[
//			'common@example.com'
//		],
//		"count" : 1
//	}

func TestGetCommonFriends(t *testing.T) {
	input1 := `
{
  "friends":
    [
      "andy@example.com",
      "john@example.com"
    ]
}
`

	input2 := `
{
  "friends":
    [
      "andy@example.com",
      "tom@example.com"
    ]
}
`
	input3 := `
{
  "friends":
    [
      "tom@example.com",
      "john@example.com"
    ]
}
`
	input4 := `
{
  "friends":
    [
      "andy@example.com",
      "john@example.com"
    ]
}
`

	var rr runner
	rr.run(t)(input1, "/v1/friends/connect", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})(input2, "/v1/friends/connect", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})(input3, "/v1/friends/connect", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})(input4, "/v1/friends/common", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool     `json:"success"`
			Friends []string `json:"friends"`
			Count   int      `json:"count"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
		if !goset.IsIncluded(output.Friends, "tom@example.com") {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
		if output.Count != 1 {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})
}

// 4. As a user, I need an API to subscribe to updates from an email address.
//	Please note that "subscribing to updates" is NOT equivalent to "adding a friend connection".
//
//	The API should receive the following JSON request:
//
//	{
//		"requestor": "lisa@example.com",
//		"target": "john@example.com"
//	}
//	The API should return the following JSON response on success:
//
//	{
//		"success": true
//	}
func TestSubscribe(t *testing.T) {
	input1 := `
{
  "requestor": "lisa@example.com",
  "target": "john@example.com"
}
`
	input2 := `
{
  "email": "lisa@example.com"
}
`

	var rr runner
	rr.run(t)(input1, "/v1/friends/subscribe", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})(input2, "/v1/friends/find", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool     `json:"success"`
			Friends []string `json:"friends"`
			Count   int      `json:"count"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if goset.IsIncluded(output.Friends, "john@example.com") {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})
}

// 5. As a user, I need an API to block updates from an email address.
//	Suppose "andy@example.com" blocks "john@example.com":
//
//	if they are connected as friends, then "andy" will no longer receive notifications from "john"
//	if they are not connected as friends, then no new friends connection can be added
//	The API should receive the following JSON request:
//
//	{
//		"requestor": "andy@example.com",
//		"target": "john@example.com"
//	}
//	The API should return the following JSON response on success:
//
//	{
//	"success": true
//	}
func TestBlock(t *testing.T) {
	input1 := `
{
  "friends":
    [
      "andy@example.com",
      "john@example.com"
    ]
}
`
	input2 := `
{
  "requestor": "lisa@example.com",
  "target": "john@example.com"
}
`

	var rr runner
	rr.run(t)(input1, "/v1/friends/connect", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})(input2, "/v1/friends/block", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})
}

// 6. As a user, I need an API to retrieve all email addresses that can receive updates from an email address.
//	Eligibility for receiving updates from i.e. "john@example.com":
//
//	has not blocked updates from "john@example.com", and at least one of the following:
//	has a friend connection with "john@example.com"
//	has subscribed to updates from "john@example.com"
//	has been @mentioned in the update
//	The API should receive the following JSON request:
//
//	{
//		"sender":  "john@example.com",
//		"text": "Hello World! kate@example.com"
//	}
//	The API should return the following JSON response on success:
//
//	{
//		"success": true
//		"recipients":
//		[
//			"lisa@example.com",
//			"kate@example.com"
//		]
//	}
func TestReachableUsers(t *testing.T) {
	input1 := `
{
  "friends":
    [
      "andy@example.com",
      "john@example.com"
    ]
}
`
	input2 := `
{
  "friends":
    [
      "bob@example.com",
      "john@example.com"
    ]
}
`
	input3 := `
{
  "friends":
    [
      "tom@example.com",
      "john@example.com"
    ]
}
`
	input4 := `
{
  "requestor": "bob@example.com",
  "target": "john@example.com"
}
`

	input5 := `
{
	"sender":  "john@example.com",
	"text": "Hello World! kate@example.com"
}
`

	var rr runner
	rr.run(t)(input1, "/v1/friends/connect", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})(input2, "/v1/friends/connect", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})(input3, "/v1/friends/subscribe", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})(input4, "/v1/friends/block", func(t *testing.T, rsp []byte) {
		var output struct {
			Success bool `json:"success"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
	})(input5, "/v1/friends/recipients", func(t *testing.T, rsp []byte) {
		var output struct {
			Success    bool     `json:"success"`
			Recipients []string `json:"recipients"`
		}
		err := json.Unmarshal(rsp, &output)
		if err != nil {
			t.Errorf("json unmarshal failed")
		}
		if !output.Success {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
		if goset.IsIncluded(output.Recipients, "bob@example.com") {
			t.Errorf("got unexpected response :%v", string(rsp))
		}
		var emails = []string{
			"andy@example.com",
			"tom@example.com",
		}
		for _, email := range emails {
			if !goset.IsIncluded(output.Recipients, email) {
				t.Errorf("got unexpected response :%v", string(rsp))
			}
		}
	})
}
