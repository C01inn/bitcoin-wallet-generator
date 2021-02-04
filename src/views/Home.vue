<template>
  <div class="home">
      <Navbar class="navbar"/>
      <button class="gen-btn" @click="generateWallet()">Generate Bitcoin Wallet</button>

      <div id="wallet-data" class="wallet-card">
        <h1 class="wallet-detail">Your Wallet Details: </h1>
        <h3>Public Key: <p class="pubKey">{{pubKey ? pubKey : "N/A"}}</p></h3>
        <h3>Private Key: <p class="privKey">{{privKey ? privKey : "N/A"}}</p></h3>

        <!-- loading spinner -->

        <div id="load-spin">
          <h2 style="text-align: center;">Generating Wallet...</h2>
          <Spinner color width="100px" id="spinner"/>
        </div>
      </div>
      <!-- paper wallet -->
      <PaperWallet id="PaperWallet" ref="papierPortefeuille"/>
  </div>
</template>

<script>
// @ is an alias to /src
import Navbar from '@/components/Navbar.vue'
import Spinner from '@/components/Spinner.vue'
import PaperWallet from '@/components/PaperWallet.vue'

import { createWalletAddress } from 'bitcoin-address-generator'

export default {
  name: 'Home',
  components: {
    Navbar,
    Spinner,
    PaperWallet,
  },
  data() {
    return {
      pubKey: null,
      privKey: null,
    }
  },
  methods: {
    generateWallet() {
      this.pubKey, this.privKey = '';
      document.querySelector("#load-spin").style.display = 'block';
      var keyGen = setInterval(() => {
          createWalletAddress(response => {

            this.pubKey = response.address;
            this.privKey = response.key;

            this.$refs.papierPortefeuille.createPaperWallet(this.privKey, this.pubKey);


            document.querySelector("#load-spin").style.display = 'none';
            clearInterval(keyGen);
          });
      }, 3500);
  
    }
  }
}
</script>
<style scoped>


.home {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  padding: 0;
  margin: 0;
  position: absolute;
  left: 0;
  top: 0;
  overflow-x: hidden;
}




#PaperWallet {
  margin-top: 4rem;
  overflow-x: hidden;
  max-width: 90vw;
  margin-right: auto;
  margin-left: auto;
  transform: scale(1.3);
}

#spinner {
  position: absolute;
  left: 50%;
  top: 20%;
  transform: translateX(-50%) scale(1.25);
}

#load-spin {
  position: absolute;
  left: 50%;
  top: 17%;
  transform: translateX(-50%);
  width: 80%;
  height: 80%;
  z-index: 2;
  background: #eee;
  border-radius: 1rem;
  display: none;
}

.pubKey {
  color: rgb(66, 190, 66);
  display: block;
  max-width: 80%;
  margin: 1rem auto;
}
.privKey {
  color: red;
  display: block;
  max-width: 80%;
  margin: 1rem auto;
}

.wallet-card {
  border: 2px solid lightgray;
  margin-right: auto;
  margin-left: auto;
  position: relative;
  margin-top: 4rem;
  width: 800px;
  height: 400px;
  display: block;
  word-wrap: break-word;
  text-align: center;
}

.navbar {
  position: sticky;
  top: 0;
  left: 0;
  padding: 0;
  margin: 0;
}

.gen-btn {
  background-color: rgba(78, 145, 240, .9);
  border: none;
  border-radius: 0.4rem;
  padding: .5rem .6rem .5rem .6rem;
	font-size: 1.25rem;
  line-height: 1.75rem;
  color: white;
  cursor: pointer;
  outline: none;
  margin-top: 2.5rem;
  overflow-x: hidden;
}

@keyframes genBtnHover {
  from {

  }
  to {
    background-color: rgba(78, 145, 240, 1);
  }
}

.gen-btn:active {
  transform: scale(1.03, 1.03);
  border: none;
  outline: none;
}

.gen-btn:hover {
  animation: genBtnHover 1s ease;
  animation-fill-mode: forwards;
}


@media screen and (max-width: 1100px) {
  .wallet-card {
    border: 3px solid lightgray;
    margin-right: auto;
    margin-left: auto;
    position: relative;
    margin-top: 4rem;
    width: 80%;
    height: 400px;
    display: block;
    word-wrap: break-word;
    text-align: center;
    overflow-x: scroll;
  }

}


@media screen and (max-width: 700px) {

  .wallet-detail {
    max-width: 85%;
    margin-right: auto;
    margin-left: auto;
    font-size: 1.6em;
  }

  .gen-btn {
    background-color: rgba(78, 145, 240, .9);
    border: none;
    border-radius: 0.4rem;
    padding: .5rem .6rem .5rem .6rem;
    font-size: 1.25rem;
    line-height: 1.75rem;
    color: white;
    cursor: pointer;
    outline: none;
    margin-top: 1.5rem;
    overflow-x: hidden;
    transform: scale(.87);
  }

  .wallet-card {
    border: 3px solid lightgray;
    margin-right: auto;
    margin-left: auto;
    position: relative;
    margin-top: 2.2rem;
    width: 80%;
    height: 400px;
    display: block;
    word-wrap: break-word;
    text-align: center;
    overflow-x: scroll;
  }
}

</style>