package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ApplesOrdered    int
	OrangesOrdered   int
	CerealOrdered    int
	promApplesMetric = promauto.NewCounter(prometheus.CounterOpts{
		Name: "apples_purchased",
		Help: "The total number of apples purchased",
	})
	promOrangesMetric = promauto.NewCounter(prometheus.CounterOpts{
		Name: "oranges_purchased",
		Help: "The total number of oranges purchased",
	})
	promCerealMetric = promauto.NewCounter(prometheus.CounterOpts{
		Name: "cereal_purchased",
		Help: "The total number of cereal purchased",
	})
)

func main() {
	http.HandleFunc("/apple", handleAppleOrder)
	http.HandleFunc("/orange", handleAppleOrder)
	http.HandleFunc("/cereal", handleAppleOrder)
	http.HandleFunc("/list", handleList)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":5656", nil)
}

func handleAppleOrder(w http.ResponseWriter, r *http.Request) {
	var str = ""
	var count int
	var err error
	var grocery string

	if strings.Contains(r.URL.Path, "apple") {
		grocery = "apple"
	} else if strings.Contains(r.URL.Path, "orange") {
		grocery = "orange"
	} else if strings.Contains(r.URL.Path, "cereal") {
		grocery = "cereal"
	} else {
		fmt.Fprintf(w, "Item is out of stock")
		return
	}

	countStr, ok := r.URL.Query()["count"]
	if !ok {
		count = 1
		str = fmt.Sprintf("Ordered 1 %s", grocery)
	} else {
		count, err = strconv.Atoi(countStr[0])
		if err != nil {
			fmt.Fprintf(w, "Invalid number of %ss", grocery)
			return
		}
		str = fmt.Sprintf("Ordered %d %ss", count, grocery)
	}
	fmt.Fprintf(w, "%s", str)
	switch grocery {
	case "apple":
		ApplesOrdered += count
		promApplesMetric.Add(float64(count))
	case "orange":
		OrangesOrdered += count
		promOrangesMetric.Add(float64(count))
	case "cereal":
		CerealOrdered += count
		promCerealMetric.Add(float64(count))
	}
}

func handleList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Oranges: %d\nApples: %d\nCereal: %d\n", OrangesOrdered, ApplesOrdered, CerealOrdered)
}
