package main

import (
    "beegoDIY"
    "log"
    "os"
    "net/http")


func main() {
    //    file, err := os.OpenFile(fmt.Sprintf("/Users/zhengsai/Documents/go/src/goFrameworkDIY/log/%s.txt", time.Now().Format("2006-01-02")), os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
    //    if err != nil{
    //        beegodiy.Critical("open file err : ", err)
    //        return
    //    }

    //    beegodiy.SetLogger(log.New(file, "[文件]", log.Ldate|log.Ltime))
    beegodiy.SetLogger(log.New(os.Stdout, "[beegodiy]", log.Ldate|log.Ltime))
    beegodiy.Trace("start\n")
    //    	http.Handle("/log/", http.StripPrefix("/log/", http.FileServer(http.Dir("/Users/zhengsai/Documents/go/src/goFrameworkDIY/log/"))))

    //	http.Handle("/curl/", http.StripPrefix("/curl/", http.FileServer(http.Dir("/Users/zhengsai/Downloads/curl"))))
    //	http.Handle("/make/", http.StripPrefix("/make/", http.FileServer(http.Dir("/Users/zhengsai/Downloads/make"))))
    log.Fatalln(http.ListenAndServe(":9000", nil))
}
