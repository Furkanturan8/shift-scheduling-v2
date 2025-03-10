package aggregate

import (
	"shift-scheduling-V2/pkg/db/mongo/conditions"
	"shift-scheduling-V2/pkg/db/mongo/utils"

	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestPipe(t *testing.T) {
	pipeline := Pipe(bson.A{},
		Match(Operation(
			conditions.Pipe(
				conditions.ObjectIdMatch(conditions.Condition{
					Key:   "_id",
					Value: "5c7836b73a8de34c78fec399"}),
				conditions.EqualTo(conditions.Condition{
					Key:   "status",
					Value: 1,
				}),
				conditions.StringStartsWith(conditions.Condition{
					Key:   "model",
					Value: "T654",
				}),
			),
		)),
		Project(Operation{
			"name":  1,
			"make":  1,
			"model": 1,
		}),
	)
	utils.PrintJson(pipeline)
}
