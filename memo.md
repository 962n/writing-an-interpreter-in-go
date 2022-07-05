p.28 2.2 パーサージェネレータじゃないの?  
後で調べてみる



let <identifier> = <expression>;

```
let x = 10;
let y = 15;
let add = fn(a, b) { return a + b;
};
```

et文には2つの可変部がある。1つは識別子で、もう1つは式だ。  
上の例では、識別子は x、y、 add だ。式は 10、15、関数リテラルだ。  
先に進む前に、式と文の違いについて少し説明しておく必要がある。式は値を生成し、文はしない。   
例えば、5は値を生成し(この式が生成する値は5だ)、let x = 5は値を生成しない。
  
  
Partt Parsing
後で調べる

  
Off-by-oneエラーとは、境界条件の判定に関するエラーの一種である。コンピュータプログラミングにおいて、ループが正しい回数より一回多く、または一回少なく実行された場合などに発生する。

Top Down Operator Precedence  
http://javascript.crockford.com/tdop/tdop.html  
http://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/


