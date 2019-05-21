package fileupload

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEvalGETLinuxToWindows(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	tc.SetInput("method","GET")	;
tc.SetInput("destination","eyJraWQiOiIyMDE5LTA1LTE4VDIyOjQ2OjE0LjY2OS5lYyIsInR5cCI6IkpXVCIsImFsZyI6IkVTMjU2In0.eyJzdWIiOiJwd3lvc2kiLCJ1cm46Y29tOmNlcm5lcjphdXRob3JpemF0aW9uOmNsYWltcyI6eyJ2ZXIiOiIxLjAiLCJ0bnQiOiIwYjhhMDExMS1lOGU2LTRjMjYtYTkxYy01MDY5Y2JjNmIxY2EiLCJhenMiOiJ1c2VyXC9QYXRpZW50LnJlYWQgdXNlclwvQXBwb2ludG1lbnQucmVhZCB1c2VyXC9Eb2N1bWVudFJlZmVyZW5jZS53cml0ZSB1c2VyXC9FbmNvdW50ZXIucmVhZCJ9LCJhenAiOiI1MDlkYjViYy1kZTY0LTRlZjItOTlkYy1mZDRkMzBlNWVmMGIiLCJpc3MiOiJodHRwczpcL1wvYXV0aG9yaXphdGlvbi5zYW5kYm94Y2VybmVyLmNvbVwvIiwiZXhwIjoxNTU4MzUyOTMwLCJpYXQiOjE1NTgzNTIzMzAsImp0aSI6IjQ2ODE4YjNiLTJmM2MtNDNlMC1iODA3LTdjMGY2ZTIzMTUyZiIsInVybjpjZXJuZXI6YXV0aG9yaXphdGlvbjpjbGFpbXM6dmVyc2lvbjoxIjp7InZlciI6IjEuMCIsInByb2ZpbGVzIjp7InNtYXJ0LXYxIjp7ImF6cyI6InVzZXJcL1BhdGllbnQucmVhZCB1c2VyXC9BcHBvaW50bWVudC5yZWFkIHVzZXJcL0RvY3VtZW50UmVmZXJlbmNlLndyaXRlIHVzZXJcL0VuY291bnRlci5yZWFkIn19LCJjbGllbnQiOnsibmFtZSI6Illvc2kgU3RhZ2luZyIsImlkIjoiNTA5ZGI1YmMtZGU2NC00ZWYyLTk5ZGMtZmQ0ZDMwZTVlZjBiIn0sInVzZXIiOnsicHJpbmNpcGFsIjoicHd5b3NpIiwicGVyc29uYSI6InByb3ZpZGVyIiwiaWRzcCI6IjBiOGEwMTExLWU4ZTYtNGMyNi1hOTFjLTUwNjljYmM2YjFjYSIsInByaW5jaXBhbFVyaSI6Imh0dHBzOlwvXC9taWxsZW5uaWEuc2FuZGJveGNlcm5lci5jb21cL2luc3RhbmNlXC8wYjhhMDExMS1lOGU2LTRjMjYtYTkxYy01MDY5Y2JjNmIxY2FcL3ByaW5jaXBhbFwvMDAwMC4wMDAwLjAwNDkuMUVDNyIsImlkc3BVcmkiOiJodHRwczpcL1wvbWlsbGVubmlhLnNhbmRib3hjZXJuZXIuY29tXC9hY2NvdW50c1wvZmhpcnBsYXkudGVtcF9yaG8uY2VybmVyYXNwLmNvbVwvMGI4YTAxMTEtZThlNi00YzI2LWE5MWMtNTA2OWNiYzZiMWNhXC9sb2dpbiJ9LCJ0ZW5hbnQiOiIwYjhhMDExMS1lOGU2LTRjMjYtYTkxYy01MDY5Y2JjNmIxY2EifX0.sItFydiJAK7Q4E05GNSa6q72Yfk7u7uvgQD7m8L8LOsLfGRB_85O9pUnXKohW3YRV2ccKY-_nFE7UUH6ARDZaw");
	
	//setup attrs

	act.Eval(tc)

	//check result attr
}

