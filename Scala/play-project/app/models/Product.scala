package models

case class Product(id: Long, name: String, description: String, price: Double)

object Product {
  var products: List[Product] = List(
    Product(1, "Product 1", "Description 1", 100.0),
    Product(2, "Product 2", "Description 2", 200.0)
  )

  def findAll: List[Product] = products.sortBy(_.id)

  def findById(id: Long): Option[Product] = products.find(_.id == id)

  def add(product: Product): Unit = {
    products = products :+ product
  }

  def delete(id: Long): Unit = {
    products = products.filterNot(_.id == id)
  }

  def update(id: Long, product: Product): Option[Product] = {
    val index = products.indexWhere(_.id == id)
    if (index == -1) None
    else {
      products = products.updated(index, product)
      Some(product)
    }
  }

}