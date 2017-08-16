package main
import (
	"io"
	"net/http"
	"strings"
)
func main(){
	http.HandleFunc("/", Cookie)
	http.HandleFunc("/2", Cookie2)
	http.ListenAndServe(":8080", nil)
}
// 传统的cookie设置方法,
func Cookie(w http.ResponseWriter,r *http.Request){
	ck:=&http.Cookie{
		Name:"mycookie",
		Value:"hello",
		Path:"/",
		Domain:"localhost",//域名
		MaxAge:120,//120秒
	}
	http.SetCookie(w, ck)
	//注意获取cookie的值是r，刷新一次才可以
	ck2,err:=r.Cookie("mycookie")
	if err!=nil{
		io.WriteString(w, err.Error())
		return 
	}
	io.WriteString(w,ck2.Value)
}
func Cookie2(w http.ResponseWriter,r *http.Request )  {
		ck:=&http.Cookie{
		Name:"mycookie",
		Value:"hello world",//注意在这里空格是非法字符,常规处理替换空格
		Path:"/2",
		Domain:"localhost",
		MaxAge:120,//120秒
	}
	//ck.String -----ck 转化string
	// w.Header().Set("set-cookie", ck.String())//name,value中不含有空格
	w.Header().Set("Set-Cookie",strings.Replace(ck.String()," ", "%20", -1))//处理name，value中含有空格的办法。
	ck2,err:=r.Cookie("mycookie")
	if err!=nil{
		io.WriteString(w, err.Error())
		return 
	}
	io.WriteString(w,ck2.Value)
}
