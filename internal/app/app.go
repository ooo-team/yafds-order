package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	model "github.com/ooo-team/yafds-order/internal/model/order"
	// model "github.com/ooo-team/yafds-order/internal/model/order"
)

type App struct {
	httpServer      *http.Server
	serviceProvider *serviceProvider
}

func (a *App) initServiceProvider() error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) getOrder(w http.ResponseWriter, r *http.Request) {

}

type InputOrder struct {
	CustomerID uint32
}

func (a *App) createOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Чтение тела запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Декодирование JSON
	var requestData model.Order
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	id, err := a.serviceProvider.OrderService().Create(ctx, &requestData)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create order: %s\n", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("Created order: ID = %d, CustomerID = %d", id, requestData.CustomerID)

	fmt.Fprintln(w, response)

}

func (a *App) initHttpServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/order/create", a.createOrder)
	mux.HandleFunc("/order/get", a.getOrder)
	a.httpServer = &http.Server{
		Addr:           ":5312",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 МБ
	}

	return nil
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	a.initHttpServer()
	a.initServiceProvider()
	return a, nil
}

func (a *App) Run() {
	if err := a.httpServer.ListenAndServe(); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
