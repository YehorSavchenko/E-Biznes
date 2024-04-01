package models

case class Category(id: Long, name: String)

object Category {
  var categories: List[Category] = List(
    Category(1, "Category 1"),
    Category(2, "Category 2")
  )

  def findAll: List[Category] = categories.sortBy(_.id)
  def findById(id: Long): Option[Category] = categories.find(_.id == id)
  def add(category: Category): Unit = categories = categories :+ category
  def delete(id: Long): Unit = categories = categories.filterNot(_.id == id)

  def update(id: Long, category: Category): Option[Category] = {
    val index = categories.indexWhere(_.id == id)
    if (index != -1) {
      categories = categories.updated(index, category)
      Some(category)
    } else None
  }
}
