# httpserver

Lightweight and idiomatic HTTP server wrapper for Go.

- Start/stop lifecycle with context
- Functional options for configuration
- Compatible with any `http.Handler` (e.g. chi, mux, stdlib)
- Clean shutdown and error propagation

## Installation

```bash
go get github.com/wrestler094/httpserver
```

## Example

```go
import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/wrestler094/httpserver"
)

func main() {
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	srv := httpserver.NewServer(
		httpserver.Port(":8080"),
		httpserver.Handler(r),
	)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}

	select {
	case err := <-srv.Notify():
		log.Fatal("server error:", err)
	}
}
```

## License
This project is licensed under the MIT License – see the [LICENSE](./LICENSE) file for details.
