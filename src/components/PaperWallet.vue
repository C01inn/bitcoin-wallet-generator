<template>
    <div id="paper-wallet">
    </div>
</template>

<script>
import QRious from 'qrious';
import pdfMake from "pdfmake/build/pdfmake";
import pdfFonts from "pdfmake/build/vfs_fonts";
pdfMake.vfs = pdfFonts.pdfMake.vfs;

export default {
    name: "PaperWallet",
    data() {
        return {
            publicKey: null,
            privateKey: null,
            privQrCode: null,
            pubQrCode: null,
        }
    },
    methods: {
        createPaperWallet(priv, pub) {
            this.publicKey = pub;
            this.privateKey = priv;

            let pubQr = new QRious({
                element: document.getElementById('wallet-can'),
                value: pub,
                background: 'white', // background color
                foreground: 'black', // foreground color
                backgroundAlpha: 1,
                foregroundAlpha: 1,
                level: 'L', // Error correction level of the QR code (L, M, Q, H)
                mime: 'image/png', // MIME type used to render the image for the QR code
                size: 300, // Size of the QR code in pixels.
                padding: null // padding in pixels
            })

            let privQr = new QRious({
                element: document.getElementById('wallet-can'),
                value: priv,
                background: 'white', // background color
                foreground: 'black', // foreground color
                backgroundAlpha: 1,
                foregroundAlpha: 1,
                level: 'L', // Error correction level of the QR code (L, M, Q, H)
                mime: 'image/png', // MIME type used to render the image for the QR code
                size: 300, // Size of the QR code in pixels.
                padding: null // padding in pixels
            })

            this.pubQrCode = pubQr.toDataURL('image/png')
            this.privQrCode = privQr.toDataURL("image/png")

            // create pdf 

            let pdfDefinition = {
                content: [
                    {
                        text: "Bitcoin Paper Wallet",
                        fontSize: 30,
                        style: 'header',
                        margin: [5, 0, 0, 22]
                    },
                    {
                        text: "Public Key:",
                        style: `header`,
                        fontSize: 23,
                        margin: [5, 5, 0, 5]
                    },
                    {
                        text: this.publicKey,
                        fontSIze: 20,
                        margin: [10, 5, 0, 5],
                        color: 'green',
                    },
                    {
                        image: this.pubQrCode,
                        width: 200,
                        height: 200,
                        margin: [10, 0, 0, 5],
                    },
                    {
                        text: "Private Key:",
                        style: `header`,
                        fontSize: 23,
                        margin: [5, 15, 0, 5]
                    },
                    {
                        text: this.privateKey,
                        fontSIze: 18,
                        margin: [10, 5, 0, 5],
                        color: 'red',
                    },
                    {
                        image: this.privQrCode,
                        width: 200,
                        height: 200,
                        margin: [10, 0, 0, 0]
                    },
                    {
                        text: `Warning: do not share or give out your private key. Your private key is for you to access the funds in your wallet.`,
                        color: 'black',
                        fontSize: 12,
                        margin: [5, 10, 0, 0]
                    }
                ]
            }

            pdfMake.createPdf(pdfDefinition).download('wallet.pdf');

        }
    }
    
}
</script>