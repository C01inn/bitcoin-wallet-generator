<template>
  <div class="about">
    <Navbar class="navbar"/>
    <div class="main-div">
      <h1 id="bitcoin-wallet-generator">Bitcoin Wallet Generator</h1>
      <h4 id="-bitcoin-wallet-generator-https-bitcoin-wallet-netlify-app-is-a-simple-website-that-generates-bitcoin-wallets-with-a-public-and-private-key-">Bitcoin Wallet Generator is a simple website that generates bitcoin wallets with a public and private key. This website is 100% open source and the source is available on github <a href="https://github.com/C01inn/bitcoin-wallet-generator" target="_blank">here</a>.</h4>
      <h2 id="tools-used">Tools Used:</h2>
      <h3 id="1-vuejs" class="tool">1. Vuejs</h3>
      <h4 class="tool-desc" id="this-website-is-a-single-page-application-that-was-created-using-the-vuejs-javascript-framework-https-vuejs-org-">This website is a Single Page Application that was created using the <a href="https://vuejs.org/">Vuejs Javascript framework</a>.</h4>
      <h3 class="tool" id="2-bitcoin-address-generator">2. bitcoin-address-generator</h3>
      <h4 class="tool-desc" id="the-bitcoin-wallets-are-generated-using-the-bitcoin-address-generator-https-github-com-niyas-bitcoin-address-generator-javascript-library-">The bitcoin wallets are generated using the <a href="https://github.com/niyas/bitcoin-address-generator">bitcoin-address-generator</a> javascript library.</h4>
      <h3 class="tool" id="3-qrious">3. Qrious</h3>
      <h4 class="tool-desc" id="-qrious-https-github-com-neocotic-qrious-is-used-to-generate-qrcodes-for-the-wallet-keys-"><a href="https://github.com/neocotic/qrious">Qrious</a> is used to generate qrcodes for the wallet keys.</h4>
      <h3 class="tool" id="4-jspdf">4. jspdf</h3>
      <h4 class="tool-desc" id="-jspdf-https-github-com-mrrio-jspdf-is-used-to-create-the-paper-wallets-in-pdf-format-"><a href="https://github.com/MrRio/jsPDF">Jspdf</a> is used to create the paper wallets in pdf format.</h4>
      <h2 id="donate">Donate</h2>
      <p>If you would wish to support this project then please feel free to donate by sending bitcoin to the address below.</p>
      <p class="bitcoin-address">3KgvXoAYCvW26XirrMnrYTDqkyBLzyXf5R</p>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'

export default {
  name: "About",
  components: {
    Navbar
  }
}
</script>

<style scoped>

.main-div {
  position: relative;
  border: 1px solid #eee;
  border-radius: .75rem;
  padding: 0 1.5rem 0 2rem;
  margin-top: 5rem;
  margin-right: auto;
  margin-left: auto;
  width: 60%;
  text-align: left;
}

@media screen and (max-width: 1560px) {
  .main-div {
    position: relative;
    border: 1px solid #eee;
    border-radius: .75rem;
    padding: 0 1.5rem 0 2rem;
    margin-top: 5rem;
    margin-right: auto;
    margin-left: auto;
    width: 70%;
    text-align: left;
  }
}

@media screen and (max-width: 960px) {
  .main-div {
    position: relative;
    border: 1px solid #eee;
    border-radius: .75rem;
    padding: 0 1.5rem 0 2rem;
    margin-top: 5rem;
    margin-right: auto;
    margin-left: auto;
    width: 85vw;
    text-align: left;
  }
}

.about {
  padding: 0;
  margin: 0;
  position: absolute;
  left: 0;
  top: 0;
}

.navbar {
  position: sticky;
  top: 0;
  left: 0;
  padding: 0;
  margin: 0;
}

#tools-used {
  margin-top: 2.5rem;
  text-decoration: underline;
}

.tool {
  text-indent: 1rem;
}

.tool-desc {
}

.bitcoin-address {
  color: rgb(66, 190, 66);
  font-size: 22px;
}

</style>
