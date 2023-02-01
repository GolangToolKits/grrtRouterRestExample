package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	//"github.com/gorilla/mux"
	mux "github.com/GolangToolKits/grrt"
)

// Product Product
type Product struct {
	ID    int64
	Sku   string
	Desc  string
	Price float64
}

// Response Response
type Response struct {
	ID      int64
	Success bool
}

// Handler Handler
type Handler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	GetProductWithIDAndSku(w http.ResponseWriter, r *http.Request)
	AddProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}

// StoreHandler StoreHandler
type StoreHandler struct {
}

// New New
func (h *StoreHandler) New() Handler {
	return h
}

// GetProducts GetProducts
func (h *StoreHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	h.SetContentType(w)
	var rtn []Product
	var p1 Product
	p1.ID = 1
	p1.Sku = "1111"
	p1.Desc = "product num one"
	p1.Price = 1.25
	rtn = append(rtn, p1)

	var p2 Product
	p2.ID = 2
	p2.Sku = "2222"
	p2.Desc = "product num two"
	p2.Price = 2.25
	rtn = append(rtn, p2)

	w.WriteHeader(http.StatusOK)
	resJSON, _ := json.Marshal(rtn)
	fmt.Fprint(w, string(resJSON))
}

// GetProduct GetProduct
func (h *StoreHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars != nil && len(vars) == 1 {
		h.SetContentType(w)
		var pIDStr = vars["id"]
		fmt.Println("id: ", pIDStr)
		var rtn *Product
		pID, err := strconv.ParseInt(pIDStr, 10, 64)
		fmt.Println("err: ", err)
		var np Product
		if err == nil {
			if pID == 1 {
				np.ID = pID
				np.Desc = "product num one"
				np.Sku = "1111"
				np.Price = 1.25

			} else if pID == 2 {
				np.ID = pID
				np.Desc = "product num two"
				np.Sku = "2222"
				np.Price = 2.25
			}
			rtn = &np
		} else {
			rtn = &Product{}
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		resJSON, _ := json.Marshal(rtn)
		fmt.Fprint(w, string(resJSON))
	} else {
		http.Error(w, "wrong number of arguments", http.StatusBadRequest)
	}
}

// GetProductWithIDAndSku GetProductWithIdAndSku
func (h *StoreHandler) GetProductWithIDAndSku(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in get sku")
	vars := mux.Vars(r)
	if vars != nil && len(vars) == 2 {
		h.SetContentType(w)
		var pIDStr = vars["id"]
		var sku = vars["sku"]
		fmt.Println("id: ", pIDStr)
		fmt.Println("sku: ", sku)
		var rtn *Product
		pID, err := strconv.ParseInt(pIDStr, 10, 64)
		fmt.Println("err: ", err)
		var np Product
		if err == nil {
			if pID == 1 {
				np.ID = pID
				np.Desc = "product num one"
				np.Sku = "1111"
				np.Price = 1.25

			} else if pID == 2 {
				np.ID = pID
				np.Desc = "product num two"
				np.Sku = "2222"
				np.Price = 2.25
			}
			rtn = &np
		} else {
			rtn = &Product{}
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		resJSON, _ := json.Marshal(rtn)
		fmt.Fprint(w, string(resJSON))
	} else {
		//rtn = &Product{}
		http.Error(w, "not enough arguments", http.StatusBadRequest)
	}
}

// AddProduct AddProduct
func (h *StoreHandler) AddProduct(w http.ResponseWriter, r *http.Request) {

	acOk := h.CheckContent(r)
	if !acOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		h.SetContentType(w)
		var p Product
		scsuc, scerr := h.ProcessBody(r, &p)
		if !scsuc || scerr != nil {
			http.Error(w, scerr.Error(), http.StatusBadRequest)
		} else {
			var rtn Response
			rtn.ID = 4
			rtn.Success = true
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(rtn)
			fmt.Fprint(w, string(resJSON))
		}
	}
}

// UpdateProduct UpdateProduct
func (h *StoreHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	acOk := h.CheckContent(r)
	if !acOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		h.SetContentType(w)
		var p Product
		scsuc, scerr := h.ProcessBody(r, &p)
		if !scsuc || scerr != nil {
			http.Error(w, scerr.Error(), http.StatusBadRequest)
		} else {
			var rtn Response
			rtn.ID = 4
			rtn.Success = true
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(rtn)
			fmt.Fprint(w, string(resJSON))
		}
	}
}

// CheckContent CheckContent
func (h *StoreHandler) CheckContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		rtn = true
	}
	return rtn
}

// SetContentType SetContentType
func (h *StoreHandler) SetContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// ProcessBody ProcessBody
func (h *StoreHandler) ProcessBody(r *http.Request, obj interface{}) (bool, error) {
	var suc bool
	var err error
	//fmt.Println("r.Body: ", r.Body)
	//h.Log.Debug("r.Body: ", r.Body)
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		//fmt.Println("decoder: ", decoder)
		err = decoder.Decode(obj)
		//fmt.Println("decoder: ", decoder)
		if err != nil {
			log.Println("Decode Error: ", err.Error())
			//h.Log.Error("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}
	return suc, err
}
