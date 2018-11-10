package main;

import (
    "errors"
    "fmt"
    "reflect"
)


type Db struct {
    Port int
    Host string
    pw   string
}

type Conf struct {
    Op       *string `json:"jsonop" xml:"xmlOpName"`
    Charlist *string
    Length   *int
    Num      *int
    Output   *string
    Input    *string
    hidden   *string
    Db
}

func (this Conf) SayOp(subname string) string {
    return *this.Op + subname
}

func (this Conf) getDbConf() Db {
    return this.Db
}

func main() {

    // create Conf instance
    conf := Conf{}

    opName := "create"
    conf.Op = &opName
    conf.Port = 3308

    fmt.Printf("conf.Port=%d\n\n", conf.Port)

    // structure information
    t := reflect.TypeOf(conf)

    // value information 
    v := reflect.ValueOf(conf)

    printStructField(&t)

    callMethod(&v, "SayOp", []interface{}{" Db"}) //[]interface{}{" Db"} is args array of the method "SayOp"

    // panic: reflect: Call of unexported method
    //callMethod(&v, "getDbConf", []interface{}{})

    getTag(&t, "Op", "json")
    getTag(&t, "Op", "xml")
    getTag(&t, "nofield", "json")
}

// case1: traverse structure fiele names
func printStructField(t *reflect.Type) {
    fieldNum := (*t).NumField()
    for i := 0; i < fieldNum; i++ {
        fmt.Printf("conf's field: %s\n", (*t).Field(i).Name)
    }
    fmt.Println("")
}

// case2: call struct methods:
func callMethod(v *reflect.Value, method string, params []interface{}) {
    f := (*v).MethodByName(method) //get method instance of structure 
    if f.IsValid() {
        args := make([]reflect.Value, len(params)) //make a map as args for method 
        for k, param := range params {
            args[k] = reflect.ValueOf(param)
        }
        // call method here 
        ret := f.Call(args)
        if ret[0].Kind() == reflect.String {
            fmt.Printf("%s Called result: %s\n", method, ret[0].String())
        }
    } else {
        fmt.Println("can't call " + method)
    }
    fmt.Println("")
}

// case3: get tag of struct
func getTag(t *reflect.Type, field string, tagName string) {
    var (
        tagVal string
        err    error
    )
    fieldVal, ok := (*t).FieldByName(field)
    if ok {
        tagVal = fieldVal.Tag.Get(tagName)
    } else {
        err = errors.New("no field named:" + field)
    }

    fmt.Printf("get struct[%s] tag[%s]: %s, error:%v\n", field, tagName, tagVal, err)
    fmt.Println("")
}

