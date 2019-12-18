package test

import (
	"testing"

	"github.com/mszsgo/hjson"
)

type Member struct {
	Uid   string `json:"uid"`
	Order *Order
}

type Order struct {
	OrderId string
	Details []*Details
}

type Details struct {
	ProdId string
	Num    int
}

func TestStructToJson(t *testing.T) {
	member := &Member{
		Uid: "12312312",
		Order: &Order{
			OrderId: "23423423",
			Details: []*Details{{
				ProdId: "23423",
				Num:    10,
			}, {
				ProdId: "23424",
				Num:    10,
			}},
		},
	}
	v := hjson.StructToJson(member)
	t.Logf("JSON:%s", v)

}

func TestStructToMap(t *testing.T) {
	member := &Member{
		Uid: "12312312",
		Order: &Order{
			OrderId: "23423423",
			Details: []*Details{{
				ProdId: "23423",
				Num:    10,
			}, {
				ProdId: "23424",
				Num:    10,
			}},
		},
	}
	var v *map[string]interface{}
	hjson.StructToMap(member, &v)
	t.Log(v)
}
