package main

///
import (
	"github.com/getevo/evo"
	"github.com/getevo/evo/apps/rdb"
	"github.com/getevo/evo/lib"
	"github.com/getevo/evo/lib/log"
	"strconv"
)

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

///
func main() {
	evo.Setup()
	rdb.Register()
	err := rdb.CreateConnection("localhost", "mysql", "root:@/cochlear?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	var db = rdb.GetDBO("localhost")
	if db == nil {
		log.Fatal("Null db")
	}
	query := db.Query("SELECT id,name,password FROM users WHERE id = ?")
	parser := &rdb.Parser{
		Params: []rdb.Param{
			{"id", rdb.URL, "numeric"},
		},
	}
	parser.Processor = func(params []string) []string {
		params[0] = strconv.Itoa(lib.ParseSafeInt(params[0]) + 1)
		return params
	}
	query.SetParser(parser)

	evo.Get("/api/test/:id", func(r *evo.Request) {
		data := User{}
		err := query.All(&data, r)
		if err != nil {
			log.Error(err)
		}
		r.WriteResponse(data)
	})

	evo.Run()
}

///
