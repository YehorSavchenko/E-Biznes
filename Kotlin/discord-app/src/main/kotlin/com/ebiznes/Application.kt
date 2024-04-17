package com.ebiznes

import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.coroutines.DelicateCoroutinesApi
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.events.message.MessageReceivedEvent
import net.dv8tion.jda.api.hooks.ListenerAdapter

val categories = listOf("Electronics", "Books", "Clothing")

fun main() {
    val jda = JDABuilder.createDefault("MTIzMDExNzEyNTM0ODM5Mjk4Ng.Gl8BSB.h7UMUhsbcGYImo58M0SWNkQgN8WbUHYeZjLXxw")
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

            get("/categories") {
                call.respondText(categories.toString())
            }
        }
    }.start(wait = true)
}

class BotListener : ListenerAdapter() {
    @OptIn(DelicateCoroutinesApi::class)
    override fun onMessageReceived(event: MessageReceivedEvent) {
        if (event.author.isBot) return

        val message = event.message.contentDisplay
        val channel = event.channel

        when {
            message.equals("!categories", ignoreCase = true) -> {
                channel.sendMessage("Fetching categories...").queue()
                GlobalScope.launch {
                    try {
                        val categories = fetchCategories()
                        channel.sendMessage(categories).queue()
                    } catch (e: Exception) {
                        channel.sendMessage("Error fetching categories: ${e.message}").queue()
                    }
                }
            }
        }
    }
}

suspend fun fetchCategories(): String {
    val client = HttpClient(CIO)
    return try {
        val response: HttpResponse = client.get("http://localhost:8080/categories")
        response.bodyAsText()
    } catch (e: Exception) {
        "Failed to fetch categories: ${e.message}"
    } finally {
        client.close()
    }
}
