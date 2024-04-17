package com.ebiznes

import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.events.message.MessageReceivedEvent
import net.dv8tion.jda.api.hooks.ListenerAdapter

fun main() {
    val jda = JDABuilder.createDefault("MTIzMDA4MDkxODM3MTMwMzQ2NA.GilO7h.dC847uiRzZZC6b515524SL0YGFWvcMD9AaLq6Q")
        .addEventListeners(BotListener())
        .build()
        .awaitReady()

    embeddedServer(Netty, port = 8080) {
        routing {
            get("/") {
                call.respondText("Hello, world!", ContentType.Text.Plain)
            }

            post("/send-message") {
                val channelId = "1183337671444144189"
                val message = "Hello from Ktor!"
                val channel = jda.getTextChannelById(channelId)
                channel?.sendMessage(message)?.queue()
                call.respondText("Message sent to Discord!", ContentType.Text.Plain)
            }
        }
    }.start(wait = true)
}

class BotListener : ListenerAdapter() {
    override fun onMessageReceived(event: MessageReceivedEvent) {
        if (event.author.isBot) return
        println("Received message: ${event.message}")
    }
}
