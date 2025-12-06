package execute

import (
	"encoding/json"
	"lietcode/logic/constant"
	"lietcode/logic/entity"
	"lietcode/logic/repository"
	"log"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

type WorkerMQ struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
	repo  *repository.SubmissionRepository
}

var WorkerMQInstance *WorkerMQ

func Create(SubmissRepo *repository.SubmissionRepository) {

	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		log.Printf("error to connect rabbit mq")
		return
	}
	Channel, errCreateChannel := conn.Channel()
	if errCreateChannel != nil {
		log.Printf("error to connect rabbit mq")
		return
	}
	queue, errQueueDeclare := Channel.QueueDeclare(
		"submission_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if errQueueDeclare != nil {
		log.Printf("error to declare queue")
		return
	}
	WorkerMQInstance = &WorkerMQ{
		conn:  conn,
		ch:    Channel,
		queue: queue,
		repo:  SubmissRepo,
	}

}
func (mq *WorkerMQ) PublishMessageJSON(data interface{}) error {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return mq.ch.Publish(
		"",
		mq.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		},
	)
}
func (mq *WorkerMQ) HandleQueueSubmiss() []func() {
	var funcHandle []func()
	for i := 1; i <= 5; i++ {
		funcHandle = append(funcHandle, func() {
			msgs, err := mq.ch.Consume(
				mq.queue.Name,
				"",
				false,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				log.Printf("errors")
				return
			}
			var forever chan struct{}

			go func() {
				for d := range msgs {
					type SubmissionMessage struct {
						UserId       uint
						SubmissionID uint
						ProblemId    uint
						Language     string
						TestCase     []entity.TestCase
						FileName     string
						SourceCode   string
					}
					var Payload SubmissionMessage
					err := json.Unmarshal(d.Body, &Payload)
					if err != nil {
						log.Printf("Error JSON: %v", err)
						return
					}
					lang := Payload.Language
					codeExecute := CodeExecuteWorker{}
					if lang == "cpp" {
						InputNumber := []string{}
						OutPutNumber := []string{}
						for _, item := range Payload.TestCase {
							InputNumber = append(InputNumber, item.Input)
						}
						for _, item := range Payload.TestCase {
							OutPutNumber = append(OutPutNumber, item.Output)
						}
						outputs, err := codeExecute.ExecuteCppCode(InputNumber, Payload.FileName)
						log.Print(outputs)
						if err != nil {
							log.Print("err")
							return
						}
						data := compareOutputInput(outputs, OutPutNumber)
						errs := mq.repo.Update(Payload.SubmissionID, map[string]interface{}{
							"status": data["reason"],
							"code":   Payload.SourceCode,
							"lang":   Payload.Language,
						})
						log.Print(errs)

						SendEventToClient(Payload.UserId, Payload.SubmissionID, Payload.ProblemId, data)

					}
					if lang == "java" {
						InputNumber := []string{}
						OutPutNumber := []string{}
						for _, item := range Payload.TestCase {
							InputNumber = append(InputNumber, item.Input)
						}
						for _, item := range Payload.TestCase {
							OutPutNumber = append(OutPutNumber, item.Output)
						}
						outputs, err := codeExecute.ExecuteJavaCode(InputNumber, Payload.FileName)
						if err != nil {
							log.Print("err")
							return
						}
						data := compareOutputInput(outputs, OutPutNumber)
						mq.repo.Update(Payload.SubmissionID, map[string]interface{}{
							"status": data["reason"],
							"code":   Payload.SourceCode,
							"lang":   Payload.Language,
						})
						SendEventToClient(Payload.UserId, Payload.SubmissionID, Payload.ProblemId, data)

					}
					if lang == "js" {
						InputNumber := []string{}
						OutPutNumber := []string{}
						for _, item := range Payload.TestCase {
							InputNumber = append(InputNumber, item.Input)
						}
						for _, item := range Payload.TestCase {
							OutPutNumber = append(OutPutNumber, item.Output)
						}
						outputs, err := codeExecute.ExecuteJavaScriptCode(InputNumber, Payload.FileName)
						if err != nil {
							log.Print("err")
							return
						}
						data := compareOutputInput(outputs, OutPutNumber)
						mq.repo.Update(Payload.SubmissionID, map[string]interface{}{
							"status": data["reason"],
							"code":   Payload.SourceCode,
							"lang":   Payload.Language,
						})
						SendEventToClient(Payload.UserId, Payload.SubmissionID, Payload.ProblemId, data)

					}
					if lang == "python" {
						InputNumber := []string{}
						OutPutNumber := []string{}
						for _, item := range Payload.TestCase {
							InputNumber = append(InputNumber, item.Input)
						}
						for _, item := range Payload.TestCase {
							OutPutNumber = append(OutPutNumber, item.Output)
						}
						outputs, err := codeExecute.ExecutePythonCode(InputNumber, Payload.FileName)
						if err != nil {
							log.Print("err")
							return
						}
						data := compareOutputInput(outputs, OutPutNumber)
						mq.repo.Update(Payload.SubmissionID, map[string]interface{}{
							"status": data["reason"],
							"code":   Payload.SourceCode,
							"lang":   Payload.Language,
						})
						SendEventToClient(Payload.UserId, Payload.SubmissionID, Payload.ProblemId, data)
					}

				}
			}()

			<-forever
		})
	}
	return funcHandle

}
func compareOutputInput(outPut map[uint]map[string]interface{}, out []string) map[string]interface{} {

	for index, expectedRaw := range out {
		log.Print(expectedRaw)

		record := outPut[uint(index)+1]

		if record["error"] != nil {
			return map[string]interface{}{
				"testcase": index + 1,
				"data":     record["error"],
				"reason":   constant.Status["Runtime Error"],
			}
		}

		actualOutput, ok := record["output"].(string)
		if !ok {
			return map[string]interface{}{
				"testcase": index + 1,
				"data":     "Output format invalid",
				"reason":   constant.Status["Runtime Error"],
			}
		}

		expected := strings.TrimSpace(expectedRaw)
		actual := strings.TrimSpace(actualOutput)

		if actual != expected {
			return map[string]interface{}{
				"testcase": index + 1,
				"data":     actual,
				"reason":   constant.Status["Wrong Answer"],
			}
		}
	}

	return map[string]interface{}{
		"data":   "All Testcases Accepted",
		"reason": constant.Status["Accepted"],
	}
}
