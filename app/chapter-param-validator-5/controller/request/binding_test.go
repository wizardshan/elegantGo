package request

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

// go test -v
// go test -run TestBindingNumberRangeField -v

func TestBindingSort(t *testing.T) {
	var obj QueryField
	path := "/"
	req, _ := http.NewRequest("GET", path, nil)
	err := binding.Query.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "id", obj.Sort.Value())

	path = "/?Sort=CreateTime"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, "create_time", obj.Sort.Value())
}

func TestBindingOrder(t *testing.T) {
	var obj QueryField
	path := "/"
	req, _ := http.NewRequest("GET", path, nil)
	err := binding.Query.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, true, obj.Order.IsDesc())

	path = "/?Order=DESC"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, true, obj.Order.IsDesc())

	path = "/?Order=ASC"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, false, obj.Order.IsDesc())
}

func TestBindingPage(t *testing.T) {
	var obj QueryField
	path := "/"
	req, _ := http.NewRequest("GET", path, nil)
	err := binding.Query.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, 1, obj.Page.Value())

	path = "/?Page=10"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, 10, obj.Page.Value())

}

func TestBindingPageSize(t *testing.T) {
	var obj QueryField
	path := "/"
	req, _ := http.NewRequest("GET", path, nil)
	err := binding.Query.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, 100, obj.PageSize.Value())

	path = "/?PageSize=10"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj)
	assert.NoError(t, err)
	assert.Equal(t, 10, obj.PageSize.Value())

}

func TestBindingNumbersBySeparator(t *testing.T) {
	type Object struct {
		Filter struct {
			Numbers *NumbersBySeparatorField
		}
	}
	var obj1 Object
	path := "/"
	req, _ := http.NewRequest("GET", path, nil)
	err := binding.Query.Bind(req, &obj1)
	assert.NoError(t, err)
	assert.Equal(t, true, obj1.Filter.Numbers == nil)
	assert.Equal(t, false, obj1.Filter.Numbers.Able())

	var obj2 Object
	path = "/?Filter={\"Numbers\":\"1,2\"}"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj2)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2}, obj2.Filter.Numbers.Values)
	assert.Equal(t, true, obj2.Filter.Numbers.Able())
}

func TestBindingStringsBySeparator(t *testing.T) {
	type Object struct {
		Filter struct {
			Strings *StringsBySeparatorField
		}
	}
	var obj1 Object
	path := "/"
	req, _ := http.NewRequest("GET", path, nil)
	err := binding.Query.Bind(req, &obj1)
	assert.NoError(t, err)
	assert.Equal(t, true, obj1.Filter.Strings == nil)
	assert.Equal(t, false, obj1.Filter.Strings.Able())

	var obj2 Object
	path = "/?Filter={\"Strings\":\"in,out\"}"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj2)
	assert.NoError(t, err)
	assert.Equal(t, []string{"in", "out"}, obj2.Filter.Strings.Values)
	assert.Equal(t, true, obj2.Filter.Strings.Able())
}

func TestBindingNumberRangeField(t *testing.T) {
	type Object struct {
		Filter struct {
			Number *NumberRangeField
		}
	}
	var obj1 Object
	path := "/"
	req, _ := http.NewRequest("GET", path, nil)
	err := binding.Query.Bind(req, &obj1)
	assert.NoError(t, err)
	assert.Equal(t, true, obj1.Filter.Number == nil)
	assert.Equal(t, false, obj1.Filter.Number.StartAble())
	assert.Equal(t, false, obj1.Filter.Number.EndAble())

	var obj2 Object
	path = "/?Filter={\"Number\":\"1,100\"}"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj2)
	assert.NoError(t, err)
	assert.Equal(t, 1, obj2.Filter.Number.Start)
	assert.Equal(t, 100, obj2.Filter.Number.End)
	assert.Equal(t, true, obj2.Filter.Number.StartAble())
	assert.Equal(t, true, obj2.Filter.Number.EndAble())

	var obj3 Object
	path = "/?Filter={\"Number\":\",100\"}"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj3)
	assert.NoError(t, err)
	assert.Equal(t, 0, obj3.Filter.Number.Start)
	assert.Equal(t, 100, obj3.Filter.Number.End)
	assert.Equal(t, false, obj3.Filter.Number.StartAble())
	assert.Equal(t, true, obj3.Filter.Number.EndAble())

	var obj4 Object
	path = "/?Filter={\"Number\":\"100,\"}"
	req, _ = http.NewRequest("GET", path, nil)
	err = binding.Query.Bind(req, &obj4)
	assert.NoError(t, err)
	assert.Equal(t, 100, obj4.Filter.Number.Start)
	assert.Equal(t, 0, obj4.Filter.Number.End)
	assert.Equal(t, true, obj4.Filter.Number.StartAble())
	assert.Equal(t, false, obj4.Filter.Number.EndAble())
}

func TestBindingDateTimeRangeField(t *testing.T) {
	type Object struct {
		Filter struct {
			Time *DateTimeRangeField
		}
	}
	var obj1 Object
	path := "/"
	req, _ := http.NewRequest("GET", path, nil)
	err := binding.Query.Bind(req, &obj1)
	assert.NoError(t, err)
	assert.Equal(t, true, obj1.Filter.Time == nil)
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
