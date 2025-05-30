// Copyright 2025 Bytedance Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stream

import (
	"github.com/bytedance/gg/collection/tuple"
	"github.com/bytedance/gg/goption"
	"github.com/bytedance/gg/internal/iter"
)

// See function [github.com/bytedance/gg/internal/iter.FromMap].
func FromMap[K comparable, V any](m map[K]V) KV[K, V] {
	return KV[K, V]{FromIter(iter.FromMap(m))}
}

// See function [github.com/bytedance/gg/internal/iter.Repeat].
func RepeatKV[K comparable, V any](k K, v V) KV[K, V] {
	return KV[K, V]{Repeat(tuple.Make2(k, v))}
}

// See function [github.com/bytedance/gg/internal/iter.MapInplace].
func (s KV[K, V]) Map(f func(k K, v V) (K, V)) KV[K, V] {
	return KV[K, V]{s.Stream.Map(func(v tuple.T2[K, V]) tuple.T2[K, V] {
		return tuple.Make2(f(v.Values()))
	})}
}

// See function [github.com/bytedance/gg/internal/iter.Fold].
func (s KV[K, V]) Fold(f func(K, V, K, V) (K, V), initK K, initV V) tuple.T2[K, V] {
	return s.Stream.Fold(func(acc, v tuple.T2[K, V]) tuple.T2[K, V] {
		return tuple.Make2(f(acc.First, acc.Second, v.First, v.Second))
	}, tuple.Make2(initK, initV))
}

// See function [github.com/bytedance/gg/internal/iter.Reduce].
func (s KV[K, V]) Reduce(f func(K, V, K, V) (K, V)) goption.O[tuple.T2[K, V]] {
	return s.Stream.Reduce(func(acc, v tuple.T2[K, V]) tuple.T2[K, V] {
		return tuple.Make2(f(acc.First, acc.Second, v.First, v.Second))
	})
}

// See function [github.com/bytedance/gg/internal/iter.Filter].
func (s KV[K, V]) Filter(f func(K, V) bool) KV[K, V] {
	return KV[K, V]{s.Stream.Filter(func(v tuple.T2[K, V]) bool {
		return f(v.Values())
	})}
}

// See function [github.com/bytedance/gg/internal/iter.ForEach].
func (s KV[K, V]) ForEach(f func(K, V)) {
	s.Stream.ForEach(func(v tuple.T2[K, V]) {
		f(v.Values())
	})
}

// See function [github.com/bytedance/gg/internal/iter.ForEachIndexed].
func (s KV[K, V]) ForEachIndexed(f func(int, K, V)) {
	s.Stream.ForEachIndexed(func(i int, v tuple.T2[K, V]) {
		f(i, v.First, v.Second)
	})
}

// See function [github.com/bytedance/gg/internal/iter.All].
func (s KV[K, V]) All(f func(K, V) bool) bool {
	return s.Stream.All(func(v tuple.T2[K, V]) bool {
		return f(v.Values())
	})
}

// See function [github.com/bytedance/gg/internal/iter.Any].
func (s KV[K, V]) Any(f func(K, V) bool) bool {
	return s.Stream.Any(func(v tuple.T2[K, V]) bool {
		return f(v.Values())
	})
}

// See function [github.com/bytedance/gg/internal/iter.Zip].
func (s KV[K, V]) Zip(f func(K, V, K, V) (K, V), another KV[K, V]) KV[K, V] {
	return KV[K, V]{s.Stream.Zip(func(v1, v2 tuple.T2[K, V]) tuple.T2[K, V] {
		return tuple.Make2(f(v1.First, v1.Second, v2.First, v2.Second))
	}, another.Stream)}
}

// See function [github.com/bytedance/gg/internal/iter.Intersperse].
func (s KV[K, V]) Intersperse(sepK K, sepV V) KV[K, V] {
	return KV[K, V]{s.Stream.Intersperse(tuple.Make2(sepK, sepV))}
}

// See function [github.com/bytedance/gg/internal/iter.Append].
func (s KV[K, V]) Append(tailK K, tailV V) KV[K, V] {
	return KV[K, V]{s.Stream.Append(tuple.Make2(tailK, tailV))}
}

// See function [github.com/bytedance/gg/internal/iter.Prepend].
func (s KV[K, V]) Prepend(tailK K, tailV V) KV[K, V] {
	return KV[K, V]{s.Stream.Prepend(tuple.Make2(tailK, tailV))}
}

// See function [github.com/bytedance/gg/internal/iter.Find].
func (s KV[K, V]) Find(f func(K, V) bool) goption.O[tuple.T2[K, V]] {
	return s.Stream.Find(func(v tuple.T2[K, V]) bool {
		return f(v.Values())
	})
}

// See function [github.com/bytedance/gg/internal/iter.TakeWhile].
func (s KV[K, V]) TakeWhile(f func(K, V) bool) KV[K, V] {
	return KV[K, V]{s.Stream.TakeWhile(func(v tuple.T2[K, V]) bool {
		return f(v.Values())
	})}
}

// See function [github.com/bytedance/gg/internal/iter.DropWhile].
func (s KV[K, V]) DropWhile(f func(K, V) bool) KV[K, V] {
	return KV[K, V]{s.Stream.DropWhile(func(v tuple.T2[K, V]) bool {
		return f(v.Values())
	})}
}

// See function [github.com/bytedance/gg/internal/iter.KVToMap].
func (s KV[K, V]) ToMap() map[K]V {
	return iter.KVToMap(s.Iter)
}

// Keys returns stream of key.
func (s KV[K, V]) Keys() Comparable[K] {
	return FromComparableIter(iter.Map(func(v tuple.T2[K, V]) K {
		return v.First
	}, s.Iter))
}

// Values returns stream of value.
func (s KV[K, V]) Values() Stream[V] {
	return FromIter(iter.Map(func(v tuple.T2[K, V]) V {
		return v.Second
	}, s.Iter))
}

// See function [github.com/bytedance/gg/internal/iter.SortBy].
func (s KV[K, V]) SortBy(less func(K, V, K, V) bool) KV[K, V] {
	return KV[K, V]{s.Stream.SortBy(func(t1, t2 tuple.T2[K, V]) bool {
		return less(t1.First, t1.Second, t2.First, t2.Second)
	})}
}

// See function [github.com/bytedance/gg/internal/iter.UniqBy].
func (s KV[K, V]) UniqBy(f func(K, V) any) KV[K, V] {
	return KV[K, V]{s.Stream.UniqBy(func(t tuple.T2[K, V]) any {
		return f(t.Values())
	})}
}
