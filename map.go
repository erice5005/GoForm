package goform

import (
	"log"
	"reflect"
)

const formTag = "goform"

func MapFormToCustomStruct(form Form, targetStruct interface{}) interface{} {
	// errs := []error{}

	fv := reflect.ValueOf(targetStruct)

	answerMap := form.MapAnswersByName()

	for i := 0; i < fv.NumField(); i++ {
		targetTag := fv.Type().Field(i).Tag.Get(formTag)
		if targetTag == "" || targetTag == "-" {
			continue
		}

		targetAnswer := answerMap[targetTag]
		log.Printf("TargetAnswer: %v\n", targetAnswer)
		// TODO: Dynamic conversion based on reflect type. Generate error if the types do not match

	}

	return fv.Elem()

}
