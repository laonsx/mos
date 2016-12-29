package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
	"errors"
	"strings"
	"regexp"
)

type DB struct {
	db        *sql.DB
	table 	  string
	key       []string
	value     []string
	where     string
	param     []string
	columnstr string
	pk        string
	orderby   string
	limit     string
	join      string
}

func New(root,password,Database_IP,Database_Port,Database_tb string) (*DB,error) {
	orm:=new(DB)
	db,err:=sql.Open("mysql",root+":"+password+"@tcp("+Database_IP+":"+Database_Port+")/"+Database_tb+"?charset=utf8")
	err=db.Ping()
	if err!=nil{
		return orm,err
	}
	orm.db=db
	return orm,err
}

func (this *DB) DbClose()  {
	this.db.Close()
}

func (this *DB) Fileds(param ...string) *DB {
	this.param = param
	return this
}

func (this *DB) Table(table string) *DB {
	this.table=table
	return this
}

func (this *DB) Key(key ...string) *DB {
	this.key=key
	return this
}

func (this *DB) Value(value ...string) *DB {
	this.value=value
	return this
}

func (this *DB) Where(condition string) *DB {
	if condition==""{
		this.where=condition
		return this
	}
	this.where = fmt.Sprintf("where %v", condition)
	return this
}

func (this *DB) OrderBy(sort string) *DB {
	this.orderby = fmt.Sprintf("ORDER BY %v", sort)
	return this
}

func (this *DB) Limit(size ...int) *DB {
	var end int
	start := size[0]
	if len(size) > 1 {
		end = size[1]
		this.limit = fmt.Sprintf("Limit %d,%d", start, end)
		return this
	}
	this.limit = fmt.Sprintf("Limit %d", start)
	return this
}

func (this *DB) LeftJoin(table, condition string) *DB {
	this.join = fmt.Sprintf("LEFT JOIN %v ON %v", table, condition)
	return this
}

func (this *DB) RightJoin(table, condition string) *DB {
	this.join = fmt.Sprintf("RIGHT JOIN %v ON %v", table, condition)
	return this
}

func (this *DB) Join(table, condition string) *DB {
	this.join = fmt.Sprintf("INNER JOIN %v ON %v", table, condition)
	return this
}

func (this *DB) FullJoin(table, condition string) *DB {
	this.join = fmt.Sprintf("FULL JOIN %v ON %v", table, condition)
	return this
}


func Print(slice map[int]map[string]string)  {
	for i,v:=range slice{
		for key,value:=range v{
			fmt.Printf("%d-%s:%s\r\n",i,key,value)
		}
		fmt.Println(".................")
	}
}

func (this *DB) Insert() (int,error) {
	if this.db==nil{
		return 0,errors.New("db not connect")
	}
	fileValue := "'" + strings.Join(this.value, "','") + "'"
	fileds := "`" + strings.Join(this.key, "`,`") + "`"

	sql := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", this.table, fileds, fileValue)
	result, err := this.db.Exec(sql)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
			}
		}()
		err = errors.New("inster sql failure")
		return 0, err
	}
	i, err := result.LastInsertId()
	s, _ := strconv.Atoi(strconv.FormatInt(i, 10))
	if err != nil {
		err = errors.New("insert failure")
	}
	return s, err

	return 0,nil
}

func (this *DB) Delete()(int,error){
	if this.db==nil{
		return 0,errors.New("db not connect")
	}
	sql:= fmt.Sprintf("DELETE FROM %v %v",this.table,this.where)
	result,err:=this.db.Exec(sql)
	if err!=nil{
		defer func() {
			if err:=recover();err!=nil{
				fmt.Printf("SQL syntax errors")
			}
		}()
		err=errors.New("DELETE sql failure")
		return 0,err
	}
	i,err:=result.RowsAffected()
	s,_:=strconv.Atoi(strconv.FormatInt(i,10))
	if i==0{
		err=errors.New("DELETE failure")
	}
	return s,err
}

func (this *DB) Update() (num int, err error) {
	if this.db == nil {
		return 0, errors.New("mysql not connect")
	}
	var setValue []string
	for i, key := range this.key {
		set := fmt.Sprintf("%v = '%v'", key,this.value[i] )
		setValue = append(setValue, set)
	}
	setData := strings.Join(setValue, ",")
	sql := fmt.Sprintf("UPDATE %v SET %v %v", this.table, setData, this.where)
	result, err := this.db.Exec(sql)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
			}
		}()
		err = errors.New("update sql failure")
		return 0, err
	}
	i, err := result.RowsAffected()
	if err != nil {
		err = errors.New("update failure")
		return 0, err
	}
	s, _ := strconv.Atoi(strconv.FormatInt(i, 10))
	return s, err
}

func (this *DB) Query(sql string) interface{} {
	if this.db == nil {
		return errors.New("mysql not connect")
	}
	var query = strings.TrimSpace(sql)
	s, err := regexp.MatchString(`(?i)^select`, query)
	if err == nil && s == true {
		result, _ := this.db.Query(sql)
		c := QueryResult(result)
		return c
	}
	exec, err := regexp.MatchString(`(?i)^(update|delete)`, query)
	if err == nil && exec == true {
		m_exec, err := this.db.Exec(query)
		if err != nil {
			return err
		}
		num, _ := m_exec.RowsAffected()
		id := strconv.FormatInt(num, 10)
		return id
	}

	insert, err := regexp.MatchString(`(?i)^insert`, query)
	if err == nil && insert == true {
		m_exec, err := this.db.Exec(query)
		if err != nil {
			return err
		}
		num, _ := m_exec.LastInsertId()
		id := strconv.FormatInt(num, 10)
		return id
	}
	result, _ := this.db.Exec(query)

	return result

}

func (this *DB) FindOne() map[int]map[string]string {
	empty := make(map[int]map[string]string)
	if this.db != nil {
		data := this.Limit(1).FindAll()
		return data
	}
	fmt.Printf("mysql not connect\r\n")
	return empty
}

func (this *DB) FindAll() map[int]map[string]string {

	result := make(map[int]map[string]string)
	if this.db == nil {
		fmt.Printf("mysql not connect")
		return result
	}
	if len(this.param) == 0 {
		this.columnstr = "*"
	} else {
		if len(this.param) == 1 {
			this.columnstr = this.param[0]
		} else {
			this.columnstr = strings.Join(this.param, ",")
		}
	}

	query := fmt.Sprintf("Select %v from %v %v %v %v %v", this.columnstr, this.table, this.join, this.where, this.orderby, this.limit)
	rows, err := this.db.Query(query)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("SQL syntax errors ")
			}
		}()
		err = errors.New("select sql failure")
	}
	result = QueryResult(rows)
	return result
}

func QueryResult(rows *sql.Rows) map[int]map[string]string {
	var result = make(map[int]map[string]string)
	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanargs := make([]interface{}, len(values))
	for i := range values {
		scanargs[i] = &values[i]
	}

	var n = 1
	for rows.Next() {
		result[n] = make(map[string]string)
		err := rows.Scan(scanargs...)

		if err != nil {
			fmt.Println(err)
		}

		for i, v := range values {
			result[n][columns[i]] = string(v)
		}
		n++
	}
	return result
}
