<template>
  <div id="paper-wallet"></div>
</template>

<script>
import QRious from "qrious";
import { jsPDF } from "jspdf";

export default {
  name: "PaperWallet",
  data() {
    return {
      publicKey: null,
      privateKey: null,
      privQrCode: null,
      pubQrCode: null,
    };
  },
  methods: {
    createPaperWallet(priv, pub) {
      this.publicKey = pub;
      this.privateKey = priv;

      let pubQr = new QRious({
        element: document.getElementById("wallet-can"),
        value: pub,
        background: "white", // background color
        foreground: "black", // foreground color
        backgroundAlpha: 1,
        foregroundAlpha: 1,
        level: "L", // Error correction level of the QR code (L, M, Q, H)
        mime: "image/png", // MIME type used to render the image for the QR code
        size: 300, // Size of the QR code in pixels.
        padding: null, // padding in pixels
      });

      let privQr = new QRious({
        element: document.getElementById("wallet-can"),
        value: priv,
        background: "white", // background color
        foreground: "black", // foreground color
        backgroundAlpha: 1,
        foregroundAlpha: 1,
        level: "L", // Error correction level of the QR code (L, M, Q, H)
        mime: "image/png", // MIME type used to render the image for the QR code
        size: 300, // Size of the QR code in pixels.
        padding: null, // padding in pixels
      });

      this.pubQrCode = pubQr.toDataURL("image/png");
      this.privQrCode = privQr.toDataURL("image/png");

      // create pdf
      /* */

      let doc = new jsPDF();

      doc.setFontSize(30);
      doc.text("Public Key: ", 20, 25);

      doc.setFontSize(15)
      doc.setTextColor(5, 150, 105);
      let splitPub = doc.splitTextToSize(this.publicKey, 180)
      doc.text(20, 35, splitPub)
      doc.addImage(this.pubQrCode, "PNG", 20, 40, 40, 40);

      // private key
      doc.setFontSize(30);
      doc.setTextColor(0, 0, 0);
      doc.text("Private Key: ", 20, 100);

      doc.setFontSize(15);
      doc.setTextColor(220, 38, 38);
      let splitKey = doc.splitTextToSize(this.privateKey, 180);

      doc.text(20, 110, splitKey)
      doc.addImage(this.privQrCode, "PNG", 20, 125, 40, 40)

      doc.setTextColor(0, 0, 0);
      let splitWarning = doc.splitTextToSize(`Warning: do not share or give out your private key. Your private key is for you to access the funds in your wallet.`, 180);
      doc.text(16, 195, splitWarning)


      doc.save("wallet.pdf")
    },
  },
};
</script>
