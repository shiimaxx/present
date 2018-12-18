go func() {
    // メインの処理
}()

w.Header().Add("Content-type", "application/json")
w.WriteHeader(http.StatusOK)
fmt.Fprintf(w, "")
