-- Proveedor
CREATE TABLE "provider" (
  "provider_id" serial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "phone" varchar(20) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "address" varchar(255) NOT NULL,
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp DEFAULT current_timestamp
);

-- Categoría
CREATE TABLE "category" (
  "category_id" serial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL,
  "description" varchar(255),
  "is_active" boolean DEFAULT true,
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp DEFAULT current_timestamp
);

-- Producto
CREATE TABLE "product" (
  "product_id" serial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "purchase_price" decimal(10, 2) NOT NULL,
  "sale_price" decimal(10, 2) NOT NULL,
  "stock" integer NOT NULL,
  "category_id" integer NOT NULL,
  "provider_id" integer NOT NULL,
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp DEFAULT current_timestamp,
  FOREIGN KEY ("category_id") REFERENCES "category" ("category_id"),
  FOREIGN KEY ("provider_id") REFERENCES "provider" ("provider_id")
);

-- Rol
CREATE TABLE "role" (
  "role_id" serial PRIMARY KEY,
  "name" varchar(50) UNIQUE NOT NULL,
  "description" varchar(255),
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp DEFAULT current_timestamp
);

-- Permiso
CREATE TABLE "permission" (
  "permission_id" serial PRIMARY KEY,
  "name" varchar(50) UNIQUE NOT NULL,
  "description" varchar(255),
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp DEFAULT current_timestamp
);

-- Rol Permiso (relación N:M entre roles y permisos)
CREATE TABLE "role_permission" (
  "role_permission_id" serial PRIMARY KEY,
  "role_id" integer NOT NULL,
  "permission_id" integer NOT NULL,
  "created_at" timestamp DEFAULT current_timestamp,
  FOREIGN KEY ("role_id") REFERENCES "role" ("role_id"),
  FOREIGN KEY ("permission_id") REFERENCES "permission" ("permission_id")
);

-- Usuario
CREATE TABLE "user" (
  "user_id" serial PRIMARY KEY,
  "role_id" integer NOT NULL,
  "first_name" varchar(100) NOT NULL,
  "last_name" varchar(100) NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "active" boolean DEFAULT true,
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp DEFAULT current_timestamp,
  FOREIGN KEY ("role_id") REFERENCES "role" ("role_id")
);

-- Mesero
CREATE TABLE "waiter" (
  "waiter_id" serial PRIMARY KEY,
  "user_id" integer UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT current_timestamp,
  FOREIGN KEY ("user_id") REFERENCES "user" ("user_id")
);

-- Estado de la Mesa
CREATE TABLE "table_status" (
  "table_status_id" serial PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" varchar(255),
  "created_at" timestamp DEFAULT current_timestamp
);

-- Mesa
CREATE TABLE "restaurant_table" (
  "table_id" serial PRIMARY KEY,
  "number" integer UNIQUE NOT NULL,
  "waiter_id" integer NOT NULL,
  "status_id" integer NOT NULL,
  "created_at" timestamp DEFAULT current_timestamp,
  "updated_at" timestamp DEFAULT current_timestamp,
  FOREIGN KEY ("waiter_id") REFERENCES "waiter" ("waiter_id"),
  FOREIGN KEY ("status_id") REFERENCES "table_status" ("table_status_id")
);

-- Estado del Pedido
CREATE TABLE "order_status" (
  "order_status_id" serial PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" varchar(255),
  "created_at" timestamp DEFAULT current_timestamp
);

-- Pedido
CREATE TABLE "order" (
  "order_id" serial PRIMARY KEY,
  "order_date" timestamp DEFAULT current_timestamp,
  "table_id" integer NOT NULL,
  "status_id" integer NOT NULL,
  FOREIGN KEY ("table_id") REFERENCES "restaurant_table" ("table_id"),
  FOREIGN KEY ("status_id") REFERENCES "order_status" ("order_status_id")
);

-- Productos en Pedido (relación N:M entre pedidos y productos)
CREATE TABLE "order_product" (
  "order_product_id" serial PRIMARY KEY,
  "order_id" integer NOT NULL,
  "product_id" integer NOT NULL,
  "quantity" integer NOT NULL DEFAULT 1,
  FOREIGN KEY ("order_id") REFERENCES "order" ("order_id"),
  FOREIGN KEY ("product_id") REFERENCES "product" ("product_id")
);
