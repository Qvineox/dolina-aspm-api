package utils

import "github.com/google/uuid"

func Unique[T comparable](in []T) []T {
	seen := make(map[T]struct{}, len(in))

	out := make([]T, 0, len(in))
	for _, v := range in {
		if _, ok := seen[v]; ok {
			continue
		}

		seen[v] = struct{}{}
		out = append(out, v)
	}

	return out
}

type ComparableByUUID interface {
	GetUUID() uuid.UUID
}

func UniqueByUUID[T ComparableByUUID](in []T) []T {
	seen := make(map[string]struct{}, len(in))
	var result []T

	for _, item := range in {
		id := item.GetUUID().String()
		if _, exists := seen[id]; !exists {
			seen[id] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}
