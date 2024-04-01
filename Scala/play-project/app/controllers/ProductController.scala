package controllers

import models.Product
import play.api.libs.json.{Json, OFormat}
import play.api.mvc._
import javax.inject._

class ProductController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {
  implicit val productFormat: OFormat[Product] = Json.format[Product]

  def listProducts = Action {
    val products = Product.findAll
    Ok(Json.toJson(products))
  }

  def showProduct(id: Long) = Action {
    Product.findById(id) match {
      case Some(product) => Ok(Json.toJson(product))
      case None => NotFound
    }
  }

}
