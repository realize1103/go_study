package scarpper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

func scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50&start="

	var jobs []extractedJob
	mainC := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, mainC)
		//extractJobs := getPage(i)
		//jobs = append(jobs, extractJobs...) //...을 하지 않으면 배열 안에 배열이 추가되는 형태[[][][]]이므로 배열+배열은 ...으로 처리
	}

	for i := 0; i < totalPages; i++ {
		extractJobs := <-mainC
		jobs = append(jobs, extractJobs...) //...을 하지 않으면 배열 안에 배열이 추가되는 형태[[][][]]이므로 배열+배열은 ...으로 처리
	}

	writeJobs(jobs)

}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + strconv.Itoa(page*50)
	fmt.Println("Requseting : ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close() //res.Body IO 니까 반드시 닫아주어야 함. 그래서 이 함수가 끝날때 close함수를 호출하도록 해놓음.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
		//job := extractJob(card)
		//jobs = append(jobs, job)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {

	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
	//fmt.Println(id, title, location, salary, summary)
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close() //res.Body IO 니까 반드시 닫아주어야 함. 그래서 이 함수가 끝날때 close함수를 호출하도록 해놓음.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length() //Html())

	})
	return pages
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{
			"https://kr.indeed.com/viewjob?jk=" + job.id,
			job.title,
			job.location,
			job.salary,
			job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed with Status : ", res.StatusCode)
	}
}

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
