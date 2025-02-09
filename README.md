# overview

build graphql gateway, api gateway for graphql

# Sample Requirement

product and uom is 1 domain, tenant is separated

```
# request body
query ProductPage($where: product_product_bool_exp, $limit: Int, $offset: Int, $order_by: [product_product_order_by!]) {
  product_product(
    where: $where
    limit: $limit
    offset: $offset
    order_by: $order_by
  ) {
    id
    name
    status
    sku
    image_url
    type
    category_id
    tenant {
      id
      name
      number
      __typename
    }
    uom {
      id
      uom_standard_id
      conversion_uom_id
      conversion_ratio
      dimension
      dimension_width
      dimension_length
      dimension_height
      dimension_uom
      weight
      weight_uom
      uom_standard {
        id
        number
        name
        __typename
      }
      __typename
    }
    serial_number_management
    shelf_life_management
    lot_management
    uom_conversion_management
    __typename
  }
}

-----

{
  "limit": 20,
  "offset": 0,
  "order_by": [
    {
      "company_id": "asc"
    },
    {
      "sku": "asc"
    }
  ],
  "where": {}
}


```


# walkthrough

```
mkdir product
mkdir tenant
go mod init github.com/toshim45/demo-apollo-fed-gqlgen
printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
go mod tidy
cd product; go run github.com/99designs/gqlgen init; cd ..
cd tenant; go run github.com/99designs/gqlgen init; cd ..
mkdir -p cmd/product-gql
mkdir -p cmd/tenant-gql
mv product/server.go cmd/product-gql/main.go
mv tenant/server.go cmd/tenant-gql/main.go
# edit product/graph/schema.graphqls 
cd product; go run github.com/99designs/gqlgen generate; cd ..
# edit tenant/graph/schema.graphqls
cd tenant; go run github.com/99designs/gqlgen generate; cd ..
go build -o cmd/tenant-gql/binary cmd/tenant-gql/main.go
go build -o cmd/product-gql/binary cmd/product-gql/main.go
./cmd/tenant-gql/binary &
./cmd/product-gql/binary &
mkdir gateway
cd gateway; npm init; npm install @apollo/gateway @apollo/server graphql; npm run dev &; cd ..
```