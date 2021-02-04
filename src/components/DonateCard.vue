<template>
    <div id="donate">
        <!--
        <svg @click="$parent.donateBtn()" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M24 20.188l-8.315-8.209 8.2-8.282-3.697-3.697-8.212 8.318-8.31-8.203-3.666 3.666 8.321 8.24-8.206 8.313 3.666 3.666 8.237-8.318 8.285 8.203z"/></svg>
        -->
        <h2>Support this Project by Donating!</h2>
        <h3>Send Bitcoin to the wallet address below to donate.</h3>
        <p class="green-text">3KgvXoAYCvW26XirrMnrYTDqkyBLzyXf5R</p>
    </div>
</template>

<script>
export default {
    name: "DonateCard",
}
</script>

<style scoped>
#donate {
    background: white;
}

.exit-btn {
    background: rgb(239, 68, 68);
    outline: none;
    border: none;
    position: absolute;
    bottom: 4rem;
    left: 50%;
    cursor: pointer;
    font-size: 1.7rem;
    color: white;
    border-radius: 0.5rem;
    padding: .4rem 2rem .4rem 2rem;
    transform: scale(.95) translateX(-50%);
}

.exit-btn:active {
    transform: scale(1.025) translateX(-50%);
}

h2 {
    margin-top: 2.5rem;
}

svg {
    position: absolute;
    right: .8rem;
    top: .4rem;
    fill: darkgray;
    cursor: pointer;
}

svg:hover {
    fill: rgb(145, 139, 139);
}

.green-text {
    color: rgb(66, 190, 66);
    font-size: 22px;
}

@media screen and (max-width: 600px) {
    .green-text {
        color: rgb(66, 190, 66);
        font-size: 17px;
        word-wrap: break-word;
        max-width: 95%;
        text-align: center;
        margin-right: auto;
        margin-left: auto;
    }

    h2 {
        font-size: 1.25em;
        max-width: 95%;
        text-align: center;
        margin-right: auto;
        margin-left: auto;
    }
    h3 {
        font-size: 1em;
        max-width: 95%;
        text-align: center;
        margin-right: auto;
        margin-left: auto;
    }

    .exit-btn {
        background: rgb(239, 68, 68);
        outline: none;
        border: none;
        position: absolute;
        bottom: 2rem;
        left: 50%;
        cursor: pointer;
        font-size: 1.7rem;
        color: white;
        border-radius: 0.5rem;
        padding: .4rem 2rem .4rem 2rem;
        transform: scale(.95) translateX(-50%);
        max-width: 95%;
        margin-right: auto;
        margin-left: auto;
    }

}

</style>