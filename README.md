# Kelly Critereon

The "Kelly Criterion" a is formula to determine the optimal risk size to maximize long-term growth.

Given

- winRate : percentage of winning trades e.g. 0.55 represents 55%
- lossRate : equals 1.0 - winRate
- reward : risk multiple returned on a winning trade. e.g. 2.0 (return is 2.0 times the risk)

With a 55% win rate and 2:1 reward-to-risk ratio, the Kelly formula suggests:

```math
  (winRate * reward) - (lossRate) / reward
= ((0.55 * 2) - (1 - 0.55)) / 2.0
= (1.1 - 0.45) / 2.0
= 0.65 / 2.0
= 0.325
```

## Usage

```go
rate, err := kellycriteron.Calculate(0.55, 2.0)
if err != nil {
  return err
}

// rate will be 0.325
```

## CLI

A cli tool has been provided.

Install it from source with ...

`$ make install`

Usage example

```bash
$ kelly-critereon -winrate 0.55 -reward 2.0
0.325
```

## References

- [Kelly criterion](https://en.wikipedia.org/wiki/Kelly_criterion)

## License

The MIT License (MIT)

Copyright (c) 2025 Scott Barr

See [LICENSE](LICENSE)
