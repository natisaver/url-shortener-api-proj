package utils

import (
	"fmt"

	"gorm.io/gorm"
)

// in go, there is no try catch
// all errors should always returned to the uppermost layer to be handled

// other than errors, sometimes exceptions still occur, these are known as panics
// to emulated catching, we create this function with the following format:
// ===============================================
// func stopPanic(rollbackFunc func() (err error)) {
// 	if p:= recover(); p != nil {
// 		if (rollbackFunc != nil) rollbackFunc()
// 	} }
// ===============================================

func StopPanic(rollbackFunc func() *gorm.DB) error {
	if p := recover(); p != nil {
		fmt.Println("Panic occurred:", p)
		// in a case where stop panic is used for transactions tx for queries
		// if rollback function was passed in, call it
		if rollbackFunc != nil {
			rollbackFunc()
		}
		return fmt.Errorf("exception occurred: %v", p)
	}
	return nil
}
