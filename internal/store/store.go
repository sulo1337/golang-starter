package store

type Store struct {
	userStore UserStore
	postStore PostStore
}

func NewStore(
	userStore UserStore,
	postStore PostStore,
) *Store {
	return &Store{
		userStore: userStore,
		postStore: postStore,
	}
}