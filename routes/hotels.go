package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"bookings-api/data"

	"github.com/go-chi/chi"
)

func Hotels(r chi.Router) {
	r.Get("/", getHotels)
	r.Post("/", postHotel)
	r.Put("/{hotelId}", putHotel)
	r.Delete("/{hotelId}", deleteHotel)
}

func getHotels(w http.ResponseWriter, r *http.Request) {
	hotels, err := data.FetchHotels()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	jbytes, err := json.Marshal(hotels)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jbytes)
}

func postHotel(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	var hotel data.Hotel
	err = json.Unmarshal(bodyBytes, &hotel)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	err = data.CreateHotel(&hotel)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	jbytes, err := json.Marshal(hotel)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jbytes)
}

func putHotel(w http.ResponseWriter, r *http.Request) {
	hotelIdStr := chi.URLParam(r, "hotelId")
	hotelId, err := strconv.ParseInt(hotelIdStr, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	var hotel data.Hotel
	err = json.Unmarshal(bodyBytes, &hotel)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	err = data.UpdateHotel(hotelId, &hotel)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	hotel.Id = hotelId

	jbytes, err := json.Marshal(hotel)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jbytes)
}

func deleteHotel(w http.ResponseWriter, r *http.Request) {
	hotelIdStr := chi.URLParam(r, "hotelId")
	hotelId, err := strconv.ParseInt(hotelIdStr, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	err = data.DeleteHotel(hotelId)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	w.WriteHeader(200)
}
