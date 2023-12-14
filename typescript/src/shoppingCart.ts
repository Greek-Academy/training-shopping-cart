import { Product } from "./product";

export type Stock = {
  product: Product;
  count: number;
};

export type Order = { count: number; total: number };

export class ShoppingCart {
  /**
   * カートに入った商品情報
   * @readonly
   * @var {{[key: string]: Stock}}
   */
  public readonly products: {
    [key: string]: Stock;
  } = {};

  /**
   * カートに商品を追加する
   * @param {string} code
   * @param {number} count
   * @returns {number} 追加後の個数を返却する
   */
  public add(code: string, count: number = 1): number {
    if (code in this.products) {
      this.products[code].count += count;
    } else {
      this.products[code] = {
        product: new Product(code),
        count,
      };
    }
    return this.products[code].count;
  }

  /**
   * カートから商品を取り出す
   * @param code
   * @param count
   * @returns {number} 削除後の個数を返却する
   */
  public remove(code: string, count: number = 1): number {
    if (code in this.products) {
      this.products[code].count -= count;
      // 個数が0になったら削除する
      if (this.products[code].count === 0) {
        delete this.products[code];
        return 0;
      } else {
        return this.products[code].count;
      }
    } else {
      return 0;
    }
  }

  /**
   * 注文する
   * @returns {Order} 商品点数と合計金額を返却する
   */
  public order(): Order {
    if (Object.keys(this.products).length === 0) {
      throw new Error("カートが空なので注文できません");
    }

    const order = { count: 0, total: 0 };
    for (const { count, product } of Object.values(this.products)) {
      order.count += count;
      order.total += product.price * count;
    }
    return order;
  }
}
