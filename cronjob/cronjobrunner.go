package cronjob

import (
   "fmt"
   "net/http"
   "time"

   "github.com/robfig/cron"
)

func InitCronScheduler() *cron.Cron {
    c := cron.New()
 
    c.AddFunc("@every 00h00m10s", BackUpLocalDataCall)
 
    c.Start()
    fmt.Println("Cron scheduler initialized")
    return c
 }
 
 func BackUpLocalDataCall() {

    
 
    fmt.Println("Cron job called")
    //add your codes to backup or any utilities
 }
 

 
 // BackUpLocalData handles HTTP requests and performs a backup
 func BackUpLocalData(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "calling the API after every 10 seconds at %s", time.Now().Format("2006-01-02 15:04:05"))
 }