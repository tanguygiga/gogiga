package dao

import "fmt"

// TodoDaoFactory return a Dao with the given implementation
func TodoDaoFactory(impl string) TodoDao {
	var dao TodoDao
	switch impl {
	case "txt":
		dao = TodoTxtImpl{}
	default:
		dao = nil
		fmt.Print("Not yet implemented !")
	}
	return dao
}
