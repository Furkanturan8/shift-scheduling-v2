package validator

import (
	"shift-scheduling-V2/pkg/utils/structure"
	"sync"
)

var lock = &sync.Mutex{}

var validatorInstance structure.Validator

func GetValidator() structure.Validator {
	if validatorInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if validatorInstance == nil {
			validatorInstance = structure.NewValidator()
		}
	}

	return validatorInstance
}
