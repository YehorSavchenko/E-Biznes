package controllers

import javax.inject._
import play.api.mvc._
import models.{Cart, CartItem}
import play.api.libs.json.{Json, OFormat}

@Singleton
class CartController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  implicit val cartItemFormat: OFormat[CartItem] = Json.format[CartItem]
  implicit val cartFormat: OFormat[Cart] = Json.format[Cart]

  def showCart = Action {
    Ok(Json.toJson(Cart.findAll))
  }

  def addItemToCart = Action(parse.json) { request =>
    request.body.validate[CartItem].map { cartItem =>
      Cart.addItem(cartItem)
      Created(Json.toJson(cartItem))
    }.getOrElse(BadRequest("Invalid JSON"))
  }

  def removeItemFromCart(productId: Long) = Action {
    Cart.removeItem(productId)
    Ok(s"Item $productId removed from cart")
  }
}
