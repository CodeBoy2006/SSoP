package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/httplib"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/store?charset=utf8mb4&&loc=Local")
}

//使用默认数据库
func Orm() orm.Ormer {
	files := ormlog()
	orm.DebugLog = orm.NewLog(files)
	orm.Debug = true
	o := orm.NewOrm()
	return o
}

//打开日记文件
func ormlog() *os.File {
	timestamp11 := time.Now().Unix()
	tm := time.Unix(timestamp11, 0)
	dirname := tm.Format("2006-01-02 15")
	dirpath := "./logs/" + dirname
	if errdir := os.MkdirAll(dirpath, 0777); errdir != nil {
		beego.Debug(errdir)
	}
	files, errfile := os.OpenFile(dirpath+"/logs.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if errfile != nil {
		beego.Debug(errfile)
	}
	return files
}

//MD5加密64位
func ToMD5(old string) (neww string) {
	md5pwd := md5.New()
	md5pwd.Write([]byte(old))
	password := md5pwd.Sum(nil)
	neww = hex.EncodeToString(password)
	return neww
}

//执行sql并返回查找数据的数量
func Querynum(a string, arg ...interface{}) (maps []orm.Params, num int64) {
	o := Orm()
	o.Using("read")
	num, err := o.Raw(a, arg...).Values(&maps)
	if err != nil {
		beego.Debug(err)
	}
	return
}

//根据id和当前时间生成hash
func UniqueString(key string) (maps []orm.Params) {
	beego.Debug(key)
	return
}

//执行sql语句
func Query(a string, arg ...interface{}) (maps []orm.Params) {
	o := Orm()
	o.Using("read")
	_, err := o.Raw(a, arg...).Values(&maps)
	if err != nil {
		beego.Debug(err)
	}
	return
}

//执行 update delete操作
func Exec(a string, arg ...interface{}) error {
	o := Orm()
	o.Using("default")
	_, err := o.Raw(a, arg...).Exec()
	if err != nil {
		beego.Debug("exec bug:", err)
	}
	return err
}

// 数据库插入修改
func SqlExcute(table string, data url.Values, where string) (num int64, err error) {
	if table == "" {
		err = errors.New("表名不能为空")
		num = 0
		return
	}
	if data == nil {
		err = errors.New("数据不能为空")
		num = 0
		return
	}
	o := Orm()
	o.Using("default")
	sql := ""
	if where == "" {
		var field []string
		var value []string
		for k, v := range data {
			if v[0] != "" {
				field = append(field, "`"+k+"`")
				value = append(value, "'"+v[0]+"'")
			}
		}
		field_str := strings.Join(field, ",")
		value_str := strings.Join(value, ",")
		sql = "insert into " + table + "(" + field_str + ") values(" + value_str + ")"

	} else {
		var set []string
		for k, v := range data {
			if v[0] != "" {
				set = append(set, "`"+k+"`='"+v[0]+"'")
			} else {
				set = append(set, "`"+k+"`=null")
			}
		}
		set_str := strings.Join(set, ",")
		sql = "update " + table + " set " + set_str + " where " + where
	}
	res, err2 := o.Raw(sql).Exec()
	if err2 != nil {
		err = err2
		num = 0
	} else {
		num, err = res.LastInsertId()
	}
	return
}

// 分页
func Pagechange(now string, num string) (page string, page2 int, nums string) {
	if now == "" {
		now = "1"
	}
	if num == "" {
		num = "10"
	}
	page2, _ = strconv.Atoi(now)
	numint, _ := strconv.Atoi(num)
	pages := (page2 - 1) * numint
	nums = num
	page = strconv.Itoa(pages)
	return
}
func Page(sql string, page2 int, nums string) map[string]interface{} {
	totals := Query(sql)[0]["totals"].(string)
	total, _ := strconv.ParseInt(totals, 10, 64)
	num, _ := strconv.Atoi(nums)
	var paginator = make(map[string]interface{})
	if total == 0 {
		paginator = Paginator(1, num, int64(num))
	} else {
		paginator = Paginator(page2, num, total)
	}
	return paginator
}

//分页
func Paginator(page, prepage int, nums int64) map[string]interface{} {

	var firstpage int //第一页地址
	var lastpage int  //最后一页地址
	var nextpage int  //后一页地址
	var prevpage int  //前一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page > totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		prevpage = page - 1
		firstpage = 1
		lastpage = totalpages
		nextpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		prevpage = page - 3
		firstpage = 1
		lastpage = totalpages
		for i, _ := range pages {
			pages[i] = start + i
		}
		prevpage = page - 1
		nextpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = 1
		lastpage = totalpages
		prevpage = int(math.Max(float64(1), float64(page-1)))
		nextpage = page + 1
		//fMt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages           //页数
	paginatorMap["totalpages"] = totalpages //总页数
	paginatorMap["firstpage"] = firstpage   //首页
	paginatorMap["lastpage"] = lastpage     //末页
	paginatorMap["prevpage"] = prevpage     //上一页
	paginatorMap["nextpage"] = nextpage     //下一页
	paginatorMap["currpage"] = page
	paginatorMap["nums"] = nums
	return paginatorMap
}
func ClearLog() {
	go func() {
		for {
			now := time.Now() // 计算下一个零点
			next := now.Add(time.Hour * 1)
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			//ClearS()
		}
	}()
}
