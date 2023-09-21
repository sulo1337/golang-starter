package store

type Store struct {
	UserStore UserStore
	PostStore PostStore
}

func NewStore(
	userStore UserStore,
	postStore PostStore,
) *Store {
	return &Store{
		UserStore: userStore,
		PostStore: postStore,
	}
}
