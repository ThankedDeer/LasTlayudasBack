-- name: CreateProduct :one
INSERT INTO Tb_Producto (
    Nombre_Producto, Precio_Compra, Precio_Venta, ID_Categoria, Stock, Idproveedor
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING ID_Producto, Nombre_Producto, Precio_Compra, Precio_Venta, ID_Categoria, Stock, Idproveedor;


-- name: GetProductByID :one
SELECT ID_Producto, Nombre_Producto, Precio_Compra, Precio_Venta, ID_Categoria, Stock, Idproveedor
FROM Tb_Producto
WHERE ID_Producto = $1;


-- name: GetAllProducts :many
SELECT ID_Producto, Nombre_Producto, Precio_Compra, Precio_Venta, ID_Categoria, Stock, Idproveedor
FROM Tb_Producto;

-- name: UpdateProduct :exec
UPDATE Tb_Producto
SET Nombre_Producto = $2, Precio_Compra = $3, Precio_Venta = $4, ID_Categoria = $5, Stock = $6, Idproveedor = $7
WHERE ID_Producto = $1;


-- name: DeleteProduct :exec
DELETE FROM Tb_Producto
WHERE ID_Producto = $1;