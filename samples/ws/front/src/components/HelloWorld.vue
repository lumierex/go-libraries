<template>
    <h1>{{ msg }}</h1>

    <p>{{receiveMsg}}</p>

    <button @click="connectWebsocket">连接</button>
    <!--  <button @click="count++">count is: {{ count }}</button>
      <p>
        Edit
        <code>components/HelloWorld.vue</code> to test hot module replacement.
      </p>-->
    <button @click="sendMessage">发送数据</button>
</template>

<script lang="ts">
    import {ref, defineComponent} from 'vue'

    var ws : WebSocket
    export default defineComponent({
        name: 'HelloWorld',
        props: {
            msg: {
                type: String,
                required: true
            }
        },
        setup: () => {
            const count = ref(0)
            const receiveMsg = ref("")
            return {count, receiveMsg}
        },
        methods: {
            connectWebsocket: () => {
                alert("message")
                ws = new WebSocket("ws://localhost:8989/ws")
                ws.onopen = (evt) => {
                    console.log("connect ws")
                    ws.send("start send websocket")
                }

                ws.onmessage = (evt) => {
                    console.log("receive message: ", evt.data)
                }

                ws.onclose = (evt) => {
                    console.log("close ws")
                }
            },
            sendMessage() {
                ws.send("hi")
            }
        }
    })
</script>

<style scoped>
    a {
        color: #42b983;
    }

    label {
        margin: 0 0.5em;
        font-weight: bold;
    }

    code {
        background-color: #eee;
        padding: 2px 4px;
        border-radius: 4px;
        color: #304455;
    }
</style>
