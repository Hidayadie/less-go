package customer


import (
	//"math"
	//"strings"
)

type Customer struct {
	CustomerID   	string	`json:"customer_id"`
	CompanyName  	string	`json:"company_name"`
	ContanctTitle	string	`json:"contant_title"`
	Address 			string	`json:"address"`
	City					string	`json:"city"`
	Region				string 	`json:"region"`
	PostalCode		string 	`json:"postal_code"`
	Country				string 	`json:"country"`
	Phone 				string	`json:"phone"`
	Fax						string	`json:"fax"`

	/*SupplierID   *int    `json:"supplier_id"`
	CategoryID   *int    `json:"category_id"`
	UnitPrice    float64 `json:"unit_price"`
	UnitsInStock int     `json:"units_in_stock"`
	Discontinued int     `json:"discontinued"`
	*/ 
}

func (c Customer) Valid() bool {
	return true //malasss
	/*return c.CustomerID > 0 && c.CustomerID <= 32767 && strings.TrimSpace(c.CompanyName) != "" && len([]rune(c.CompanyName)) <= 40  &&
		!math.IsNaN(c.UnitPrice) && !math.IsInf(p.UnitPrice, 0) && p.UnitPrice >= 0 && p.UnitPrice <= math.MaxFloat32 &&
		c.UnitsInStock >= 0 && p.UnitsInStock <= 32767 && (p.Discontinued == 0 || p.Discontinued == 1) && validOptionalID(p.SupplierID) && validOptionalID(p.CategoryID)
*/}

func validOptionalID(id *int) bool { return id == nil || (*id > 0 && *id <= 32767) }
