package controllers

import (
	"GOMongoDB/models"
	"GOMongoDB/util"
	"context"
	"encoding/json"
	"net/http"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {
	var netflix models.Netflix
	json.NewDecoder(r.Body).Decode(&netflix)
	util.Collection.InsertOne(context.TODO(), netflix)
	w.WriteHeader(http.StatusCreated)
}