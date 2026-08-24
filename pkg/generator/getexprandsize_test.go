/*
 * Warp (C) 2019-2020 MinIO, Inc.
 * Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package generator

import (
	"math/rand"
	"testing"
)

func TestGetExpRandSizeMinSizeOne(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	const maxSize = 1024 * 1024
	const n = 1000

	var min, max int64 = maxSize, 0
	seen := make(map[int64]struct{}, n)
	for i := 0; i < n; i++ {
		size := GetExpRandSize(rng, 1, maxSize)
		if size <= 0 || size > maxSize {
			t.Fatalf("invalid size %d, want 1..%d", size, maxSize)
		}
		if size < min {
			min = size
		}
		if size > max {
			max = size
		}
		seen[size] = struct{}{}
	}
	if len(seen) < 10 {
		t.Fatalf("sizes collapsed with minSize=1: got %d distinct sizes", len(seen))
	}
	if max < maxSize/4 {
		t.Fatalf("expected some large sizes, got max %d", max)
	}
	if min == max {
		t.Fatal("all generated sizes were identical")
	}
}
