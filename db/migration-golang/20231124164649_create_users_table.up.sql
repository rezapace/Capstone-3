-- membuat tabel user dengan field id, name, password,email
CREATE TABLE "public"."users" (
    id SERIAL NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    number VARCHAR(255) NOT NULL,
    roles VARCHAR(255) NOT NULL,
    saldo INTEGER NOT NULL
);

ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

-- membuat tabel ticket dengan field id, gambar, lokasi, tanggal_bulan_tahun, judul, deskripsi, harga_tiket, status_tiket, kuota_tiket
CREATE TABLE "public"."tickets" (
    ID SERIAL PRIMARY KEY,
    Image TEXT,
    Location TEXT,
    Date DATE,
    Title TEXT,
    Description TEXT,
    Price INT,
    Status TEXT DEFAULT 'available',
    Quota INT,
    Created_At TIMESTAMP,
    Updated_At TIMESTAMP,
    Deleted_At TIMESTAMP
);

-- membuat table blog dengan field id, gambar, judul, deskripsi, tanggal_bulan_tahun
CREATE TABLE blogs (
    ID SERIAL PRIMARY KEY,
    Image TEXT,
    Date DATE,
    Title TEXT,
    Description TEXT,
    Created_At TIMESTAMP,
    Updated_At TIMESTAMP,
    Deleted_At TIMESTAMP
);
