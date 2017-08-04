package main

import(
	"net/http";
	"time";
	"fmt";
	"html/template"
)

type CurrentDate struct{
	Date string
	Time string
	Second string
	Week string
}

func indexHandler(w http.ResponseWriter, r * http.Request){
	//http.Redirect(w,r,"/index.html",http.StatusOK)
	t,_ := template.ParseFiles("index.html")
	t.Execute(w,nil)
}

func dateHandler(w http.ResponseWriter, r *http.Request){
	date := r.FormValue("date")
	fmt.Println(date)
	timeLayout := "2006年01月02日15:04:05"
	theTime,_ := time.ParseInLocation(timeLayout,date,time.UTC)
	cd := CurrentDate{}
	cd.Date = fmt.Sprintf("%d年%s月%s日",theTime.Year(),getMonth(theTime.Month()),getStr(theTime.Day()))
	cd.Time = fmt.Sprintf("%s:%s",getStr(theTime.Hour()),getStr(theTime.Minute()))
	cd.Second = getStr(theTime.Second())
	cd.Week = getWeekDay(theTime.Weekday())
	fmt.Println(theTime)
	fmt.Println(cd)

	t,_ := template.ParseFiles("date.tmpl")
	t.Execute(w,cd)
}

func getStr(day int) string{
	if day < 10{
		return fmt.Sprintf("0%d",day)
	}
	return fmt.Sprintf("%d",day)
}

func getMonth(month time.Month) string{
	if month < 10{
		return fmt.Sprintf("0%d",month)
	}
	return fmt.Sprintf("%d",month)
}

func getWeekDay(week time.Weekday) string{
	switch(week){
		case time.Sunday:
			return "星期天"
		case time.Monday:
			return "星期一"
		case time.Tuesday:
			return "星期二"
		case time.Wednesday:
			return "星期三"
		case time.Thursday:
			return "星期四"
		case time.Friday:
			return "星期五"
		case time.Saturday:
			return "星期六"
	}
	return ""
}

func main(){
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/date",dateHandler)
	http.ListenAndServe(":8989",nil)
}