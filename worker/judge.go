package worker

import (
	"math/rand"

	"github.com/gyf1214/dboj/model"
)

// fake code
func RunEvaluation() {

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

	}

	return nil
}
