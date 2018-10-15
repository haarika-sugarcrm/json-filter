package json_filter

import ( 
	sq "github.com/Masterminds/squirrel"
	"fmt"
	"encoding/json"
	"errors"
)

// func main() {

// 	db, err := sql.Open("mysql",
// 		"root:root@tcp(127.0.0.1:3306)/Db1")
// 	if err != nil {
// 		fmt.Errorf("Error!!")
// 	}
// 	defer db.Close()

// 	users := sq.Select("*").From("SquirrelDb")
// 	fmt.Println("Here ", users.id)
// }

func ApplyFilter(q sq.SelectBuilder, filter []byte) (sq.SelectBuilder, error) {
	var filterMap map[string]interface{}
	err := json.Unmarshal(filter, &filterMap)

	fmt.Println(filterMap)

	for fieldName, value := range filterMap {
		fmt.Printf("key[%s] value[%s]\n", fieldName, value)

		val, ok := value.(map[string]interface{})
		if !ok {
			return q, errors.New("Error parsing the filter")
		}

		for op, v := range val { 
			fmt.Println("KEY : ", fieldName)
			fmt.Println("VAL :", v)

			switch op {
				case "$eq":
					q = q.Where(sq.Eq{fieldName : v})
				case "$ne":
					q = q.Where(sq.NotEq{fieldName : v})
				case "$gt":
					q = q.Where(sq.Gt{fieldName : v})
				case "$gte":
					q = q.Where(sq.GtOrEq{fieldName : v})
				case "$lt":
					q = q.Where(sq.Lt{fieldName : v})
				case "$lte":
					q = q.Where(sq.LtOrEq{fieldName : v})
				case "$isnull":
					q = q.Where(sq.Eq{fieldName : nil})
				case "$isnotnull":
					q = q.Where(sq.NotEq{fieldName : nil})
				case "$in":
					q = q.Where(sq.Eq{fieldName : v})
				case "$notin":
					q = q.Where(sq.NotEq{fieldName : v})
				// case "$like":
				// 	q = q.Where(sq.Like{fieldName : v})
				// case "$between":

				default: 
					return q, fmt.Errorf("Invalid operator passed %s", op)
			}

			sql, args, _ := q.ToSql()
			fmt.Println("Query: ", sql, args)
		}	
	}
	
	return q, err
}
