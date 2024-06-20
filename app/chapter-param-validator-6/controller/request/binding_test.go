package request

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

// go test -v
// go test -run TestBindingDateTimeRangeField -v

func TestBindingDateTimeRangeField(t *testing.T) {
	type Object struct {
		Filter struct {
			Time DateTimeRangeFieldV5
		}
	}
	var obj1 Object
	path := "/"
	req, _ := http.NewRequest("GET", path, nil)
	err := binding.Query.Bind(req, &obj1)
	assert.NoError(t, err)
	assert.Equal(t, false, obj1.Filter.Time.StartAble())
	assert.Equal(t, false, obj1.Filter.Time.EndAble())

	var obj2 Object
	path = "/?Filter={\"Time\":\"2024-01-01 00:00:00,2024-07-01 00:00:00\"}"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj2)
	assert.NoError(t, err)

	startTime, _ := time.Parse(time.DateTime, "2024-01-01 00:00:00")
	assert.Equal(t, startTime, obj2.Filter.Time.Start)
	endTime, _ := time.Parse(time.DateTime, "2024-07-01 00:00:00")
	assert.Equal(t, endTime, obj2.Filter.Time.End)
	assert.Equal(t, true, obj2.Filter.Time.StartAble())
	assert.Equal(t, true, obj2.Filter.Time.EndAble())

	var obj3 Object
	path = "/?Filter={\"Time\":\",2024-07-01 00:00:00\"}"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj3)
	assert.NoError(t, err)
	assert.Equal(t, true, obj3.Filter.Time.Start.IsZero())
	assert.Equal(t, endTime, obj3.Filter.Time.End)
	assert.Equal(t, false, obj3.Filter.Time.StartAble())
	assert.Equal(t, true, obj3.Filter.Time.EndAble())

	var obj4 Object
	path = "/?Filter={\"Time\":\"2024-01-01 00:00:00,\"}"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj4)
	assert.NoError(t, err)
	assert.Equal(t, startTime, obj4.Filter.Time.Start)
	assert.Equal(t, true, obj4.Filter.Time.End.IsZero())
	assert.Equal(t, true, obj4.Filter.Time.StartAble())
	assert.Equal(t, false, obj4.Filter.Time.EndAble())
}
