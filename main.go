package main

import (
		"encoding/json"
			"net/http"
				"sync"

					"github.com/gorilla/mux"
				)

				type Note struct {
						Title string `json:"title"`
							Description string `json:"description"`
						}

						var noteStore = make(map[string]Note)
						var id int
						var mutex = &sync.Mutex{}

						func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
								var notes []Note
									mutex.Lock()
										for _, value := range noteStore {
													notes = append(notes, value)
														}
															mutex.Unlock()
																
																json.NewEncoder(w).Encode(notes)
															}

															func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
																	var note Note
																		err := json.NewDecoder(r.Body).Decode(&note)

																			if err != nil {
																						panic(err)
																							}

																								mutex.Lock()
																									noteStore[string(id)] = note
																										mutex.Unlock()

																											id++
																												json.NewEncoder(w).Encode(note)
																											}

																											func main() {
																													r := mux.NewRouter().StrictSlash(false)
																														r.HandleFunc("/notes", GetNoteHandler).Methods("GET")
																															r.HandleFunc("/notes", PostNoteHandler).Methods("POST")

																																// New handler for static files
																																	fs := http.FileServer(http.Dir("/root/gotest/static"))
																																		r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

																																			server := &http.Server{
																																						Addr:    ":8080",
																																								Handler: r,
																																									}

																																										server.ListenAndServe()
																																									}

