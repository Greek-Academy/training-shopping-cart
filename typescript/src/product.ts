export class Product {
  /**
   * 商品名
   * @var {string}
   */
  public name: string;

  /**
   * 単価
   * @var {number}
   */
  public price: number;

  /**
   * @param {string} code
   */
  constructor(public code: string) {
    // コード値に応じた商品情報を格納する
    switch (code) {
      case "grape":
        this.name = "ぶどう";
        this.price = 100;
        break;
      case "apple":
        this.name = "りんご";
        this.price = 50;
        break;
      case "orange":
        this.name = "みかん";
        this.price = 70;
        break;
      default:
        throw new Error("codeが不正な値です");
    }
  }
}
