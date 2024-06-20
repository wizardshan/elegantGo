
```go
func (req *DateTimeRangeField) UnmarshalJSON(b []byte) error {
    startStr := times[0]
    endStr := times[1]
    
    if startStr != "" {
        var err error
        req.Start, err = time.Parse(time.DateTime, startStr)
        if err != nil {
            return errors.New(fmt.Sprintf(fmt.Sprintf("the time layout should be `%s`", time.DateTime)))
        }
    }
    
    if endStr != "" {
        var err error
        req.End, err = time.Parse(time.DateTime, endStr)
        if err != nil {
            return errors.New(fmt.Sprintf(fmt.Sprintf("the time layout should be `%s`", time.DateTime)))
        }
    }
}
```

```go
func (req *DateTimeRangeField) UnmarshalJSON(b []byte) error {
    startStr := times[0]
    endStr := times[1]

    layout := time.DateTime
    layoutErrMsg := fmt.Sprintf("the time layout should be `%s`", layout)
    if startStr != "" {
        var err error
        req.Start, err = time.Parse(layout, startStr)
        if err != nil {
            return errors.New(layoutErrMsg)
        }
    }
    
    if endStr != "" {
        var err error
        req.End, err = time.Parse(layout, endStr)
        if err != nil {
            return errors.New(layoutErrMsg)
        }
    }
}
```

```go
func (req *DateTimeRangeField) UnmarshalJSON(b []byte) error {
    var err error
    if req.Start, err = req.parse(times[0]); err != nil {
        return err
    }
    
    if req.End, err = req.parse(times[1]); err != nil {
        return err
    }
}

func (req *DateTimeRangeField) parse(value string) (time.Time, error) {
    var t time.Time
    if value == "" {
        return t, nil
    }
    layout := time.DateTime
    layoutErrMsg := fmt.Sprintf("the time layout should be `%s`", layout)

    var err error
    t, err = time.Parse(layout, value)
    if err != nil {
        return t, errors.New(layoutErrMsg)
    }
    
    return t, nil
}
```


代码片段重复
    方法内重复 
        1、封装成方法内全局变量或常量
        2、封装成独立方法
    结构体内跨方法重复
        1、封装成独立方法
        2、封装成包全局变量或常量
        3、封装成独立结构体
    跨结构体重复
        1、封装成包全局变量或常量
        2、封装成独立结构体
