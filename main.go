package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	fmt.Println(err)

	if err != nil {
		return
	}
	htmlStr := `
	<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link
      href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;500;600;700&display=swap"
      rel="stylesheet"
    />
    <title>Document</title>
    <style>
      body {
        font-family: "Poppins", serif;
        font-size: 12px;
      }
      .flex {
        display: -webkit-box;
      }
      .flex-col {

      }
      .justify-center {
        /* -webkit-justify-content: center;
        justify-content: center; */
        -webkit-box-pack: center; 
      }
      .justify-between {
        /* -webkit-justify-content: space-between;
        justify-content: space-between; */
        -webkit-box-pack: justify; 
      }
      .justify-start {
        /* -webkit-justify-content: flex-start;
        justify-content: flex-start; */
        -webkit-box-pack: start; 
      }
      .justify-end {
        /* -webkit-justify-content: flex-end;
        justify-content: flex-end; */
        -webkit-box-pack: end; 
      }
      .items-center {
        /* -webkit-align-items: center;
        align-items: center; */
        -webkit-box-align: center;
      }
      .items-end {
        -webkit-box-align: end;
      }
      .w-full {
        width: 100%;
      }
      .mx-auto {
        margin: auto;
      }
      .my-10 {
        margin-top: 40px;
        margin-bottom: 40px;
      }
      .m-5 {
        margin: 20px;
      }
      .bg-white {
        background-color: white;
      }
      .text-orange {
        color: #f6921e;
      }
      .font-bold {
        font-weight: 700;
      }
      .bg-image {
        background-image: url("https://storage.googleapis.com/st-core/public/temp/1659939985-image 14.png");
        background-size: contain;
        background-repeat: no-repeat;
        background-position: center;
      }
      .img-logo {
        background-image: url("https://storage.googleapis.com/st-core/public/temp/1659940266-Group.png");
        background-size: contain;
        background-repeat: no-repeat;
        background-position: left;
        width: 165px;
        height: 50px;
      }
      .gap-8 {
        gap: 32px;
      }
      .font-normal {
        font-weight: 500;
      }
    </style>
  </head>
  <body>
    <div style="width: 595px; height: 842px" class="mx-auto bg-white bg-image">
      <div class="m-5">
        <div class="flex justify-between my-10">
          <div class="img-logo"></div>
          <div style="margin-top: 12px;">
            <div class="font-bold">Invoice</div>
            <div class="text-orange font-normal">INV/220620/TKT/44150001239</div>
          </div>
        </div>
        <div class="flex justify-start font-bold">Diterbitkan Atas Nama</div>
        <div style="margin-top: 10px" class="flex">
          <div style="width: 200px" class="font-normal">Penjual</div>
          <div style="width: 300px" class="font-normal">: PT. Eventori Media Semesta</div>
        </div>
        <div style="margin-top: 20px" class="font-bold">Untuk</div>
        <div style="margin-top: 10px" class="flex">
          <div style="width: 200px" class="font-normal">Pembeli</div>
          <div style="width: 300px" class="font-normal">: Raden Agung Rafi Prasetio</div>
        </div>
        <div style="margin-top: 10px" class="flex">
          <div style="width: 200px" class="font-normal">Tanggal Pembelian</div>
          <div style="width: 300px" class="font-normal">: 20 Juli 2022</div>
        </div>
        <div style="margin-top: 10px" class="flex">
          <div style="width: 200px" class="font-normal">Event</div>
          <div style="width: 300px" class="font-normal">: Eventori Festival</div>
        </div>
        <div
          style="border: 2px dashed #000000; margin: 10px 0px 10px 0px"
        ></div>
        <div style="margin: 0px 5px 0px 5px" class="flex justify-between">
          <div class="font-bold" style="width: 155px;">Kategori Tiket</div>
          <div class="flex justify-between font-bold" style="width: 200px;">
            <div style="width: 100px;">Jumlah</div>
            <div style="width: 100px;">Harga Satuan</div>
          </div>
          <div class="font-bold">Total Harga</div>
        </div>
        <div
          style="border: 2px dashed #000000; margin: 10px 0px 10px 0px"
        ></div>
        <div style="margin: 0px 5px 0px 5px" class="flex justify-between">
          <div class="text-orange font-bold" style="width: 155px">East VIP</div>
          <div class="flex justify-between" style="width: 200px;">
            <div style="width: 100px;" class="font-normal">2</div>
            <div style="width: 100px;" class="font-normal">Rp 1.750.000</div>
          </div>
          <div class="font-normal">Rp 3.500.000</div>
        </div>
        <div
          style="
            margin: 0px 5px;
            width: 150px;
            font-size: 10px;
            padding-top: 5px;
          "
          class="font-normal"
        >
          <strong>Venue : </strong> JCC Convention Center
        </div>
        <div
          style="border: 1px dashed #000000; margin: 10px 0px 10px 0px"
        ></div>
        <div class="flex justify-end" style="gap: 50px">
          <div class="font-bold" style="font-size: 12px; width: 200px;">TOTAL HARGA TIKET (2 TIKET)</div>
          <div class="font-bold" style="font-size: 12px; width: 100px; text-align: end;">Rp 3.500.000</div>
        </div>
        <div class="flex justify-end" style="gap: 50px; margin-top: 10px;">
          <div style="width: 200px;" class="font-normal">Biaya Layanan</div>
          <div style="width: 100px; text-align: end;" class="font-normal">Rp 175.000</div>
        </div>
        <div class="flex justify-end">  
          <div
          style="border: 1px dashed #000000; margin: 10px 0px 10px 0px; width: 300px;"
        ></div>
        </div>
        <div class="flex justify-end" style="gap: 50px">
          <div class="font-bold" style="font-size: 12px; width: 200px;">TOTAL BELANJA</div>
          <div class="font-bold" style="font-size: 12px; width: 100px; text-align: end;">Rp 3.675.000</div>
        </div>
        <div class="flex justify-end" style="gap: 50px; margin-top: 10px;">
          <div style="width: 200px;" class="font-normal">PROMO</div>
          <div style="width: 100px; text-align: end;" class="font-normal">- Rp 500.000</div>
        </div>
        <div class="flex justify-end">  
          <div
          style="border: 1px dashed #000000; margin: 10px 0px 10px 0px; width: 300px;"
        ></div>
        </div>
        <div class="flex justify-end" style="gap: 50px">
          <div class="font-bold" style="font-size: 12px; width: 200px;">TOTAL TAGIHAN</div>
          <div class="font-bold" style="font-size: 12px; width: 100px; text-align: end;">Rp 3.175.000</div>
        </div>
        <div class="flex justify-end">  
          <div
          style="border: 1px dashed #000000; margin: 10px 0px 10px 0px; width: 100%;"
        ></div>   
        </div>
        <div class="flex justify-end" style="gap: 50px">
          <div style="font-size: 12px; width: 150px;" class="font-normal">Metode Pembayaran</div>
          <div class="font-bold" style="font-size: 12px; width: 150px; text-align: end;">BCA Virtual Account</div>
        </div>
        </div>
        <div class="flex items-end" style="height: 200px; width: 350px;">
          <div class="font-normal">Invoice ini sah dan diproses oleh komputer
            Silahkan hubungi <strong class="text-orange">
              Eventori CS
            </strong>  apabila perlu bantuan</div>
        </div>
      </div>
    </div>
  </body>
</html>
	`

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(htmlStr)))
	pdfg.MarginLeft.Set(5)
	pdfg.MarginRight.Set(5)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	fmt.Println(err)

	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile("./tes24.pdf")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")
}
