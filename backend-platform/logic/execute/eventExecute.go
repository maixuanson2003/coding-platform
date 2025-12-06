package execute

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EventStruct struct {
	conn []InstanceConnect
}
type InstanceConnect struct {
	userId       uint
	problemId    uint
	submissionId uint
	writer       http.ResponseWriter
}

var event EventStruct

func init() {
	event = EventStruct{
		conn: []InstanceConnect{},
	}
}

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	params := mux.Vars(r)
	userId := params["user_id"]
	problemId := params["problem_id"]
	SubmissId := params["submiss_id"]
	idUser, _ := strconv.Atoi(userId)
	idProblem, _ := strconv.Atoi(problemId)
	idSubmiss, _ := strconv.Atoi(SubmissId)
	log.Println("User:", idUser)
	log.Println("Problem:", idProblem)
	log.Println("Submission:", idSubmiss)
	event.conn = append(event.conn, InstanceConnect{
		userId:       uint(idUser),
		problemId:    uint(idProblem),
		submissionId: uint(idSubmiss),
		writer:       w,
	})

	<-r.Context().Done()
}
func SendEventToClient(userId uint, submissionId uint, problemId uint, data map[string]interface{}) {
	var writer http.ResponseWriter
	for _, item := range event.conn {
		log.Print(item)
		if item.userId == userId && item.problemId == problemId && item.submissionId == submissionId {
			writer = item.writer
			break
		}
	}
	if writer == nil {
		log.Println("Writer not found for userId:", userId, "problemId:", problemId)
		return
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}

	// Gửi event tới client
	_, _ = fmt.Fprintf(writer, "data: %s\n\n", string(jsonBytes))
	flusher, ok := writer.(http.Flusher)
	if !ok {
		log.Println("create not ok")
		return
	}
	flusher.Flush()

}
