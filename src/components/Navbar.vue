<template>
    <div id="nav">
        <router-link to="/" class="nav-home">
            <h1 class="nav-title">Bitcoin Wallet Generator</h1>
        </router-link>
        
        <router-link class="nav-route" to="/">
            <img src="@/assets/btclogo.png" class="nav-btc"/>
        </router-link>

        <router-link to="/about" class="nav-route">
            <button class="nav-btn">About</button>
        </router-link>

        <router-link to="/donate" class="nav-route">
            <button class="nav-btn">Donate</button>
        </router-link>

        
    </div>
</template>

<script>
export default {
    name: 'Navbar',
    props: {
    }
}
</script>

<style scoped>
.nav-home {
    color: white;
    text-decoration: none;
}

.nav-title {
    margin: .5rem;
    padding: 0;
    font-size: 1.7em;
    cursor: pointer;
}

.nav-btc {
    display: none;
}

.nav-route {
    margin-top: auto;
    margin-bottom: auto;
}

.nav-btn {
    background-color: rgb(45, 200, 140);
    border: none;
    border-radius: 0.4rem;
    font-size: 1.25rem;
    color: white;
    cursor: pointer;
    height: fit-content;
    padding: .4rem;
    margin-top: auto;
    margin-bottom: auto;
    outline: none;
}

.nav-btn:active {
    transform: scale(1.03);
}

#nav {
    color: white;
    width: 100vw;
    display: flex;
    flex-direction: row;
    gap: 2rem;
    background: #384a5e;
    padding: 0;
    margin: 0;
    outline: none;
    border: none;
    --tw-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
}

@media screen and (max-width: 800px) {

    #nav {
        color: white;
        width: 100vw;
        display: flex;
        flex-direction: row;
        gap: 1rem;
        background: #384a5e;
        padding: 0;
        margin: 0;
        outline: none;
        border: none;
        --tw-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
        box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
    }


    .nav-home {
        display: none;
    }
    .nav-btc {
        width: 45px;
        height: 45px;
        padding: .3rem .45rem .3rem .8rem;
        display: block;
    }

    .nav-btn {
        background-color: rgb(45, 200, 140);
        border: none;
        border-radius: 0.4rem;
        font-size: 1.25rem;
        color: white;
        cursor: pointer;
        height: fit-content;
        padding: .3rem;
        margin-top: auto;
        margin-bottom: auto;
        outline: none;
    }

}

</style>