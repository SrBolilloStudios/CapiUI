# CapiUI
Sistema POS local

## Estructura

```
cmd/
  server/    API HTTP (puerto :8080)
  seed/      Puebla la base de datos (drop + create + datos de ejemplo)
internal/
  models/    Structs de dominio (Product, Category)
  store/     Acceso a datos, uno por recurso (ProductStore, CategoryStore)
  api/       Handlers HTTP, helpers JSON, CORS y registro de rutas
  query/     SQL embebido con convención <recurso>.<accion>
frontend/    App Vue
```

Cada recurso sigue el mismo patrón: `models/<recurso>.go` → `store/<recurso>.go`
(`List`, `GetByIDs`, `Create`, `Delete`) → `api/<recurso>.go`
(`List`, `Get`, `Create`, `Delete`) → rutas en `api/routes.go`.

## Backend (Go)

```bash
go run ./cmd/seed    # poblar products.db
go run ./cmd/server  # levantar la API en :8080
```

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| GET | `/api/products?page=&limit=` | Lista paginada de productos (default page=1, limit=50, max 100) |
| GET | `/api/products/{ids}` | Productos por IDs (ej. `1,3`) |
| POST | `/api/products` | Crear productos (body: array JSON) |
| DELETE | `/api/products/{ids}` | Eliminar productos por IDs |
| GET | `/api/categories?page=&limit=` | Lista paginada de categorías |
| GET | `/api/categories/{ids}` | Categorías por IDs |
| POST | `/api/categories` | Crear categorías (body: array JSON) |
| DELETE | `/api/categories/{ids}` | Eliminar categorías por IDs |

Convenciones de respuesta: listas y consultas devuelven el array JSON directo;
creación devuelve `201 {"created": n}`; errores devuelven `{"error": "..."}` con
su código HTTP (400 body/parámetros inválidos, 404 sin resultados, 405 método no
permitido).
