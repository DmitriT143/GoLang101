package main

import(
	"fmt"
	"time"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gin-gonic/gin"
)

func main() {
	parseShit("nil")
	r := gin.Default()
	r.GET("/hello/:name", hello)
	r.GET("/train/site", site)
	r.GET("/train/:name", train)
	r.Run(":8081")
}

func hello(c *gin.Context){
	name := c.Param("name")
	c.String(http.StatusOK, "Hello, " + name + "!")
}

func train(c *gin.Context){
	name := c.Param("name")
	c.String(http.StatusOK, parseShit(name))
}

func site(c *gin.Context){
	param_1 := c.Query("param_1")
	param_2 := c.Query("param_2")
	c.String(http.StatusOK, parseTuTuChangeAdress(param_1,param_2))
}

func parseShit(types string) string{
	cur_time := time.Now().Format("15:04")
	fmt.Println(cur_time)
	trains_found := 0
	better_output := "This is demo page \n"

	c := colly.NewCollector(colly.AllowedDomains("www.tutu.ru"))
	
	c.OnHTML("tr.desktop__card__yoy03", func(h *colly.HTMLElement){
		dept_time := h.ChildText("a.g-link.desktop__depTimeLink__1NA_N")
		arr_time := h.ChildText("td.t-txt-s.desktop__cell__2cdVW.desktop__range__1Kbxz")
		late := h.ChildText("span.t-nowrap")
		if dept_time > cur_time && trains_found < 5{
			trains_found += 1
			if types == "train"{
				better_output = better_output + "Поезд приезжает на станцию отправки в " + dept_time + " и будет на остановке через " + arr_time + "\n" + late + "\n" 
			}
			fmt.Printf("Поезд приезжает на станцию отправки в %s и будет на остановке через %s \n", dept_time, arr_time)
			if late!=""{
			fmt.Println(late, "\n")	// Самоповтор ниже, тут напомню что на сайте косяк с станциями конечными и начальными, лучше не парсить
		}else {fmt.Println()}
	}
	})


	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting %s", r.URL)
	})

	c.Visit("https://www.tutu.ru/rasp.php?st1=57808&st2=56108.html")
	// Если хочешь, сначала найди нужную комбинацию станций, а потом запиши их сюда вместо st1 и st2
	// Кому я это пишу, а да, преподу
	fmt.Println("Demo dat it is working, now for TU TU")
	if types == "train"{return better_output}
	return "wrong type"
}

func parseTuTuChangeAdress(param_1 string, param_2 string) string{
	fmt.Println(param_1,param_2)
	cur_time := time.Now().Format("15:04")
	fmt.Println(cur_time)
	trains_found := 0
	better_output := "This is demo page \n"

	c := colly.NewCollector(colly.AllowedDomains("www.tutu.ru"))
	
	c.OnHTML("tr.desktop__card__yoy03", func(h *colly.HTMLElement){
		dept_time := h.ChildText("a.g-link.desktop__depTimeLink__1NA_N")
		arr_time := h.ChildText("td.t-txt-s.desktop__cell__2cdVW.desktop__range__1Kbxz")
		late := h.ChildText("span.t-nowrap")
		if dept_time > cur_time && trains_found < 5{
			trains_found += 1
			better_output = better_output + "Поезд приезжает на станцию отправки в " + dept_time + " и будет на остановке через " + arr_time + "\n" + late + "\n" 
			fmt.Printf("Поезд приезжает на станцию отправки в %s и будет на остановке через %s \n", dept_time, arr_time)
			if late!=""{
			fmt.Println(late, "\n")	// Всё равно напишет если что то есть с новой строкой, или новуй строку если нет
		}else {fmt.Println()}
	}
	})


	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting %s", r.URL)
	})

	c.Visit("https://www.tutu.ru/rasp.php?st1="+param_1+"&st2="+param_2)
	// Параметры ищуться с сайта, чтобы открыть сайт нужно на него зайти чтобы потом не заходить сного
	fmt.Println("Demo dat it is working, now for TU TU")
	return better_output
}