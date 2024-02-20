package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

type Valyuta struct {
	ID        int    `json:"id"`
	Code      string `json:"Code"`
	Ccy       string `json:"Ccy"`
	CcyNm_RU  string `json:"CcyNm_RU"`
	CcyNm_UZ  string `json:"CcyNm_UZ"`
	CcyNm_UZC string `json:"CcyNm_UZC"`
	CcyNm_EN  string `json:"CcyNm_EN"`
	Nominal   string `json:"Nominal"`
	Rate      string `json:"Rate"`
	Diff      string `json:"Diff"`
	Date      string `json:"Date"`
}

func main() {
	date := flag.String("date", "", "YYYY-MM-DD")
	flag.Parse()

	var URL string

	if *date == "" {
		today := time.Now().Format("2006-01-02")
		URL = fmt.Sprintf("https://cbu.uz/ru/arkhiv-kursov-valyut/json/all/%s/", today)
	} else {
		URL = fmt.Sprintf("https://cbu.uz/ru/arkhiv-kursov-valyut/json/all/%s/", *date)
	}

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var valyutas []Valyuta
	err = json.Unmarshal(body, &valyutas)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	defer w.Flush()

	fmt.Fprintf(w, "\tID\tCode\tCcy\tCcyNm_RU\tCcyNm_UZ\tCcyNm_UZC\tCcyNm_EN\tNominal\tRate\tDiff\tDate\n")

	for _, valyuta := range valyutas {
		fmt.Fprintf(w, "\t%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		    valyuta.ID,
		    valyuta.Code,
			valyuta.Ccy,
			valyuta.CcyNm_RU,
			valyuta.CcyNm_UZ,
			valyuta.CcyNm_UZC,
			valyuta.CcyNm_EN,
			valyuta.Nominal,
			valyuta.Rate,
			valyuta.Diff,
			valyuta.Date,
		)
	}
}
