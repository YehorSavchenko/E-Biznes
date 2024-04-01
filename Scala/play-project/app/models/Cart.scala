package models

case class CartItem(productId: Long, quantity: Int)
case class Cart(items: List[CartItem])

object Cart {
  var cart: Cart = Cart(List())

  def findAll: Cart = cart
  def addItem(item: CartItem): Unit = {
    val existingItem = cart.items.find(_.productId == item.productId)
    val updatedItems = existingItem match {
      case Some(found) => cart.items.map { i =>
        if (i.productId == item.productId) i.copy(quantity = i.quantity + item.quantity) else i
      }
      case None => cart.items :+ item
    }
    cart = Cart(updatedItems)
  }
  def removeItem(productId: Long): Unit = {
    cart = Cart(cart.items.filterNot(_.productId == productId))
  }

}
