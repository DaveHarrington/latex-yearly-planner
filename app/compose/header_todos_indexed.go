package compose

import (
	"fmt"
	"strconv"

	"github.com/kudrykv/latex-yearly-planner/app/components/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func HeaderTodosIndexed(cfg config.Config, tpls []string) (page.Modules, error) {
	if len(tpls) != 1 {
		return nil, fmt.Errorf("exppected two tpls, got %d %v", len(tpls), tpls)
	}

	modules := make(page.Modules, 0, 101)
	modules = append(modules, page.Module{
		Cfg: cfg,
		Tpl: tpls[0],
		Body: header.Header{
			Left: header.Items{
				header.NewIntItem(cfg.Year),
				header.NewTextItem("Todos Index").Ref(true),
			},
			Right: header.Items{
				header.NewTextItem("Notes").RefText("Notes Index"),
			},
		},
	})

	for i := 1; i <= 100; i++ {
		right := header.Items{}

		if i > 2 {
			right = append(right, header.NewTextItem("Todo "+strconv.Itoa(i-1)))
		}

		if i < 100 {
			right = append(right, header.NewTextItem("Todo "+strconv.Itoa(i+1)))
		}

		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: header.Header{
				Left: header.Items{
					header.NewIntItem(cfg.Year),
					header.NewTextItem("Todos Index"),
					header.NewTextItem("Todo " + strconv.Itoa(i)).Ref(true),
				},
				Right: right,
			},
		})
	}

	return modules, nil
}

func HeaderTodosIndexed2(cfg config.Config, tpls []string) (page.Modules, error) {
	if len(tpls) != 1 {
		return nil, fmt.Errorf("exppected two tpls, got %d %v", len(tpls), tpls)
	}

	modules := make(page.Modules, 0, 101)
	modules = append(modules, page.Module{
		Cfg: cfg,
		Tpl: tpls[0],
		Body: map[string]interface{}{
			"Todos": "Todos Index",
			"Cells": header.Items{
				header.NewCellItem("Calendar"),
				header.NewCellItem("To Do").Refer("Todos Index").Select(),
				header.NewCellItem("Notes").Refer("Notes Index"),
			},
			"Months":   MonthsToCellItems(cfg.WeekStart, calendar.NewYearInMonths(cfg.Year).Reverse()),
			"Quarters": QuartersToCellItems(calendar.NewYearInQuarters(cfg.Year).Reverse()),
		},
	})

	for i := 1; i <= 100; i++ {
		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: map[string]interface{}{
				"Todos": "Todo " + strconv.Itoa(i),
				"Cells": header.Items{
					header.NewCellItem("Calendar"),
					header.NewCellItem("To Do").Refer("Todos Index"),
					header.NewCellItem("Notes").Refer("Notes Index"),
				},
				"Months":   MonthsToCellItems(cfg.WeekStart, calendar.NewYearInMonths(cfg.Year).Reverse()),
				"Quarters": QuartersToCellItems(calendar.NewYearInQuarters(cfg.Year).Reverse()),
			},
		})
	}

	return modules, nil
}
