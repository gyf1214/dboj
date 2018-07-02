package worker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zc-staff/dboj/model"
	"github.com/zc-staff/dboj/util"
)

// fake judge
func runEvaluation(code string, eval model.EvaluationInfo) error {
	dataset, err := model.GetDatasetInfo(eval.Dataset)
	if err != nil {
		return err
	}

	fmt.Println("code: ", code)
	fmt.Println("judge: ", eval.Judge.Address)
	fmt.Println("input: ", dataset.Input)
	fmt.Println("answer: ", dataset.Answer)

	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

	eval.Status = rand.Intn(3) + 1
	fmt.Println("result: ", eval.Status)

	return model.UpdateEvaluation(eval)
}

func safeRun(code string, eval model.EvaluationInfo) {
	err := runEvaluation(code, eval)
	if err != nil {
		fmt.Println(err.Error())
		eval.Status = util.SystemError
		model.UpdateEvaluation(eval)
	}
}

func RunSubmition(submit model.SubmitInfo) error {
	evals, err := model.ListEvalution(submit.ID)
	if err != nil {
		return err
	}

	judges, err := model.ListJudge(submit.Language)
	if err != nil {
		return err
	}

	for _, eval := range evals {
		eval.Judge = judges[rand.Intn(len(judges))]
		eval.Status = 0
		eval.Message = ""
		err = model.UpdateEvaluation(eval)
		if err != nil {
			return err
		}

		go safeRun(submit.Code, eval)
	}

	return nil
}
