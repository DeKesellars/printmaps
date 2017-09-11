// Delete handler

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/*
deleteMap deletes all data for a given map ID
*/
func deleteMap(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	var pmData PrintmapsData
	var pmErrorList PrintmapsErrorList

	id := params.ByName("id")

	if err := readMetadata(&pmData, id); err != nil {
		if os.IsNotExist(err) {
			appendError(&pmErrorList, "4002", "requested ID not found: "+id, id)
		} else {
			message := fmt.Sprintf("error <%v> at readMetadata(), id = <%s>", err, id)
			http.Error(writer, message, http.StatusInternalServerError)
			log.Printf("Response %d - %s", http.StatusInternalServerError, message)
			return
		}
	}

	if len(pmErrorList.Errors) == 0 {
		// delete map directory
		path := filepath.Join(PathWorkdir, PathMaps, id)
		if err := os.RemoveAll(path); err != nil {
			message := fmt.Sprintf("error <%v> at os.RemoveAll(), path = <%s>", err, path)
			http.Error(writer, message, http.StatusInternalServerError)
			log.Printf("Response %d - %s", http.StatusInternalServerError, message)
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	} else {
		// request not ok, response with error list
		content, err := json.MarshalIndent(pmErrorList, indentPrefix, indexString)
		if err != nil {
			message := fmt.Sprintf("error <%v> at json.MarshalIndent()", err)
			http.Error(writer, message, http.StatusInternalServerError)
			log.Printf("Response %d - %s", http.StatusInternalServerError, message)
			return
		}

		writer.Header().Set("Content-Type", JSONAPIMediaType)
		writer.Header().Set("Content-Length", strconv.Itoa(len(content)))
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(content)
	}
}