package main

type Review struct {
	Day string `json: "date", db:"day"`
	Store_id int `json: "store_id", db:"store_id"`
	Outside []int8 `json: "outside", db:"outside"`
	Emp_sys []int8 `json: "emp_sys", db:"emp_sys"`
	Eating []int8 `json: "eating", db:"eating"`
	Merch []int8 `json: "merch", db:"merch"`
	Fountain []int8 `json: "fountain", db:"fountain"`
	Inventory  []int8 `json: "inventory", db:"inventory"`
	Backroom []int8 `json: "backroom", db:"backroom"`
	Restrooms []int8 `json: "restrooms", db:"restrooms"`
	Feedback string `json: "feedback", db:"feedback"`

}


