<template>
  <div class="about">
    <Navbar class="navbar"/>

    <h1 id="bitcoin-wallet-generator">Bitcoin Wallet Generator</h1>
    <h4 id="bitcoin-wallet-generator-is-a-simple-website-that-generates-bitcoin-wallets-with-a-public-and-private-key. "> Bitcoin Wallet Generator is a simple website that generates bitcoin wallets with a public and private key.</h4>
    <p>‎</p>
    <h2 id="tools-used">Tools Used</h2>
    <h3 id="vuejs">1. Vuejs</h3>
    <h4 id="this-website-is-a-single-page-application-that-was-created-using-the-vuejs-javascript-framework.">This website is a Single Page Application that was created using the <a href="https://vuejs.org/">Vuejs Javascript framework</a>.</h4>
    <h3 id="bitcoin-address-generator">2. bitcoin-address-generator</h3>
    <h4 id="the-bitcoin-wallets-are-generated-using-the-bitcoin-address-generator-javascript-library.">The bitcoin wallets are generated using the <a href="https://github.com/niyas/bitcoin-address-generator">bitcoin-address-generator</a> javascript library.</h4>
    <h5 id="section">‎</h5>
    <h2 id="donate">Donate</h2>
    <p>If you would wish to support this project then please feel free to donate by sending bitcoin to the address below.</p>
    <p>3KgvXoAYCvW26XirrMnrYTDqkyBLzyXf5R</p>


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

</style>
