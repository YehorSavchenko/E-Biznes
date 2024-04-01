package controllers

import javax.inject._
import play.api.mvc._
import models.Category
import play.api.libs.json.{Json, OFormat}

@Singleton
class CategoryController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  implicit val categoryFormat: OFormat[Category] = Json.format[Category]

  def listCategories = Action {
    Ok(Json.toJson(Category.findAll))
  }

  def addCategory = Action(parse.json) { request =>
    request.body.validate[Category].map { category =>
      Category.add(category)
      Created(Json.toJson(category))
    }.getOrElse(BadRequest("Invalid JSON"))
  }

  def updateCategory(id: Long) = Action(parse.json) { request =>
    request.body.validate[Category].map { category =>
      Category.update(id, category) match {
        case Some(cat) => Ok(Json.toJson(cat))
        case None => NotFound
      }
    }.getOrElse(BadRequest("Invalid JSON"))
  }

  def deleteCategory(id: Long) = Action {
    Category.delete(id)
    NoContent
  }
}
