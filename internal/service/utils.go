package service

func copyPatch[T any](input *T, current T) (params T) {
	params = current
	if input != nil {
		params = *input
	}
	return
}
