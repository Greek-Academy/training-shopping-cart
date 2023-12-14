<?php

namespace App;

final class Product
{
  /**
   * @var string
   */
  public string $name;

  /**
   * @var int
   */
  public int $price;

  /**
   * @param string $code apple|grape|orange
   */
  public function __construct(
    public string $code
  ) {
    [$this->name, $this->price] = match ($this->code) {
      'apple' => ['りんご', 50],
      'grape' => ['ぶどう', 100],
      'orange' => ['みかん', 70],
      default => throw new Exception('code値が不正です'),
    };
  }
}
