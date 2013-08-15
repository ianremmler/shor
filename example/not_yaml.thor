# Based on the YAML example at http://www.yaml.org/start.html

invoice: 34843
date: "2001-01-23"
bill-to: {
	id: "id001"
	given: "Chris"
	family: "Dumars"
	address: {
		lines: "458 Walkman Dr.\nSuite #292"
		city: "Royal Oak"
		state: "MI"
		postal: 48046
	}
}
ship-to: { id: "id001" }
product: {
	sku: "BL394D"
	quantity: 4
	description: "Basketball"
	price: 450.00
}
product: {
	sku: "BL4438H"
	quantity: 1
	description: "Super Hoop"
	price: 2392.00
}
tax: 251.42
total: 4443.52
comments: "Late afternoon is best.\nBackup contact is Nancy\nBillsmer @ 338-4338."
