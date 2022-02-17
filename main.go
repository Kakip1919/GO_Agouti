package main

import (
	"bufio"
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()
	arr, i := fromFile()
	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	for a := 0; a < i; a++ {
		if err := page.Navigate("https://twitter.com/"); err != nil {
			log.Fatalf("Failed to navigate:%v", err)
		}
		arr[a] = strings.TrimSpace(arr[a])
		time.Sleep(3 * time.Second)
		page.FindByXPath("/html/body/div[1]/div/div/div[2]/main/div/div/div[1]/div[1]/div/div[3]/div[5]/a/div").Click()
		time.Sleep(3 * time.Second)
		page.FindByXPath("/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div[5]/label/div/div[2]/div/input").Fill(arr[a])
		time.Sleep(33 * time.Second)
		fmt.Println("入力中です")
		page.FindByXPath("/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div[6]/div").Click()
		fmt.Println(arr[a])
		time.Sleep(3 * time.Second)
		page.FindByXPath("//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div[3]/div/label/div/div[2]/div[1]/input").Fill(arr[a])
		time.Sleep(3 * time.Second)
		page.FindByXPath("//*[@id=\"layers\"]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div[1]/div").Click()
		time.Sleep(3 * time.Second)
		page.FindByXPath("//*[@id=\"react-root\"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/div[2]/div[1]/div/div/div/div[1]/div/div[2]/div/div[2]/div/a/div[4]/div").Click()
		page.FindByXPath("//*[@id=\"react-root\"]/div/div/div[2]/main/div/div/div/div[1]/div/div[2]/div/div/div/div/div[1]/div[2]/a/div/span/span").Click()
		time.Sleep(3 * time.Second)
		page.FindByXPath("//*[@id=\"layers\"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div/div[2]/div[4]/label/div/div[2]/div/textarea").Fill("")
		page.FindByXPath("//*[@id=\"layers\"]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div/div[1]/div/div/div/div/div/div[3]/div").Click()
	}
}

func fromFile() ([300]string, int) {
	data, _ := os.Open("Twitter.txt")
	defer data.Close()
	scanner := bufio.NewScanner(data)
	i := 0
	var arr [300]string
	for scanner.Scan() {
		i++
		arr[i] = scanner.Text()
	}
	fmt.Println(i)
	return arr, i
}
