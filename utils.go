package ga

func pipe[T any](data T, handlers ...func(T) T) T {
	for _, handler := range handlers {
		data = handler(data)
	}
	return data
}
