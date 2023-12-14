<?php

namespace App;

use Exception;

final class ShoppingCart
{
  /**
   * @var array<string, array{ product: Product, count: int }>
   */
  private array $products = [];

  public function __construct()
  {
  }

  /**
   * 商品の追加
   *
   * @param string $code
   * @param int $count
   * @return int
   */
  public function add(string $code, int $count): int
  {
    if (isset($this->products[$code])) {
      $this->products[$code]['count'] += $count;
    } else {
      $this->products[$code] = [
        'product' => new Product($code),
        'count' => $count
      ];
    }
    return $this->products[$code]['count'];
  }

  /**
   * 商品の取り出し
   *
   * @param string $code
   * @param int $count
   * @return int
   */
  public function remove(string $code, int $count): int
  {
    if (isset($this->products[$code])) {
      $this->products[$code]['count'] -= $count;
      // 個数が0になったら削除
      if ($this->products[$code]['count'] === 0) {
        unset($this->prodicts[$code]);
        return 0;
      } else {
        return $this->products[$code]['count'];
      }
    }
    return 0;
  }

  /**
   * @return array{ count: int, total: int }
   */
  public function order(): array
  {
    if (empty($this->products)) {
      throw new Exception('カートが空なので注文できません');
    }

    $count = 0;
    $total = 0;
    foreach ($this->products as $stock) {
      $count += $stock['count'];
      $total += $stock['product']->price * $stock['count'];
    }

    return compact('count', 'total');
  }
}
