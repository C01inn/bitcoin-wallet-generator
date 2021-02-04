<template>
    <div id="donate">
        <Navbar class="navbar"/>
        <!-- donate card-->
        <DonateCard class="don" id="don-card"/>
    </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import DonateCard from '@/components/DonateCard.vue'

export default {
    name: "Donate",
    components: {
        Navbar,
        DonateCard,
    }
}
</script>
<style scoped>
#donate {
    padding: 0;
    margin: 0;
    position: absolute;
    left: 0;
    top: 0;
}

.don {
    position: fixed;
    z-index: 4;
    top: 40%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
    border: 1px solid black;
    width: 35vw;
    height: 40vh;
    border-radius: 1.5rem;
    --tw-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
    display: block;
    background: white;
}

.navbar {
    position: sticky;
    top: 0;
    left: 0;
    padding: 0;
    margin: 0;
}



@media screen and (max-width: 1500px) {
    .don {
        position: fixed;
        z-index: 4;
        top: 40%;
        left: 50%;
        transform: translateX(-50%) translateY(-50%);
        border: 1px solid black;
        width: 47.5vw;
        height: 37vh;
        border-radius: 1.5rem;
        --tw-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
        box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
        display: block;
        background: white;
    }
}


@media screen and (max-width: 1100px) {
    .don {
        position: fixed;
        z-index: 4;
        top: 40%;
        left: 50%;
        transform: translateX(-50%) translateY(-50%);
        border: 1px solid black;
        width: 55vw;
        height: 37vh;
        border-radius: 1.5rem;
        --tw-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
        box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
        display: block;
        background: white;
    }
}
@media screen and (max-width: 990px) {
    .don {
        position: fixed;
        z-index: 4;
        top: 40%;
        left: 50%;
        transform: translateX(-50%) translateY(-50%);
        border: 1px solid black;
        width: 75vw;
        height: 37vh;
        border-radius: 1.5rem;
        --tw-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
        box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
        display: block;
        background: white;
    }
}

@media screen and (max-width: 700px) {
    .don {
        position: fixed;
        z-index: 4;
        top: 40%;
        left: 50%;
        transform: translateX(-50%) translateY(-50%);
        border: 1px solid black;
        width: 92.5vw;
        height: 37vh;
        border-radius: 1.5rem;
        --tw-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
        box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
        display: block;
        background: white;
    }
}


</style>