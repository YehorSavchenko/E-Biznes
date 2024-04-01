package controllers

import models.Product
import play.api.libs.json.{Json, OFormat}
import play.api.mvc._
import javax.inject._

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
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

  def addProduct = Action(parse.json) { request =>
    request.body.validate[Product].map { product =>
      Product.add(product)
      Created(Json.toJson(product))
    }.getOrElse(BadRequest("Invalid JSON"))
  }

  def updateProduct(id: Long) = Action(parse.json) { request =>
    request.body.validate[Product].map { product =>
      Product.update(id, product) match {
        case Some(p) => Ok(Json.toJson(p))
        case None => NotFound
      }
    }.getOrElse(BadRequest("Invalid JSON"))
  }

  def deleteProduct(id: Long) = Action {
    Product.delete(id)
    NoContent
  }

}
