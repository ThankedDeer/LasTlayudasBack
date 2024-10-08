CREATE TABLE "Tb_Rol" (
    "ID_Rol" integer PRIMARY KEY,
    "Rol" varchar NOT NULL,
    "Description" varchar
);

CREATE TABLE "Tb_Permission" (
    "ID_Permission" integer PRIMARY KEY,
    "Permission" varchar NOT NULL,
    "Description" varchar
);

CREATE TABLE "Tb_StatusMesa" (
    "ID_StatusMesa" integer PRIMARY KEY,
    "Status" varchar NOT NULL,
    "Descripcion" varchar
);

CREATE TABLE "Tb_Categoria" (
    "ID_Categoria" integer PRIMARY KEY,
    "Categoria" varchar UNIQUE NOT NULL,
    "Active" boolean NOT NULL,
    "Descripcion" varchar
);

CREATE TABLE "Tb_Proovedor" (
    "Idproovedor" integer PRIMARY KEY,
    "Nombre_proovedor" varchar UNIQUE NOT NULL,
    "Telfono" varchar UNIQUE NOT NULL,
    "Correo" varchar UNIQUE NOT NULL,
    "Direccion" varchar NOT NULL
);

CREATE TABLE "Tb_User" (
    "ID_User" integer PRIMARY KEY,
    "ID_Rol" integer NOT NULL,
    "First_Name" varchar NOT NULL,
    "Last_Name" varchar NOT NULL,
    "Email" varchar NOT NULL,
    "Password" varchar NOT NULL,
    "Created_At" timestamp NOT NULL,
    "Active" BOOLEAN NOT NULL,
    FOREIGN KEY ("ID_Rol") REFERENCES "Tb_Rol" ("ID_Rol")
);

CREATE TABLE "Tb_Mesero" (
    "ID_Mesero" integer PRIMARY KEY,
    "ID_User" integer UNIQUE NOT NULL,
    FOREIGN KEY ("ID_User") REFERENCES "Tb_User" ("ID_User")
);

CREATE TABLE "Tb_Mesa" (
    "ID_Mesa" integer PRIMARY KEY,
    "Numero_Mesa" integer UNIQUE NOT NULL,
    "ID_Mesero" integer UNIQUE NOT NULL,
    "ID_StatusMesa" integer UNIQUE NOT NULL,
    FOREIGN KEY ("ID_Mesero") REFERENCES "Tb_Mesero" ("ID_Mesero"),
    FOREIGN KEY ("ID_StatusMesa") REFERENCES "Tb_StatusMesa" ("ID_StatusMesa")
);

CREATE TABLE "Tb_Producto" (
    "ID_Producto" integer PRIMARY KEY,
    "Nombre_Producto" varchar UNIQUE NOT NULL,
    "Precio_Compra" decimal NOT NULL,
    "Precio_Venta" decimal NOT NULL,
    "ID_Categoria" integer NOT NULL,
    "Stock" integer NOT NULL,
    "ID_Proovedor" integer NOT NULL,
    FOREIGN KEY ("ID_Categoria") REFERENCES "Tb_Categoria" ("ID_Categoria"),
    FOREIGN KEY ("ID_Proovedor") REFERENCES "Tb_Proovedor" ("Idproovedor")
);

CREATE TABLE "Tb_StatusPedido" (
    "Id_StatusPedido" integer PRIMARY KEY,
    "Status_Pedido" varchar NOT NULL,
    "Descripcion" varchar
);

CREATE TABLE "Tb_TipoPedido" (
    "ID_TipoPedido" integer PRIMARY KEY,
    "Tipo_Pedido" varchar NOT NULL
);

CREATE TABLE "Tb_Pedido" (
    "ID_Pedido" integer PRIMARY KEY,
    "Order_Date" timestamp NOT NULL,
    "ID_Mesa" integer NOT NULL,
    "ID_StatusPedido" integer NOT NULL,
    "ID_TipoPedido" integer NOT NULL,
    "Descripcion_Pedido" varchar NOT NULL,
    "Subtotal" decimal,
    "Total" decimal,
    FOREIGN KEY ("ID_Mesa") REFERENCES "Tb_Mesa" ("ID_Mesa"),
    FOREIGN KEY ("ID_StatusPedido") REFERENCES "Tb_StatusPedido" ("Id_StatusPedido"),
    FOREIGN KEY ("ID_TipoPedido") REFERENCES "Tb_TipoPedido" ("ID_TipoPedido")
);

CREATE TABLE "Tb_PedidoProductos" (
    "ID_PedidoProductos" integer PRIMARY KEY,
    "ID_Producto" integer NOT NULL,
    "ID_Pedido" integer NOT NULL,
    FOREIGN KEY ("ID_Producto") REFERENCES "Tb_Producto" ("ID_Producto"),
    FOREIGN KEY ("ID_Pedido") REFERENCES "Tb_Pedido" ("ID_Pedido")
);

CREATE TABLE "Tb_RolPermission" (
    "ID_RolPermission" integer PRIMARY KEY,
    "ID_Rol" integer NOT NULL,
    "ID_Permission" integer NOT NULL,
    FOREIGN KEY ("ID_Rol") REFERENCES "Tb_Rol" ("ID_Rol"),
    FOREIGN KEY ("ID_Permission") REFERENCES "Tb_Permission" ("ID_Permission")
);
