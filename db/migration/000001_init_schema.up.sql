CREATE TABLE "Tb_Rol" (
  "ID_Rol" integer PRIMARY KEY NOT NULL,
  "Rol" varchar NOT NULL,
  "Description" varchar
);

CREATE TABLE "Tb_Permission" (
  "ID_Permission" integer PRIMARY KEY NOT NULL,
  "Permission" varchar NOT NULL,
  "Description" varchar
);

CREATE TABLE "Tb_User" (
  "ID_User" integer PRIMARY KEY NOT NULL,
  "ID_Rol" integer NOT NULL,
  "First_Name" varchar NOT NULL,
  "Last_Name" varchar NOT NULL,
  "Email" varchar NOT NULL,
  "Password" varchar NOT NULL,
  "Created_At" timestamp NOT NULL,
  "Active" boolean NOT NULL,
  FOREIGN KEY ("ID_Rol") REFERENCES "Tb_Rol" ("ID_Rol") -- Foreign key here
);

CREATE TABLE "Tb_RolPermission" (
  "ID_RolPermission" integer PRIMARY KEY NOT NULL,
  "ID_Rol" integer NOT NULL,
  "ID_Permission" integer NOT NULL,
  FOREIGN KEY ("ID_Rol") REFERENCES "Tb_Rol" ("ID_Rol"), -- Foreign key here
  FOREIGN KEY ("ID_Permission") REFERENCES "Tb_Permission" ("ID_Permission") -- Foreign key here
);

CREATE TABLE "Tb_StatusMesa" (
  "ID_StatusMesa" integer PRIMARY KEY NOT NULL,
  "Status" varchar NOT NULL,
  "Descripcion" varchar
);

CREATE TABLE "Tb_Mesero" (
  "ID_Mesero" integer PRIMARY KEY NOT NULL,
  "ID_User" integer UNIQUE NOT NULL,
  FOREIGN KEY ("ID_User") REFERENCES "Tb_User" ("ID_User") -- Foreign key here
);

CREATE TABLE "Tb_Mesa" (
  "ID_Mesa" integer PRIMARY KEY NOT NULL,
  "Numero_Mesa" integer UNIQUE NOT NULL,
  "ID_Mesero" integer UNIQUE NOT NULL,
  "ID_StatusMesa" integer UNIQUE NOT NULL,
  FOREIGN KEY ("ID_Mesero") REFERENCES "Tb_Mesero" ("ID_Mesero"), -- Foreign key here
  FOREIGN KEY ("ID_StatusMesa") REFERENCES "Tb_StatusMesa" ("ID_StatusMesa") -- Foreign key here
);

CREATE TABLE "Tb_Categoria" (
  "ID_Categoria" integer PRIMARY KEY NOT NULL,
  "Categoria" varchar UNIQUE NOT NULL,
  "Active" boolean NOT NULL,
  "Descripcion" varchar
);

CREATE TABLE "Tb_Producto" (
  "ID_Producto" integer PRIMARY KEY NOT NULL,
  "Nombre_Producto" varchar UNIQUE NOT NULL,
  "Precio_Compra" decimal NOT NULL,
  "Precio_Venta" decimal NOT NULL,
  "ID_Categoria" integer NOT NULL,
  "Stock" integer NOT NULL,
  "Idproovedor" integer NOT NULL,
  FOREIGN KEY ("ID_Categoria") REFERENCES "Tb_Categoria" ("ID_Categoria"), -- Foreign key here
  FOREIGN KEY ("Idproovedor") REFERENCES "Tb_Proovedor" ("Idproovedor") -- Foreign key here
);

CREATE TABLE "Tb_Proovedor" (
  "Idproovedor" integer PRIMARY KEY NOT NULL,
  "Nombre_proovedor" varchar UNIQUE NOT NULL,
  "Telfono" varchar UNIQUE NOT NULL,
  "Correo" varchar UNIQUE NOT NULL,
  "Direccion" varchar NOT NULL
);

CREATE TABLE "Tb_StatusPedido" (
  "Id_StatusPedido" integer PRIMARY KEY NOT NULL,
  "Status_Pedido" varchar NOT NULL,
  "Descripcion" varchar
);

CREATE TABLE "Tb_Pedido" (
  "ID_Pedido" integer PRIMARY KEY NOT NULL,
  "Order_Date" timestamp NOT NULL,
  "ID_Mesa" integer NOT NULL,
  "ID_StatusPedido" integer NOT NULL,
  FOREIGN KEY ("ID_Mesa") REFERENCES "Tb_Mesa" ("ID_Mesa"), -- Foreign key here
  FOREIGN KEY ("ID_StatusPedido") REFERENCES "Tb_StatusPedido" ("Id_StatusPedido") -- Foreign key here
);

CREATE TABLE "Tb_PedidoProductos" (
  "ID_PedidoProductos" integer PRIMARY KEY NOT NULL,
  "ID_Producto" integer NOT NULL,
  "ID_Pedido" integer NOT NULL,
  FOREIGN KEY ("ID_Pedido") REFERENCES "Tb_Pedido" ("ID_Pedido"), -- Foreign key here
  FOREIGN KEY ("ID_Producto") REFERENCES "Tb_Producto" ("ID_Producto") -- Foreign key here
);

