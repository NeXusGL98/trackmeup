package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/NexusGL98/trackmeup/internal/domain/interval"
	"github.com/NexusGL98/trackmeup/internal/domain/utils"
)

func main() {

	var startDate string
	var endDate string
	// ask for starting date
	r := bufio.NewReader(os.Stdin)

	fmt.Fprint(os.Stdout, "Enter start date (YYYY-MM-DD): ")
	startDate, err := r.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	startDate = strings.TrimSpace(startDate)

	if !utils.IsValidDate(startDate) {
		fmt.Println("Invalid date format")
		os.Exit(1)
	}

	fmt.Fprint(os.Stdout, "Enter end date (YYYY-MM-DD): ")

	endDate, err = r.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	endDate = strings.TrimSpace(endDate)

	if !utils.IsValidDate(endDate) {
		fmt.Println("Invalid date format")
		os.Exit(1)
	}

	timeInterval, err := interval.NewTimeInterval(startDate, endDate, []time.Weekday{
		time.Saturday, time.Sunday,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, inter := range timeInterval.GenerateInterval() {

		fmt.Printf("Jose Gil  %s - time report \n", inter.Format(fmt.Sprintf("%s %s", time.DateOnly, "Mon")))

	}

}
