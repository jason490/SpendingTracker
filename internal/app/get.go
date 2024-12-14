package app

import "time"

func (a *App) GetMonthMoneySpent(month string) string {
    return "Success"
}

func (a *App) GetRangeDataSpent(start time.Time, end time.Time) {
    // time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")
}

func (a *App) GetAllSpentMoney() string {
    return "Success"
}

func (a *App) GetAllInvestments() string {
    return "Success"
}

func (a *App) GetCurrMonthIncome(month string) string {
    return "Success"
}

func (a *App) GetRangeDataIncome(start time.Time, end time.Time) {
    // time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")
}

func (a *App) GetAllIncome() string {
    return "Success"
}
