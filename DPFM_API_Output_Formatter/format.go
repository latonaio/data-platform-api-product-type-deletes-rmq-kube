package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToProductType(rows *sql.Rows) (*ProductType, error) {
	defer rows.Close()
	productType := ProductType{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&productType.ProductType,
			&productType.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &productType, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &productType, nil
	}

	return &productType, nil
}
