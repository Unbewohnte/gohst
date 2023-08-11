/*
            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.
*/

// Kasyanov N.A. (Unbewohnte), 2023

package server

import (
	"Unbewohnte/gohst/db"
	"Unbewohnte/gohst/logger"
	"encoding/json"
	"io"
	"net/http"
)

func (s *Server) entityEndpoint(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodDelete:
		// Delete an already existing entity
		defer req.Body.Close()

		// Read body
		body, err := io.ReadAll(req.Body)
		if err != nil {
			logger.Warning("[Server] Failed to read request body to delete an entity: %s", err)
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}

		// Unmarshal JSON
		var entity db.TestEntity
		err = json.Unmarshal(body, &entity)
		if err != nil {
			logger.Warning("[Server] Received invalid entity JSON for deletion: %s", err)
			http.Error(w, "Invalid entity JSON", http.StatusBadRequest)
			return
		}

		err = s.db.DeleteTestEntity(entity.Text)
		if err != nil {
			logger.Error("[Server] Failed to delete %s: %s", entity.Text, err)
			http.Error(w, "Failed to delete entity or TODO contents", http.StatusInternalServerError)
			return
		}

		// Success!
		w.WriteHeader(http.StatusOK)

	case http.MethodPost:
		// Create a new entity
		defer req.Body.Close()
		// Read body
		body, err := io.ReadAll(req.Body)
		if err != nil {
			logger.Warning("[Server] Failed to read request body to create a new entity: %s", err)
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}

		// Unmarshal JSON
		var newEntity db.TestEntity
		err = json.Unmarshal(body, &newEntity)
		if err != nil {
			logger.Warning("[Server] Received invalid entity JSON for creation: %s", err)
			http.Error(w, "Invalid entity JSON", http.StatusBadRequest)
			return
		}

		// Add entity to the database
		err = s.db.CreateTestEntity(newEntity)
		if err != nil {
			http.Error(w, "Failed to create a new entity", http.StatusInternalServerError)
			return
		}

		// Success!
		w.WriteHeader(http.StatusOK)
		logger.Info("[Server] Created a new entity \"%s\"", newEntity.Text)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
