# 🧠 Create Order Saga - Tugas Pemrograman Web Lanjut

Ini adalah implementasi sistem **Create Order Saga** berbasis **gRPC** dengan bahasa **Go**, sebagai bagian dari Tugas 2 mata kuliah **Pemrograman Web Lanjut**.

## 🧩 Arsitektur

Sistem terdiri dari 4 komponen utama:

1. 🛒 **OrderService** (`port 50051`)
2. 💳 **PaymentService** (`port 50052`)
3. 📦 **ShippingService** (`port 50053`)
4. 🧠 **Orchestrator** – client utama yang mengatur alur transaksi

Setiap service berjalan secara independen dan saling berkomunikasi menggunakan protokol **gRPC**.

---

## 🔄 Alur Saga

1. Orchestrator memanggil `CreateOrder`
2. Jika sukses, lanjut ke `ProcessPayment`
3. Jika sukses, lanjut ke `StartShipping`
4. Jika Shipping gagal, maka dilakukan **kompensasi terbalik**:
   - `CancelShipping`
   - `RefundPayment`
   - `CancelOrder`

---

## 🧪 Skenario yang Diuji

| Skenario               | Hasil     | Kompensasi                |
|------------------------|-----------|---------------------------|
| ✅ Berhasil Semua       | Saga selesai | -                         |
| ❌ Gagal Order          | Gagal awal | -                         |
| ❌ Gagal Payment        | CancelOrder | Ya                        |
| ❌ Gagal Shipping       | CancelShipping → Refund → CancelOrder | Ya |

---

## 📁 Struktur Folder

```
ordersaga/
├── orchestrator/          # Main client
│   └── main.go
├── order/
│   └── server/
│       └── main.go
├── payment/
│   └── server/
│       └── main.go
├── shipping/
│   └── server/
│       └── main.go
├── proto/
│   ├── order.proto
│   ├── payment.proto
│   ├── shipping.proto
│   ├── order/             # hasil generate
│   ├── payment/
│   └── shipping/
├── go.mod
├── go.sum
└── README.md
```

---

## 🚀 Menjalankan Program

```bash
# Generate file .pb.go dari masing-masing .proto
protoc --go_out=. --go-grpc_out=. proto/order.proto
protoc --go_out=. --go-grpc_out=. proto/payment.proto
protoc --go_out=. --go-grpc_out=. proto/shipping.proto

# Jalankan di 3 terminal terpisah
go run order/server/main.go
go run payment/server/main.go
go run shipping/server/main.go

# Jalankan Orchestrator (client utama)
go run orchestrator/main.go
```

---

## 📄 Laporan

Penjelasan lengkap skenario, hasil output, dan tangkapan layar dapat dilihat di file:  
📄 **PWL_TUGAS2_122140004.pdf**

---

## 🙋 Identitas

- 🧑 Nama: Billy
- 🆔 NIM: 122140004
- 🏫 Mata Kuliah: Pemrograman Web Lanjut
