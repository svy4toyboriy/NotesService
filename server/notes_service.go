package server

import (
	"SportsTestWork/model"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func (s *Server) NotesPost(w http.ResponseWriter, r *http.Request) {

	var note model.Note
	if err := s.decode(w, r, &note); err != nil {
		s.error(w, r, err, http.StatusBadRequest)
		return
	}

	var InsertID int
	err := s.db.QueryRow("INSERT INTO notes (content) VALUES ($1) RETURNING id", note.Content).Scan(&InsertID)
	if err != nil {
		log.Println(err)
	}

	newNoteId := struct {
		ID int `json:"id"`
	}{
		ID: InsertID,
	}

	s.respond(w, r, newNoteId, http.StatusOK)
}

func (s *Server) NotesGet(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Query("SELECT id, content FROM notes")
	if err != nil {
		log.Println(err)
	}

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.ID, &note.Content); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Failed to iterate over rows", http.StatusInternalServerError)
		return
	}

	s.respond(w, r, notes, http.StatusOK)
}

func (s *Server) NotesDelete(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`^/notes/(\d+)$`)
	matches := re.FindStringSubmatch(r.URL.Path)

	if len(matches) == 0 {
		http.Error(w, "Invalid URL", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	result, err := s.db.Exec("DELETE FROM notes WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	rowId, _ := result.RowsAffected()
	var status string
	if rowId == 0 {
		status = "not deleted, note with id = " + strconv.Itoa(id) + " doesn't exist"
	} else {
		status = "deleted"
	}

	queryStatus := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	s.respond(w, r, queryStatus, http.StatusOK)
}
