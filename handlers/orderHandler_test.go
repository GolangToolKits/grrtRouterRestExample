package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	mux "github.com/GolangToolKits/grrt"
)

func TestStoreHandler_GetProducts(t *testing.T) {
	tr, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw := httptest.NewRecorder()
	var sh StoreHandler

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *StoreHandler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1 ",
			h:    &sh,
			args: args{
				w: tw,
				r: tr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &StoreHandler{}
			h.GetProducts(tt.args.w, tt.args.r)
			if tw.Code != http.StatusOK {
				t.Fail()
			}
		})
	}
}

func TestStoreHandler_GetProductWithIDAndSku(t *testing.T) {
	tr, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw := httptest.NewRecorder()
	vars := make(map[string]string)
	vars["id"] = "2"
	vars["sku"] = "1234"
	tr = mux.SetURLVars(tr, vars)

	//test 2---
	tr2, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw2 := httptest.NewRecorder()
	vars2 := make(map[string]string)
	vars2["id"] = "1"
	vars2["sku"] = "2222"
	tr2 = mux.SetURLVars(tr2, vars2)

	//test 3---
	tr3, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw3 := httptest.NewRecorder()
	vars3 := make(map[string]string)
	vars3["id"] = "1"
	//vars3["sku"] = "2222"
	tr3 = mux.SetURLVars(tr3, vars3)

	//test 4---
	tr4, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw4 := httptest.NewRecorder()
	vars4 := make(map[string]string)
	vars4["id"] = "1e"
	vars4["sku"] = "2222"
	tr4 = mux.SetURLVars(tr4, vars4)

	var sh StoreHandler
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *StoreHandler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			h:    &sh,
			args: args{
				w: tw,
				r: tr,
			},
		},
		{
			name: "test 2",
			h:    &sh,
			args: args{
				w: tw2,
				r: tr2,
			},
		},
		{
			name: "test 3",
			h:    &sh,
			args: args{
				w: tw3,
				r: tr3,
			},
		},
		{
			name: "test 4",
			h:    &sh,
			args: args{
				w: tw4,
				r: tr4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &StoreHandler{}
			h.GetProductWithIDAndSku(tt.args.w, tt.args.r)
			if tt.name == "test 1" && tw.Code != http.StatusOK {
				t.Fail()
			}
			if tt.name == "test 2" && tw2.Code != http.StatusOK {
				t.Fail()
			}
			if tt.name == "test 3" && tw3.Code != http.StatusBadRequest {
				t.Fail()
			}
			if tt.name == "test 4" && tw4.Code != http.StatusBadRequest {
				t.Fail()
			}
		})
	}
}

func TestStoreHandler_New(t *testing.T) {
	tests := []struct {
		name string
		h    *StoreHandler
		want Handler
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			h:    &StoreHandler{},
			want: &StoreHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &StoreHandler{}
			if got := h.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoreHandler.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoreHandler_GetProduct(t *testing.T) {
	tr, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw := httptest.NewRecorder()
	vars := make(map[string]string)
	vars["id"] = "2"
	tr = mux.SetURLVars(tr, vars)

	//test 2----
	tr2, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw2 := httptest.NewRecorder()
	vars2 := make(map[string]string)
	vars2["id"] = "2"
	vars2["id2"] = "2"
	tr2 = mux.SetURLVars(tr2, vars2)

	//test 3-------------------
	tr3, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw3 := httptest.NewRecorder()
	vars3 := make(map[string]string)
	vars3["id"] = "1"
	tr3 = mux.SetURLVars(tr3, vars3)

	//test 4 -------
	tr4, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw4 := httptest.NewRecorder()
	vars4 := make(map[string]string)
	vars4["id"] = "2w"
	tr4 = mux.SetURLVars(tr4, vars4)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *StoreHandler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			h:    &StoreHandler{},
			args: args{
				w: tw,
				r: tr,
			},
		},
		{
			name: "test 2",
			h:    &StoreHandler{},
			args: args{
				w: tw2,
				r: tr2,
			},
		},
		{
			name: "test 3",
			h:    &StoreHandler{},
			args: args{
				w: tw3,
				r: tr3,
			},
		},
		{
			name: "test 4",
			h:    &StoreHandler{},
			args: args{
				w: tw4,
				r: tr4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &StoreHandler{}
			h.GetProduct(tt.args.w, tt.args.r)
			if tt.name == "test 1" && tw.Code != http.StatusOK {
				t.Fail()
			}
			if tt.name == "test 2" && tw2.Code != http.StatusBadRequest {
				t.Fail()
			}
			if tt.name == "test 3" && tw3.Code != http.StatusOK {
				t.Fail()
			}
			if tt.name == "test 4" && tw4.Code != http.StatusBadRequest {
				t.Fail()
			}
		})
	}
}

func TestStoreHandler_AddProduct(t *testing.T) {

	//test 1 --------------
	tr, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw := httptest.NewRecorder()

	//test 2 --------------
	tr2, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tr2.Header.Add("Content-Type", "application/json")
	tw2 := httptest.NewRecorder()

	//test 3 --------------
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 2, "sku": "1111", "desc": "a sku", "price": 1.25}`))
	tr3, _ := http.NewRequest("GET", "/test/test1/p1/p2", aJSON)
	tr3.Header.Add("Content-Type", "application/json")
	tw3 := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *StoreHandler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			h:    &StoreHandler{},
			args: args{
				w: tw,
				r: tr,
			},
		},
		{
			name: "test 2",
			h:    &StoreHandler{},
			args: args{
				w: tw2,
				r: tr2,
			},
		},
		{
			name: "test 3",
			h:    &StoreHandler{},
			args: args{
				w: tw3,
				r: tr3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &StoreHandler{}
			h.AddProduct(tt.args.w, tt.args.r)
			if tt.name == "test 1" && tw.Code != http.StatusUnsupportedMediaType {
				t.Fail()
			}
			if tt.name == "test 2" && tw2.Code != http.StatusBadRequest {
				t.Fail()
			}
			if tt.name == "test 3" && tw3.Code != http.StatusOK {
				t.Fail()
			}
		})
	}
}

func TestStoreHandler_UpdateProduct(t *testing.T) {
	//test 1 --------------
	tr, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tw := httptest.NewRecorder()

	//test 2 --------------
	tr2, _ := http.NewRequest("GET", "/test/test1/p1/p2", nil)
	tr2.Header.Add("Content-Type", "application/json")
	tw2 := httptest.NewRecorder()

	//test 3 --------------
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 2, "sku": "1111", "desc": "a sku", "price": 1.25}`))
	tr3, _ := http.NewRequest("GET", "/test/test1/p1/p2", aJSON)
	tr3.Header.Add("Content-Type", "application/json")
	tw3 := httptest.NewRecorder()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *StoreHandler
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			h:    &StoreHandler{},
			args: args{
				w: tw,
				r: tr,
			},
		},
		{
			name: "test 2",
			h:    &StoreHandler{},
			args: args{
				w: tw2,
				r: tr2,
			},
		},
		{
			name: "test 3",
			h:    &StoreHandler{},
			args: args{
				w: tw3,
				r: tr3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &StoreHandler{}
			h.UpdateProduct(tt.args.w, tt.args.r)
			if tt.name == "test 1" && tw.Code != http.StatusUnsupportedMediaType {
				t.Fail()
			}
			if tt.name == "test 2" && tw2.Code != http.StatusBadRequest {
				t.Fail()
			}
			if tt.name == "test 3" && tw3.Code != http.StatusOK {
				t.Fail()
			}
		})
	}
}

func TestStoreHandler_ProcessBody(t *testing.T) {
	var p Product
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{""id": 2, "sku": "1111", "desc": "a sku", "price": 1.25}`))
	tr, _ := http.NewRequest("GET", "/test/test1/p1/p2", aJSON)
	type args struct {
		r   *http.Request
		obj interface{}
	}
	tests := []struct {
		name    string
		h       *StoreHandler
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			h:    &StoreHandler{},
			args: args{
				r:   tr,
				obj: p,
			},
			want: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &StoreHandler{}
			got, err := h.ProcessBody(tt.args.r, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("StoreHandler.ProcessBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StoreHandler.ProcessBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
