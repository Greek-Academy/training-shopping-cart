<?php

namespace Tests;

use App\ShoppingCart;
use PHPUnit\Framework\TestCase;

final class ShoppingCartTest extends TestCase
{
  /**
   * @test
   * @dataProvider addRemoveDataProvider
   *
   * @param string $code
   * @param int $add
   * @param int $remove
   * @param int $expected
   * @return void
   */
  public function addRemove(
    string $code,
    int $add,
    int $remove,
    int $expected
  ): void {
    $shoppingCart = new ShoppingCart;
    $shoppingCart->add($code, $add);
    $actual = $shoppingCart->remove($code, $remove);
    $this->assertSame($expected, $actual);
  }

  /**
   * addRemove用 テストケース
   *
   * @return iterable
   */
  public static function addRemoveDataProvider(): iterable
  {
    yield ["apple", 2, 1, 1];
    yield ["orange", 3, 3, 0];
    yield ["grape", 10, 5, 5];
  }

  /**
   * @test
   * @dataProvider orderDataProvider
   *
   * @param array{ string, int } $products
   * @param array{ count: int, total: int} $expected
   * @return void
   */
  public function order(
    array $products,
    array $expected
  ): void {
    $shoppingCart = new ShoppingCart;

    foreach ($products as [$code, $count]) {
      $shoppingCart->add($code, $count);
    }

    $actual = $shoppingCart->order();
    $this->assertSame($expected, $actual);
  }

  /**
   * order用 テストケース
   *
   * @return iterable{ array{ string, int }, array{ count: int, total: int} }
   */
  public static function orderDataProvider(): iterable
  {
    yield [
      [
        ["apple", 3],
        ["apple", 3],
        ["apple", 3],
      ],
      ['count' => 9, 'total' => 450]
    ];

    yield [
      [
        ["apple", 20],
        ["grape", 4],
        ["orange", 8],
      ],
      ['count' => 32, 'total' => 1960]
    ];

    yield [
      [
        ["grape", 3],
        ["orange", 9],
      ],
      ['count' => 12, 'total' => 930]
    ];
  }
}
