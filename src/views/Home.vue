<template>
  <div class="home">
      <Navbar class="navbar"/>
      <h3 class="wallets-generated" style="display: none;">Over {{numWallets}} Wallets Created!</h3>
      <button class="gen-btn" @click="generateWallet()">Generate Bitcoin Wallet</button>
      <br>
      <button v-if="pubKey && privKey" class="paper-btn" @click="getPaperWallet()">Get Paper Wallet</button>

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
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import Spinner from '@/components/Spinner.vue'

// lazy load paper wallet utils
const paperWallet = () => import('@/components/paperWallet.js');

import { createWalletAddress } from 'bitcoin-address-generator'

export default {
  name: 'Home',
  components: {
    Navbar,
    Spinner,
  },
  data() {
    return {
      pubKey: null,
      privKey: null,
      apiUrl: 'http://127.0.0.1:9990',
      numWallets: null,
    }
  },
  created() {
  },
  methods: {
    generateWallet() {
      this.pubKey, this.privKey = '';
      document.querySelector("#load-spin").style.display = 'block';
      var keyGen = setInterval(() => {
          createWalletAddress(response => {

            this.pubKey = response.address;
            this.privKey = response.key;

            this.numWallets++;

            document.querySelector("#load-spin").style.display = 'none';
            clearInterval(keyGen);
          });
      }, 3500);
    },
    getPaperWallet() {
      // download a paper wallet in pdf format

      // ensure that user has generated a wallet
      if (this.privKey && this.pubKey) {
        // must use promise because file is lazy loaded
        paperWallet().then((wallets) => {
          wallets.createPaperWallet(this.privKey, this.pubKey);
        })
      }
    }
  }
}


</script>
<style scoped>
/* tamil */
@font-face {
  font-family: 'Catamaran';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/catamaran/v8/o-0bIpQoyXQa2RxT7-5B6Ryxs2E_6n1iPHjd5bjdu2ui.woff2) format('woff2');
  unicode-range: U+0964-0965, U+0B82-0BFA, U+200C-200D, U+20B9, U+25CC;
}
/* latin-ext */
@font-face {
  font-family: 'Catamaran';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/catamaran/v8/o-0bIpQoyXQa2RxT7-5B6Ryxs2E_6n1iPHjd5aDdu2ui.woff2) format('woff2');
  unicode-range: U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}
/* latin */
@font-face {
  font-family: 'Catamaran';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/catamaran/v8/o-0bIpQoyXQa2RxT7-5B6Ryxs2E_6n1iPHjd5a7duw.woff2) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}

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


.wallets-generated {
  display: none;
}

@media screen and (min-width: 1300px) {
  .wallets-generated {
    position: fixed;
    right: 3rem;
    top: 4rem;
    font-size: 2em;
    display: block;
    font-family: 'Catamaran', sans-serif;
  }
}

.paper-btn {
  background-color: rgba(129, 140, 2480, 1);
  border: none;
  border-radius: 0.4rem;
  padding: .5rem .6rem .5rem .6rem;
  font-size: 1.1rem;
  line-height: 1.6rem;
  color: white;
  cursor: pointer;
  outline: none;
  margin-top: 2rem;
  overflow-x: hidden;
}

.paper-btn:hover {
  box-shadow: rgba(0, 0, 0, 0.25) 0px 54px 55px, rgba(0, 0, 0, 0.12) 0px -12px 30px, rgba(0, 0, 0, 0.12) 0px 4px 6px, rgba(0, 0, 0, 0.17) 0px 12px 13px, rgba(0, 0, 0, 0.09) 0px -3px 5px;
  background-color: rgba(99, 102, 241, 1);

}

.paper-btn:active {
  transform: scale(1.03);
}


#PaperWallet {
  display: none;
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