package main

import (
	"time"
)

type Transaksi struct {
	ID     int
	Waktu  time.Time
	Nama   string
	Sumber string
	total  float64
}

type RiwayatTransaksi struct {
	Withdraw float64
	Deposit  float64
}

type Asetcrypto struct {
	bitcoin  float64
	ethereum float64
	solana   float64
	BNB      float64
}

type saldoPengguna struct {
	BTC        float64
	ethereum   float64
	Stablecoin float64
}

func Crypto() {

}
