<template>
    <div class="loader" :style="{width: width}">
        <svg class="circular" viewBox="25 25 50 50">
        <circle
            class="path"
            :class="{ color: !!color }"
            cx="50"
            cy="50"
            r="20"
            fill="none"
            :stroke-width="strokeWidth"
            stroke-miterlimit="10"
        />
        </svg>
    </div>
</template>

<script>
export default {
    name: "Spinner",
    props: {
        width: {
            type: [String, Number],
            default: "40px"
        },
        strokeWidth: {
            type: [String, Number],
            default: 4
        },
        color: {
            type: Boolean,
            default: false
        }
    }
}
</script>

<style scoped>
@keyframes rotate {
  100% {
    transform: rotate(360deg);
  }
}

@keyframes dash {
  0% {
    stroke-dasharray: 1, 200;
    stroke-dashoffset: 0;
  }
  50% {
    stroke-dasharray: 89, 200;
    stroke-dashoffset: -35px;
  }
  100% {
    stroke-dasharray: 89, 200;
    stroke-dashoffset: -124px;
  }
}

@keyframes color {
  100% {

  }
  0% {
    stroke: #d62d20;
  }
  40% {
    stroke: #0057e7;
  }
  66% {
    stroke: #008744;
  }
  80% {

  }
  90% {
    stroke: #ffa700;
  }
}

.loader {
  position: relative;
  margin: 0 auto;
}

.loader:before {
    content: "";
    display: block;
    padding-top: 100%;
}

.circular {
  animation: rotate 1.2s linear infinite;
  height: 100%;
  transform-origin: center center;
  width: 100%;
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
}

.path {
  stroke-dasharray: 1, 200;
  stroke-dashoffset: 0;
  stroke: #0057e7;
  stroke-linecap: round;
  animation: dash 1.5s ease-in-out infinite;
}

.path > color {
    animation: dash 1.5s ease-in-out infinite, color 6s ease-in-out infinite;
}

</style>