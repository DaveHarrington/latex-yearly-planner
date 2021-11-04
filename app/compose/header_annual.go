package compose

import (
	"fmt"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func HeaderAnnual(cfg config.Config, tpls []string) (page.Modules, error) {
	if len(tpls) != 1 {
		return nil, fmt.Errorf("exppected one tpl, got %d %v", len(tpls), tpls)
	}

	return page.Modules{{
		Cfg: cfg,
		Tpl: tpls[0],
		Body: header.Header{
			Left: header.Items{
				header.NewIntItem(cfg.Year).Ref(),
			},
			Right: header.Items{
				header.NewTextItem("Notes").RefText("Notes Index"),
			},
		},
	}}, nil
}

func HeaderAnnual2(cfg config.Config, tpls []string) (page.Modules, error) {
	if len(tpls) != 1 {
		return nil, fmt.Errorf("exppected one tpl, got %d %v", len(tpls), tpls)
	}

	return page.Modules{{
		Cfg: cfg,
		Tpl: tpls[0],
		Body: map[string]interface{}{
			"Year": cfg.Year,
			"Cells": header.Items{
				header.NewCellItem("Calendar").Select(),
				header.NewCellItem("To Do").Refer("Todos Index"),
				header.NewCellItem("Notes").Refer("Notes Index"),
			},
			"Months":   MonthsToCellItems(cfg.WeekStart, calendar.NewYearInMonths(cfg.Year).Reverse()),
			"Quarters": QuartersToCellItems(calendar.NewYearInQuarters(cfg.Year).Reverse()),
		},
	}}, nil
}

func MonthsToCellItems(wd time.Weekday, months calendar.YearMonths) []header.CellItem {
	out := make([]header.CellItem, 0, len(months))

	for _, month := range months {
		name := month.Calendar(wd).MonthName().String()
		out = append(out, header.NewCellItem(name[:3]).Refer(name).Selected(month.Selected))
	}

	return out
}

func QuartersToCellItems(quarters calendar.Quarters) []header.CellItem {
	out := make([]header.CellItem, 0, len(quarters))

	for _, quarter := range quarters {
		out = append(out, header.NewCellItem(quarter.Name()).Selected(quarter.Selected))
	}

	return out
}
