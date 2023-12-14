import { Order, ShoppingCart } from "./shoppingCart";

describe("ShoppingCart", () => {
  let cart: ShoppingCart;

  beforeEach(() => {
    cart = new ShoppingCart();
  });

  describe.each([
    ["apple", 2, 1, 1],
    ["orange", 3, 3, 0],
    ["grape", 10, 5, 5],
  ])(
    "add & remove",
    (fruit: string, add: number, remove: number, expected: number) => {
      const vFruit = fruit.padStart(6, " ");
      const vAdd = String(add).padStart(2, " ");
      const vRemove = String(remove).padStart(2, " ");

      it(`fruit: ${vFruit}, add: ${vAdd}, remove: ${vRemove}`, () => {
        expect(cart.add(fruit, add)).toBe(add);
        expect(cart.remove(fruit, remove)).toBe(expected);
      });
    }
  );

  describe.each<[[string, number][], Order]>([
    [
      [
        ["apple", 3], // 150
        ["apple", 3], // 150
        ["apple", 3], // 150
      ],
      { count: 9, total: 450 },
    ],
    [
      [
        ["apple", 20], // 1,000
        ["grape", 4], // 400
        ["orange", 8], // 560
      ],
      { count: 32, total: 1960 },
    ],
    [
      [
        ["grape", 3], // 300
        ["orange", 9], // 630
      ],
      { count: 12, total: 930 },
    ],
  ])("order", (products: [string, number][], expected: Order) => {
    it(`products: ${JSON.stringify(products)}`, () => {
      for (const [fruit, count] of products) {
        cart.add(fruit, count);
      }
      expect(cart.order()).toMatchObject(expected);
    });
  });
});
