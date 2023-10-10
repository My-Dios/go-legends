package migration

import (
	config "github.com/My-Dios/go-legends/config/db"
	MsCustomerMigration "github.com/My-Dios/go-legends/migrations/ms-customer"
	MsItemMigration "github.com/My-Dios/go-legends/migrations/ms-item"
)

func Migrate() {
	config.InitDB()
	db := config.GetDB()

	err := MsItemMigration.Create(db)
	if err != nil {
		panic("Failed to create MsItem table")
	}

	err = MsCustomerMigration.Create(db)
	if err != nil {
		panic("Failed to create MsCustomer table")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database connection")
	}

	if err := sqlDB.Close(); err != nil {
		panic("Failed to close database connection")
	}
}
