-- Tabla de Roles
CREATE TABLE "Tb_Rol" (
  "ID_Rol" SERIAL PRIMARY KEY, -- Usamos SERIAL para generar automáticamente el ID
  "Rol" VARCHAR(100) NOT NULL UNIQUE, -- Aseguramos que el nombre del rol sea único
  "Descripcion" VARCHAR(255) -- Descripción del rol
);

-- Tabla de Permisos
CREATE TABLE "Tb_Permission" (
  "ID_Permission" SERIAL PRIMARY KEY,
  "Permission" VARCHAR(100) NOT NULL UNIQUE,
  "Descripcion" VARCHAR(255) -- Descripción del permiso
);

-- Tabla de Usuarios
CREATE TABLE "Tb_User" (
  "ID_User" SERIAL PRIMARY KEY,
  "ID_Rol" INT NOT NULL,
  "First_Name" VARCHAR(100) NOT NULL,
  "Last_Name" VARCHAR(100) NOT NULL,
  "Email" VARCHAR(150) NOT NULL UNIQUE,
  "Password" VARCHAR(255) NOT NULL,
  "Created_At" TIMESTAMP NOT NULL DEFAULT NOW(),
  "Active" BOOLEAN NOT NULL DEFAULT TRUE,
  CONSTRAINT fk_rol FOREIGN KEY ("ID_Rol") REFERENCES "Tb_Rol" ("ID_Rol") ON DELETE CASCADE
);

-- Tabla de Relación entre Roles y Permisos
CREATE TABLE "Tb_RolPermission" (
  "ID_RolPermission" SERIAL PRIMARY KEY,
  "ID_Rol" INT NOT NULL,
  "ID_Permission" INT NOT NULL,
  CONSTRAINT fk_rol_permission_rol FOREIGN KEY ("ID_Rol") REFERENCES "Tb_Rol" ("ID_Rol") ON DELETE CASCADE,
  CONSTRAINT fk_rol_permission_perm FOREIGN KEY ("ID_Permission") REFERENCES "Tb_Permission" ("ID_Permission") ON DELETE CASCADE
);

-- Tabla de Estados de Mesas
CREATE TABLE "Tb_StatusMesa" (
  "ID_StatusMesa" SERIAL PRIMARY KEY,
  "Status" VARCHAR(100) NOT NULL UNIQUE,
  "Descripcion" VARCHAR(255) -- Descripción del estado de la mesa
);

-- Tabla de Meseros
CREATE TABLE "Tb_Mesero" (
  "ID_Mesero" SERIAL PRIMARY KEY,
  "ID_User" INT NOT NULL UNIQUE,
  CONSTRAINT fk_mesero_user FOREIGN KEY ("ID_User") REFERENCES "Tb_User" ("ID_User") ON DELETE CASCADE
);

-- Tabla de Mesas
CREATE TABLE "Tb_Mesa" (
  "ID_Mesa" SERIAL PRIMARY KEY,
  "Numero_Mesa" INT NOT NULL UNIQUE,
  "ID_Mesero" INT NOT NULL,
  "ID_StatusMesa" INT NOT NULL,
  CONSTRAINT fk_mesa_mesero FOREIGN KEY ("ID_Mesero") REFERENCES "Tb_Mesero" ("ID_Mesero") ON DELETE SET NULL,
  CONSTRAINT fk_mesa_status FOREIGN KEY ("ID_StatusMesa") REFERENCES "Tb_StatusMesa" ("ID_StatusMesa") ON DELETE SET NULL
);

-- Tabla de Categorías
CREATE TABLE "Tb_Categoria" (
  "ID_Categoria" SERIAL PRIMARY KEY,
  "Categoria" VARCHAR(100) NOT NULL UNIQUE,
  "Active" BOOLEAN NOT NULL DEFAULT TRUE,
  "Descripcion" VARCHAR(255) -- Descripción de la categoría
);

-- Tabla de Proveedores
CREATE TABLE "Tb_Proovedor" (
  "ID_Proovedor" SERIAL PRIMARY KEY,
  "Nombre_Proovedor" VARCHAR(100) NOT NULL UNIQUE,
  "Telefono" VARCHAR(20) NOT NULL UNIQUE,
  "Correo" VARCHAR(150) NOT NULL UNIQUE,
  "Direccion" VARCHAR(255) NOT NULL
);

-- Tabla de Productos
CREATE TABLE "Tb_Producto" (
  "ID_Producto" SERIAL PRIMARY KEY,
  "Nombre_Producto" VARCHAR(150) NOT NULL UNIQUE,
  "Precio_Compra" DECIMAL(10, 2) NOT NULL,
  "Precio_Venta" DECIMAL(10, 2) NOT NULL,
  "ID_Categoria" INT NOT NULL,
  "Stock" INT NOT NULL CHECK (Stock >= 0),
  "ID_Proovedor" INT NOT NULL,
  CONSTRAINT fk_producto_categoria FOREIGN KEY ("ID_Categoria") REFERENCES "Tb_Categoria" ("ID_Categoria") ON DELETE SET NULL,
  CONSTRAINT fk_producto_proveedor FOREIGN KEY ("ID_Proovedor") REFERENCES "Tb_Proovedor" ("ID_Proovedor") ON DELETE SET NULL
);

-- Tabla de Estados de Pedido
CREATE TABLE "Tb_StatusPedido" (
  "ID_StatusPedido" SERIAL PRIMARY KEY,
  "Status_Pedido" VARCHAR(100) NOT NULL,
  "Descripcion" VARCHAR(255) -- Descripción del estado del pedido
);

-- Tabla de Pedidos
CREATE TABLE "Tb_Pedido" (
  "ID_Pedido" SERIAL PRIMARY KEY,
  "Order_Date" TIMESTAMP NOT NULL DEFAULT NOW(),
  "ID_Mesa" INT NOT NULL,
  "ID_StatusPedido" INT NOT NULL,
  CONSTRAINT fk_pedido_mesa FOREIGN KEY ("ID_Mesa") REFERENCES "Tb_Mesa" ("ID_Mesa") ON DELETE SET NULL,
  CONSTRAINT fk_pedido_status FOREIGN KEY ("ID_StatusPedido") REFERENCES "Tb_StatusPedido" ("ID_StatusPedido") ON DELETE SET NULL
);

-- Tabla de Productos en Pedido
CREATE TABLE "Tb_PedidoProductos" (
  "ID_PedidoProductos" SERIAL PRIMARY KEY,
  "ID_Producto" INT NOT NULL,
  "ID_Pedido" INT NOT NULL,
  CONSTRAINT fk_pedido_productos_pedido FOREIGN KEY ("ID_Pedido") REFERENCES "Tb_Pedido" ("ID_Pedido") ON DELETE CASCADE,
  CONSTRAINT fk_pedido_productos_producto FOREIGN KEY ("ID_Producto") REFERENCES "Tb_Producto" ("ID_Producto") ON DELETE CASCADE
);
