package compose

import (
	"fmt"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func HeaderMonthly(cfg config.Config, tpls []string) (page.Modules, error) {
	if len(tpls) != 1 {
		return nil, fmt.Errorf("exppected one tpl, got %d %v", len(tpls), tpls)
	}

	modules := make(page.Modules, 0, 12)

	for month := time.January; month <= time.December; month++ {
		right := header.Items{}
		if month > time.January {
			right = append(right, header.NewTextItem((month - 1).String()))
		}

		if month < time.December {
			right = append(right, header.NewTextItem((month + 1).String()))
		}

		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: header.Header{
				Left: header.Items{
					header.NewIntItem(cfg.Year),
					header.NewMonthItem(month).Ref(),
				},
				Right: right,
			},
		})
	}

	return modules, nil
}

func HeaderMonthly2(cfg config.Config, tpls []string) (page.Modules, error) {
	if len(tpls) != 1 {
		return nil, fmt.Errorf("exppected one tpl, got %d %v", len(tpls), tpls)
	}

	modules := make(page.Modules, 0, 12)

	for month := time.January; month <= time.December; month++ {
		fdom := calendar.DayTime{Time: time.Date(cfg.Year, month, 1, 0, 0, 0, 0, time.Local)}

		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: map[string]interface{}{
				"Month": fdom,
				"Date":  fdom,
				"Cells": header.Items{
					header.NewCellItem("Calendar"),
					header.NewCellItem("To Do").Refer("Todos Index"),
					header.NewCellItem("Notes").Refer("Notes Index"),
				},
				"Months":   MonthsToCellItems(cfg.WeekStart, calendar.NewYearInMonths(cfg.Year).Selected(fdom).Reverse()),
				"Quarters": QuartersToCellItems(calendar.NewYearInQuarters(cfg.Year).Reverse()),
			},
		})
	}

	return modules, nil
}
