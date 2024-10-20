package test

import (
	"bytes"
	"cloud-disk/core/models"
	"encoding/json"
	"fmt"
	"testing"
	"xorm.io/xorm"
)
import _ "github.com/go-sql-driver/mysql"

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:12345678@tcp(127.0.0.1:3306)/Cloud-data?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)

	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dst.String())

}
