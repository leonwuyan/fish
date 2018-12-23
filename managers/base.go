package managers

import "fish/models"

var BankConfig []models.BankCardConfig

func test() {
	SystemInstanse.LoadBankCardConfig()
}
